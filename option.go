// Generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.

package wx_miniprogram

import (
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"context"
)

type Option func(*Options)

type FnRequest func(context.Context, *http.Client,*http.Request) (*http.Response, error)
type FnResponse func(context.Context, *http.Response, interface{}) error

var (
	ErrNil = errors.New("resp nil")
	ErrNot200 = errors.New("resp not 200")
)

type Options struct {
	// do request
	DoRequest FnRequest
	// do response
	DoResponse FnResponse
	// addr
	addr string
	// client
	client *http.Client
}

func newOptions(opts ...Option) *Options {
	opt := Options{
		client: http.DefaultClient,
		DoRequest: doRequest,
		DoResponse: doResponse,
	}
	for _, o := range opts {
		o(&opt)
	}
	return &opt
}

func buildOptions(opt *Options, opts ...Option) *Options {
	res := newOptions(opts...)
	if len(res.addr) == 0 {
		res.addr = opt.addr
	}
	return res
}

func doRequest(_ context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func doResponse(_ context.Context, resp *http.Response, a interface{}) error {
	if resp == nil {
		return ErrNil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ErrNot200
	}
	defer func(){ _ = resp.Body.Close()}()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, a)
}

func WithDoRequest(fn FnRequest) Option {
	return func(o *Options) {
		o.DoRequest = fn
	}
}

func WithDoResponse(fn FnResponse) Option {
	return func(o *Options) {
		o.DoResponse = fn
	}
}

func WithClient(c *http.Client) Option {
	return func(o *Options) {
		o.client = c
	}
}

// addr must start with https:// or http://
func WithAddr(addr string) Option {
	return func(o *Options) {
		o.addr = addr
	}
}

