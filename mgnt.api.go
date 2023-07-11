// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.0). DO NOT EDIT.
// source: wx-miniprogram/mgnt.proto

package wx_miniprogram

import (
	context "context"
	fmt "fmt"
	io "io"
	json "encoding/json"
	bytes "bytes"
	http "net/http"
	strings "strings"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = context.Background
var _ = http.NewRequest
var _ = io.Copy
var _ = bytes.Compare
var _ = json.Marshal
var _ = strings.Compare


// Client API for Mgnt service

type MgntService interface {
	// ClearQuota  重置 API 调用次数 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuota.html
	ClearQuota(ctx context.Context, in *ClearQuotaReq, opts ...Option) (*ClearQuotaRes, error)
	// GetApiQuota  查询API调用额度 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getApiQuota.html
	GetApiQuota(ctx context.Context, in *GetApiQuotaReq, opts ...Option) (*GetApiQuotaRes, error)
	// GetRidInfo  查询rid信息 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getRidInfo.html
	GetRidInfo(ctx context.Context, in *GetRidInfoReq, opts ...Option) (*GetRidInfoRes, error)
	// ClearQuotaV2  使用AppSecret重置 API 调用次数 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuotaByAppSecret.html
	ClearQuotaV2(ctx context.Context, in *ClearQuotaV2Req, opts ...Option) (*ClearQuotaV2Res, error)
}

type mgntService struct {
	// opts
	opts *Options
}

func NewMgntService(opts ...Option) MgntService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://api.weixin.qq.com"
	}
	return &mgntService {
		opts: opt,
	}
}


func (c *mgntService) ClearQuota(ctx context.Context, in *ClearQuotaReq, opts ...Option) (*ClearQuotaRes, error) {
	var res ClearQuotaRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/clear_quota?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	var body io.Reader
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(resp, &res)
	return &res, err 

}

func (c *mgntService) GetApiQuota(ctx context.Context, in *GetApiQuotaReq, opts ...Option) (*GetApiQuotaRes, error) {
	var res GetApiQuotaRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/openapi/quota/get?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	var body io.Reader
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(resp, &res)
	return &res, err 

}

func (c *mgntService) GetRidInfo(ctx context.Context, in *GetRidInfoReq, opts ...Option) (*GetRidInfoRes, error) {
	var res GetRidInfoRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/openapi/rid/get?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	var body io.Reader
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(resp, &res)
	return &res, err 

}

func (c *mgntService) ClearQuotaV2(ctx context.Context, in *ClearQuotaV2Req, opts ...Option) (*ClearQuotaV2Res, error) {
	var res ClearQuotaV2Res
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/cgi-bin/clear_quota/v2", opt.addr)

	// body
	var body io.Reader
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(resp, &res)
	return &res, err 

}