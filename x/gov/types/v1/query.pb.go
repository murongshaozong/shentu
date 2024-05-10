// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1/query.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryCertVotedRequest struct {
	// proposal_id defines the unique id of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *QueryCertVotedRequest) Reset()         { *m = QueryCertVotedRequest{} }
func (m *QueryCertVotedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCertVotedRequest) ProtoMessage()    {}
func (*QueryCertVotedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d1918b998d725b, []int{0}
}
func (m *QueryCertVotedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertVotedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertVotedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertVotedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertVotedRequest.Merge(m, src)
}
func (m *QueryCertVotedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertVotedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertVotedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertVotedRequest proto.InternalMessageInfo

func (m *QueryCertVotedRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// QueryProposalResponse is the response type for the Query/Proposal RPC method.
type QueryCertVotedResponse struct {
	CertVoted bool `protobuf:"varint,1,opt,name=cert_voted,json=certVoted,proto3" json:"cert_voted,omitempty"`
}

func (m *QueryCertVotedResponse) Reset()         { *m = QueryCertVotedResponse{} }
func (m *QueryCertVotedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCertVotedResponse) ProtoMessage()    {}
func (*QueryCertVotedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d1918b998d725b, []int{1}
}
func (m *QueryCertVotedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertVotedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertVotedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertVotedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertVotedResponse.Merge(m, src)
}
func (m *QueryCertVotedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertVotedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertVotedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertVotedResponse proto.InternalMessageInfo

func (m *QueryCertVotedResponse) GetCertVoted() bool {
	if m != nil {
		return m.CertVoted
	}
	return false
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// voting_params defines the parameters related to voting.
	VotingParams v1.VotingParams `protobuf:"bytes,1,opt,name=voting_params,json=votingParams,proto3" json:"voting_params"`
	// deposit_params defines the parameters related to deposit.
	DepositParams v1.DepositParams `protobuf:"bytes,2,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params"`
	// tally_params defines the parameters related to tally.
	TallyParams v1.TallyParams `protobuf:"bytes,3,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params"`
	// custom_params defines the parameters related to custom.
	CustomParams CustomParams `protobuf:"bytes,4,opt,name=custom_params,json=customParams,proto3" json:"custom_params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d1918b998d725b, []int{2}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetVotingParams() v1.VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return v1.VotingParams{}
}

func (m *QueryParamsResponse) GetDepositParams() v1.DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return v1.DepositParams{}
}

func (m *QueryParamsResponse) GetTallyParams() v1.TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return v1.TallyParams{}
}

func (m *QueryParamsResponse) GetCustomParams() CustomParams {
	if m != nil {
		return m.CustomParams
	}
	return CustomParams{}
}

func init() {
	proto.RegisterType((*QueryCertVotedRequest)(nil), "shentu.gov.v1.QueryCertVotedRequest")
	proto.RegisterType((*QueryCertVotedResponse)(nil), "shentu.gov.v1.QueryCertVotedResponse")
	proto.RegisterType((*QueryParamsResponse)(nil), "shentu.gov.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("shentu/gov/v1/query.proto", fileDescriptor_54d1918b998d725b) }

var fileDescriptor_54d1918b998d725b = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0xc7, 0xbb, 0x58, 0x90, 0x4e, 0xc1, 0xc3, 0xf3, 0x07, 0xb0, 0x62, 0xc1, 0xe5, 0xa7, 0xd1,
	0xec, 0xa4, 0x48, 0xc4, 0x44, 0xbd, 0x50, 0x63, 0x42, 0xbc, 0x60, 0x63, 0x38, 0x78, 0x69, 0x96,
	0x76, 0x5c, 0x9a, 0xb4, 0x3b, 0xcb, 0xce, 0xec, 0x46, 0xac, 0xbd, 0x90, 0x18, 0xe3, 0xc5, 0x98,
	0x68, 0xe2, 0x5f, 0xe2, 0xff, 0xc0, 0x91, 0xc4, 0x8b, 0x27, 0x63, 0xc0, 0xf8, 0x77, 0x98, 0x99,
	0x9d, 0x29, 0xdd, 0x4d, 0x59, 0xc0, 0xd3, 0x6e, 0xe6, 0x7d, 0xdf, 0xf7, 0x7d, 0x66, 0xe6, 0xcd,
	0x0c, 0x9a, 0x62, 0x3b, 0xc4, 0xe3, 0x21, 0x76, 0x69, 0x84, 0xa3, 0x32, 0xde, 0x0d, 0x49, 0xb0,
	0x67, 0xfb, 0x01, 0xe5, 0x14, 0xc6, 0xe3, 0x90, 0xed, 0xd2, 0xc8, 0x8e, 0xca, 0xe6, 0x35, 0x97,
	0xba, 0x54, 0x46, 0xb0, 0xf8, 0x8b, 0x45, 0xe6, 0xb4, 0x4b, 0xa9, 0xdb, 0x22, 0xd8, 0xf1, 0x9b,
	0xd8, 0xf1, 0x3c, 0xca, 0x1d, 0xde, 0xa4, 0x1e, 0x53, 0xd1, 0x89, 0x3a, 0x65, 0x6d, 0xca, 0xb4,
	0xbb, 0x70, 0x8a, 0x03, 0x53, 0xc9, 0x40, 0x5f, 0x59, 0x73, 0x22, 0x49, 0xd4, 0xcb, 0xb1, 0x1e,
	0xa2, 0xeb, 0x2f, 0x84, 0xae, 0x42, 0x02, 0xbe, 0x45, 0x39, 0x69, 0x54, 0xc9, 0x6e, 0x48, 0x18,
	0x87, 0x19, 0x54, 0xf4, 0x03, 0xea, 0x53, 0xe6, 0xb4, 0x6a, 0xcd, 0xc6, 0xa4, 0x31, 0x6b, 0x2c,
	0xe7, 0xab, 0x48, 0x0f, 0x6d, 0x34, 0xac, 0x35, 0x74, 0x23, 0x9d, 0xc9, 0x7c, 0xea, 0x31, 0x02,
	0xb7, 0x10, 0xaa, 0x93, 0x80, 0xd7, 0x22, 0x31, 0x2a, 0x33, 0x47, 0xab, 0x85, 0xba, 0x96, 0x59,
	0xdf, 0x87, 0xd0, 0x55, 0x99, 0xb9, 0xe9, 0x04, 0x4e, 0x9b, 0xf5, 0xd2, 0x9e, 0xa1, 0xf1, 0x88,
	0xf2, 0xa6, 0xe7, 0xd6, 0x7c, 0x19, 0x90, 0x99, 0xc5, 0x95, 0x9b, 0x76, 0x3c, 0x2d, 0xb5, 0x64,
	0xf6, 0x96, 0xd4, 0xc4, 0xb9, 0xeb, 0xf9, 0x83, 0x5f, 0x33, 0xb9, 0xea, 0x58, 0xd4, 0x37, 0x06,
	0x1b, 0xe8, 0x4a, 0x83, 0xf8, 0x94, 0x35, 0xb9, 0x36, 0x1a, 0x92, 0x46, 0xd3, 0x29, 0xa3, 0xa7,
	0xb1, 0x28, 0xe1, 0x34, 0xde, 0xe8, 0x1f, 0x84, 0x0a, 0x1a, 0xe3, 0x4e, 0xab, 0xb5, 0xa7, 0x8d,
	0x2e, 0x49, 0x23, 0x33, 0x65, 0xf4, 0x52, 0x48, 0x12, 0x36, 0x45, 0x7e, 0x32, 0x24, 0xe6, 0x55,
	0x0f, 0x19, 0xa7, 0x6d, 0xed, 0x92, 0x57, 0xf3, 0x4a, 0xb4, 0x82, 0x5d, 0x91, 0x9a, 0xe4, 0xbc,
	0xea, 0x7d, 0x63, 0x2b, 0x7f, 0x0b, 0x68, 0x58, 0xae, 0x1b, 0xbc, 0x37, 0xd0, 0xe8, 0xa6, 0xda,
	0x09, 0x98, 0x4b, 0xd1, 0xc4, 0x4b, 0xab, 0xa2, 0x6a, 0x37, 0xcd, 0xf9, 0x6c, 0x51, 0xbc, 0x03,
	0x96, 0xbd, 0xff, 0xe3, 0xcf, 0x97, 0xa1, 0x65, 0x58, 0xc4, 0xc9, 0x4e, 0xd2, 0xbb, 0xce, 0x70,
	0xa7, 0xaf, 0x27, 0xba, 0xf0, 0x16, 0x15, 0xb4, 0x07, 0x83, 0xcc, 0x12, 0x4c, 0x83, 0x2c, 0x9c,
	0xa1, 0x52, 0x24, 0xb3, 0x92, 0xc4, 0x84, 0xc9, 0xd3, 0x48, 0xe0, 0x83, 0x81, 0xf2, 0xa2, 0x9f,
	0x60, 0x66, 0x90, 0xa3, 0x88, 0xe8, 0x92, 0xb3, 0xa7, 0x0b, 0x54, 0xb5, 0xc7, 0xb2, 0xda, 0x03,
	0x58, 0x3d, 0xdf, 0xbc, 0xb1, 0x68, 0x6c, 0x86, 0x3b, 0xe2, 0x13, 0x74, 0x61, 0xdf, 0x40, 0xc3,
	0xc2, 0x8e, 0xc1, 0xa9, 0x95, 0x7a, 0xd3, 0xbf, 0x9d, 0xa1, 0x50, 0x30, 0xab, 0x12, 0xc6, 0x86,
	0x7b, 0x17, 0x81, 0x81, 0x77, 0x68, 0x44, 0xb5, 0xdb, 0xc0, 0x12, 0xfa, 0xa8, 0xc5, 0x14, 0x56,
	0xaa, 0xf5, 0x06, 0x9c, 0x46, 0xeb, 0xae, 0xc4, 0x58, 0x80, 0xb9, 0x34, 0x86, 0x94, 0xe1, 0x4e,
	0xfc, 0xad, 0xf1, 0x3d, 0x9f, 0x74, 0xe1, 0x9b, 0x81, 0x2e, 0xab, 0xe3, 0x04, 0xd6, 0xa0, 0xfa,
	0x2a, 0xa8, 0x01, 0xe6, 0x32, 0x35, 0x8a, 0xa0, 0x22, 0x09, 0x9e, 0xc0, 0xa3, 0x73, 0x2e, 0x84,
	0x3a, 0xba, 0x0c, 0x77, 0xd4, 0x1f, 0x0d, 0xba, 0xf0, 0xc9, 0x40, 0xa3, 0xca, 0x98, 0x41, 0x56,
	0x59, 0x96, 0x79, 0x54, 0x4e, 0x44, 0x0a, 0x6e, 0x4d, 0xc2, 0x95, 0x01, 0x5f, 0x10, 0x0e, 0xbe,
	0x1a, 0xa8, 0x28, 0x2f, 0x8c, 0x2a, 0x61, 0x61, 0x8b, 0xc3, 0xe2, 0xa0, 0x72, 0x7d, 0x02, 0x8d,
	0xb5, 0x74, 0xa6, 0xee, 0x3f, 0xfb, 0x47, 0x5e, 0x55, 0xf0, 0xd1, 0x40, 0x85, 0xde, 0x4d, 0x0e,
	0xf3, 0x83, 0x1a, 0x24, 0xfd, 0x44, 0x98, 0x0b, 0x67, 0xa8, 0x14, 0x10, 0x96, 0x40, 0x77, 0x60,
	0x09, 0x27, 0x1f, 0xa1, 0x93, 0x37, 0x22, 0x49, 0xb4, 0xfe, 0xfc, 0xe0, 0xa8, 0x64, 0x1c, 0x1e,
	0x95, 0x8c, 0xdf, 0x47, 0x25, 0xe3, 0xf3, 0x71, 0x29, 0x77, 0x78, 0x5c, 0xca, 0xfd, 0x3c, 0x2e,
	0xe5, 0x5e, 0x95, 0xdd, 0x26, 0xdf, 0x09, 0xb7, 0xed, 0x3a, 0x6d, 0x2b, 0xb3, 0xd7, 0x34, 0xf4,
	0x1a, 0xf2, 0x79, 0xd4, 0xee, 0x6f, 0xa4, 0xbf, 0xe8, 0x4b, 0x86, 0xa3, 0xf2, 0xf6, 0x88, 0x7c,
	0xe7, 0xee, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x69, 0xe0, 0x9b, 0x64, 0x94, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(ctx context.Context, in *v1.QueryProposalRequest, opts ...grpc.CallOption) (*v1.QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(ctx context.Context, in *v1.QueryProposalsRequest, opts ...grpc.CallOption) (*v1.QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(ctx context.Context, in *v1.QueryVoteRequest, opts ...grpc.CallOption) (*v1.QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(ctx context.Context, in *v1.QueryVotesRequest, opts ...grpc.CallOption) (*v1.QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(ctx context.Context, in *v1.QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(ctx context.Context, in *v1.QueryDepositRequest, opts ...grpc.CallOption) (*v1.QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(ctx context.Context, in *v1.QueryDepositsRequest, opts ...grpc.CallOption) (*v1.QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(ctx context.Context, in *v1.QueryTallyResultRequest, opts ...grpc.CallOption) (*v1.QueryTallyResultResponse, error)
	// Proposal queries proposal details based on ProposalID.
	CertVoted(ctx context.Context, in *QueryCertVotedRequest, opts ...grpc.CallOption) (*QueryCertVotedResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proposal(ctx context.Context, in *v1.QueryProposalRequest, opts ...grpc.CallOption) (*v1.QueryProposalResponse, error) {
	out := new(v1.QueryProposalResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Proposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proposals(ctx context.Context, in *v1.QueryProposalsRequest, opts ...grpc.CallOption) (*v1.QueryProposalsResponse, error) {
	out := new(v1.QueryProposalsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Proposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vote(ctx context.Context, in *v1.QueryVoteRequest, opts ...grpc.CallOption) (*v1.QueryVoteResponse, error) {
	out := new(v1.QueryVoteResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Votes(ctx context.Context, in *v1.QueryVotesRequest, opts ...grpc.CallOption) (*v1.QueryVotesResponse, error) {
	out := new(v1.QueryVotesResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Votes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *v1.QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposit(ctx context.Context, in *v1.QueryDepositRequest, opts ...grpc.CallOption) (*v1.QueryDepositResponse, error) {
	out := new(v1.QueryDepositResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposits(ctx context.Context, in *v1.QueryDepositsRequest, opts ...grpc.CallOption) (*v1.QueryDepositsResponse, error) {
	out := new(v1.QueryDepositsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/Deposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TallyResult(ctx context.Context, in *v1.QueryTallyResultRequest, opts ...grpc.CallOption) (*v1.QueryTallyResultResponse, error) {
	out := new(v1.QueryTallyResultResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/TallyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CertVoted(ctx context.Context, in *QueryCertVotedRequest, opts ...grpc.CallOption) (*QueryCertVotedResponse, error) {
	out := new(QueryCertVotedResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1.Query/CertVoted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(context.Context, *v1.QueryProposalRequest) (*v1.QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(context.Context, *v1.QueryProposalsRequest) (*v1.QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(context.Context, *v1.QueryVoteRequest) (*v1.QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(context.Context, *v1.QueryVotesRequest) (*v1.QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(context.Context, *v1.QueryParamsRequest) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(context.Context, *v1.QueryDepositRequest) (*v1.QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(context.Context, *v1.QueryDepositsRequest) (*v1.QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(context.Context, *v1.QueryTallyResultRequest) (*v1.QueryTallyResultResponse, error)
	// Proposal queries proposal details based on ProposalID.
	CertVoted(context.Context, *QueryCertVotedRequest) (*QueryCertVotedResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proposal(ctx context.Context, req *v1.QueryProposalRequest) (*v1.QueryProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposal not implemented")
}
func (*UnimplementedQueryServer) Proposals(ctx context.Context, req *v1.QueryProposalsRequest) (*v1.QueryProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposals not implemented")
}
func (*UnimplementedQueryServer) Vote(ctx context.Context, req *v1.QueryVoteRequest) (*v1.QueryVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedQueryServer) Votes(ctx context.Context, req *v1.QueryVotesRequest) (*v1.QueryVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Votes not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *v1.QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Deposit(ctx context.Context, req *v1.QueryDepositRequest) (*v1.QueryDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedQueryServer) Deposits(ctx context.Context, req *v1.QueryDepositsRequest) (*v1.QueryDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposits not implemented")
}
func (*UnimplementedQueryServer) TallyResult(ctx context.Context, req *v1.QueryTallyResultRequest) (*v1.QueryTallyResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TallyResult not implemented")
}
func (*UnimplementedQueryServer) CertVoted(ctx context.Context, req *QueryCertVotedRequest) (*QueryCertVotedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CertVoted not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Proposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposal(ctx, req.(*v1.QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Proposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposals(ctx, req.(*v1.QueryProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vote(ctx, req.(*v1.QueryVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Votes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Votes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Votes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Votes(ctx, req.(*v1.QueryVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*v1.QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposit(ctx, req.(*v1.QueryDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/Deposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposits(ctx, req.(*v1.QueryDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TallyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryTallyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TallyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/TallyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TallyResult(ctx, req.(*v1.QueryTallyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CertVoted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCertVotedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CertVoted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1.Query/CertVoted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CertVoted(ctx, req.(*QueryCertVotedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.gov.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proposal",
			Handler:    _Query_Proposal_Handler,
		},
		{
			MethodName: "Proposals",
			Handler:    _Query_Proposals_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Query_Vote_Handler,
		},
		{
			MethodName: "Votes",
			Handler:    _Query_Votes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Query_Deposit_Handler,
		},
		{
			MethodName: "Deposits",
			Handler:    _Query_Deposits_Handler,
		},
		{
			MethodName: "TallyResult",
			Handler:    _Query_TallyResult_Handler,
		},
		{
			MethodName: "CertVoted",
			Handler:    _Query_CertVoted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/gov/v1/query.proto",
}

func (m *QueryCertVotedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertVotedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertVotedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCertVotedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertVotedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertVotedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertVoted {
		i--
		if m.CertVoted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CustomParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCertVotedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	return n
}

func (m *QueryCertVotedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CertVoted {
		n += 2
	}
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VotingParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.DepositParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.CustomParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCertVotedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertVotedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertVotedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertVotedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertVotedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertVotedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertVoted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CertVoted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CustomParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
