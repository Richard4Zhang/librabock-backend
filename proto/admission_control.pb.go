// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admission_control.proto

package libra

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Additional statuses that are possible from admission control in addition
// to VM statuses.
type AdmissionControlStatus int32

const (
	AdmissionControlStatus_Accepted    AdmissionControlStatus = 0
	AdmissionControlStatus_Blacklisted AdmissionControlStatus = 1
	AdmissionControlStatus_Rejected    AdmissionControlStatus = 2
)

var AdmissionControlStatus_name = map[int32]string{
	0: "Accepted",
	1: "Blacklisted",
	2: "Rejected",
}

var AdmissionControlStatus_value = map[string]int32{
	"Accepted":    0,
	"Blacklisted": 1,
	"Rejected":    2,
}

func (x AdmissionControlStatus) String() string {
	return proto.EnumName(AdmissionControlStatus_name, int32(x))
}

func (AdmissionControlStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

// -----------------------------------------------------------------------------
// ---------------- Submit transaction
// -----------------------------------------------------------------------------
// The request for transaction submission.
type SubmitTransactionRequest struct {
	// Transaction signed by wallet.
	SignedTxn            *SignedTransaction `protobuf:"bytes,1,opt,name=signed_txn,json=signedTxn,proto3" json:"signed_txn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetSignedTxn() *SignedTransaction {
	if m != nil {
		return m.SignedTxn
	}
	return nil
}

// The response for transaction submission.
//
// How does a client know if their transaction was included?
// A response from the transaction submission only means that the transaction
// was successfully added to mempool, but not that it is guaranteed to be
// included in the chain.  Each transaction should include an expiration time in
// the signed transaction.  Let's call this T0.  As a client, I submit my
// transaction to a validator. I now need to poll for the transaction I
// submitted.  I can use the query that takes my account and sequence number. If
// I receive back that the transaction is completed, I will verify the proofs to
// ensure that this is the transaction I expected.  If I receive a response that
// my transaction is not yet completed, I must check the latest timestamp in the
// ledgerInfo that I receive back from the query.  If this time is greater than
// T0, I can be certain that my transaction will never be included.  If this
// time is less than T0, I need to continue polling.
type SubmitTransactionResponse struct {
	// The status of a transaction submission can either be a VM status, or
	// some other admission control/mempool specific status e.g. Blacklisted.
	//
	// Types that are valid to be assigned to Status:
	//	*SubmitTransactionResponse_VmStatus
	//	*SubmitTransactionResponse_AcStatus
	//	*SubmitTransactionResponse_MempoolStatus
	Status isSubmitTransactionResponse_Status `protobuf_oneof:"status"`
	// Public key(id) of the validator that processed this transaction
	ValidatorId          []byte   `protobuf:"bytes,4,opt,name=validator_id,json=validatorId,proto3" json:"validator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionResponse) Reset()         { *m = SubmitTransactionResponse{} }
func (m *SubmitTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionResponse) ProtoMessage()    {}
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{1}
}

func (m *SubmitTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionResponse.Unmarshal(m, b)
}
func (m *SubmitTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionResponse.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionResponse.Merge(m, src)
}
func (m *SubmitTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionResponse.Size(m)
}
func (m *SubmitTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionResponse proto.InternalMessageInfo

type isSubmitTransactionResponse_Status interface {
	isSubmitTransactionResponse_Status()
}

type SubmitTransactionResponse_VmStatus struct {
	VmStatus *VMStatus `protobuf:"bytes,1,opt,name=vm_status,json=vmStatus,proto3,oneof"`
}

type SubmitTransactionResponse_AcStatus struct {
	AcStatus AdmissionControlStatus `protobuf:"varint,2,opt,name=ac_status,json=acStatus,proto3,enum=admission_control.AdmissionControlStatus,oneof"`
}

type SubmitTransactionResponse_MempoolStatus struct {
	MempoolStatus MempoolAddTransactionStatus `protobuf:"varint,3,opt,name=mempool_status,json=mempoolStatus,proto3,enum=mempool.MempoolAddTransactionStatus,oneof"`
}

func (*SubmitTransactionResponse_VmStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_AcStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_MempoolStatus) isSubmitTransactionResponse_Status() {}

func (m *SubmitTransactionResponse) GetStatus() isSubmitTransactionResponse_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SubmitTransactionResponse) GetVmStatus() *VMStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_VmStatus); ok {
		return x.VmStatus
	}
	return nil
}

