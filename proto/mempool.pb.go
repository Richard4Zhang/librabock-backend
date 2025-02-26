// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mempool.proto

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

type AddTransactionWithValidationRequest struct {
	// Transaction from a wallet
	SignedTxn *SignedTransaction `protobuf:"bytes,1,opt,name=signed_txn,json=signedTxn,proto3" json:"signed_txn,omitempty"`
	// Max amount of gas required to execute the transaction
	// Without running the program, it is very difficult to determine this number,
	// so we use the max gas specified by the signed transaction.
	// This field is still included separately from the signed transaction so that
	// if we have a better methodology in the future, we can more accurately
	// specify the max gas.
	MaxGasCost uint64 `protobuf:"varint,2,opt,name=max_gas_cost,json=maxGasCost,proto3" json:"max_gas_cost,omitempty"`
	// Latest sequence number of the involved account from state db.
	LatestSequenceNumber uint64 `protobuf:"varint,3,opt,name=latest_sequence_number,json=latestSequenceNumber,proto3" json:"latest_sequence_number,omitempty"`
	// Latest account balance of the involved account from state db.
	AccountBalance       uint64   `protobuf:"varint,4,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTransactionWithValidationRequest) Reset()         { *m = AddTransactionWithValidationRequest{} }
func (m *AddTransactionWithValidationRequest) String() string { return proto.CompactTextString(m) }
func (*AddTransactionWithValidationRequest) ProtoMessage()    {}
func (*AddTransactionWithValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{0}
}

func (m *AddTransactionWithValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTransactionWithValidationRequest.Unmarshal(m, b)
}
func (m *AddTransactionWithValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTransactionWithValidationRequest.Marshal(b, m, deterministic)
}
func (m *AddTransactionWithValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTransactionWithValidationRequest.Merge(m, src)
}
func (m *AddTransactionWithValidationRequest) XXX_Size() int {
	return xxx_messageInfo_AddTransactionWithValidationRequest.Size(m)
}
func (m *AddTransactionWithValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTransactionWithValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTransactionWithValidationRequest proto.InternalMessageInfo

func (m *AddTransactionWithValidationRequest) GetSignedTxn() *SignedTransaction {
	if m != nil {
		return m.SignedTxn
	}
	return nil
}

func (m *AddTransactionWithValidationRequest) GetMaxGasCost() uint64 {
	if m != nil {
		return m.MaxGasCost
	}
	return 0
}

func (m *AddTransactionWithValidationRequest) GetLatestSequenceNumber() uint64 {
	if m != nil {
		return m.LatestSequenceNumber
	}
	return 0
}

func (m *AddTransactionWithValidationRequest) GetAccountBalance() uint64 {
	if m != nil {
		return m.AccountBalance
	}
	return 0
}

type AddTransactionWithValidationResponse struct {
	// The ledger version at the time of the transaction submitted. The submitted
	// transaction will have version bigger than this 'current_version'
	CurrentVersion uint64 `protobuf:"varint,1,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	// The result of the transaction submission
	Status               MempoolAddTransactionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=mempool.MempoolAddTransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AddTransactionWithValidationResponse) Reset()         { *m = AddTransactionWithValidationResponse{} }
func (m *AddTransactionWithValidationResponse) String() string { return proto.CompactTextString(m) }
func (*AddTransactionWithValidationResponse) ProtoMessage()    {}
func (*AddTransactionWithValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{1}
}

