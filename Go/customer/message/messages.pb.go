// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: messages.proto

package message

import (
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

type ReceiveOrder_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ItemId         string  `protobuf:"bytes,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity       int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AccountBalance float64 `protobuf:"fixed64,4,opt,name=accountBalance,proto3" json:"accountBalance,omitempty"`
	PricePerItem   float64 `protobuf:"fixed64,5,opt,name=pricePerItem,proto3" json:"pricePerItem,omitempty"`
}

func (x *ReceiveOrder_Request) Reset() {
	*x = ReceiveOrder_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveOrder_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveOrder_Request) ProtoMessage() {}

func (x *ReceiveOrder_Request) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveOrder_Request.ProtoReflect.Descriptor instead.
func (*ReceiveOrder_Request) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveOrder_Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReceiveOrder_Request) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *ReceiveOrder_Request) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ReceiveOrder_Request) GetAccountBalance() float64 {
	if x != nil {
		return x.AccountBalance
	}
	return 0
}

func (x *ReceiveOrder_Request) GetPricePerItem() float64 {
	if x != nil {
		return x.PricePerItem
	}
	return 0
}

type Payment_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AccountNumber  string  `protobuf:"bytes,2,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	AccountBalance float64 `protobuf:"fixed64,3,opt,name=accountBalance,proto3" json:"accountBalance,omitempty"`
}

func (x *Payment_Request) Reset() {
	*x = Payment_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment_Request) ProtoMessage() {}

func (x *Payment_Request) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment_Request.ProtoReflect.Descriptor instead.
func (*Payment_Request) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Payment_Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Payment_Request) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Payment_Request) GetAccountBalance() float64 {
	if x != nil {
		return x.AccountBalance
	}
	return 0
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    *Message `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	ReceiverId string   `protobuf:"bytes,2,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Notification) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Notification) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Action  string `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
	OrderId string `protobuf:"bytes,3,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Message) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x77, 0x0a, 0x0f, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x55, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_goTypes = []interface{}{
	(*ReceiveOrder_Request)(nil), // 0: messages.ReceiveOrder_Request
	(*Payment_Request)(nil),      // 1: messages.Payment_Request
	(*Notification)(nil),         // 2: messages.Notification
	(*Message)(nil),              // 3: messages.Message
}
var file_messages_proto_depIdxs = []int32{
	3, // 0: messages.Notification.Message:type_name -> messages.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveOrder_Request); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment_Request); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