func (m *SubmitTransactionResponse) GetAcStatus() AdmissionControlStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_AcStatus); ok {
		return x.AcStatus
	}
	return AdmissionControlStatus_Accepted
}

func (m *SubmitTransactionResponse) GetMempoolStatus() MempoolAddTransactionStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_MempoolStatus); ok {
		return x.MempoolStatus
	}
	return MempoolAddTransactionStatus_Valid
}

func (m *SubmitTransactionResponse) GetValidatorId() []byte {
	if m != nil {
		return m.ValidatorId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitTransactionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitTransactionResponse_VmStatus)(nil),
		(*SubmitTransactionResponse_AcStatus)(nil),
		(*SubmitTransactionResponse_MempoolStatus)(nil),
	}
}

func init() {
	proto.RegisterEnum("admission_control.AdmissionControlStatus", AdmissionControlStatus_name, AdmissionControlStatus_value)
	proto.RegisterType((*SubmitTransactionRequest)(nil), "admission_control.SubmitTransactionRequest")
	proto.RegisterType((*SubmitTransactionResponse)(nil), "admission_control.SubmitTransactionResponse")
}

func init() { proto.RegisterFile("admission_control.proto", fileDescriptor_d13d6ff9aac892b4) }

var fileDescriptor_d13d6ff9aac892b4 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x9b, 0x82, 0x56, 0xed, 0xb4, 0x6c, 0xb7, 0x56, 0x05, 0xa1, 0xa7, 0x52, 0x38, 0x94,
	0x3f, 0xca, 0xa1, 0x1c, 0x38, 0x77, 0x11, 0xd2, 0x22, 0x6d, 0x2f, 0x6d, 0xe1, 0x1a, 0x79, 0xed,
	0xa1, 0x18, 0x12, 0xdb, 0xd8, 0xd3, 0xb0, 0x3c, 0x06, 0x4f, 0xc9, 0x6b, 0xa0, 0x26, 0xce, 0x6e,
	0xa1, 0x01, 0xed, 0x29, 0xc9, 0x6f, 0xbe, 0x7c, 0xe3, 0xcc, 0x04, 0x1e, 0x71, 0x99, 0x2b, 0xef,
	0x95, 0xd1, 0xa9, 0x30, 0x9a, 0x9c, 0xc9, 0x12, 0xeb, 0x0c, 0x19, 0x36, 0x3c, 0x2a, 0x8c, 0x47,
	0x5b, 0xa4, 0xf4, 0xbb, 0xa2, 0xcf, 0xa9, 0x75, 0xc6, 0x7c, 0xaa, 0xc0, 0xf1, 0x28, 0xc7, 0xdc,
	0x1a, 0x93, 0xa5, 0x9e, 0x38, 0xed, 0x7c, 0x48, 0x87, 0xe4, 0xb8, 0xf6, 0x5c, 0x90, 0x32, 0x3a,
	0x44, 0x83, 0x22, 0x4f, 0xd1, 0x39, 0xe3, 0x02, 0x33, 0x5d, 0x43, 0xbc, 0xde, 0x5d, 0xe5, 0x8a,
	0x36, 0xb7, 0xec, 0x0a, 0xbf, 0xed, 0xd0, 0x13, 0x7b, 0x03, 0xe0, 0xd5, 0x56, 0xa3, 0x4c, 0xe9,
	0x5a, 0xc7, 0xd1, 0x24, 0x9a, 0xf5, 0xe6, 0x71, 0x42, 0x3f, 0x2c, 0xfa, 0x64, 0x5d, 0x16, 0x0e,
	0x5f, 0xea, 0x56, 0xec, 0xe6, 0x5a, 0x4f, 0x7f, 0xb6, 0xe1, 0x71, 0x83, 0xd5, 0x5b, 0xa3, 0x3d,
	0xb2, 0x04, 0xba, 0x45, 0x1e, 0x4e, 0x1a, 0xac, 0x83, 0x60, 0xfd, 0xb8, 0x5c, 0x97, 0xf1, 0x45,
	0x6b, 0xd5, 0x29, 0xf2, 0xea, 0x9e, 0x5d, 0x40, 0x97, 0x8b, 0x9a, 0x6f, 0x4f, 0xa2, 0xd9, 0xe9,
	0xfc, 0x79, 0x72, 0x3c, 0xb2, 0x45, 0x9d, 0xbc, 0xad, 0x82, 0x5b, 0x13, 0x17, 0xc1, 0xb4, 0x84,
	0xd3, 0x3f, 0x07, 0x15, 0xdf, 0x2b, 0x75, 0xcf, 0x92, 0x10, 0x27, 0xcb, 0xea, 0xba, 0x90, 0x87,
	0x9f, 0x76, 0x63, 0x7a, 0x10, 0xb0, 0xa0, 0x7b, 0x02, 0xfd, 0x82, 0x67, 0x4a, 0x72, 0x32, 0x2e,
	0x55, 0x32, 0xbe, 0x3f, 0x89, 0x66, 0xfd, 0x55, 0xef, 0x26, 0x7b, 0x2f, 0xcf, 0x3b, 0x70, 0x52,
	0x75, 0x7a, 0xf1, 0x0e, 0x1e, 0x36, 0x9f, 0x90, 0xf5, 0xa1, 0xb3, 0x10, 0x02, 0x2d, 0xa1, 0x3c,
	0x6b, 0xb1, 0x01, 0xf4, 0xce, 0x33, 0x2e, 0xbe, 0x66, 0xca, 0xef, 0x83, 0x68, 0x5f, 0x5e, 0xe1,
	0x17, 0x14, 0xfb, 0xa7, 0xf6, 0xfc, 0x57, 0x04, 0x67, 0x7f, 0x7b, 0x98, 0x85, 0xe1, 0xd1, 0xb8,
	0xd9, 0xcb, 0x86, 0x19, 0xfd, 0x6b, 0xd5, 0xe3, 0x57, 0x77, 0x83, 0xab, 0x0d, 0x4e, 0x5b, 0x8c,
	0xc3, 0xe8, 0x83, 0x95, 0x9c, 0x70, 0x63, 0x2e, 0x39, 0xa1, 0xa7, 0x4b, 0x94, 0x5b, 0x74, 0x6c,
	0x1a, 0x16, 0xd9, 0x54, 0xac, 0x7b, 0x3d, 0xfd, 0x2f, 0x53, 0xb7, 0xb8, 0x3a, 0x29, 0x7f, 0xd0,
	0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x22, 0x50, 0x5a, 0xd7, 0x1e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdmissionControlClient is the client API for AdmissionControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdmissionControlClient interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error)
}

