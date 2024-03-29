// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: wx-miniprogram/user.proto

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


// Client API for User service

type UserService interface {
	// GetPluginOpenPId  获取插件用户openpid https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/basic-info/getPluginOpenPId.html
	GetPluginOpenPId(ctx context.Context, in *GetPluginOpenPIdReq, opts ...Option) (*GetPluginOpenPidRes, error)
	// CheckEncryptedData  检查加密信息 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/basic-info/checkEncryptedData.html
	CheckEncryptedData(ctx context.Context, in *CheckEncryptedDataReq, opts ...Option) (*CheckEncryptedDataRes, error)
	// GetPaidUnionid  支付后获取 Unionid https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/basic-info/getPaidUnionid.html
	GetPaidUnionid(ctx context.Context, in *GetPaidUnionidReq, opts ...Option) (*GetPaidUnionidRes, error)
	// GetUserEncryptKey  获取用户encryptKey https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/internet/getUserEncryptKey.html
	GetUserEncryptKey(ctx context.Context, in *GetUserEncryptKeyReq, opts ...Option) (*GetUserEncryptKeyRes, error)
	// GetPhoneNumber  获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
	GetPhoneNumber(ctx context.Context, in *GetPhoneNumberReq, opts ...Option) (*GetPhoneNumberRes, error)
}

type userService struct {
	// opts
	opts *Options
}

func NewUserService(opts ...Option) UserService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://api.weixin.qq.com"
	}
	return &userService {
		opts: opt,
	}
}


func (c *userService) GetPluginOpenPId(ctx context.Context, in *GetPluginOpenPIdReq, opts ...Option) (*GetPluginOpenPidRes, error) {
	var res GetPluginOpenPidRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/getpluginopenpid?access_token=%v", opt.addr, in.GetAccessToken())

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

func (c *userService) CheckEncryptedData(ctx context.Context, in *CheckEncryptedDataReq, opts ...Option) (*CheckEncryptedDataRes, error) {
	var res CheckEncryptedDataRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/business/checkencryptedmsg?access_token=%v", opt.addr, in.GetAccessToken())

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

func (c *userService) GetPaidUnionid(ctx context.Context, in *GetPaidUnionidReq, opts ...Option) (*GetPaidUnionidRes, error) {
	var res GetPaidUnionidRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/getpaidunionid?access_token=%v", opt.addr, in.GetAccessToken())

	// body
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetMchId() != "" {
		params.Add("mch_id", fmt.Sprintf("%v", in.GetMchId()))
	}
	if in.GetOpenid() != "" {
		params.Add("openid", fmt.Sprintf("%v", in.GetOpenid()))
	}
	if in.GetOutTradeNo() != "" {
		params.Add("out_trade_no", fmt.Sprintf("%v", in.GetOutTradeNo()))
	}
	if in.GetTransactionId() != "" {
		params.Add("transaction_id", fmt.Sprintf("%v", in.GetTransactionId()))
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

func (c *userService) GetUserEncryptKey(ctx context.Context, in *GetUserEncryptKeyReq, opts ...Option) (*GetUserEncryptKeyRes, error) {
	var res GetUserEncryptKeyRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/business/getuserencryptkey?access_token=%v", opt.addr, in.GetAccessToken())

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

func (c *userService) GetPhoneNumber(ctx context.Context, in *GetPhoneNumberReq, opts ...Option) (*GetPhoneNumberRes, error) {
	var res GetPhoneNumberRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/wxa/business/getuserphonenumber?access_token=%v", opt.addr, in.GetAccessToken())

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
