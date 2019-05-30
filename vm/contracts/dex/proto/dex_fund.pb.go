// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dex_fund.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Available            []byte   `protobuf:"bytes,2,opt,name=Available,proto3" json:"Available,omitempty"`
	Locked               []byte   `protobuf:"bytes,3,opt,name=Locked,proto3" json:"Locked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Account) GetAvailable() []byte {
	if m != nil {
		return m.Available
	}
	return nil
}

func (m *Account) GetLocked() []byte {
	if m != nil {
		return m.Locked
	}
	return nil
}

type Fund struct {
	Address              []byte     `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Accounts             []*Account `protobuf:"bytes,2,rep,name=Accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Fund) Reset()         { *m = Fund{} }
func (m *Fund) String() string { return proto.CompactTextString(m) }
func (*Fund) ProtoMessage()    {}
func (*Fund) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{1}
}

func (m *Fund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fund.Unmarshal(m, b)
}
func (m *Fund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fund.Marshal(b, m, deterministic)
}
func (m *Fund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fund.Merge(m, src)
}
func (m *Fund) XXX_Size() int {
	return xxx_messageInfo_Fund.Size(m)
}
func (m *Fund) XXX_DiscardUnknown() {
	xxx_messageInfo_Fund.DiscardUnknown(m)
}

var xxx_messageInfo_Fund proto.InternalMessageInfo

func (m *Fund) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Fund) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type FundSettle struct {
	IsTradeToken         bool     `protobuf:"varint,1,opt,name=isTradeToken,proto3" json:"isTradeToken,omitempty"`
	IncAvailable         []byte   `protobuf:"bytes,2,opt,name=IncAvailable,proto3" json:"IncAvailable,omitempty"`
	ReduceLocked         []byte   `protobuf:"bytes,3,opt,name=ReduceLocked,proto3" json:"ReduceLocked,omitempty"`
	ReleaseLocked        []byte   `protobuf:"bytes,4,opt,name=ReleaseLocked,proto3" json:"ReleaseLocked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundSettle) Reset()         { *m = FundSettle{} }
func (m *FundSettle) String() string { return proto.CompactTextString(m) }
func (*FundSettle) ProtoMessage()    {}
func (*FundSettle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{2}
}

func (m *FundSettle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundSettle.Unmarshal(m, b)
}
func (m *FundSettle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundSettle.Marshal(b, m, deterministic)
}
func (m *FundSettle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundSettle.Merge(m, src)
}
func (m *FundSettle) XXX_Size() int {
	return xxx_messageInfo_FundSettle.Size(m)
}
func (m *FundSettle) XXX_DiscardUnknown() {
	xxx_messageInfo_FundSettle.DiscardUnknown(m)
}

var xxx_messageInfo_FundSettle proto.InternalMessageInfo

func (m *FundSettle) GetIsTradeToken() bool {
	if m != nil {
		return m.IsTradeToken
	}
	return false
}

func (m *FundSettle) GetIncAvailable() []byte {
	if m != nil {
		return m.IncAvailable
	}
	return nil
}

func (m *FundSettle) GetReduceLocked() []byte {
	if m != nil {
		return m.ReduceLocked
	}
	return nil
}

func (m *FundSettle) GetReleaseLocked() []byte {
	if m != nil {
		return m.ReleaseLocked
	}
	return nil
}

type UserFundSettle struct {
	Address              []byte        `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	FundSettles          []*FundSettle `protobuf:"bytes,2,rep,name=fundSettles,proto3" json:"fundSettles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserFundSettle) Reset()         { *m = UserFundSettle{} }
func (m *UserFundSettle) String() string { return proto.CompactTextString(m) }
func (*UserFundSettle) ProtoMessage()    {}
func (*UserFundSettle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{3}
}

func (m *UserFundSettle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFundSettle.Unmarshal(m, b)
}
func (m *UserFundSettle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFundSettle.Marshal(b, m, deterministic)
}
func (m *UserFundSettle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFundSettle.Merge(m, src)
}
func (m *UserFundSettle) XXX_Size() int {
	return xxx_messageInfo_UserFundSettle.Size(m)
}
func (m *UserFundSettle) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFundSettle.DiscardUnknown(m)
}

var xxx_messageInfo_UserFundSettle proto.InternalMessageInfo

