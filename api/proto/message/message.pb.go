// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// ReceiverType indicates who is the receiver of this message.
type ReceiverType int32

const (
	ReceiverType_NEWNODE   ReceiverType = 0
	ReceiverType_LEADER    ReceiverType = 1
	ReceiverType_VALIDATOR ReceiverType = 2
	ReceiverType_CLIENT    ReceiverType = 3
)

var ReceiverType_name = map[int32]string{
	0: "NEWNODE",
	1: "LEADER",
	2: "VALIDATOR",
	3: "CLIENT",
}

var ReceiverType_value = map[string]int32{
	"NEWNODE":   0,
	"LEADER":    1,
	"VALIDATOR": 2,
	"CLIENT":    3,
}

func (x ReceiverType) String() string {
	return proto.EnumName(ReceiverType_name, int32(x))
}

func (ReceiverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

// ServiceType indicates which service used to generate this message.
type ServiceType int32

const (
	ServiceType_CONSENSUS ServiceType = 0
	ServiceType_STAKING   ServiceType = 1
	ServiceType_DRAND     ServiceType = 2
)

var ServiceType_name = map[int32]string{
	0: "CONSENSUS",
	1: "STAKING",
	2: "DRAND",
}

var ServiceType_value = map[string]int32{
	"CONSENSUS": 0,
	"STAKING":   1,
	"DRAND":     2,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

// MessageType indicates what is the type of this message.
type MessageType int32

const (
	MessageType_NEWNODE_BEACON_STAKING MessageType = 0
	MessageType_ANNOUNCE               MessageType = 1
	MessageType_PREPARE                MessageType = 2
	MessageType_PREPARED               MessageType = 3
	MessageType_COMMIT                 MessageType = 4
	MessageType_COMMITTED              MessageType = 5
	MessageType_DRAND_INIT             MessageType = 6
	MessageType_DRAND_COMMIT           MessageType = 7
)

var MessageType_name = map[int32]string{
	0: "NEWNODE_BEACON_STAKING",
	1: "ANNOUNCE",
	2: "PREPARE",
	3: "PREPARED",
	4: "COMMIT",
	5: "COMMITTED",
	6: "DRAND_INIT",
	7: "DRAND_COMMIT",
}

var MessageType_value = map[string]int32{
	"NEWNODE_BEACON_STAKING": 0,
	"ANNOUNCE":               1,
	"PREPARE":                2,
	"PREPARED":               3,
	"COMMIT":                 4,
	"COMMITTED":              5,
	"DRAND_INIT":             6,
	"DRAND_COMMIT":           7,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

type LotteryRequest_Type int32

const (
	LotteryRequest_ENTER  LotteryRequest_Type = 0
	LotteryRequest_RESULT LotteryRequest_Type = 1
)

var LotteryRequest_Type_name = map[int32]string{
	0: "ENTER",
	1: "RESULT",
}

var LotteryRequest_Type_value = map[string]int32{
	"ENTER":  0,
	"RESULT": 1,
}

func (x LotteryRequest_Type) String() string {
	return proto.EnumName(LotteryRequest_Type_name, int32(x))
}

func (LotteryRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3, 0}
}

// This is universal message for all communication protocols.
// There are different Requests for different message types.
// As we introduce a new type of message just add a new MessageType and new type of request in Message.
//
// The request field will be either one of the structure corresponding to the MessageType type.
type Message struct {
	ReceiverType ReceiverType `protobuf:"varint,1,opt,name=receiver_type,json=receiverType,proto3,enum=message.ReceiverType" json:"receiver_type,omitempty"`
	ServiceType  ServiceType  `protobuf:"varint,2,opt,name=service_type,json=serviceType,proto3,enum=message.ServiceType" json:"service_type,omitempty"`
	Type         MessageType  `protobuf:"varint,3,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	Signature    []byte       `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*Message_Staking
	//	*Message_Consensus
	//	*Message_Drand
	//	*Message_LotteryRequest
	Request              isMessage_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetReceiverType() ReceiverType {
	if m != nil {
		return m.ReceiverType
	}
	return ReceiverType_NEWNODE
}

func (m *Message) GetServiceType() ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return ServiceType_CONSENSUS
}

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_NEWNODE_BEACON_STAKING
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isMessage_Request interface {
	isMessage_Request()
}

type Message_Staking struct {
	Staking *StakingRequest `protobuf:"bytes,5,opt,name=staking,proto3,oneof"`
}

type Message_Consensus struct {
	Consensus *ConsensusRequest `protobuf:"bytes,6,opt,name=consensus,proto3,oneof"`
}

type Message_Drand struct {
	Drand *DrandRequest `protobuf:"bytes,7,opt,name=drand,proto3,oneof"`
}

type Message_LotteryRequest struct {
	LotteryRequest *LotteryRequest `protobuf:"bytes,8,opt,name=lottery_request,json=lotteryRequest,proto3,oneof"`
}

func (*Message_Staking) isMessage_Request() {}

func (*Message_Consensus) isMessage_Request() {}

func (*Message_Drand) isMessage_Request() {}

func (*Message_LotteryRequest) isMessage_Request() {}

func (m *Message) GetRequest() isMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Message) GetStaking() *StakingRequest {
	if x, ok := m.GetRequest().(*Message_Staking); ok {
		return x.Staking
	}
	return nil
}

func (m *Message) GetConsensus() *ConsensusRequest {
	if x, ok := m.GetRequest().(*Message_Consensus); ok {
		return x.Consensus
	}
	return nil
}

func (m *Message) GetDrand() *DrandRequest {
	if x, ok := m.GetRequest().(*Message_Drand); ok {
		return x.Drand
	}
	return nil
}

func (m *Message) GetLotteryRequest() *LotteryRequest {
	if x, ok := m.GetRequest().(*Message_LotteryRequest); ok {
		return x.LotteryRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Staking)(nil),
		(*Message_Consensus)(nil),
		(*Message_Drand)(nil),
		(*Message_LotteryRequest)(nil),
	}
}

