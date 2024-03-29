// 微信小程序接口-接口调用凭证 weixin-miniprogram

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: wx-miniprogram/token.proto

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

type GetAccessTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantType string `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"` // 填写 client_credential
	Appid     string `protobuf:"bytes,2,opt,name=appid,proto3" json:"appid,omitempty"`                          // 小程序唯一凭证，即 AppID，可在「微信公众平台 - 设置 - 开发设置」页中获得。（需要已经成为开发者，且帐号没有异常状态）
	Secret    string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`                        //小程序唯一凭证密钥，即 AppSecret，获取方式同 appid
}

func (x *GetAccessTokenReq) Reset() {
	*x = GetAccessTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_miniprogram_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenReq) ProtoMessage() {}

func (x *GetAccessTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_wx_miniprogram_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenReq.ProtoReflect.Descriptor instead.
func (*GetAccessTokenReq) Descriptor() ([]byte, []int) {
	return file_wx_miniprogram_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccessTokenReq) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *GetAccessTokenReq) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetAccessTokenReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type GetAccessTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int32  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`      // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}

func (x *GetAccessTokenRes) Reset() {
	*x = GetAccessTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_miniprogram_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenRes) ProtoMessage() {}

func (x *GetAccessTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_wx_miniprogram_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenRes.ProtoReflect.Descriptor instead.
func (*GetAccessTokenRes) Descriptor() ([]byte, []int) {
	return file_wx_miniprogram_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessTokenRes) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetAccessTokenRes) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_wx_miniprogram_token_proto protoreflect.FileDescriptor

var file_wx_miniprogram_token_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x78, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x71, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x55, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x32, 0x84, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x71, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x71, 0x71, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x63, 0x67, 0x69, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wx_miniprogram_token_proto_rawDescOnce sync.Once
	file_wx_miniprogram_token_proto_rawDescData = file_wx_miniprogram_token_proto_rawDesc
)

func file_wx_miniprogram_token_proto_rawDescGZIP() []byte {
	file_wx_miniprogram_token_proto_rawDescOnce.Do(func() {
		file_wx_miniprogram_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_wx_miniprogram_token_proto_rawDescData)
	})
	return file_wx_miniprogram_token_proto_rawDescData
}

var file_wx_miniprogram_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wx_miniprogram_token_proto_goTypes = []interface{}{
	(*GetAccessTokenReq)(nil), // 0: api.weixin.qq.com.GetAccessTokenReq
	(*GetAccessTokenRes)(nil), // 1: api.weixin.qq.com.GetAccessTokenRes
}
var file_wx_miniprogram_token_proto_depIdxs = []int32{
	0, // 0: api.weixin.qq.com.TokenService.GetAccessToken:input_type -> api.weixin.qq.com.GetAccessTokenReq
	1, // 1: api.weixin.qq.com.TokenService.GetAccessToken:output_type -> api.weixin.qq.com.GetAccessTokenRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wx_miniprogram_token_proto_init() }
func file_wx_miniprogram_token_proto_init() {
	if File_wx_miniprogram_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wx_miniprogram_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenReq); i {
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
		file_wx_miniprogram_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenRes); i {
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
			RawDescriptor: file_wx_miniprogram_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wx_miniprogram_token_proto_goTypes,
		DependencyIndexes: file_wx_miniprogram_token_proto_depIdxs,
		MessageInfos:      file_wx_miniprogram_token_proto_msgTypes,
	}.Build()
	File_wx_miniprogram_token_proto = out.File
	file_wx_miniprogram_token_proto_rawDesc = nil
	file_wx_miniprogram_token_proto_goTypes = nil
	file_wx_miniprogram_token_proto_depIdxs = nil
}