func (m *UserFundSettle) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserFundSettle) GetFundSettles() []*FundSettle {
	if m != nil {
		return m.FundSettles
	}
	return nil
}

type UserFeeSettle struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	BaseFee              []byte   `protobuf:"bytes,2,opt,name=BaseFee,proto3" json:"BaseFee,omitempty"`
	BrokerFee            []byte   `protobuf:"bytes,3,opt,name=BrokerFee,proto3" json:"BrokerFee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFeeSettle) Reset()         { *m = UserFeeSettle{} }
func (m *UserFeeSettle) String() string { return proto.CompactTextString(m) }
func (*UserFeeSettle) ProtoMessage()    {}
func (*UserFeeSettle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{4}
}

func (m *UserFeeSettle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFeeSettle.Unmarshal(m, b)
}
func (m *UserFeeSettle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFeeSettle.Marshal(b, m, deterministic)
}
func (m *UserFeeSettle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFeeSettle.Merge(m, src)
}
func (m *UserFeeSettle) XXX_Size() int {
	return xxx_messageInfo_UserFeeSettle.Size(m)
}
func (m *UserFeeSettle) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFeeSettle.DiscardUnknown(m)
}

var xxx_messageInfo_UserFeeSettle proto.InternalMessageInfo

func (m *UserFeeSettle) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserFeeSettle) GetBaseFee() []byte {
	if m != nil {
		return m.BaseFee
	}
	return nil
}

func (m *UserFeeSettle) GetBrokerFee() []byte {
	if m != nil {
		return m.BrokerFee
	}
	return nil
}

type FeeSettle struct {
	Token                []byte           `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Broker               []byte           `protobuf:"bytes,2,opt,name=Broker,proto3" json:"Broker,omitempty"`
	UserFeeSettles       []*UserFeeSettle `protobuf:"bytes,3,rep,name=userFeeSettles,proto3" json:"userFeeSettles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FeeSettle) Reset()         { *m = FeeSettle{} }
func (m *FeeSettle) String() string { return proto.CompactTextString(m) }
func (*FeeSettle) ProtoMessage()    {}
func (*FeeSettle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{5}
}

func (m *FeeSettle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeeSettle.Unmarshal(m, b)
}
func (m *FeeSettle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeeSettle.Marshal(b, m, deterministic)
}
func (m *FeeSettle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeSettle.Merge(m, src)
}
func (m *FeeSettle) XXX_Size() int {
	return xxx_messageInfo_FeeSettle.Size(m)
}
func (m *FeeSettle) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeSettle.DiscardUnknown(m)
}

var xxx_messageInfo_FeeSettle proto.InternalMessageInfo

func (m *FeeSettle) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *FeeSettle) GetBroker() []byte {
	if m != nil {
		return m.Broker
	}
	return nil
}

func (m *FeeSettle) GetUserFeeSettles() []*UserFeeSettle {
	if m != nil {
		return m.UserFeeSettles
	}
	return nil
}

type SettleActions struct {
	TradeToken           []byte            `protobuf:"bytes,1,opt,name=TradeToken,proto3" json:"TradeToken,omitempty"`
	QuoteToken           []byte            `protobuf:"bytes,2,opt,name=QuoteToken,proto3" json:"QuoteToken,omitempty"`
	FundActions          []*UserFundSettle `protobuf:"bytes,3,rep,name=fundActions,proto3" json:"fundActions,omitempty"`
	FeeActions           []*UserFeeSettle  `protobuf:"bytes,4,rep,name=feeActions,proto3" json:"feeActions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SettleActions) Reset()         { *m = SettleActions{} }
func (m *SettleActions) String() string { return proto.CompactTextString(m) }
func (*SettleActions) ProtoMessage()    {}
func (*SettleActions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{6}
}

func (m *SettleActions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleActions.Unmarshal(m, b)
}
func (m *SettleActions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleActions.Marshal(b, m, deterministic)
}
func (m *SettleActions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleActions.Merge(m, src)
}
func (m *SettleActions) XXX_Size() int {
	return xxx_messageInfo_SettleActions.Size(m)
}
func (m *SettleActions) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleActions.DiscardUnknown(m)
}

var xxx_messageInfo_SettleActions proto.InternalMessageInfo

func (m *SettleActions) GetTradeToken() []byte {
	if m != nil {
		return m.TradeToken
	}
	return nil
}

func (m *SettleActions) GetQuoteToken() []byte {
	if m != nil {
		return m.QuoteToken
	}
	return nil
}