type Response struct {
	ReceiverType ReceiverType `protobuf:"varint,1,opt,name=receiver_type,json=receiverType,proto3,enum=message.ReceiverType" json:"receiver_type,omitempty"`
	ServiceType  ServiceType  `protobuf:"varint,2,opt,name=service_type,json=serviceType,proto3,enum=message.ServiceType" json:"service_type,omitempty"`
	Type         MessageType  `protobuf:"varint,3,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*Response_LotteryResponse
	Response             isResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetReceiverType() ReceiverType {
	if m != nil {
		return m.ReceiverType
	}
	return ReceiverType_NEWNODE
}

func (m *Response) GetServiceType() ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return ServiceType_CONSENSUS
}

func (m *Response) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_NEWNODE_BEACON_STAKING
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_LotteryResponse struct {
	LotteryResponse *LotteryResponse `protobuf:"bytes,4,opt,name=lottery_response,json=lotteryResponse,proto3,oneof"`
}

func (*Response_LotteryResponse) isResponse_Response() {}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetLotteryResponse() *LotteryResponse {
	if x, ok := m.GetResponse().(*Response_LotteryResponse); ok {
		return x.LotteryResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_LotteryResponse)(nil),
	}
}

type LotteryResponse struct {
	Players              []string `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Balances             []uint64 `protobuf:"varint,3,rep,packed,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LotteryResponse) Reset()         { *m = LotteryResponse{} }
func (m *LotteryResponse) String() string { return proto.CompactTextString(m) }
func (*LotteryResponse) ProtoMessage()    {}
func (*LotteryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *LotteryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LotteryResponse.Unmarshal(m, b)
}
func (m *LotteryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LotteryResponse.Marshal(b, m, deterministic)
}
func (m *LotteryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotteryResponse.Merge(m, src)
}
func (m *LotteryResponse) XXX_Size() int {
	return xxx_messageInfo_LotteryResponse.Size(m)
}
func (m *LotteryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LotteryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LotteryResponse proto.InternalMessageInfo

func (m *LotteryResponse) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *LotteryResponse) GetBalances() []uint64 {
	if m != nil {
		return m.Balances
	}
	return nil
}

type LotteryRequest struct {
	Type                 LotteryRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=message.LotteryRequest_Type" json:"type,omitempty"`
	PrivateKey           string              `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LotteryRequest) Reset()         { *m = LotteryRequest{} }
func (m *LotteryRequest) String() string { return proto.CompactTextString(m) }
func (*LotteryRequest) ProtoMessage()    {}
func (*LotteryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *LotteryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LotteryRequest.Unmarshal(m, b)
}
func (m *LotteryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LotteryRequest.Marshal(b, m, deterministic)
}
func (m *LotteryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotteryRequest.Merge(m, src)
}
func (m *LotteryRequest) XXX_Size() int {
	return xxx_messageInfo_LotteryRequest.Size(m)
}
func (m *LotteryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LotteryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LotteryRequest proto.InternalMessageInfo

func (m *LotteryRequest) GetType() LotteryRequest_Type {
	if m != nil {
		return m.Type
	}
	return LotteryRequest_ENTER
}

func (m *LotteryRequest) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

// Staking Request from new node to beacon node.
type StakingRequest struct {
	Transaction          []byte   `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakingRequest) Reset()         { *m = StakingRequest{} }