type admissionControlClient struct {
	cc *grpc.ClientConn
}

func NewAdmissionControlClient(cc *grpc.ClientConn) AdmissionControlClient {
	return &admissionControlClient{cc}
}

func (c *admissionControlClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *admissionControlClient) UpdateToLatestLedger(ctx context.Context, in *UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*UpdateToLatestLedgerResponse, error) {
	out := new(UpdateToLatestLedgerResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/UpdateToLatestLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdmissionControlServer is the server API for AdmissionControl service.
type AdmissionControlServer interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(context.Context, *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error)
}

// UnimplementedAdmissionControlServer can be embedded to have forward compatible implementations.
type UnimplementedAdmissionControlServer struct {
}

func (*UnimplementedAdmissionControlServer) SubmitTransaction(ctx context.Context, req *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (*UnimplementedAdmissionControlServer) UpdateToLatestLedger(ctx context.Context, req *UpdateToLatestLedgerRequest) (*UpdateToLatestLedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToLatestLedger not implemented")
}

func RegisterAdmissionControlServer(s *grpc.Server, srv AdmissionControlServer) {
	s.RegisterService(&_AdmissionControl_serviceDesc, srv)
}

func _AdmissionControl_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdmissionControl_UpdateToLatestLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateToLatestLedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/UpdateToLatestLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, req.(*UpdateToLatestLedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdmissionControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admission_control.AdmissionControl",
	HandlerType: (*AdmissionControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _AdmissionControl_SubmitTransaction_Handler,
		},
		{
			MethodName: "UpdateToLatestLedger",
			Handler:    _AdmissionControl_UpdateToLatestLedger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admission_control.proto",
}
