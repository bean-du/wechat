package wechat

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"math"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
)

const (
	minRead           = 16 * 1024
	defaultRetryCount = 0
)

type HttpClient struct {
	conf      *Config
	client    *http.Client
	dialer    *net.Dialer
	transport *http.Transport
	retryCount int
	retry Retriable
}

func NewHttpClient(c *Config) *HttpClient  {
	dialer := &net.Dialer{
		Timeout: c.Dial,
		KeepAlive: c.KeepAlive,
	}
	transport := &http.Transport{
		DialContext: dialer.DialContext,
		MaxConnsPerHost: c.MaxConn,
		MaxIdleConnsPerHost: c.MaxIdle,
		IdleConnTimeout: c.KeepAlive,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	_ = http2.ConfigureTransport(transport)
	bo := NewConstantBackoff(c.BackoffInterval)
	return &HttpClient{
		conf: c,
		client: &http.Client{
			Transport: transport,
		},
		retryCount: defaultRetryCount,
		retry: NewRetrier(bo),
	}
}

func (c *HttpClient) SetRetryCount(count int) {
	c.retryCount = count
}


// Get makes a HTTP GET request to provided URL with context passed in
func (c *HttpClient) Get(ctx context.Context, url string, headers http.Header, res interface{}) (err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.New("GET - request creation failed"+ err.Error())
	}

	request.Header = headers

	return c.Do(ctx, request, res)
}

// Post makes a HTTP POST request to provided URL with context passed in
func (c *HttpClient) Post(ctx context.Context, url, contentType string, headers http.Header, param, res interface{}) (err error) {
	request, err := http.NewRequest(http.MethodPost, url, reqBody(contentType, param))
	if err != nil {
		return errors.New("POST - request creation failed " + err.Error())
	}
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request, res)
}

// Put makes a HTTP PUT request to provided URL with context passed in
func (c *HttpClient) Put(ctx context.Context, url, contentType string, headers http.Header, param, res interface{}) (err error) {
	request, err := http.NewRequest(http.MethodPut, url, reqBody(contentType, param))
	if err != nil {
		return errors.New("PUT - request creation failed" + err.Error())
	}

	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request, res)
}

// Patch makes a HTTP PATCH request to provided URL with context passed in
func (c *HttpClient) Patch(ctx context.Context, url, contentType string, headers http.Header, param, res interface{}) (err error) {
	request, err := http.NewRequest(http.MethodPatch, url, reqBody(contentType, param))
	if err != nil {
		return errors.New("PATCH - request creation failed " + err.Error())
	}

	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request, res)
}

// Delete makes a HTTP DELETE request to provided URL with context passed in
func (c *HttpClient) Delete(ctx context.Context, url, contentType string, headers http.Header, param, res interface{}) (err error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return errors.New("DELETE - request creation failed" + err.Error())
	}

	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", contentType)
	request.Header = headers

	return c.Do(ctx, request, res)
}

// Do makes an HTTP request with the native `http.Do` interface and context passed in
func (c *HttpClient) Do(ctx context.Context, req *http.Request, res interface{}) (err error) {
	for i := 0; i <= c.retryCount; i++ {
		if err = c.request(ctx, req, res); err != nil {
			backoffTime := c.retry.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		break
	}
	return
}

func (c *HttpClient) request(ctx context.Context, req *http.Request, res interface{}) (err error) {
	var (
		response *http.Response
		bs       []byte
		cancel   func()
	)
	ctx, cancel = context.WithTimeout(ctx, c.conf.Timeout)
	defer cancel()
	response, err = c.client.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		}
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= http.StatusInternalServerError {
		err = errors.New(fmt.Sprintf("response.StatusCode %d", response.StatusCode))
		return
	}
	if bs, err = readAll(response.Body, minRead); err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	return
}

func reqBody(contentType string, param interface{}) (body io.Reader) {
	var err error
	if contentType == MIMEPOSTForm {
		enc, ok := param.(string)
		if ok {
			body = strings.NewReader(enc)
		}
	}
	if contentType == MIMEJSON {
		buff := new(bytes.Buffer)
		err = json.NewEncoder(buff).Encode(param)
		if err != nil {
			return
		}
		body = buff
	}
	return
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

// Backoff interface defines contract for backoff strategies
type Backoff interface {
	Next(retry int) time.Duration
}

type constantBackoff struct {
	backoffInterval time.Duration
}

// NewConstantBackoff returns an instance of ConstantBackoff
func NewConstantBackoff(backoffInterval time.Duration) Backoff {
	return &constantBackoff{backoffInterval: backoffInterval}
}

// Next returns next time for retrying operation with constant strategy
func (cb *constantBackoff) Next(retry int) time.Duration {
	if retry <= 0 {
		return 0 * time.Millisecond
	}

	return time.Duration(cb.backoffInterval) * 1 << uint(retry)
}

type exponentialBackoff struct {
	exponentFactor float64
	initialTimeout float64
	maxTimeout     float64
}

// Retriable defines contract for retriers to implement
type Retriable interface {
	NextInterval(retry int) time.Duration
	Do(fn RetryFunc, retries int) (err error)
}

type RetryFunc func() (err error)

// RetriableFunc is an adapter to allow the use of ordinary functions
// as a Retriable
type RetriableFunc func(retry int) time.Duration

func (f RetriableFunc) Do(fn RetryFunc, retries int) (err error) {
	return fn()
}

//NextInterval calls f(retry)
func (f RetriableFunc) NextInterval(retry int) time.Duration {
	return f(retry)
}

func (f RetryFunc) Do() (err error) {
	return f()
}

type retrier struct {
	backoff Backoff
}

// NewRetrier returns retrier with some backoff strategy
func NewRetrier(backoff Backoff) Retriable {
	return &retrier{
		backoff: backoff,
	}
}

// NewRetrierFunc returns a retrier with a retry function defined
func NewRetrierFunc(f RetriableFunc) Retriable {
	return f
}

// NextInterval returns next retriable time
func (r *retrier) NextInterval(retry int) time.Duration {
	return r.backoff.Next(retry)
}

type noRetrier struct {
}

// NewNoRetrier returns a null object for retriable
func NewNoRetrier() Retriable {
	return &noRetrier{}
}

// NextInterval returns next retriable time, always 0
func (r *noRetrier) NextInterval(retry int) time.Duration {
	return 0 * time.Millisecond
}

func (r *retrier) Do(fn RetryFunc, retries int) (err error) {
	for i := 0; i <= retries; i++ {
		if err = fn(); err != nil {
			time.Sleep(r.NextInterval(i))
			continue
		}
		break
	}
	return nil
}

func (r *noRetrier) Do(fn RetryFunc, retries int) (err error) {
	for i := 0; i <= retries; i++ {
		if err = fn(); err != nil {
			time.Sleep(r.NextInterval(i))
			continue
		}
		break
	}
	return nil
}

// NewExponentialBackoff returns an instance of ExponentialBackoff
func NewExponentialBackoff(initialTimeout, maxTimeout time.Duration, exponentFactor float64) Backoff {
	return &exponentialBackoff{
		exponentFactor: exponentFactor,
		initialTimeout: float64(initialTimeout / time.Millisecond),
		maxTimeout:     float64(maxTimeout / time.Millisecond),
	}
}

// Next returns next time for retrying operation with exponential strategy
func (eb *exponentialBackoff) Next(retry int) time.Duration {
	if retry <= 0 {
		return 0 * time.Millisecond
	}

	return time.Duration(math.Min(eb.initialTimeout+math.Pow(eb.exponentFactor, float64(retry)), eb.maxTimeout)) * time.Millisecond
}