func (m *StakingRequest) String() string { return proto.CompactTextString(m) }
func (*StakingRequest) ProtoMessage()    {}
func (*StakingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *StakingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingRequest.Unmarshal(m, b)
}
func (m *StakingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingRequest.Marshal(b, m, deterministic)
}
func (m *StakingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingRequest.Merge(m, src)
}
func (m *StakingRequest) XXX_Size() int {
	return xxx_messageInfo_StakingRequest.Size(m)
}
func (m *StakingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakingRequest proto.InternalMessageInfo

func (m *StakingRequest) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *StakingRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ConsensusRequest struct {
	ConsensusId          uint32   `protobuf:"varint,1,opt,name=consensus_id,json=consensusId,proto3" json:"consensus_id,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	SenderAddress        string   `protobuf:"bytes,3,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusRequest) Reset()         { *m = ConsensusRequest{} }
func (m *ConsensusRequest) String() string { return proto.CompactTextString(m) }
func (*ConsensusRequest) ProtoMessage()    {}
func (*ConsensusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *ConsensusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusRequest.Unmarshal(m, b)
}
func (m *ConsensusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusRequest.Marshal(b, m, deterministic)
}
func (m *ConsensusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusRequest.Merge(m, src)
}
func (m *ConsensusRequest) XXX_Size() int {
	return xxx_messageInfo_ConsensusRequest.Size(m)
}
func (m *ConsensusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusRequest proto.InternalMessageInfo

func (m *ConsensusRequest) GetConsensusId() uint32 {
	if m != nil {
		return m.ConsensusId
	}
	return 0
}

func (m *ConsensusRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ConsensusRequest) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *ConsensusRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DrandRequest struct {
	SenderAddress        string   `protobuf:"bytes,1,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrandRequest) Reset()         { *m = DrandRequest{} }
func (m *DrandRequest) String() string { return proto.CompactTextString(m) }
func (*DrandRequest) ProtoMessage()    {}
func (*DrandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *DrandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrandRequest.Unmarshal(m, b)
}
func (m *DrandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrandRequest.Marshal(b, m, deterministic)
}
func (m *DrandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrandRequest.Merge(m, src)
}
func (m *DrandRequest) XXX_Size() int {
	return xxx_messageInfo_DrandRequest.Size(m)
}
func (m *DrandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DrandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DrandRequest proto.InternalMessageInfo

func (m *DrandRequest) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *DrandRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *DrandRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.ReceiverType", ReceiverType_name, ReceiverType_value)
	proto.RegisterEnum("message.ServiceType", ServiceType_name, ServiceType_value)
	proto.RegisterEnum("message.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("message.LotteryRequest_Type", LotteryRequest_Type_name, LotteryRequest_Type_value)
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*Response)(nil), "message.Response")
	proto.RegisterType((*LotteryResponse)(nil), "message.LotteryResponse")
	proto.RegisterType((*LotteryRequest)(nil), "message.LotteryRequest")
	proto.RegisterType((*StakingRequest)(nil), "message.StakingRequest")
	proto.RegisterType((*ConsensusRequest)(nil), "message.ConsensusRequest")
	proto.RegisterType((*DrandRequest)(nil), "message.DrandRequest")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0x41, 0x6f, 0x9b, 0x4a,
	0x10, 0xc7, 0x8d, 0xb1, 0x8d, 0x19, 0xb0, 0xb3, 0x6f, 0xf5, 0xde, 0x0b, 0x2f, 0x4a, 0xf4, 0x5c,
	0xa4, 0x4a, 0x56, 0xa4, 0x46, 0x91, 0x53, 0xa9, 0x6a, 0x6f, 0xd8, 0xa0, 0x04, 0xc5, 0xc1, 0xd1,
	0x9a, 0xb4, 0x47, 0x44, 0xcc, 0x2a, 0x41, 0xa1, 0xe0, 0xb2, 0x24, 0x92, 0xaf, 0xfd, 0x02, 0x3d,
	0xf5, 0x9b, 0xf4, 0xe3, 0xf5, 0x50, 0xb1, 0x60, 0x83, 0xdd, 0x54, 0xbd, 0xf6, 0xc6, 0xfc, 0x67,
	0x7f, 0x33, 0xb3, 0x33, 0x3b, 0x36, 0xf4, 0x3e, 0x52, 0xc6, 0xfc, 0x3b, 0x7a, 0xb2, 0x4c, 0x93,
	0x2c, 0xc1, 0x52, 0x69, 0xea, 0xdf, 0x44, 0x90, 0xae, 0x8a, 0x6f, 0xfc, 0x0e, 0x7a, 0x29, 0x5d,
	0xd0, 0xf0, 0x89, 0xa6, 0x5e, 0xb6, 0x5a, 0x52, 0x4d, 0x18, 0x08, 0xc3, 0xfe, 0xe8, 0x9f, 0x93,
	0x35, 0x4b, 0x4a, 0xaf, 0xbb, 0x5a, 0x52, 0xa2, 0xa6, 0x35, 0x0b, 0xbf, 0x01, 0x95, 0xd1, 0xf4,
	0x29, 0x5c, 0xd0, 0x02, 0x6d, 0x72, 0xf4, 0xef, 0x0d, 0x3a, 0x2f, 0x9c, 0x9c, 0x54, 0x58, 0x65,
	0xe0, 0x21, 0xb4, 0x38, 0x20, 0xee, 0x00, 0x65, 0x51, 0x1c, 0xe0, 0x27, 0xf0, 0x21, 0xc8, 0x2c,
	0xbc, 0x8b, 0xfd, 0xec, 0x31, 0xa5, 0x5a, 0x6b, 0x20, 0x0c, 0x55, 0x52, 0x09, 0xf8, 0x0c, 0x24,
	0x96, 0xf9, 0x0f, 0x61, 0x7c, 0xa7, 0xb5, 0x07, 0xc2, 0x50, 0x19, 0xed, 0x57, 0xb9, 0x0b, 0x9d,
	0xd0, 0x4f, 0x8f, 0x94, 0x65, 0x17, 0x0d, 0xb2, 0x3e, 0x89, 0xdf, 0x82, 0xbc, 0x48, 0x62, 0x46,
	0x63, 0xf6, 0xc8, 0xb4, 0x0e, 0xc7, 0xfe, 0xdb, 0x60, 0x93, 0xb5, 0xa7, 0x02, 0xab, 0xd3, 0xf8,
	0x15, 0xb4, 0x83, 0xd4, 0x8f, 0x03, 0x4d, 0xe2, 0x58, 0xd5, 0x24, 0x33, 0x57, 0x2b, 0xa4, 0x38,
	0x85, 0xc7, 0xb0, 0x17, 0x25, 0x59, 0x46, 0xd3, 0x95, 0x97, 0x16, 0x3e, 0xad, 0xbb, 0x53, 0xe6,
	0xb4, 0xf0, 0x57, 0x68, 0x3f, 0xda, 0x52, 0xc6, 0x32, 0x48, 0x25, 0xab, 0x7f, 0x17, 0xa0, 0x4b,
	0x28, 0x5b, 0xe6, 0xe5, 0xfc, 0xe9, 0x73, 0xb3, 0x00, 0x55, 0x57, 0x2f, 0x4a, 0xe6, 0xe3, 0x53,
	0x46, 0xda, 0xcf, 0x77, 0x2f, 0xfc, 0x17, 0x0d, 0xb2, 0x17, 0x6d, 0x4b, 0x63, 0x80, 0xee, 0x1a,
	0xd7, 0xcf, 0x61, 0x6f, 0x87, 0xc0, 0x1a, 0x48, 0xcb, 0xc8, 0x5f, 0xd1, 0x94, 0x69, 0xc2, 0x40,
	0x1c, 0xca, 0x64, 0x6d, 0xe2, 0x03, 0xe8, 0xde, 0xfa, 0x91, 0x1f, 0x2f, 0x28, 0xd3, 0xc4, 0x81,
	0x38, 0x6c, 0x91, 0x8d, 0xad, 0x7f, 0x16, 0xa0, 0xbf, 0xdd, 0x77, 0x7c, 0x5a, 0x5e, 0xac, 0x68,
	0xe2, 0xe1, 0x2f, 0xc6, 0x73, 0x52, 0xbb, 0xe0, 0xff, 0xa0, 0x2c, 0xd3, 0xf0, 0xc9, 0xcf, 0xa8,
	0xf7, 0x40, 0x57, 0xbc, 0x85, 0x32, 0x81, 0x52, 0xba, 0xa4, 0x2b, 0xfd, 0x08, 0x5a, 0xbc, 0x67,
	0x32, 0xb4, 0x2d, 0xc7, 0xb5, 0x08, 0x6a, 0x60, 0x80, 0x0e, 0xb1, 0xe6, 0x37, 0x53, 0x17, 0x09,
	0xfa, 0x25, 0xf4, 0xb7, 0x9f, 0x28, 0x1e, 0x80, 0x92, 0xa5, 0x7e, 0xcc, 0xfc, 0x45, 0x16, 0x26,
	0x31, 0x2f, 0x45, 0x25, 0x75, 0x09, 0xef, 0x83, 0x14, 0x27, 0x01, 0xf5, 0xc2, 0xa0, 0xcc, 0xd7,
	0xc9, 0x4d, 0x3b, 0xd0, 0xbf, 0x0a, 0x80, 0x76, 0x5f, 0x2e, 0x7e, 0x01, 0xea, 0xe6, 0xe5, 0xe6,
	0x48, 0x1e, 0xb0, 0x47, 0x94, 0x8d, 0x66, 0x07, 0xf8, 0x08, 0xe0, 0x36, 0x4a, 0x16, 0x0f, 0xde,
	0xbd, 0xcf, 0xee, 0x79, 0x4c, 0x95, 0xc8, 0x5c, 0xb9, 0xf0, 0xd9, 0x3d, 0x7e, 0x09, 0x7d, 0x46,
	0xe3, 0x80, 0xa6, 0x9e, 0x1f, 0x04, 0x29, 0x65, 0x8c, 0x0f, 0x5e, 0x26, 0xbd, 0x42, 0x35, 0x0a,
	0x91, 0x4f, 0xc1, 0x5f, 0x45, 0x89, 0x1f, 0x94, 0x1b, 0xba, 0x36, 0xf5, 0x18, 0xd4, 0xfa, 0x66,
	0x3c, 0x13, 0x50, 0x78, 0x2e, 0xe0, 0x6f, 0xca, 0xaa, 0xe5, 0x13, 0xb7, 0xf2, 0x1d, 0x8f, 0x41,
	0xad, 0x3f, 0x7b, 0xac, 0x80, 0xe4, 0x58, 0x1f, 0x9c, 0x99, 0x69, 0x15, 0xdd, 0x9f, 0x5a, 0x86,
	0x69, 0x11, 0x24, 0xe0, 0x1e, 0xc8, 0xef, 0x8d, 0xa9, 0x6d, 0x1a, 0xee, 0x8c, 0xa0, 0x66, 0xee,
	0x9a, 0x4c, 0x6d, 0xcb, 0x71, 0x91, 0x78, 0xfc, 0x1a, 0x94, 0xda, 0xfb, 0xcf, 0x4f, 0x4e, 0x66,
	0xce, 0xdc, 0x72, 0xe6, 0x37, 0x73, 0xd4, 0xc8, 0x23, 0xce, 0x5d, 0xe3, 0xd2, 0x76, 0xce, 0x91,
	0x90, 0x8f, 0xd6, 0x24, 0x86, 0x63, 0xa2, 0xe6, 0xf1, 0x17, 0x01, 0x94, 0xda, 0x16, 0xe0, 0x03,
	0xf8, 0xb7, 0xcc, 0xec, 0x8d, 0x2d, 0x63, 0x32, 0x73, 0xbc, 0x35, 0xd6, 0xc0, 0x2a, 0x74, 0x0d,
	0xc7, 0x99, 0xdd, 0x38, 0x13, 0x0b, 0x09, 0x79, 0xc4, 0x6b, 0x62, 0x5d, 0x1b, 0xc4, 0x42, 0xcd,
	0xdc, 0x55, 0x1a, 0x26, 0x12, 0x79, 0x59, 0xb3, 0xab, 0x2b, 0xdb, 0x45, 0xad, 0xa2, 0x8e, 0xfc,
	0xdb, 0xb5, 0x4c, 0xd4, 0xc6, 0x7d, 0x00, 0x9e, 0xda, 0xb3, 0x1d, 0xdb, 0x45, 0x1d, 0x8c, 0x40,
	0x2d, 0xec, 0x12, 0x90, 0x46, 0x06, 0xf4, 0x26, 0x51, 0x48, 0xe3, 0xac, 0xbc, 0x0d, 0x3e, 0x05,
	0xe9, 0x3a, 0x4d, 0x16, 0x79, 0x83, 0xd1, 0xee, 0xe6, 0x1e, 0xfc, 0x55, 0xfb, 0xdd, 0x28, 0xf7,
	0xad, 0x71, 0xdb, 0xe1, 0xff, 0x1b, 0x67, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0x2a, 0xe5,
	0x7e, 0x48, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.ClientService/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	Process(context.Context, *Message) (*Response, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.ClientService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Process(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _ClientService_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
