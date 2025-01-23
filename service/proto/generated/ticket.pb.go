// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: service/proto/ticket.proto

package generated

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

type SubmitTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User      *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid int32  `protobuf:"varint,4,opt,name=pricePaid,proto3" json:"pricePaid,omitempty"`
}

func (x *SubmitTicketRequest) Reset() {
	*x = SubmitTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTicketRequest) ProtoMessage() {}

func (x *SubmitTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTicketRequest.ProtoReflect.Descriptor instead.
func (*SubmitTicketRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitTicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SubmitTicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SubmitTicketRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SubmitTicketRequest) GetPricePaid() int32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

type GetDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetDetailsRequest) Reset() {
	*x = GetDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailsRequest) ProtoMessage() {}

func (x *GetDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetDetailsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *GetDetailsRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUsersInSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *GetUsersInSectionRequest) Reset() {
	*x = GetUsersInSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersInSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersInSectionRequest) ProtoMessage() {}

func (x *GetUsersInSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersInSectionRequest.ProtoReflect.Descriptor instead.
func (*GetUsersInSectionRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersInSectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ModifyUserSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewSeat string `protobuf:"bytes,2,opt,name=newSeat,proto3" json:"newSeat,omitempty"`
}

func (x *ModifyUserSeatRequest) Reset() {
	*x = ModifyUserSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyUserSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyUserSeatRequest) ProtoMessage() {}

func (x *ModifyUserSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyUserSeatRequest.ProtoReflect.Descriptor instead.
func (*ModifyUserSeatRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ModifyUserSeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifyUserSeatRequest) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TicketReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From          string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User          *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid     int32  `protobuf:"varint,4,opt,name=pricePaid,proto3" json:"pricePaid,omitempty"`
	AllocatedSeat string `protobuf:"bytes,5,opt,name=allocatedSeat,proto3" json:"allocatedSeat,omitempty"`
}

func (x *TicketReceipt) Reset() {
	*x = TicketReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketReceipt) ProtoMessage() {}

func (x *TicketReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketReceipt.ProtoReflect.Descriptor instead.
func (*TicketReceipt) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *TicketReceipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketReceipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TicketReceipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TicketReceipt) GetPricePaid() int32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *TicketReceipt) GetAllocatedSeat() string {
	if x != nil {
		return x.AllocatedSeat
	}
	return ""
}

type UsersInSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string          `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
	Users   []*UserWithSeat `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersInSection) Reset() {
	*x = UsersInSection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersInSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersInSection) ProtoMessage() {}

func (x *UsersInSection) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersInSection.ProtoReflect.Descriptor instead.
func (*UsersInSection) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *UsersInSection) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *UsersInSection) GetUsers() []*UserWithSeat {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserWithSeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AllocatedSeat string `protobuf:"bytes,2,opt,name=allocatedSeat,proto3" json:"allocatedSeat,omitempty"`
}

func (x *UserWithSeat) Reset() {
	*x = UserWithSeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWithSeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithSeat) ProtoMessage() {}

