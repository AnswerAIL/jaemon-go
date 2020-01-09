// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Sex                  int64    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_eeb1e187a22552a6, []int{0}
}


func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetSex() int64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_eeb1e187a22552a6, []int{1}
}

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "common.User")
	proto.RegisterType((*Users)(nil), "common.Users")
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_entity_eeb1e187a22552a6) }

var fileDescriptor_entity_eeb1e187a22552a6 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcd, 0x2b, 0xc9,
	0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf,
	0x53, 0x0a, 0xe0, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0x92, 0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d,
	0xf2, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05,
	0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10, 0x39, 0x18, 0x5f, 0x48, 0x80, 0x8b,
	0xb9, 0x38, 0xb5, 0x42, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc4, 0x54, 0xd2, 0xe6, 0x62,
	0x05, 0x99, 0x58, 0x2c, 0xa4, 0xc4, 0xc5, 0x0a, 0x32, 0xa2, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83,
	0xdb, 0x88, 0x47, 0x0f, 0x62, 0xa5, 0x1e, 0x48, 0x36, 0x08, 0x22, 0x95, 0xc4, 0x06, 0x76, 0x8d,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x02, 0xad, 0x18, 0x5a, 0x9d, 0x00, 0x00, 0x00,
}