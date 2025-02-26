// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mempool_status.proto

package libra

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MempoolAddTransactionStatus int32

const (
	// Transaction was sent to Mempool
	MempoolAddTransactionStatus_Valid MempoolAddTransactionStatus = 0
	// The sender does not have enough balance for the transaction.
	MempoolAddTransactionStatus_InsufficientBalance MempoolAddTransactionStatus = 1
	// Sequence number is old, etc.
	MempoolAddTransactionStatus_InvalidSeqNumber MempoolAddTransactionStatus = 2
	// Mempool is full (reached max global capacity)
	MempoolAddTransactionStatus_MempoolIsFull MempoolAddTransactionStatus = 3
	// Account reached max capacity per account
	MempoolAddTransactionStatus_TooManyTransactions MempoolAddTransactionStatus = 4
	// Invalid update. Only gas price increase is allowed
	MempoolAddTransactionStatus_InvalidUpdate MempoolAddTransactionStatus = 5
)

var MempoolAddTransactionStatus_name = map[int32]string{
	0: "Valid",
	1: "InsufficientBalance",
	2: "InvalidSeqNumber",
	3: "MempoolIsFull",
	4: "TooManyTransactions",
	5: "InvalidUpdate",
}

var MempoolAddTransactionStatus_value = map[string]int32{
	"Valid":               0,
	"InsufficientBalance": 1,
	"InvalidSeqNumber":    2,
	"MempoolIsFull":       3,
	"TooManyTransactions": 4,
	"InvalidUpdate":       5,
}

func (x MempoolAddTransactionStatus) String() string {
	return proto.EnumName(MempoolAddTransactionStatus_name, int32(x))
}

func (MempoolAddTransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cad4a86f8a5465be, []int{0}
}

func init() {
	proto.RegisterEnum("mempool.MempoolAddTransactionStatus", MempoolAddTransactionStatus_name, MempoolAddTransactionStatus_value)
}

func init() { proto.RegisterFile("mempool_status.proto", fileDescriptor_cad4a86f8a5465be) }

var fileDescriptor_cad4a86f8a5465be = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcd, 0x2d,
	0xc8, 0xcf, 0xcf, 0x89, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x8a, 0x6a, 0x4d, 0x63, 0xe4, 0x92, 0xf6, 0x85, 0xb0, 0x1d, 0x53, 0x52, 0x42,
	0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0xc1, 0xca, 0x85, 0x38, 0xb9,
	0x58, 0xc3, 0x12, 0x73, 0x32, 0x53, 0x04, 0x18, 0x84, 0xc4, 0xb9, 0x84, 0x3d, 0xf3, 0x8a, 0x4b,
	0xd3, 0xd2, 0x32, 0x93, 0x33, 0x53, 0xf3, 0x4a, 0x9c, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0x05,
	0x18, 0x85, 0x44, 0xb8, 0x04, 0x3c, 0xf3, 0xca, 0x40, 0xaa, 0x82, 0x53, 0x0b, 0xfd, 0x4a, 0x73,
	0x93, 0x52, 0x8b, 0x04, 0x98, 0x84, 0x04, 0xb9, 0x78, 0xa1, 0x06, 0x7b, 0x16, 0xbb, 0x95, 0xe6,
	0xe4, 0x08, 0x30, 0x83, 0x4c, 0x08, 0xc9, 0xcf, 0xf7, 0x4d, 0xcc, 0xab, 0x44, 0xb2, 0xa8, 0x58,
	0x80, 0x05, 0xa4, 0x16, 0x6a, 0x42, 0x68, 0x41, 0x4a, 0x62, 0x49, 0xaa, 0x00, 0x6b, 0x12, 0x1b,
	0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x40, 0xfc, 0xb6, 0xc0, 0x00, 0x00,
	0x00,
}
