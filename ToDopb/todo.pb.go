// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0--rc3
// source: ToDopb/todo.proto

package ToDopb

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

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToDopb_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_ToDopb_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_ToDopb_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type ToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *ToDo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *ToDoResponse) Reset() {
	*x = ToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToDopb_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoResponse) ProtoMessage() {}

func (x *ToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ToDopb_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoResponse.ProtoReflect.Descriptor instead.
func (*ToDoResponse) Descriptor() ([]byte, []int) {
	return file_ToDopb_todo_proto_rawDescGZIP(), []int{1}
}

func (x *ToDoResponse) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToDopb_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ToDopb_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ToDopb_todo_proto_rawDescGZIP(), []int{2}
}

type ToDoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ToDoId) Reset() {
	*x = ToDoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToDopb_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoId) ProtoMessage() {}

func (x *ToDoId) ProtoReflect() protoreflect.Message {
	mi := &file_ToDopb_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoId.ProtoReflect.Descriptor instead.
func (*ToDoId) Descriptor() ([]byte, []int) {
	return file_ToDopb_todo_proto_rawDescGZIP(), []int{3}
}

func (x *ToDoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_ToDopb_todo_proto protoreflect.FileDescriptor

var file_ToDopb_todo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x54, 0x6f, 0x44, 0x6f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x62, 0x0a, 0x04, 0x54, 0x6f, 0x44,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x2e, 0x0a,
	0x0c, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x0a, 0x06, 0x54, 0x6f, 0x44, 0x6f, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xc6, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0a,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x73, 0x12, 0x0b, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x30,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x6e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x54, 0x6f,
	0x44, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ToDopb_todo_proto_rawDescOnce sync.Once
	file_ToDopb_todo_proto_rawDescData = file_ToDopb_todo_proto_rawDesc
)

func file_ToDopb_todo_proto_rawDescGZIP() []byte {
	file_ToDopb_todo_proto_rawDescOnce.Do(func() {
		file_ToDopb_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ToDopb_todo_proto_rawDescData)
	})
	return file_ToDopb_todo_proto_rawDescData
}

var file_ToDopb_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ToDopb_todo_proto_goTypes = []interface{}{
	(*ToDo)(nil),         // 0: todo.ToDo
	(*ToDoResponse)(nil), // 1: todo.ToDoResponse
	(*Empty)(nil),        // 2: todo.empty
	(*ToDoId)(nil),       // 3: todo.ToDoId
}
var file_ToDopb_todo_proto_depIdxs = []int32{
	0, // 0: todo.ToDoResponse.todo:type_name -> todo.ToDo
	0, // 1: todo.ToDoService.CreateToDo:input_type -> todo.ToDo
	2, // 2: todo.ToDoService.ListToDos:input_type -> todo.empty
	3, // 3: todo.ToDoService.CheckUncheck:input_type -> todo.ToDoId
	3, // 4: todo.ToDoService.DeleteToDo:input_type -> todo.ToDoId
	1, // 5: todo.ToDoService.CreateToDo:output_type -> todo.ToDoResponse
	1, // 6: todo.ToDoService.ListToDos:output_type -> todo.ToDoResponse
	1, // 7: todo.ToDoService.CheckUncheck:output_type -> todo.ToDoResponse
	2, // 8: todo.ToDoService.DeleteToDo:output_type -> todo.empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ToDopb_todo_proto_init() }
func file_ToDopb_todo_proto_init() {
	if File_ToDopb_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ToDopb_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_ToDopb_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoResponse); i {
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
		file_ToDopb_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_ToDopb_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoId); i {
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
			RawDescriptor: file_ToDopb_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ToDopb_todo_proto_goTypes,
		DependencyIndexes: file_ToDopb_todo_proto_depIdxs,
		MessageInfos:      file_ToDopb_todo_proto_msgTypes,
	}.Build()
	File_ToDopb_todo_proto = out.File
	file_ToDopb_todo_proto_rawDesc = nil
	file_ToDopb_todo_proto_goTypes = nil
	file_ToDopb_todo_proto_depIdxs = nil
}