func (m *AddTransactionWithValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTransactionWithValidationResponse.Unmarshal(m, b)
}
func (m *AddTransactionWithValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTransactionWithValidationResponse.Marshal(b, m, deterministic)
}
func (m *AddTransactionWithValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTransactionWithValidationResponse.Merge(m, src)
}
func (m *AddTransactionWithValidationResponse) XXX_Size() int {
	return xxx_messageInfo_AddTransactionWithValidationResponse.Size(m)
}
func (m *AddTransactionWithValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTransactionWithValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddTransactionWithValidationResponse proto.InternalMessageInfo

func (m *AddTransactionWithValidationResponse) GetCurrentVersion() uint64 {
	if m != nil {
		return m.CurrentVersion
	}
	return 0
}

func (m *AddTransactionWithValidationResponse) GetStatus() MempoolAddTransactionStatus {
	if m != nil {
		return m.Status
	}
	return MempoolAddTransactionStatus_Valid
}

// -----------------------------------------------------------------------------
// ---------------- GetBlock
// -----------------------------------------------------------------------------
type GetBlockRequest struct {
	MaxBlockSize         uint64                  `protobuf:"varint,1,opt,name=max_block_size,json=maxBlockSize,proto3" json:"max_block_size,omitempty"`
	Transactions         []*TransactionExclusion `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{2}
}

func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockRequest.Unmarshal(m, b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(m, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockRequest.Size(m)
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetMaxBlockSize() uint64 {
	if m != nil {
		return m.MaxBlockSize
	}
	return 0
}

func (m *GetBlockRequest) GetTransactions() []*TransactionExclusion {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type GetBlockResponse struct {
	Block                *SignedTransactionsBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetBlockResponse) Reset()         { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()    {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{3}
}

func (m *GetBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockResponse.Unmarshal(m, b)
}
func (m *GetBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponse.Merge(m, src)
}
func (m *GetBlockResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockResponse.Size(m)
}
func (m *GetBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponse proto.InternalMessageInfo

func (m *GetBlockResponse) GetBlock() *SignedTransactionsBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

type TransactionExclusion struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionExclusion) Reset()         { *m = TransactionExclusion{} }
func (m *TransactionExclusion) String() string { return proto.CompactTextString(m) }
func (*TransactionExclusion) ProtoMessage()    {}
func (*TransactionExclusion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{4}
}

func (m *TransactionExclusion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionExclusion.Unmarshal(m, b)
}
func (m *TransactionExclusion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionExclusion.Marshal(b, m, deterministic)
}
func (m *TransactionExclusion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionExclusion.Merge(m, src)
}
func (m *TransactionExclusion) XXX_Size() int {
	return xxx_messageInfo_TransactionExclusion.Size(m)
}
func (m *TransactionExclusion) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionExclusion.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionExclusion proto.InternalMessageInfo

func (m *TransactionExclusion) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *TransactionExclusion) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

// -----------------------------------------------------------------------------
// ---------------- CommitTransactions
// -----------------------------------------------------------------------------
type CommitTransactionsRequest struct {
	Transactions []*CommittedTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// agreed monotonic timestamp microseconds since the epoch for a committed block
	// used by Mempool to GC expired transactions
	BlockTimestampUsecs  uint64   `protobuf:"varint,2,opt,name=block_timestamp_usecs,json=blockTimestampUsecs,proto3" json:"block_timestamp_usecs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitTransactionsRequest) Reset()         { *m = CommitTransactionsRequest{} }
func (m *CommitTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*CommitTransactionsRequest) ProtoMessage()    {}
func (*CommitTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{5}
}

func (m *CommitTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitTransactionsRequest.Unmarshal(m, b)
}
func (m *CommitTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *CommitTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitTransactionsRequest.Merge(m, src)
}
func (m *CommitTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_CommitTransactionsRequest.Size(m)
}
func (m *CommitTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitTransactionsRequest proto.InternalMessageInfo

func (m *CommitTransactionsRequest) GetTransactions() []*CommittedTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *CommitTransactionsRequest) GetBlockTimestampUsecs() uint64 {
	if m != nil {
		return m.BlockTimestampUsecs
	}
	return 0
}

type CommitTransactionsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitTransactionsResponse) Reset()         { *m = CommitTransactionsResponse{} }
func (m *CommitTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*CommitTransactionsResponse) ProtoMessage()    {}
func (*CommitTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{6}
}

func (m *CommitTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitTransactionsResponse.Unmarshal(m, b)
}
func (m *CommitTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *CommitTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitTransactionsResponse.Merge(m, src)
}
func (m *CommitTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_CommitTransactionsResponse.Size(m)
}
func (m *CommitTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitTransactionsResponse proto.InternalMessageInfo

type CommittedTransaction struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	IsRejected           bool     `protobuf:"varint,3,opt,name=is_rejected,json=isRejected,proto3" json:"is_rejected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommittedTransaction) Reset()         { *m = CommittedTransaction{} }
func (m *CommittedTransaction) String() string { return proto.CompactTextString(m) }
func (*CommittedTransaction) ProtoMessage()    {}
func (*CommittedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{7}
}

func (m *CommittedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommittedTransaction.Unmarshal(m, b)
}
func (m *CommittedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommittedTransaction.Marshal(b, m, deterministic)
}
func (m *CommittedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommittedTransaction.Merge(m, src)
}
func (m *CommittedTransaction) XXX_Size() int {
	return xxx_messageInfo_CommittedTransaction.Size(m)
}
func (m *CommittedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_CommittedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_CommittedTransaction proto.InternalMessageInfo

func (m *CommittedTransaction) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *CommittedTransaction) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *CommittedTransaction) GetIsRejected() bool {
	if m != nil {
		return m.IsRejected
	}
	return false
}

// -----------------------------------------------------------------------------
// ---------------- HealthCheck
// -----------------------------------------------------------------------------
type HealthCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{8}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

type HealthCheckResponse struct {
	// Indicate whether Mempool is in healthy condition.
	IsHealthy            bool     `protobuf:"varint,1,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a84c3667d8c2093a, []int{9}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetIsHealthy() bool {
	if m != nil {
		return m.IsHealthy
	}
	return false
}

func init() {
	proto.RegisterType((*AddTransactionWithValidationRequest)(nil), "mempool.AddTransactionWithValidationRequest")
	proto.RegisterType((*AddTransactionWithValidationResponse)(nil), "mempool.AddTransactionWithValidationResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "mempool.GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "mempool.GetBlockResponse")
	proto.RegisterType((*TransactionExclusion)(nil), "mempool.TransactionExclusion")
	proto.RegisterType((*CommitTransactionsRequest)(nil), "mempool.CommitTransactionsRequest")
	proto.RegisterType((*CommitTransactionsResponse)(nil), "mempool.CommitTransactionsResponse")
	proto.RegisterType((*CommittedTransaction)(nil), "mempool.CommittedTransaction")
	proto.RegisterType((*HealthCheckRequest)(nil), "mempool.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "mempool.HealthCheckResponse")
}

func init() { proto.RegisterFile("mempool.proto", fileDescriptor_a84c3667d8c2093a) }

var fileDescriptor_a84c3667d8c2093a = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0x76, 0x01, 0xf9, 0xf1, 0x8a, 0x45, 0x87, 0x4a, 0xca, 0x0a, 0x4a, 0x16, 0x12, 0x39, 0x68,
	0x0f, 0x95, 0xc4, 0x8b, 0x17, 0x20, 0x06, 0x62, 0xa2, 0x87, 0x2d, 0xc2, 0x71, 0x32, 0xdd, 0xbe,
	0xd0, 0xd1, 0xdd, 0x99, 0xba, 0x6f, 0xd6, 0x14, 0x12, 0xff, 0x04, 0x2f, 0xfe, 0x83, 0xfe, 0x1b,
	0x1e, 0xcd, 0xce, 0x4c, 0x97, 0xb6, 0x54, 0x42, 0xe2, 0x69, 0x33, 0xef, 0xfb, 0xde, 0x7c, 0x33,
	0xdf, 0xf7, 0x66, 0xe1, 0x51, 0x86, 0xd9, 0x40, 0xeb, 0xb4, 0x35, 0xc8, 0xb5, 0xd1, 0x6c, 0xc9,
	0x2f, 0xc3, 0x27, 0x26, 0x17, 0x8a, 0x44, 0x62, 0xa4, 0x56, 0x0e, 0x0b, 0x1b, 0x1e, 0xe3, 0x64,
	0x84, 0x29, 0xc8, 0x55, 0xa3, 0xdf, 0x01, 0xec, 0x1e, 0xf6, 0x7a, 0x67, 0x37, 0xf4, 0x0b, 0x69,
	0xfa, 0xe7, 0x22, 0x95, 0x3d, 0x51, 0xae, 0x62, 0xfc, 0x56, 0x20, 0x19, 0xf6, 0x16, 0x80, 0xe4,
	0xa5, 0xc2, 0x1e, 0x37, 0x43, 0xd5, 0x0c, 0x76, 0x82, 0xfd, 0x5a, 0xbb, 0xd9, 0x32, 0x57, 0x03,
	0xa4, 0x56, 0xc7, 0x02, 0x63, 0x5b, 0xc4, 0x2b, 0x8e, 0x7b, 0x36, 0x54, 0x6c, 0x07, 0x56, 0x33,
	0x31, 0xe4, 0x97, 0x82, 0x78, 0xa2, 0xc9, 0x34, 0xe7, 0x76, 0x82, 0xfd, 0x85, 0x18, 0x32, 0x31,
	0x3c, 0x11, 0x74, 0xac, 0xc9, 0xb0, 0x03, 0xd8, 0x48, 0x85, 0x41, 0x32, 0x9c, 0x4a, 0x31, 0x95,
	0x20, 0x57, 0x45, 0xd6, 0xc5, 0xbc, 0x39, 0x6f, 0xb9, 0x0d, 0x87, 0x76, 0x3c, 0xf8, 0xc9, 0x62,
	0xec, 0x25, 0xac, 0x89, 0x24, 0xd1, 0x85, 0x32, 0xbc, 0x2b, 0x52, 0xa1, 0x12, 0x6c, 0x2e, 0x58,
	0x7a, 0xdd, 0x97, 0x8f, 0x5c, 0x35, 0xfa, 0x19, 0xc0, 0xde, 0xdd, 0x37, 0xa4, 0x81, 0x56, 0x84,
	0xe5, 0x8e, 0x49, 0x91, 0xe7, 0xa8, 0x0c, 0xff, 0x8e, 0x39, 0x49, 0xed, 0xee, 0xb9, 0x10, 0xd7,
	0x7d, 0xf9, 0xdc, 0x55, 0xd9, 0x3b, 0x58, 0x74, 0x1e, 0xda, 0xcb, 0xd4, 0xdb, 0x7b, 0xad, 0x51,
	0x0a, 0x1f, 0xdd, 0x77, 0x52, 0xae, 0x63, 0xb9, 0xb1, 0xef, 0x89, 0xae, 0x61, 0xed, 0x04, 0xcd,
	0x51, 0xaa, 0x93, 0xaf, 0x23, 0x73, 0xf7, 0xa0, 0x5e, 0x7a, 0xd4, 0x2d, 0x6b, 0x9c, 0xe4, 0x35,
	0x7a, 0xe1, 0xd2, 0x39, 0x4b, 0xec, 0xc8, 0x6b, 0x64, 0x87, 0xb0, 0x3a, 0x96, 0x6a, 0x29, 0x3e,
	0xbf, 0x5f, 0x6b, 0x6f, 0x57, 0xe2, 0x63, 0x92, 0xef, 0x87, 0x49, 0x5a, 0x94, 0x67, 0x8d, 0x27,
	0x5a, 0xa2, 0x53, 0x78, 0x7c, 0xa3, 0xed, 0xaf, 0x7d, 0x00, 0x0f, 0xad, 0xb0, 0x0f, 0xf5, 0xf9,
	0xbf, 0x42, 0x25, 0xd7, 0xe6, 0xc8, 0xd1, 0x05, 0x34, 0x66, 0xe9, 0xb1, 0x0d, 0x58, 0x24, 0x54,
	0x3d, 0xcc, 0xed, 0x76, 0xab, 0xb1, 0x5f, 0x95, 0xe6, 0x4e, 0xa7, 0xeb, 0x26, 0xa1, 0x4e, 0x13,
	0xb9, 0x46, 0xbf, 0x02, 0xd8, 0x3c, 0xd6, 0x59, 0x26, 0xcd, 0xb8, 0xf6, 0xc8, 0xa9, 0x69, 0x0f,
	0x82, 0x29, 0x0f, 0x5c, 0xa7, 0x99, 0x9c, 0xc6, 0x89, 0x16, 0xd6, 0x86, 0xa7, 0xce, 0x68, 0x23,
	0x33, 0x24, 0x23, 0xb2, 0x01, 0x2f, 0x08, 0x13, 0xf2, 0xe7, 0x59, 0xb7, 0xe0, 0xd9, 0x08, 0xfb,
	0x5c, 0x42, 0xd1, 0x16, 0x84, 0xb3, 0xce, 0xe4, 0x1c, 0x8c, 0x86, 0xd0, 0x98, 0xa5, 0xfb, 0xdf,
	0x5e, 0xb0, 0x17, 0x50, 0x93, 0xc4, 0x73, 0xfc, 0x82, 0x89, 0xc1, 0x9e, 0x7d, 0x0e, 0xcb, 0x31,
	0x48, 0x8a, 0x7d, 0x25, 0x6a, 0x00, 0x3b, 0x45, 0x91, 0x9a, 0xfe, 0x71, 0x1f, 0xab, 0x71, 0x8a,
	0x0e, 0x60, 0x7d, 0xa2, 0xea, 0x83, 0xde, 0x06, 0x90, 0xc4, 0xfb, 0x16, 0xb9, 0xb2, 0x47, 0x5a,
	0x8e, 0x57, 0x24, 0x39, 0xea, 0x55, 0xfb, 0xcf, 0x1c, 0x2c, 0xf9, 0xf9, 0x65, 0x3f, 0x60, 0xeb,
	0xae, 0x27, 0xc3, 0x5e, 0x55, 0x86, 0xdf, 0xe3, 0xdf, 0x11, 0xbe, 0xbe, 0x27, 0xdb, 0xdb, 0xf9,
	0x80, 0x1d, 0xc2, 0xf2, 0x68, 0x4c, 0x59, 0xb3, 0x6a, 0x9e, 0x7a, 0x35, 0xe1, 0xe6, 0x0c, 0xa4,
	0xda, 0x82, 0x03, 0xbb, 0x9d, 0x18, 0x8b, 0xa6, 0x06, 0x65, 0xc6, 0x88, 0x85, 0xbb, 0x77, 0x72,
	0x2a, 0x81, 0x0f, 0x50, 0x1b, 0x33, 0x99, 0x3d, 0xab, 0xba, 0x6e, 0x07, 0x12, 0x6e, 0xcd, 0x06,
	0x47, 0x7b, 0x75, 0x17, 0xed, 0xbf, 0xf8, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xd3,
	0x81, 0x61, 0xce, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MempoolClient is the client API for Mempool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MempoolClient interface {
	// Adds a new transaction to the mempool with validation against existing
	// transactions in the mempool.  Note that this function performs additional
	// validation that AC is unable to perform. (because AC knows only about a
	// single transaction, but mempool potentially knows about multiple pending
	// transactions)
	AddTransactionWithValidation(ctx context.Context, in *AddTransactionWithValidationRequest, opts ...grpc.CallOption) (*AddTransactionWithValidationResponse, error)
	// Fetch ordered block of transactions
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	// Remove committed transactions from Mempool
	CommitTransactions(ctx context.Context, in *CommitTransactionsRequest, opts ...grpc.CallOption) (*CommitTransactionsResponse, error)
	// Check the health of mempool
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type mempoolClient struct {
	cc *grpc.ClientConn
}

func NewMempoolClient(cc *grpc.ClientConn) MempoolClient {
	return &mempoolClient{cc}
}

func (c *mempoolClient) AddTransactionWithValidation(ctx context.Context, in *AddTransactionWithValidationRequest, opts ...grpc.CallOption) (*AddTransactionWithValidationResponse, error) {
	out := new(AddTransactionWithValidationResponse)
	err := c.cc.Invoke(ctx, "/mempool.Mempool/AddTransactionWithValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mempoolClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/mempool.Mempool/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mempoolClient) CommitTransactions(ctx context.Context, in *CommitTransactionsRequest, opts ...grpc.CallOption) (*CommitTransactionsResponse, error) {
	out := new(CommitTransactionsResponse)
	err := c.cc.Invoke(ctx, "/mempool.Mempool/CommitTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mempoolClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/mempool.Mempool/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MempoolServer is the server API for Mempool service.
type MempoolServer interface {
	// Adds a new transaction to the mempool with validation against existing
	// transactions in the mempool.  Note that this function performs additional
	// validation that AC is unable to perform. (because AC knows only about a
	// single transaction, but mempool potentially knows about multiple pending
	// transactions)
	AddTransactionWithValidation(context.Context, *AddTransactionWithValidationRequest) (*AddTransactionWithValidationResponse, error)
	// Fetch ordered block of transactions
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	// Remove committed transactions from Mempool
	CommitTransactions(context.Context, *CommitTransactionsRequest) (*CommitTransactionsResponse, error)
	// Check the health of mempool
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedMempoolServer can be embedded to have forward compatible implementations.
type UnimplementedMempoolServer struct {
}

func (*UnimplementedMempoolServer) AddTransactionWithValidation(ctx context.Context, req *AddTransactionWithValidationRequest) (*AddTransactionWithValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransactionWithValidation not implemented")
}
func (*UnimplementedMempoolServer) GetBlock(ctx context.Context, req *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedMempoolServer) CommitTransactions(ctx context.Context, req *CommitTransactionsRequest) (*CommitTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitTransactions not implemented")
}
func (*UnimplementedMempoolServer) HealthCheck(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}

func RegisterMempoolServer(s *grpc.Server, srv MempoolServer) {
	s.RegisterService(&_Mempool_serviceDesc, srv)
}

func _Mempool_AddTransactionWithValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransactionWithValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolServer).AddTransactionWithValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.Mempool/AddTransactionWithValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolServer).AddTransactionWithValidation(ctx, req.(*AddTransactionWithValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mempool_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.Mempool/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mempool_CommitTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolServer).CommitTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.Mempool/CommitTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolServer).CommitTransactions(ctx, req.(*CommitTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mempool_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.Mempool/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mempool_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mempool.Mempool",
	HandlerType: (*MempoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransactionWithValidation",
			Handler:    _Mempool_AddTransactionWithValidation_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Mempool_GetBlock_Handler,
		},
		{
			MethodName: "CommitTransactions",
			Handler:    _Mempool_CommitTransactions_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Mempool_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mempool.proto",
}
