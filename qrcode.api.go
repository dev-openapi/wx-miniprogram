// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: wx-miniprogram/qrcode.proto

package wx_miniprogram

import (
	context "context"
	fmt "fmt"
	io "io"
	json "encoding/json"
	bytes "bytes"
	http "net/http"
	strings "strings"
	url "net/url"
	multipart "mime/multipart"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = context.Background
var _ = http.NewRequest
var _ = io.Copy
var _ = bytes.Compare
var _ = json.Marshal
var _ = strings.Compare
var _ = fmt.Errorf
var _ = url.Parse
var _ = multipart.ErrMessageTooLarge


// Client API for Qrcode service

type QrcodeService interface {
	// GetQrcode  获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
	GetQrcode(ctx context.Context, in *GetQrcodeReq, opts ...Option) (*GetQrcodeRes, error)
	// GetUnlimitedQrcode  获取不限制的小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
	GetUnlimitedQrcode(ctx context.Context, in *GetUnlimitedQrcodeReq, opts ...Option) (*GetUnlimitedQrcodeRes, error)
	// CreateQrcode  获取小程序二维码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
	CreateQrcode(ctx context.Context, in *CreateQrcodeReq, opts ...Option) (*CreateQrcodeRes, error)
	// QueryScheme  查询小程序scheme码
	QueryScheme(ctx context.Context, in *QuerySchemeReq, opts ...Option) (*QuerySchemeRes, error)
	// GenerateScheme  获取scheme码
	GenerateScheme(ctx context.Context, in *GenerateSchemeReq, opts ...Option) (*GenerateSchemeRes, error)
}

type qrcodeService struct {
	// opts
	opts *Options
}

func NewQrcodeService(opts ...Option) QrcodeService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://api.weixin.qq.com"
	}
	return &qrcodeService {
		opts: opt,
	}
}


func (c *qrcodeService) GetQrcode(ctx context.Context, in *GetQrcodeReq, opts ...Option) (*GetQrcodeRes, error) {
	var res GetQrcodeRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/getwxacode?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *qrcodeService) GetUnlimitedQrcode(ctx context.Context, in *GetUnlimitedQrcodeReq, opts ...Option) (*GetUnlimitedQrcodeRes, error) {
	var res GetUnlimitedQrcodeRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/getwxacodeunlimit?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *qrcodeService) CreateQrcode(ctx context.Context, in *CreateQrcodeReq, opts ...Option) (*CreateQrcodeRes, error) {
	var res CreateQrcodeRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/wxaapp/createwxaqrcode?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *qrcodeService) QueryScheme(ctx context.Context, in *QuerySchemeReq, opts ...Option) (*QuerySchemeRes, error) {
	var res QuerySchemeRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/queryscheme?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *qrcodeService) GenerateScheme(ctx context.Context, in *GenerateSchemeReq, opts ...Option) (*GenerateSchemeRes, error) {
	var res GenerateSchemeRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/generatescheme?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}
