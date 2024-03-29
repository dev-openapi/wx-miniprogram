// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: wx-miniprogram/token.proto

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


// Client API for Token service

type TokenService interface {
	// GetAccessToken  获取接口调用凭据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
	GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...Option) (*GetAccessTokenRes, error)
}

type tokenService struct {
	// opts
	opts *Options
}

func NewTokenService(opts ...Option) TokenService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://api.weixin.qq.com"
	}
	return &tokenService {
		opts: opt,
	}
}


func (c *tokenService) GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...Option) (*GetAccessTokenRes, error) {
	var res GetAccessTokenRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/token", opt.addr)

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