func (m *SettleActions) GetFundActions() []*UserFundSettle {
	if m != nil {
		return m.FundActions
	}
	return nil
}

func (m *SettleActions) GetFeeActions() []*UserFeeSettle {
	if m != nil {
		return m.FeeActions
	}
	return nil
}

type FeeSumAccount struct {
	Token             []byte `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	BaseAmount        []byte `protobuf:"bytes,2,opt,name=BaseAmount,proto3" json:"BaseAmount,omitempty"`
	InviteBonusAmount []byte `protobuf:"bytes,3,opt,name=InviteBonusAmount,proto3" json:"InviteBonusAmount,omitempty"`
	// rolled amount : 99% part of last period BaseAmount rolled to this period +
	// new market fee of current period +
	// new inviter fee
	DividendPoolAmount   []byte   `protobuf:"bytes,4,opt,name=DividendPoolAmount,proto3" json:"DividendPoolAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeeSumAccount) Reset()         { *m = FeeSumAccount{} }
func (m *FeeSumAccount) String() string { return proto.CompactTextString(m) }
func (*FeeSumAccount) ProtoMessage()    {}
func (*FeeSumAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{7}
}

func (m *FeeSumAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeeSumAccount.Unmarshal(m, b)
}
func (m *FeeSumAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeeSumAccount.Marshal(b, m, deterministic)
}
func (m *FeeSumAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeSumAccount.Merge(m, src)
}
func (m *FeeSumAccount) XXX_Size() int {
	return xxx_messageInfo_FeeSumAccount.Size(m)
}
func (m *FeeSumAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeSumAccount.DiscardUnknown(m)
}

var xxx_messageInfo_FeeSumAccount proto.InternalMessageInfo

func (m *FeeSumAccount) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *FeeSumAccount) GetBaseAmount() []byte {
	if m != nil {
		return m.BaseAmount
	}
	return nil
}

func (m *FeeSumAccount) GetInviteBonusAmount() []byte {
	if m != nil {
		return m.InviteBonusAmount
	}
	return nil
}

func (m *FeeSumAccount) GetDividendPoolAmount() []byte {
	if m != nil {
		return m.DividendPoolAmount
	}
	return nil
}

type FeeSumByPeriod struct {
	Fees                 []*FeeSumAccount `protobuf:"bytes,1,rep,name=fees,proto3" json:"fees,omitempty"`
	LastValidPeriod      uint64           `protobuf:"varint,2,opt,name=lastValidPeriod,proto3" json:"lastValidPeriod,omitempty"`
	FeeDivided           bool             `protobuf:"varint,3,opt,name=FeeDivided,proto3" json:"FeeDivided,omitempty"`
	MinedVxDivided       bool             `protobuf:"varint,4,opt,name=MinedVxDivided,proto3" json:"MinedVxDivided,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FeeSumByPeriod) Reset()         { *m = FeeSumByPeriod{} }
func (m *FeeSumByPeriod) String() string { return proto.CompactTextString(m) }
func (*FeeSumByPeriod) ProtoMessage()    {}
func (*FeeSumByPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{8}
}

func (m *FeeSumByPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeeSumByPeriod.Unmarshal(m, b)
}
func (m *FeeSumByPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeeSumByPeriod.Marshal(b, m, deterministic)
}
func (m *FeeSumByPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeSumByPeriod.Merge(m, src)
}
func (m *FeeSumByPeriod) XXX_Size() int {
	return xxx_messageInfo_FeeSumByPeriod.Size(m)
}
func (m *FeeSumByPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeSumByPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_FeeSumByPeriod proto.InternalMessageInfo

func (m *FeeSumByPeriod) GetFees() []*FeeSumAccount {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *FeeSumByPeriod) GetLastValidPeriod() uint64 {
	if m != nil {
		return m.LastValidPeriod
	}
	return 0
}

func (m *FeeSumByPeriod) GetFeeDivided() bool {
	if m != nil {
		return m.FeeDivided
	}
	return false
}

func (m *FeeSumByPeriod) GetMinedVxDivided() bool {
	if m != nil {
		return m.MinedVxDivided
	}
	return false
}

type UserFees struct {
	Fees                 []*UserFeeWithPeriod `protobuf:"bytes,1,rep,name=fees,proto3" json:"fees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserFees) Reset()         { *m = UserFees{} }
func (m *UserFees) String() string { return proto.CompactTextString(m) }
func (*UserFees) ProtoMessage()    {}
func (*UserFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{9}
}

func (m *UserFees) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFees.Unmarshal(m, b)
}
func (m *UserFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFees.Marshal(b, m, deterministic)
}
func (m *UserFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFees.Merge(m, src)
}
func (m *UserFees) XXX_Size() int {
	return xxx_messageInfo_UserFees.Size(m)
}
func (m *UserFees) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFees.DiscardUnknown(m)
}

