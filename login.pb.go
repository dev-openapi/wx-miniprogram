// 微信小程序接口-小程序登录 weixin-miniprogram

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: wx-miniprogram/login.proto

package wx_miniprogram

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code2SessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`                          // 小程序 appId
	Secret    string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`                        // 小程序 appSecret
	JsCode    string `protobuf:"bytes,3,opt,name=js_code,json=jsCode,proto3" json:"js_code,omitempty"`          // 登录时获取的 code，可通过wx.login获取
	GrantType string `protobuf:"bytes,4,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"` // 授权类型，此处只需填写 authorization_code
}

func (x *Code2SessionReq) Reset() {
	*x = Code2SessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_miniprogram_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionReq) ProtoMessage() {}

func (x *Code2SessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_wx_miniprogram_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionReq.ProtoReflect.Descriptor instead.
func (*Code2SessionReq) Descriptor() ([]byte, []int) {
	return file_wx_miniprogram_login_proto_rawDescGZIP(), []int{0}
}

func (x *Code2SessionReq) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Code2SessionReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Code2SessionReq) GetJsCode() string {
	if x != nil {
		return x.JsCode
	}
	return ""
}

func (x *Code2SessionReq) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

type Code2SessionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errcode    int32  `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"`                        // 错误码
	Errmsg     string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`                           // 错误信息
	Openid     string `protobuf:"bytes,3,opt,name=openid,proto3" json:"openid,omitempty"`                           // 用户唯一标识
	Unionid    string `protobuf:"bytes,4,opt,name=unionid,proto3" json:"unionid,omitempty"`                         // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
	SessionKey string `protobuf:"bytes,5,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"` // 会话密钥
}

func (x *Code2SessionRes) Reset() {
	*x = Code2SessionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_miniprogram_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionRes) ProtoMessage() {}

func (x *Code2SessionRes) ProtoReflect() protoreflect.Message {
	mi := &file_wx_miniprogram_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionRes.ProtoReflect.Descriptor instead.
func (*Code2SessionRes) Descriptor() ([]byte, []int) {
	return file_wx_miniprogram_login_proto_rawDescGZIP(), []int{1}
}

func (x *Code2SessionRes) GetErrcode() int32 {
	if x != nil {
		return x.Errcode
	}
	return 0
}

func (x *Code2SessionRes) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Code2SessionRes) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *Code2SessionRes) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *Code2SessionRes) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

var File_wx_miniprogram_login_proto protoreflect.FileDescriptor

var file_wx_miniprogram_login_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x78, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x71, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a,
	0x0f, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6a, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x32,
	0x83, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x73, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x71, 0x71,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69,
	0x6e, 0x2e, 0x71, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x73, 0x6e, 0x73, 0x2f, 0x6a, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f,
	0x77, 0x78, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wx_miniprogram_login_proto_rawDescOnce sync.Once
	file_wx_miniprogram_login_proto_rawDescData = file_wx_miniprogram_login_proto_rawDesc
)

func file_wx_miniprogram_login_proto_rawDescGZIP() []byte {
	file_wx_miniprogram_login_proto_rawDescOnce.Do(func() {
		file_wx_miniprogram_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_wx_miniprogram_login_proto_rawDescData)
	})
	return file_wx_miniprogram_login_proto_rawDescData
}

var file_wx_miniprogram_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wx_miniprogram_login_proto_goTypes = []interface{}{
	(*Code2SessionReq)(nil), // 0: api.weixin.qq.com.Code2SessionReq
	(*Code2SessionRes)(nil), // 1: api.weixin.qq.com.Code2SessionRes
}
var file_wx_miniprogram_login_proto_depIdxs = []int32{
	0, // 0: api.weixin.qq.com.LoginService.Code2Session:input_type -> api.weixin.qq.com.Code2SessionReq
	1, // 1: api.weixin.qq.com.LoginService.Code2Session:output_type -> api.weixin.qq.com.Code2SessionRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wx_miniprogram_login_proto_init() }
func file_wx_miniprogram_login_proto_init() {
	if File_wx_miniprogram_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wx_miniprogram_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wx_miniprogram_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wx_miniprogram_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wx_miniprogram_login_proto_goTypes,
		DependencyIndexes: file_wx_miniprogram_login_proto_depIdxs,
		MessageInfos:      file_wx_miniprogram_login_proto_msgTypes,
	}.Build()
	File_wx_miniprogram_login_proto = out.File
	file_wx_miniprogram_login_proto_rawDesc = nil
	file_wx_miniprogram_login_proto_goTypes = nil
	file_wx_miniprogram_login_proto_depIdxs = nil
}