func (x *UserWithSeat) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithSeat.ProtoReflect.Descriptor instead.
func (*UserWithSeat) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *UserWithSeat) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserWithSeat) GetAllocatedSeat() string {
	if x != nil {
		return x.AllocatedSeat
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ModifyUserSeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ModifyUserSeatResponse) Reset() {
	*x = ModifyUserSeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyUserSeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyUserSeatResponse) ProtoMessage() {}

func (x *ModifyUserSeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyUserSeatResponse.ProtoReflect.Descriptor instead.
func (*ModifyUserSeatResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *ModifyUserSeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ModifyUserSeatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_service_proto_ticket_proto protoreflect.FileDescriptor

var file_service_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x79, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x22,
	0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x34, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x29, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x47, 0x0a, 0x15, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x77, 0x53, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x53, 0x65, 0x61, 0x74, 0x22, 0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x99, 0x01, 0x0a,
	0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x65, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x22, 0x56, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x56, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x61, 0x74,
	0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4c, 0x0a, 0x16, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xf9, 0x02, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_ticket_proto_rawDescOnce sync.Once
	file_service_proto_ticket_proto_rawDescData = file_service_proto_ticket_proto_rawDesc
)

func file_service_proto_ticket_proto_rawDescGZIP() []byte {
	file_service_proto_ticket_proto_rawDescOnce.Do(func() {
		file_service_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_ticket_proto_rawDescData)
	})
	return file_service_proto_ticket_proto_rawDescData
}

var file_service_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_proto_ticket_proto_goTypes = []interface{}{
	(*SubmitTicketRequest)(nil),      // 0: ticket.SubmitTicketRequest
	(*GetDetailsRequest)(nil),        // 1: ticket.GetDetailsRequest
	(*GetUsersInSectionRequest)(nil), // 2: ticket.GetUsersInSectionRequest
	(*RemoveUserRequest)(nil),        // 3: ticket.RemoveUserRequest
	(*ModifyUserSeatRequest)(nil),    // 4: ticket.ModifyUserSeatRequest
	(*User)(nil),                     // 5: ticket.User
	(*TicketReceipt)(nil),            // 6: ticket.TicketReceipt
	(*UsersInSection)(nil),           // 7: ticket.UsersInSection
	(*UserWithSeat)(nil),             // 8: ticket.UserWithSeat
	(*RemoveUserResponse)(nil),       // 9: ticket.RemoveUserResponse
	(*ModifyUserSeatResponse)(nil),   // 10: ticket.ModifyUserSeatResponse
}
var file_service_proto_ticket_proto_depIdxs = []int32{
	5,  // 0: ticket.SubmitTicketRequest.user:type_name -> ticket.User
	5,  // 1: ticket.TicketReceipt.user:type_name -> ticket.User
	8,  // 2: ticket.UsersInSection.users:type_name -> ticket.UserWithSeat
	5,  // 3: ticket.UserWithSeat.user:type_name -> ticket.User
	0,  // 4: ticket.TicketService.SubmitRequest:input_type -> ticket.SubmitTicketRequest
	1,  // 5: ticket.TicketService.GetDetails:input_type -> ticket.GetDetailsRequest
	2,  // 6: ticket.TicketService.GetUsersInSection:input_type -> ticket.GetUsersInSectionRequest
	3,  // 7: ticket.TicketService.RemoveUser:input_type -> ticket.RemoveUserRequest
	4,  // 8: ticket.TicketService.ModifyUserSeat:input_type -> ticket.ModifyUserSeatRequest
	6,  // 9: ticket.TicketService.SubmitRequest:output_type -> ticket.TicketReceipt
	6,  // 10: ticket.TicketService.GetDetails:output_type -> ticket.TicketReceipt
	7,  // 11: ticket.TicketService.GetUsersInSection:output_type -> ticket.UsersInSection
	9,  // 12: ticket.TicketService.RemoveUser:output_type -> ticket.RemoveUserResponse
	10, // 13: ticket.TicketService.ModifyUserSeat:output_type -> ticket.ModifyUserSeatResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_ticket_proto_init() }
func file_service_proto_ticket_proto_init() {
	if File_service_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitTicketRequest); i {
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
		file_service_proto_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailsRequest); i {
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
		file_service_proto_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersInSectionRequest); i {
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
		file_service_proto_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserRequest); i {
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
		file_service_proto_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyUserSeatRequest); i {
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
		file_service_proto_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_service_proto_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketReceipt); i {
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
		file_service_proto_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersInSection); i {
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
		file_service_proto_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWithSeat); i {
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
		file_service_proto_ticket_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserResponse); i {
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
		file_service_proto_ticket_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyUserSeatResponse); i {
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
			RawDescriptor: file_service_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_ticket_proto_goTypes,
		DependencyIndexes: file_service_proto_ticket_proto_depIdxs,
		MessageInfos:      file_service_proto_ticket_proto_msgTypes,
	}.Build()
	File_service_proto_ticket_proto = out.File
	file_service_proto_ticket_proto_rawDesc = nil
	file_service_proto_ticket_proto_goTypes = nil
	file_service_proto_ticket_proto_depIdxs = nil
}
