// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/apps/nft_transfer/v1/nft_transfer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type NonFungibleTokenPacketData struct {
	// the class to which the NFT to be transferred belongs
	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	// the nft id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the address defined by NFT outside the chain
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// the nft sender
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the nft receiver
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// identify whether it is far away from the source chain
	AwayFromOrigin bool `protobuf:"varint,6,opt,name=away_from_origin,json=awayFromOrigin,proto3" json:"away_from_origin,omitempty"`
}

func (m *NonFungibleTokenPacketData) Reset()         { *m = NonFungibleTokenPacketData{} }
func (m *NonFungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*NonFungibleTokenPacketData) ProtoMessage()    {}
func (*NonFungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_323a0d65d2518331, []int{0}
}
func (m *NonFungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonFungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonFungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonFungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonFungibleTokenPacketData.Merge(m, src)
}
func (m *NonFungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *NonFungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonFungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_NonFungibleTokenPacketData proto.InternalMessageInfo

func (m *NonFungibleTokenPacketData) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetAwayFromOrigin() bool {
	if m != nil {
		return m.AwayFromOrigin
	}
	return false
}

func init() {
	proto.RegisterType((*NonFungibleTokenPacketData)(nil), "tibc.apps.nft_transfer.v1.NonFungibleTokenPacketData")
}

func init() {
	proto.RegisterFile("tibc/apps/nft_transfer/v1/nft_transfer.proto", fileDescriptor_323a0d65d2518331)
}

var fileDescriptor_323a0d65d2518331 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0x32, 0x41,
	0x14, 0x85, 0x19, 0xf8, 0x21, 0xfc, 0x53, 0x10, 0x32, 0x21, 0x66, 0xa5, 0xd8, 0x10, 0x2b, 0x0a,
	0xdd, 0x09, 0xf1, 0x09, 0x34, 0x86, 0x52, 0x0d, 0xb1, 0xd2, 0x82, 0xcc, 0xee, 0x5e, 0xd6, 0x2b,
	0xec, 0xdc, 0xcd, 0xcc, 0x2c, 0x86, 0xb7, 0xf0, 0x59, 0x7c, 0x0a, 0x4b, 0x4a, 0x4b, 0x03, 0x2f,
	0x62, 0x66, 0x30, 0x26, 0x14, 0x76, 0xf7, 0xfb, 0xce, 0xb9, 0xcd, 0xe1, 0xe7, 0x0e, 0xd3, 0x4c,
	0xaa, 0xaa, 0xb2, 0x52, 0x2f, 0xdc, 0xdc, 0x19, 0xa5, 0xed, 0x02, 0x8c, 0x5c, 0x4f, 0x8e, 0x38,
	0xa9, 0x0c, 0x39, 0x12, 0xa7, 0xbe, 0x9d, 0xf8, 0x76, 0x72, 0x94, 0xae, 0x27, 0xc3, 0x41, 0x41,
	0x05, 0x85, 0x96, 0xf4, 0xd7, 0xe1, 0xe1, 0xec, 0x9d, 0xf1, 0xe1, 0x2d, 0xe9, 0x69, 0xad, 0x0b,
	0x4c, 0x57, 0xf0, 0x40, 0x4b, 0xd0, 0xf7, 0x2a, 0x5b, 0x82, 0xbb, 0x51, 0x4e, 0x89, 0x01, 0x6f,
	0x67, 0x2b, 0x65, 0x6d, 0xc4, 0x46, 0x6c, 0xfc, 0x7f, 0x76, 0x00, 0xd1, 0xe3, 0x4d, 0xcc, 0xa3,
	0x66, 0x50, 0x4d, 0xcc, 0x45, 0x9f, 0xb7, 0x6a, 0x83, 0x51, 0x2b, 0x08, 0x7f, 0x8a, 0x13, 0xde,
	0xb1, 0xa0, 0x73, 0x30, 0xd1, 0xbf, 0x20, 0x7f, 0x48, 0x0c, 0x79, 0xd7, 0x40, 0x06, 0xb8, 0x06,
	0x13, 0xb5, 0x43, 0xf2, 0xcb, 0x62, 0xcc, 0xfb, 0xea, 0x55, 0x6d, 0xe6, 0x0b, 0x43, 0xe5, 0x9c,
	0x0c, 0x16, 0xa8, 0xa3, 0xce, 0x88, 0x8d, 0xbb, 0xb3, 0x9e, 0xf7, 0x53, 0x43, 0xe5, 0x5d, 0xb0,
	0xd7, 0x4f, 0x1f, 0xbb, 0x98, 0x6d, 0x77, 0x31, 0xfb, 0xda, 0xc5, 0xec, 0x6d, 0x1f, 0x37, 0xb6,
	0xfb, 0xb8, 0xf1, 0xb9, 0x8f, 0x1b, 0x8f, 0x57, 0x05, 0xba, 0xe7, 0x3a, 0x4d, 0x32, 0x2a, 0x65,
	0x8a, 0x4a, 0xbf, 0x20, 0x28, 0x94, 0x7e, 0x94, 0x8b, 0x82, 0x64, 0x49, 0x79, 0xbd, 0x02, 0x2b,
	0xff, 0x98, 0xd4, 0x6d, 0x2a, 0xb0, 0x69, 0x27, 0x0c, 0x73, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x1a, 0xc5, 0xf6, 0x79, 0x01, 0x00, 0x00,
}

func (m *NonFungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonFungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonFungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AwayFromOrigin {
		i--
		if m.AwayFromOrigin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Class) > 0 {
		i -= len(m.Class)
		copy(dAtA[i:], m.Class)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Class)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonFungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	if m.AwayFromOrigin {
		n += 2
	}
	return n
}

func sovNftTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftTransfer(x uint64) (n int) {
	return sovNftTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonFungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftTransfer
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
			return fmt.Errorf("proto: NonFungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonFungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwayFromOrigin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
			m.AwayFromOrigin = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftTransfer
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
func skipNftTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
				return 0, ErrInvalidLengthNftTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftTransfer = fmt.Errorf("proto: unexpected end of group")
)