var xxx_messageInfo_UserFees proto.InternalMessageInfo

func (m *UserFees) GetFees() []*UserFeeWithPeriod {
	if m != nil {
		return m.Fees
	}
	return nil
}

type UserFeeAccount struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	BaseAmount           []byte   `protobuf:"bytes,2,opt,name=BaseAmount,proto3" json:"BaseAmount,omitempty"`
	InviteBonusAmount    []byte   `protobuf:"bytes,3,opt,name=InviteBonusAmount,proto3" json:"InviteBonusAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFeeAccount) Reset()         { *m = UserFeeAccount{} }
func (m *UserFeeAccount) String() string { return proto.CompactTextString(m) }
func (*UserFeeAccount) ProtoMessage()    {}
func (*UserFeeAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{10}
}

func (m *UserFeeAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFeeAccount.Unmarshal(m, b)
}
func (m *UserFeeAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFeeAccount.Marshal(b, m, deterministic)
}
func (m *UserFeeAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFeeAccount.Merge(m, src)
}
func (m *UserFeeAccount) XXX_Size() int {
	return xxx_messageInfo_UserFeeAccount.Size(m)
}
func (m *UserFeeAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFeeAccount.DiscardUnknown(m)
}

var xxx_messageInfo_UserFeeAccount proto.InternalMessageInfo

func (m *UserFeeAccount) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *UserFeeAccount) GetBaseAmount() []byte {
	if m != nil {
		return m.BaseAmount
	}
	return nil
}

func (m *UserFeeAccount) GetInviteBonusAmount() []byte {
	if m != nil {
		return m.InviteBonusAmount
	}
	return nil
}

type UserFeeWithPeriod struct {
	UserFees             []*UserFeeAccount `protobuf:"bytes,1,rep,name=userFees,proto3" json:"userFees,omitempty"`
	Period               uint64            `protobuf:"varint,2,opt,name=Period,proto3" json:"Period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserFeeWithPeriod) Reset()         { *m = UserFeeWithPeriod{} }
func (m *UserFeeWithPeriod) String() string { return proto.CompactTextString(m) }
func (*UserFeeWithPeriod) ProtoMessage()    {}
func (*UserFeeWithPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{11}
}

