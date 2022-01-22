// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/Silent.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SilentGetRoomSilentReq struct {
	//
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id"`
}

func (m *SilentGetRoomSilentReq) Reset()         { *m = SilentGetRoomSilentReq{} }
func (m *SilentGetRoomSilentReq) String() string { return proto.CompactTextString(m) }
func (*SilentGetRoomSilentReq) ProtoMessage()    {}
func (*SilentGetRoomSilentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_Silent_de6616e052635452, []int{0}
}
func (m *SilentGetRoomSilentReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SilentGetRoomSilentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SilentGetRoomSilentReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SilentGetRoomSilentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SilentGetRoomSilentReq.Merge(dst, src)
}
func (m *SilentGetRoomSilentReq) XXX_Size() int {
	return m.Size()
}
func (m *SilentGetRoomSilentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SilentGetRoomSilentReq.DiscardUnknown(m)
}

var xxx_messageInfo_SilentGetRoomSilentReq proto.InternalMessageInfo

func (m *SilentGetRoomSilentReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type SilentGetRoomSilentResp struct {
	//
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	//
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data *SilentGetRoomSilentResp_Data `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *SilentGetRoomSilentResp) Reset()         { *m = SilentGetRoomSilentResp{} }
func (m *SilentGetRoomSilentResp) String() string { return proto.CompactTextString(m) }
func (*SilentGetRoomSilentResp) ProtoMessage()    {}
func (*SilentGetRoomSilentResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_Silent_de6616e052635452, []int{1}
}
func (m *SilentGetRoomSilentResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SilentGetRoomSilentResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SilentGetRoomSilentResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SilentGetRoomSilentResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SilentGetRoomSilentResp.Merge(dst, src)
}
func (m *SilentGetRoomSilentResp) XXX_Size() int {
	return m.Size()
}
func (m *SilentGetRoomSilentResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SilentGetRoomSilentResp.DiscardUnknown(m)
}

var xxx_messageInfo_SilentGetRoomSilentResp proto.InternalMessageInfo

func (m *SilentGetRoomSilentResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SilentGetRoomSilentResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SilentGetRoomSilentResp) GetData() *SilentGetRoomSilentResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type SilentGetRoomSilentResp_Data struct {
	// level等级 medal粉丝勋章等级 member全部
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// 等级type=level时为用户等级要求 type为medal为粉丝勋章等级要求 type为member时1表示开启，0表示未开启
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level"`
	// 还剩多少秒时间到期, -1表示直到直播结束都不过期
	Second int64 `protobuf:"varint,3,opt,name=second,proto3" json:"second"`
}

func (m *SilentGetRoomSilentResp_Data) Reset()         { *m = SilentGetRoomSilentResp_Data{} }
func (m *SilentGetRoomSilentResp_Data) String() string { return proto.CompactTextString(m) }
func (*SilentGetRoomSilentResp_Data) ProtoMessage()    {}
func (*SilentGetRoomSilentResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_Silent_de6616e052635452, []int{1, 0}
}
func (m *SilentGetRoomSilentResp_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SilentGetRoomSilentResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SilentGetRoomSilentResp_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SilentGetRoomSilentResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SilentGetRoomSilentResp_Data.Merge(dst, src)
}
func (m *SilentGetRoomSilentResp_Data) XXX_Size() int {
	return m.Size()
}
func (m *SilentGetRoomSilentResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_SilentGetRoomSilentResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_SilentGetRoomSilentResp_Data proto.InternalMessageInfo

func (m *SilentGetRoomSilentResp_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SilentGetRoomSilentResp_Data) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SilentGetRoomSilentResp_Data) GetSecond() int64 {
	if m != nil {
		return m.Second
	}
	return 0
}

func init() {
	proto.RegisterType((*SilentGetRoomSilentReq)(nil), "banned_service.v1.SilentGetRoomSilentReq")
	proto.RegisterType((*SilentGetRoomSilentResp)(nil), "banned_service.v1.SilentGetRoomSilentResp")
	proto.RegisterType((*SilentGetRoomSilentResp_Data)(nil), "banned_service.v1.SilentGetRoomSilentResp.Data")
}
func (m *SilentGetRoomSilentReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SilentGetRoomSilentReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RoomId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSilent(dAtA, i, uint64(m.RoomId))
	}
	return i, nil
}

func (m *SilentGetRoomSilentResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SilentGetRoomSilentResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSilent(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSilent(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSilent(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SilentGetRoomSilentResp_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SilentGetRoomSilentResp_Data) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilent(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Level != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSilent(dAtA, i, uint64(m.Level))
	}
	if m.Second != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSilent(dAtA, i, uint64(m.Second))
	}
	return i, nil
}

func encodeVarintSilent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SilentGetRoomSilentReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovSilent(uint64(m.RoomId))
	}
	return n
}

