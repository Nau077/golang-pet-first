// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: note.proto

package note_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[0]
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
	return file_note_proto_rawDescGZIP(), []int{0}
}

type NoteContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Email  string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *NoteContent) Reset() {
	*x = NoteContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteContent) ProtoMessage() {}

func (x *NoteContent) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteContent.ProtoReflect.Descriptor instead.
func (*NoteContent) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{1}
}

func (x *NoteContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NoteContent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NoteContent) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NoteContent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NoteContent *NoteContent           `protobuf:"bytes,2,opt,name=note_content,json=noteContent,proto3" json:"note_content,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Record) GetNoteContent() *NoteContent {
	if x != nil {
		return x.NoteContent
	}
	return nil
}

func (x *Record) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Record) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateNoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Author *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Email  *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateNoteInfo) Reset() {
	*x = UpdateNoteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteInfo) ProtoMessage() {}

func (x *UpdateNoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteInfo.ProtoReflect.Descriptor instead.
func (*UpdateNoteInfo) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNoteInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateNoteInfo) GetTitle() *wrapperspb.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *UpdateNoteInfo) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *UpdateNoteInfo) GetAuthor() *wrapperspb.StringValue {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *UpdateNoteInfo) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

type CreateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteContent *NoteContent `protobuf:"bytes,1,opt,name=NoteContent,proto3" json:"NoteContent,omitempty"`
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNoteRequest) GetNoteContent() *NoteContent {
	if x != nil {
		return x.NoteContent
	}
	return nil
}

type CreateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateNoteResponse) Reset() {
	*x = CreateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteResponse) ProtoMessage() {}

func (x *CreateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{5}
}

func (x *CreateNoteResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateNoteInfo *UpdateNoteInfo `protobuf:"bytes,1,opt,name=UpdateNoteInfo,proto3" json:"UpdateNoteInfo,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNoteRequest) GetUpdateNoteInfo() *UpdateNoteInfo {
	if x != nil {
		return x.UpdateNoteInfo
	}
	return nil
}

type UpdateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateNoteResponse) Reset() {
	*x = UpdateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteResponse) ProtoMessage() {}

func (x *UpdateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNoteResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{9}
}

func (x *GetNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{10}
}

func (x *GetNoteResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type GetNoteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GetNoteListResponse) Reset() {
	*x = GetNoteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteListResponse) ProtoMessage() {}

func (x *GetNoteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteListResponse.ProtoReflect.Descriptor instead.
func (*GetNoteListResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{11}
}

func (x *GetNoteListResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_note_proto protoreflect.FileDescriptor

var file_note_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6e, 0x0a, 0x0b, 0x4e, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xcb, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x4e, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x58, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x5f, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x32, 0x80, 0x04, 0x0a, 0x0b, 0x4e, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22,
	0x0f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x0f, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5f,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x61, 0x75, 0x30, 0x37, 0x37,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x70, 0x65, 0x74, 0x2d, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_note_proto_rawDescOnce sync.Once
	file_note_proto_rawDescData = file_note_proto_rawDesc
)

func file_note_proto_rawDescGZIP() []byte {
	file_note_proto_rawDescOnce.Do(func() {
		file_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_note_proto_rawDescData)
	})
	return file_note_proto_rawDescData
}

var file_note_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_note_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: api.note_v1.Empty
	(*NoteContent)(nil),            // 1: api.note_v1.NoteContent
	(*Record)(nil),                 // 2: api.note_v1.Record
	(*UpdateNoteInfo)(nil),         // 3: api.note_v1.UpdateNoteInfo
	(*CreateNoteRequest)(nil),      // 4: api.note_v1.CreateNoteRequest
	(*CreateNoteResponse)(nil),     // 5: api.note_v1.CreateNoteResponse
	(*DeleteNoteRequest)(nil),      // 6: api.note_v1.DeleteNoteRequest
	(*UpdateNoteRequest)(nil),      // 7: api.note_v1.UpdateNoteRequest
	(*UpdateNoteResponse)(nil),     // 8: api.note_v1.UpdateNoteResponse
	(*GetNoteRequest)(nil),         // 9: api.note_v1.GetNoteRequest
	(*GetNoteResponse)(nil),        // 10: api.note_v1.GetNoteResponse
	(*GetNoteListResponse)(nil),    // 11: api.note_v1.GetNoteListResponse
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 13: google.protobuf.StringValue
}
var file_note_proto_depIdxs = []int32{
	1,  // 0: api.note_v1.Record.note_content:type_name -> api.note_v1.NoteContent
	12, // 1: api.note_v1.Record.created_at:type_name -> google.protobuf.Timestamp
	12, // 2: api.note_v1.Record.updated_at:type_name -> google.protobuf.Timestamp
	13, // 3: api.note_v1.UpdateNoteInfo.title:type_name -> google.protobuf.StringValue
	13, // 4: api.note_v1.UpdateNoteInfo.text:type_name -> google.protobuf.StringValue
	13, // 5: api.note_v1.UpdateNoteInfo.author:type_name -> google.protobuf.StringValue
	13, // 6: api.note_v1.UpdateNoteInfo.email:type_name -> google.protobuf.StringValue
	1,  // 7: api.note_v1.CreateNoteRequest.NoteContent:type_name -> api.note_v1.NoteContent
	3,  // 8: api.note_v1.UpdateNoteRequest.UpdateNoteInfo:type_name -> api.note_v1.UpdateNoteInfo
	2,  // 9: api.note_v1.GetNoteResponse.record:type_name -> api.note_v1.Record
	2,  // 10: api.note_v1.GetNoteListResponse.records:type_name -> api.note_v1.Record
	4,  // 11: api.note_v1.NoteService.CreateNote:input_type -> api.note_v1.CreateNoteRequest
	6,  // 12: api.note_v1.NoteService.DeleteNote:input_type -> api.note_v1.DeleteNoteRequest
	7,  // 13: api.note_v1.NoteService.UpdateNote:input_type -> api.note_v1.UpdateNoteRequest
	9,  // 14: api.note_v1.NoteService.GetNote:input_type -> api.note_v1.GetNoteRequest
	0,  // 15: api.note_v1.NoteService.GetNoteList:input_type -> api.note_v1.Empty
	5,  // 16: api.note_v1.NoteService.CreateNote:output_type -> api.note_v1.CreateNoteResponse
	0,  // 17: api.note_v1.NoteService.DeleteNote:output_type -> api.note_v1.Empty
	8,  // 18: api.note_v1.NoteService.UpdateNote:output_type -> api.note_v1.UpdateNoteResponse
	10, // 19: api.note_v1.NoteService.GetNote:output_type -> api.note_v1.GetNoteResponse
	11, // 20: api.note_v1.NoteService.GetNoteList:output_type -> api.note_v1.GetNoteListResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_note_proto_init() }
func file_note_proto_init() {
	if File_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteContent); i {
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
		file_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_note_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteInfo); i {
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
		file_note_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteRequest); i {
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
		file_note_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteResponse); i {
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
		file_note_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteRequest); i {
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
		file_note_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteRequest); i {
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
		file_note_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteResponse); i {
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
		file_note_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteRequest); i {
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
		file_note_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteResponse); i {
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
		file_note_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteListResponse); i {
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
			RawDescriptor: file_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_note_proto_goTypes,
		DependencyIndexes: file_note_proto_depIdxs,
		MessageInfos:      file_note_proto_msgTypes,
	}.Build()
	File_note_proto = out.File
	file_note_proto_rawDesc = nil
	file_note_proto_goTypes = nil
	file_note_proto_depIdxs = nil
}
