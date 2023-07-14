// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: wx-miniprogram/login.proto

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


// Client API for Login service

type LoginService interface {
	// Code2Session  小程序登录 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
	Code2Session(ctx context.Context, in *Code2SessionReq, opts ...Option) (*Code2SessionRes, error)
}

type loginService struct {
	// opts
	opts *Options
}

func NewLoginService(opts ...Option) LoginService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://api.weixin.qq.com"
	}
	return &loginService {
		opts: opt,
	}
}


func (c *loginService) Code2Session(ctx context.Context, in *Code2SessionReq, opts ...Option) (*Code2SessionRes, error) {
	var res Code2SessionRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/sns/jscode2session", opt.addr)

	// body
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAppid() != "" {
		params.Add("appid", fmt.Sprintf("%v", in.GetAppid()))
	}
	if in.GetGrantType() != "" {
		params.Add("grant_type", fmt.Sprintf("%v", in.GetGrantType()))
	}
	if in.GetJsCode() != "" {
		params.Add("js_code", fmt.Sprintf("%v", in.GetJsCode()))
	}
	if in.GetSecret() != "" {
		params.Add("secret", fmt.Sprintf("%v", in.GetSecret()))
	}	
	req.URL.RawQuery = params.Encode()
	
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