func (m *SilentGetRoomSilentResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovSilent(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSilent(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovSilent(uint64(l))
	}
	return n
}

func (m *SilentGetRoomSilentResp_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSilent(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovSilent(uint64(m.Level))
	}
	if m.Second != 0 {
		n += 1 + sovSilent(uint64(m.Second))
	}
	return n
}

func sovSilent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSilent(x uint64) (n int) {
	return sovSilent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SilentGetRoomSilentReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SilentGetRoomSilentReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SilentGetRoomSilentReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSilent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilent
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
func (m *SilentGetRoomSilentResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SilentGetRoomSilentResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SilentGetRoomSilentResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &SilentGetRoomSilentResp_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilent
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
func (m *SilentGetRoomSilentResp_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Second", wireType)
			}
			m.Second = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Second |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSilent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilent
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
func skipSilent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSilent
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
					return 0, ErrIntOverflowSilent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSilent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSilent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSilent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSilent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSilent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSilent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/Silent.proto", fileDescriptor_Silent_de6616e052635452) }

var fileDescriptor_Silent_de6616e052635452 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xe6, 0x28, 0x16, 0x79, 0x0c, 0xc4, 0x1b, 0x14, 0x09, 0x69, 0x09, 0x71, 0x40, 0x13, 0xdb,
	0x80, 0xbb, 0x43, 0x63, 0x62, 0x1c, 0x5c, 0xce, 0xcd, 0x85, 0xb4, 0xbd, 0xb3, 0x34, 0xa1, 0x3d,
	0x68, 0x8f, 0x26, 0xfe, 0x07, 0x07, 0x7f, 0x96, 0x23, 0xa3, 0x53, 0x63, 0x60, 0xeb, 0xaf, 0x30,
	0x77, 0x57, 0x27, 0x35, 0x61, 0xf9, 0xde, 0xfb, 0xbe, 0xbc, 0x7e, 0xef, 0x7b, 0x3d, 0xe8, 0x15,
	0x53, 0xf7, 0x29, 0x5e, 0xb2, 0x54, 0x38, 0xab, 0x8c, 0x0b, 0x8e, 0x4f, 0x02, 0x3f, 0x4d, 0x19,
	0x9d, 0xe7, 0x2c, 0x2b, 0xe2, 0x90, 0x39, 0xc5, 0x74, 0x70, 0x1d, 0xc5, 0x62, 0xb1, 0x09, 0x9c,
	0x90, 0x27, 0x6e, 0xc4, 0x23, 0xee, 0xaa, 0xc9, 0x60, 0xf3, 0xa2, 0x98, 0x22, 0xaa, 0xd3, 0x0e,
	0xe3, 0x5b, 0x38, 0xd5, 0x8e, 0xf7, 0x4c, 0x10, 0xce, 0x13, 0x4d, 0x08, 0x5b, 0xe3, 0x0b, 0x68,
	0x67, 0x9c, 0x27, 0xf3, 0x98, 0xf6, 0xd1, 0x08, 0x4d, 0x0c, 0xaf, 0x5b, 0x95, 0xf6, 0x8f, 0x44,
	0x4c, 0xd9, 0x3c, 0xd0, 0xf1, 0x5b, 0x13, 0xce, 0xfe, 0x34, 0xc8, 0x57, 0x78, 0x08, 0xad, 0x90,
	0x53, 0x56, 0x7f, 0x7e, 0x5c, 0x95, 0xb6, 0xe2, 0x44, 0x21, 0x3e, 0x07, 0x23, 0xc9, 0xa3, 0x7e,
	0x73, 0x84, 0x26, 0x1d, 0xaf, 0x5d, 0x95, 0xb6, 0xa4, 0x44, 0x02, 0x7e, 0x84, 0x16, 0xf5, 0x85,
	0xdf, 0x37, 0x46, 0x68, 0xd2, 0x9d, 0xb9, 0xce, 0xaf, 0x2b, 0x9d, 0x7f, 0x56, 0x3a, 0x77, 0xbe,
	0xf0, 0xf5, 0x26, 0x69, 0x40, 0x14, 0x0e, 0x62, 0x68, 0x49, 0x5d, 0xe6, 0x11, 0xaf, 0x2b, 0x9d,
	0xa7, 0xa3, 0xa7, 0x24, 0x27, 0x0a, 0xb1, 0x0d, 0x47, 0x4b, 0x56, 0xb0, 0xa5, 0x4a, 0x64, 0x78,
	0x9d, 0xaa, 0xb4, 0xb5, 0x40, 0x74, 0xc1, 0x63, 0x30, 0x73, 0x16, 0xf2, 0x94, 0xaa, 0x5c, 0x86,
	0x07, 0x55, 0x69, 0xd7, 0x0a, 0xa9, 0xeb, 0x2c, 0x03, 0x53, 0xa7, 0xc1, 0x0b, 0xe8, 0x45, 0x4c,
	0xcc, 0xd5, 0xff, 0xca, 0xb5, 0x74, 0x79, 0xe8, 0x21, 0xeb, 0xc1, 0xd5, 0xe1, 0x37, 0x7b, 0xc3,
	0x8f, 0x9d, 0x85, 0xb6, 0x3b, 0x0b, 0x7d, 0xed, 0x2c, 0xf4, 0xbe, 0xb7, 0x1a, 0xdb, 0xbd, 0xd5,
	0xf8, 0xdc, 0x5b, 0x8d, 0xe7, 0x66, 0x31, 0x0d, 0x4c, 0xf5, 0xce, 0x37, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xab, 0xb0, 0xfb, 0x98, 0x3c, 0x02, 0x00, 0x00,
}