func (m *UserFeeWithPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFeeWithPeriod.Unmarshal(m, b)
}
func (m *UserFeeWithPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFeeWithPeriod.Marshal(b, m, deterministic)
}
func (m *UserFeeWithPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFeeWithPeriod.Merge(m, src)
}
func (m *UserFeeWithPeriod) XXX_Size() int {
	return xxx_messageInfo_UserFeeWithPeriod.Size(m)
}
func (m *UserFeeWithPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFeeWithPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_UserFeeWithPeriod proto.InternalMessageInfo

func (m *UserFeeWithPeriod) GetUserFees() []*UserFeeAccount {
	if m != nil {
		return m.UserFees
	}
	return nil
}

func (m *UserFeeWithPeriod) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type BrokerFeeAccount struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Amount               []byte   `protobuf:"bytes,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerFeeAccount) Reset()         { *m = BrokerFeeAccount{} }
func (m *BrokerFeeAccount) String() string { return proto.CompactTextString(m) }
func (*BrokerFeeAccount) ProtoMessage()    {}
func (*BrokerFeeAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{12}
}

func (m *BrokerFeeAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerFeeAccount.Unmarshal(m, b)
}
func (m *BrokerFeeAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerFeeAccount.Marshal(b, m, deterministic)
}
func (m *BrokerFeeAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerFeeAccount.Merge(m, src)
}
func (m *BrokerFeeAccount) XXX_Size() int {
	return xxx_messageInfo_BrokerFeeAccount.Size(m)
}
func (m *BrokerFeeAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerFeeAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerFeeAccount proto.InternalMessageInfo

func (m *BrokerFeeAccount) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *BrokerFeeAccount) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

type BrokerFeeSum struct {
	Broker               []byte              `protobuf:"bytes,1,opt,name=Broker,proto3" json:"Broker,omitempty"`
	Fees                 []*BrokerFeeAccount `protobuf:"bytes,2,rep,name=fees,proto3" json:"fees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BrokerFeeSum) Reset()         { *m = BrokerFeeSum{} }
func (m *BrokerFeeSum) String() string { return proto.CompactTextString(m) }
func (*BrokerFeeSum) ProtoMessage()    {}
func (*BrokerFeeSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{13}
}

func (m *BrokerFeeSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerFeeSum.Unmarshal(m, b)
}
func (m *BrokerFeeSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerFeeSum.Marshal(b, m, deterministic)
}
func (m *BrokerFeeSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerFeeSum.Merge(m, src)
}
func (m *BrokerFeeSum) XXX_Size() int {
	return xxx_messageInfo_BrokerFeeSum.Size(m)
}
func (m *BrokerFeeSum) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerFeeSum.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerFeeSum proto.InternalMessageInfo

func (m *BrokerFeeSum) GetBroker() []byte {
	if m != nil {
		return m.Broker
	}
	return nil
}

func (m *BrokerFeeSum) GetFees() []*BrokerFeeAccount {
	if m != nil {
		return m.Fees
	}
	return nil
}

type BrokerFeeSumByPeriod struct {
	BrokerFees           []*BrokerFeeSum `protobuf:"bytes,1,rep,name=brokerFees,proto3" json:"brokerFees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BrokerFeeSumByPeriod) Reset()         { *m = BrokerFeeSumByPeriod{} }
func (m *BrokerFeeSumByPeriod) String() string { return proto.CompactTextString(m) }
func (*BrokerFeeSumByPeriod) ProtoMessage()    {}
func (*BrokerFeeSumByPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{14}
}

func (m *BrokerFeeSumByPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerFeeSumByPeriod.Unmarshal(m, b)
}
func (m *BrokerFeeSumByPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerFeeSumByPeriod.Marshal(b, m, deterministic)
}
func (m *BrokerFeeSumByPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerFeeSumByPeriod.Merge(m, src)
}
func (m *BrokerFeeSumByPeriod) XXX_Size() int {
	return xxx_messageInfo_BrokerFeeSumByPeriod.Size(m)
}
func (m *BrokerFeeSumByPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerFeeSumByPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerFeeSumByPeriod proto.InternalMessageInfo

func (m *BrokerFeeSumByPeriod) GetBrokerFees() []*BrokerFeeSum {
	if m != nil {
		return m.BrokerFees
	}
	return nil
}

type VxFunds struct {
	Funds                []*VxFundWithPeriod `protobuf:"bytes,1,rep,name=funds,proto3" json:"funds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VxFunds) Reset()         { *m = VxFunds{} }
func (m *VxFunds) String() string { return proto.CompactTextString(m) }
func (*VxFunds) ProtoMessage()    {}
func (*VxFunds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{15}
}

func (m *VxFunds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VxFunds.Unmarshal(m, b)
}
func (m *VxFunds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VxFunds.Marshal(b, m, deterministic)
}
func (m *VxFunds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VxFunds.Merge(m, src)
}
func (m *VxFunds) XXX_Size() int {
	return xxx_messageInfo_VxFunds.Size(m)
}
func (m *VxFunds) XXX_DiscardUnknown() {
	xxx_messageInfo_VxFunds.DiscardUnknown(m)
}

var xxx_messageInfo_VxFunds proto.InternalMessageInfo

func (m *VxFunds) GetFunds() []*VxFundWithPeriod {
	if m != nil {
		return m.Funds
	}
	return nil
}

type VxFundWithPeriod struct {
	Period               uint64   `protobuf:"varint,1,opt,name=Period,proto3" json:"Period,omitempty"`
	Amount               []byte   `protobuf:"bytes,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VxFundWithPeriod) Reset()         { *m = VxFundWithPeriod{} }
func (m *VxFundWithPeriod) String() string { return proto.CompactTextString(m) }
func (*VxFundWithPeriod) ProtoMessage()    {}
func (*VxFundWithPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{16}
}

func (m *VxFundWithPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VxFundWithPeriod.Unmarshal(m, b)
}
func (m *VxFundWithPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VxFundWithPeriod.Marshal(b, m, deterministic)
}
func (m *VxFundWithPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VxFundWithPeriod.Merge(m, src)
}
func (m *VxFundWithPeriod) XXX_Size() int {
	return xxx_messageInfo_VxFundWithPeriod.Size(m)
}
func (m *VxFundWithPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_VxFundWithPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_VxFundWithPeriod proto.InternalMessageInfo

func (m *VxFundWithPeriod) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *VxFundWithPeriod) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

type PledgeVip struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	PledgeTimes          int32    `protobuf:"varint,2,opt,name=PledgeTimes,proto3" json:"PledgeTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PledgeVip) Reset()         { *m = PledgeVip{} }
func (m *PledgeVip) String() string { return proto.CompactTextString(m) }
func (*PledgeVip) ProtoMessage()    {}
func (*PledgeVip) Descriptor() ([]byte, []int) {
	return fileDescriptor_c247044ebb68d1ae, []int{17}
}

func (m *PledgeVip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PledgeVip.Unmarshal(m, b)
}
func (m *PledgeVip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PledgeVip.Marshal(b, m, deterministic)
}
func (m *PledgeVip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PledgeVip.Merge(m, src)
}
func (m *PledgeVip) XXX_Size() int {
	return xxx_messageInfo_PledgeVip.Size(m)
}
func (m *PledgeVip) XXX_DiscardUnknown() {
	xxx_messageInfo_PledgeVip.DiscardUnknown(m)
}

var xxx_messageInfo_PledgeVip proto.InternalMessageInfo

func (m *PledgeVip) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PledgeVip) GetPledgeTimes() int32 {
	if m != nil {
		return m.PledgeTimes
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "proto.Account")
	proto.RegisterType((*Fund)(nil), "proto.Fund")
	proto.RegisterType((*FundSettle)(nil), "proto.FundSettle")
	proto.RegisterType((*UserFundSettle)(nil), "proto.UserFundSettle")
	proto.RegisterType((*UserFeeSettle)(nil), "proto.UserFeeSettle")
	proto.RegisterType((*FeeSettle)(nil), "proto.FeeSettle")
	proto.RegisterType((*SettleActions)(nil), "proto.SettleActions")
	proto.RegisterType((*FeeSumAccount)(nil), "proto.FeeSumAccount")
	proto.RegisterType((*FeeSumByPeriod)(nil), "proto.FeeSumByPeriod")
	proto.RegisterType((*UserFees)(nil), "proto.UserFees")
	proto.RegisterType((*UserFeeAccount)(nil), "proto.UserFeeAccount")
	proto.RegisterType((*UserFeeWithPeriod)(nil), "proto.UserFeeWithPeriod")
	proto.RegisterType((*BrokerFeeAccount)(nil), "proto.BrokerFeeAccount")
	proto.RegisterType((*BrokerFeeSum)(nil), "proto.BrokerFeeSum")
	proto.RegisterType((*BrokerFeeSumByPeriod)(nil), "proto.BrokerFeeSumByPeriod")
	proto.RegisterType((*VxFunds)(nil), "proto.VxFunds")
	proto.RegisterType((*VxFundWithPeriod)(nil), "proto.VxFundWithPeriod")
	proto.RegisterType((*PledgeVip)(nil), "proto.PledgeVip")
}

func init() { proto.RegisterFile("dex_fund.proto", fileDescriptor_c247044ebb68d1ae) }

var fileDescriptor_c247044ebb68d1ae = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xce, 0xd2, 0x4f, 0x0e, 0x6d, 0x5f, 0x98, 0x17, 0x70, 0x2f, 0x0c, 0x21, 0x13, 0x63, 0x1a,
	0x45, 0x12, 0xc5, 0x44, 0x2e, 0xbc, 0xb0, 0x8d, 0x21, 0x21, 0x60, 0x82, 0x53, 0xa8, 0x77, 0x92,
	0xa5, 0x73, 0xaa, 0x1b, 0xb6, 0xbb, 0xa4, 0xb3, 0x8b, 0xf8, 0x67, 0xbc, 0xf1, 0xd6, 0xdf, 0xe0,
	0x6f, 0x33, 0xf3, 0xd5, 0x9d, 0x59, 0x3e, 0xbc, 0xf3, 0xaa, 0x3d, 0xcf, 0x79, 0xf6, 0xcc, 0xf3,
	0x9c, 0x33, 0x67, 0xa0, 0xc7, 0xf1, 0xe6, 0x7c, 0x5a, 0xa4, 0x7c, 0xf7, 0x6a, 0x9e, 0xe5, 0x19,
	0x69, 0xa8, 0x1f, 0x7a, 0x06, 0xad, 0xc1, 0x64, 0x92, 0x15, 0x69, 0x4e, 0xd6, 0xa1, 0x71, 0x9a,
	0x5d, 0x62, 0x1a, 0x06, 0xdb, 0x41, 0xbf, 0xc3, 0x74, 0x40, 0x1e, 0xc3, 0xf2, 0xe0, 0x3a, 0x8a,
	0x93, 0xe8, 0x22, 0xc1, 0x70, 0x49, 0x65, 0x4a, 0x80, 0x6c, 0x42, 0xf3, 0x38, 0x9b, 0x5c, 0x22,
	0x0f, 0x6b, 0x2a, 0x65, 0x22, 0x7a, 0x0c, 0xf5, 0x83, 0x22, 0xe5, 0x24, 0x84, 0xd6, 0x80, 0xf3,
	0x39, 0x0a, 0x61, 0xaa, 0xda, 0x90, 0x3c, 0x83, 0xb6, 0x39, 0x58, 0x84, 0x4b, 0xdb, 0xb5, 0xfe,
	0xca, 0xab, 0x9e, 0x56, 0xb6, 0x6b, 0x60, 0xb6, 0xc8, 0xd3, 0x1f, 0x01, 0x80, 0x2c, 0x37, 0xc2,
	0x3c, 0x4f, 0x90, 0x50, 0xe8, 0xc4, 0xe2, 0x74, 0x1e, 0x71, 0x2c, 0xf5, 0xb6, 0x99, 0x87, 0x49,
	0xce, 0x61, 0x3a, 0xa9, 0x2a, 0xf7, 0x30, 0xc9, 0x61, 0xc8, 0x8b, 0x09, 0x7a, 0x16, 0x3c, 0x8c,
	0x3c, 0x81, 0x2e, 0xc3, 0x04, 0x23, 0x61, 0x49, 0x75, 0x45, 0xf2, 0x41, 0x7a, 0x0e, 0xbd, 0x33,
	0x81, 0x73, 0x47, 0xe3, 0xfd, 0xc6, 0xf7, 0x60, 0x65, 0xba, 0xe0, 0x59, 0xef, 0x6b, 0xc6, 0x7b,
	0x59, 0x81, 0xb9, 0x2c, 0x1a, 0x41, 0x57, 0x1d, 0x80, 0xf8, 0xd7, 0xfa, 0x21, 0xb4, 0x86, 0x91,
	0xc0, 0x03, 0xb4, 0xa6, 0x6d, 0x28, 0x47, 0x39, 0x9c, 0x67, 0x97, 0xaa, 0x8c, 0x31, 0x5b, 0x02,
	0xf4, 0x1b, 0x2c, 0x97, 0xe5, 0xef, 0xbe, 0x0b, 0x9b, 0xd0, 0xd4, 0x7c, 0x53, 0xd9, 0x44, 0xe4,
	0x2d, 0xf4, 0x0a, 0x57, 0x9d, 0x08, 0x6b, 0xca, 0xd5, 0xba, 0x71, 0xe5, 0x49, 0x67, 0x15, 0x2e,
	0xfd, 0x1d, 0x40, 0x57, 0xff, 0x1f, 0x4c, 0xf2, 0x38, 0x4b, 0x05, 0xd9, 0x02, 0xa8, 0x8c, 0xb7,
	0xc3, 0x1c, 0x44, 0xe6, 0x3f, 0x16, 0x59, 0x6e, 0xf2, 0x5a, 0x8b, 0x83, 0x90, 0x37, 0xba, 0xc5,
	0xa6, 0x9c, 0x11, 0xb3, 0xe1, 0x8a, 0xa9, 0xb4, 0xd9, 0x1e, 0xfc, 0x1a, 0x60, 0x8a, 0x56, 0x46,
	0x58, 0x7f, 0xc0, 0x84, 0xc3, 0xa3, 0x3f, 0x03, 0xe8, 0xca, 0x4c, 0x31, 0x7b, 0x78, 0x95, 0xb6,
	0x00, 0xe4, 0x28, 0x06, 0x33, 0xc9, 0xb1, 0xb2, 0x4b, 0x84, 0xec, 0xc0, 0xda, 0x61, 0x7a, 0x1d,
	0xe7, 0x38, 0xcc, 0xd2, 0x42, 0x18, 0x9a, 0x9e, 0xd3, 0xed, 0x04, 0xd9, 0x05, 0xf2, 0x3e, 0xbe,
	0x8e, 0x39, 0xa6, 0xfc, 0x24, 0xcb, 0x12, 0x43, 0xd7, 0xd7, 0xf3, 0x8e, 0x0c, 0xfd, 0x15, 0x40,
	0x4f, 0xab, 0x1c, 0x7e, 0x3f, 0xc1, 0x79, 0x9c, 0x71, 0xd2, 0x87, 0xfa, 0x14, 0x51, 0xde, 0x20,
	0xd7, 0xa8, 0x67, 0x85, 0x29, 0x06, 0xe9, 0xc3, 0x7f, 0x49, 0x24, 0xf2, 0x71, 0x94, 0xc4, 0x5c,
	0x7f, 0xac, 0xf4, 0xd7, 0x59, 0x15, 0x96, 0x26, 0x0f, 0x10, 0xf5, 0xf9, 0x7a, 0xa5, 0xda, 0xcc,
	0x41, 0xc8, 0x53, 0xe8, 0x7d, 0x88, 0x53, 0xe4, 0xe3, 0x1b, 0xcb, 0xa9, 0x2b, 0x4e, 0x05, 0xa5,
	0xfb, 0xd0, 0x36, 0x1d, 0x17, 0x64, 0xc7, 0xd3, 0x19, 0xfa, 0x03, 0xf9, 0x14, 0xe7, 0x5f, 0xf5,
	0xd9, 0x5a, 0x2b, 0xcd, 0xcd, 0x32, 0xca, 0x01, 0xfd, 0xb3, 0x71, 0xd0, 0xcf, 0xb0, 0x76, 0x4b,
	0x10, 0x79, 0x09, 0x6d, 0x73, 0xd9, 0xad, 0xf8, 0x0d, 0x5f, 0xfc, 0xe2, 0xad, 0xb3, 0x34, 0xb9,
	0x63, 0x5e, 0x83, 0x4d, 0x44, 0xdf, 0xc1, 0xea, 0x62, 0x57, 0x1f, 0xf6, 0xb5, 0x09, 0x4d, 0xcf,
	0x93, 0x89, 0xe8, 0x08, 0x3a, 0x8b, 0x0a, 0xa3, 0x62, 0xe6, 0x6c, 0x73, 0xe0, 0x6d, 0xf3, 0x73,
	0xd3, 0x6d, 0xfd, 0x32, 0x3d, 0x32, 0x82, 0xab, 0x87, 0x9b, 0x66, 0x1f, 0xc1, 0xba, 0x5b, 0x74,
	0x71, 0xb5, 0xf6, 0x00, 0x2e, 0x2c, 0x6e, 0xbd, 0xff, 0x5f, 0x2d, 0x35, 0x2a, 0x66, 0xcc, 0xa1,
	0xd1, 0x7d, 0x68, 0x8d, 0x6f, 0xe4, 0x6e, 0x0a, 0xf2, 0x02, 0x1a, 0x72, 0x31, 0xed, 0xa7, 0x56,
	0x85, 0x4e, 0x3b, 0x23, 0xd7, 0x2c, 0x3a, 0x84, 0xd5, 0x6a, 0xca, 0xe9, 0x64, 0xe0, 0x76, 0xf2,
	0xde, 0xfe, 0x1c, 0xc1, 0xf2, 0x49, 0x82, 0xfc, 0x0b, 0x8e, 0xe3, 0x2b, 0xf9, 0x56, 0x9e, 0xc6,
	0x33, 0x14, 0x79, 0x34, 0xbb, 0x52, 0xdf, 0xd7, 0x58, 0x09, 0x90, 0x6d, 0x58, 0xd1, 0x54, 0x05,
	0xa9, 0x3a, 0x0d, 0xe6, 0x42, 0x17, 0x4d, 0xa5, 0x77, 0xef, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf1, 0x60, 0x1c, 0x36, 0x77, 0x07, 0x00, 0x00,
}