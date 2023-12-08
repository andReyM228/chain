// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: one/core/daily_stats.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type DailyStats struct {
	AmountIssue    []types.Coin `protobuf:"bytes,1,rep,name=amountIssue,proto3" json:"amountIssue"`
	AmountWithdraw []types.Coin `protobuf:"bytes,2,rep,name=amountWithdraw,proto3" json:"amountWithdraw"`
	CountIssue     int32        `protobuf:"varint,3,opt,name=countIssue,proto3" json:"countIssue,omitempty"`
	CountWithdraw  int32        `protobuf:"varint,4,opt,name=countWithdraw,proto3" json:"countWithdraw,omitempty"`
}

func (m *DailyStats) Reset()         { *m = DailyStats{} }
func (m *DailyStats) String() string { return proto.CompactTextString(m) }
func (*DailyStats) ProtoMessage()    {}
func (*DailyStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b18605964cf4b27, []int{0}
}
func (m *DailyStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DailyStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DailyStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DailyStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStats.Merge(m, src)
}
func (m *DailyStats) XXX_Size() int {
	return m.Size()
}
func (m *DailyStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStats.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStats proto.InternalMessageInfo

func (m *DailyStats) GetAmountIssue() []types.Coin {
	if m != nil {
		return m.AmountIssue
	}
	return nil
}

func (m *DailyStats) GetAmountWithdraw() []types.Coin {
	if m != nil {
		return m.AmountWithdraw
	}
	return nil
}

func (m *DailyStats) GetCountIssue() int32 {
	if m != nil {
		return m.CountIssue
	}
	return 0
}

func (m *DailyStats) GetCountWithdraw() int32 {
	if m != nil {
		return m.CountWithdraw
	}
	return 0
}

func init() {
	proto.RegisterType((*DailyStats)(nil), "one.core.DailyStats")
}

func init() { proto.RegisterFile("one/core/daily_stats.proto", fileDescriptor_4b18605964cf4b27) }

var fileDescriptor_4b18605964cf4b27 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xcf, 0x4b, 0xd5,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x49, 0xcc, 0xcc, 0xa9, 0x8c, 0x2f, 0x2e, 0x49, 0x2c, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcf, 0x4b, 0xd5, 0x03, 0xc9, 0x49, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x05, 0xf5, 0x41, 0x2c, 0x88, 0xbc, 0x94, 0x5c, 0x72, 0x7e, 0x71,
	0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x99, 0x61, 0x52, 0x6a, 0x49, 0xa2, 0xa1,
	0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x5e, 0xe9, 0x1e, 0x23, 0x17, 0x97, 0x0b, 0xc8, 0xd4, 0x60,
	0x90, 0xa1, 0x42, 0x8e, 0x5c, 0xdc, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x9e, 0xc5, 0xc5, 0xa5,
	0xa9, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x10, 0x43, 0xf4, 0x40, 0x86, 0xe8,
	0x41, 0x0d, 0xd1, 0x73, 0xce, 0xcf, 0xcc, 0x73, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0x59,
	0x8f, 0x90, 0x3b, 0x17, 0x1f, 0x84, 0x1b, 0x9e, 0x59, 0x92, 0x91, 0x52, 0x94, 0x58, 0x2e, 0xc1,
	0x44, 0x9c, 0x29, 0x68, 0xda, 0x84, 0xe4, 0xb8, 0xb8, 0x92, 0x11, 0x4e, 0x61, 0x56, 0x60, 0xd4,
	0x60, 0x0d, 0x42, 0x12, 0x11, 0x52, 0xe1, 0xe2, 0x4d, 0x46, 0xb1, 0x87, 0x05, 0xac, 0x04, 0x55,
	0xd0, 0x49, 0xeb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0x40, 0xc1,
	0x5a, 0x01, 0x09, 0xd8, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x98, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x33, 0x4b, 0x9a, 0x5f, 0x71, 0x01, 0x00, 0x00,
}

func (m *DailyStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DailyStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DailyStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CountWithdraw != 0 {
		i = encodeVarintDailyStats(dAtA, i, uint64(m.CountWithdraw))
		i--
		dAtA[i] = 0x20
	}
	if m.CountIssue != 0 {
		i = encodeVarintDailyStats(dAtA, i, uint64(m.CountIssue))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AmountWithdraw) > 0 {
		for iNdEx := len(m.AmountWithdraw) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AmountWithdraw[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDailyStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AmountIssue) > 0 {
		for iNdEx := len(m.AmountIssue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AmountIssue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDailyStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDailyStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovDailyStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DailyStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AmountIssue) > 0 {
		for _, e := range m.AmountIssue {
			l = e.Size()
			n += 1 + l + sovDailyStats(uint64(l))
		}
	}
	if len(m.AmountWithdraw) > 0 {
		for _, e := range m.AmountWithdraw {
			l = e.Size()
			n += 1 + l + sovDailyStats(uint64(l))
		}
	}
	if m.CountIssue != 0 {
		n += 1 + sovDailyStats(uint64(m.CountIssue))
	}
	if m.CountWithdraw != 0 {
		n += 1 + sovDailyStats(uint64(m.CountWithdraw))
	}
	return n
}

func sovDailyStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDailyStats(x uint64) (n int) {
	return sovDailyStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DailyStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDailyStats
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
			return fmt.Errorf("proto: DailyStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DailyStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIssue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDailyStats
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
				return ErrInvalidLengthDailyStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDailyStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountIssue = append(m.AmountIssue, types.Coin{})
			if err := m.AmountIssue[len(m.AmountIssue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountWithdraw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDailyStats
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
				return ErrInvalidLengthDailyStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDailyStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountWithdraw = append(m.AmountWithdraw, types.Coin{})
			if err := m.AmountWithdraw[len(m.AmountWithdraw)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountIssue", wireType)
			}
			m.CountIssue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDailyStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountIssue |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountWithdraw", wireType)
			}
			m.CountWithdraw = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDailyStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountWithdraw |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDailyStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDailyStats
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
func skipDailyStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDailyStats
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
					return 0, ErrIntOverflowDailyStats
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
					return 0, ErrIntOverflowDailyStats
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
				return 0, ErrInvalidLengthDailyStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDailyStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDailyStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDailyStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDailyStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDailyStats = fmt.Errorf("proto: unexpected end of group")
)
