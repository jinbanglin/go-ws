// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package ws_proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import message "github.com/jinbanglin/micro/message"

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

type PingReq struct {
	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_2278b7251541c7f6, []int{0}
}
func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(dst, src)
}
func (m *PingReq) XXX_Size() int {
	return m.Size()
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

func (m *PingReq) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PongRsp struct {
	Message    *message.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId     string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Appid      string           `protobuf:"bytes,3,opt,name=appid,proto3" json:"appid,omitempty"`
	RoomId     string           `protobuf:"bytes,4,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	ServerId   string           `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerName string           `protobuf:"bytes,6,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (m *PongRsp) Reset()         { *m = PongRsp{} }
func (m *PongRsp) String() string { return proto.CompactTextString(m) }
func (*PongRsp) ProtoMessage()    {}
func (*PongRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_2278b7251541c7f6, []int{1}
}
func (m *PongRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PongRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PongRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PongRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongRsp.Merge(dst, src)
}
func (m *PongRsp) XXX_Size() int {
	return m.Size()
}
func (m *PongRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PongRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PongRsp proto.InternalMessageInfo

func (m *PongRsp) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PongRsp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PongRsp) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *PongRsp) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *PongRsp) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *PongRsp) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type RpcReq struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomId string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Seq    string `protobuf:"bytes,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Packet []byte `protobuf:"bytes,4,opt,name=packet,proto3" json:"packet,omitempty"`
}

func (m *RpcReq) Reset()         { *m = RpcReq{} }
func (m *RpcReq) String() string { return proto.CompactTextString(m) }
func (*RpcReq) ProtoMessage()    {}
func (*RpcReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_2278b7251541c7f6, []int{2}
}
func (m *RpcReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RpcReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcReq.Merge(dst, src)
}
func (m *RpcReq) XXX_Size() int {
	return m.Size()
}
func (m *RpcReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcReq.DiscardUnknown(m)
}

var xxx_messageInfo_RpcReq proto.InternalMessageInfo

func (m *RpcReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RpcReq) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *RpcReq) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *RpcReq) GetPacket() []byte {
	if m != nil {
		return m.Packet
	}
	return nil
}

type RpcRsp struct {
	Message *message.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *RpcRsp) Reset()         { *m = RpcRsp{} }
func (m *RpcRsp) String() string { return proto.CompactTextString(m) }
func (*RpcRsp) ProtoMessage()    {}
func (*RpcRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_2278b7251541c7f6, []int{3}
}
func (m *RpcRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RpcRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcRsp.Merge(dst, src)
}
func (m *RpcRsp) XXX_Size() int {
	return m.Size()
}
func (m *RpcRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RpcRsp proto.InternalMessageInfo

func (m *RpcRsp) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*PingReq)(nil), "ws_proto.PingReq")
	proto.RegisterType((*PongRsp)(nil), "ws_proto.PongRsp")
	proto.RegisterType((*RpcReq)(nil), "ws_proto.RpcReq")
	proto.RegisterType((*RpcRsp)(nil), "ws_proto.RpcRsp")
}
func (m *PingReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ping) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Ping)))
		i += copy(dAtA[i:], m.Ping)
	}
	return i, nil
}

func (m *PongRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PongRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Message != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.Message.Size()))
		n1, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.Appid) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Appid)))
		i += copy(dAtA[i:], m.Appid)
	}
	if len(m.RoomId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.RoomId)))
		i += copy(dAtA[i:], m.RoomId)
	}
	if len(m.ServerId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.ServerId)))
		i += copy(dAtA[i:], m.ServerId)
	}
	if len(m.ServerName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.ServerName)))
		i += copy(dAtA[i:], m.ServerName)
	}
	return i, nil
}

func (m *RpcReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.RoomId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.RoomId)))
		i += copy(dAtA[i:], m.RoomId)
	}
	if len(m.Seq) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Seq)))
		i += copy(dAtA[i:], m.Seq)
	}
	if len(m.Packet) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Packet)))
		i += copy(dAtA[i:], m.Packet)
	}
	return i, nil
}

func (m *RpcRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Message != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.Message.Size()))
		n2, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PingReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ping)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *PongRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.Appid)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.RoomId)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.ServerId)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.ServerName)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *RpcReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.RoomId)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.Seq)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.Packet)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *RpcRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func sovProto(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PingReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: PingReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *PongRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: PongRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PongRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &message.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Appid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Appid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *RpcReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: RpcReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seq = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Packet = append(m.Packet[:0], dAtA[iNdEx:postIndex]...)
			if m.Packet == nil {
				m.Packet = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *RpcRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: RpcRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &message.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
				return 0, ErrInvalidLengthProto
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProto
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
				next, err := skipProto(dAtA[start:])
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
	ErrInvalidLengthProto = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto.proto", fileDescriptor_proto_2278b7251541c7f6) }

var fileDescriptor_proto_2278b7251541c7f6 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0x8e, 0x7b, 0x49, 0xda, 0xd3, 0x0e, 0x95, 0xf5, 0xeb, 0x27, 0x2a, 0x22, 0xa0, 0x0c, 0x88,
	0x29, 0x81, 0xb2, 0x30, 0xb3, 0x75, 0x00, 0x55, 0x5e, 0x18, 0xab, 0x34, 0xb1, 0x8c, 0x01, 0x5f,
	0x1a, 0xa7, 0xf0, 0x1a, 0xbc, 0x0d, 0xaf, 0xc0, 0xd8, 0x91, 0x11, 0xb5, 0x2f, 0x82, 0x62, 0xa7,
	0x52, 0xcb, 0xc4, 0x92, 0x9c, 0xef, 0x72, 0xce, 0xf9, 0x6c, 0xc3, 0x40, 0x97, 0xaa, 0x52, 0x89,
	0xfd, 0xe2, 0xde, 0x9b, 0x99, 0xdb, 0x6a, 0x7c, 0xc5, 0x78, 0xf5, 0xb8, 0x5a, 0x24, 0xb9, 0x12,
	0xe9, 0x13, 0x97, 0x8b, 0x4c, 0xb2, 0x17, 0x2e, 0x53, 0xc1, 0xf3, 0x52, 0xa5, 0x82, 0x1a, 0x93,
	0x31, 0xba, 0xfb, 0xbb, 0xe6, 0xf8, 0x04, 0x82, 0x19, 0x97, 0x8c, 0xd0, 0x25, 0xc6, 0xd0, 0xd1,
	0x5c, 0xb2, 0x10, 0x9d, 0xa1, 0x8b, 0x3e, 0xb1, 0x75, 0xfc, 0x81, 0x20, 0x98, 0x29, 0xc9, 0x88,
	0xd1, 0xf8, 0x1c, 0x82, 0xa6, 0xd7, 0x5a, 0x06, 0x93, 0x61, 0x22, 0x0c, 0x4b, 0xee, 0x1c, 0x47,
	0x76, 0x22, 0x3e, 0x82, 0x60, 0x65, 0x68, 0x39, 0xe7, 0x45, 0xd8, 0xb2, 0xa3, 0xfc, 0x1a, 0x4e,
	0x0b, 0xfc, 0x0f, 0xba, 0x99, 0xd6, 0xbc, 0x08, 0xdb, 0x96, 0x76, 0xa0, 0xb6, 0x97, 0x4a, 0x89,
	0xda, 0xde, 0x71, 0xf6, 0x1a, 0x4e, 0x0b, 0x7c, 0x0c, 0x7d, 0x43, 0xcb, 0x57, 0x37, 0xa9, 0x6b,
	0xa5, 0x9e, 0x23, 0xa6, 0x05, 0x3e, 0x85, 0x41, 0x23, 0xca, 0x4c, 0xd0, 0xd0, 0xb7, 0x32, 0x38,
	0xea, 0x3e, 0x13, 0x34, 0x2e, 0xc0, 0x27, 0x3a, 0xaf, 0xcf, 0xb5, 0x97, 0x07, 0x1d, 0xe4, 0xd9,
	0xdb, 0xdc, 0x3a, 0xd8, 0x3c, 0x82, 0xb6, 0xa1, 0xcb, 0x26, 0x66, 0x5d, 0xe2, 0xff, 0xe0, 0xeb,
	0x2c, 0x7f, 0xa6, 0x95, 0xcd, 0x38, 0x24, 0x0d, 0x8a, 0x2f, 0xdd, 0x96, 0xbf, 0xdf, 0xce, 0xe4,
	0x06, 0xba, 0x0f, 0x86, 0xe8, 0x1c, 0xa7, 0x10, 0x10, 0xba, 0x5c, 0x51, 0x53, 0xe1, 0x51, 0xb2,
	0x7b, 0xc2, 0xc4, 0x65, 0x1e, 0xff, 0x62, 0x8c, 0x8e, 0xbd, 0xdb, 0xf0, 0x73, 0x13, 0xa1, 0xf5,
	0x26, 0x42, 0xdf, 0x9b, 0x08, 0xbd, 0x6f, 0x23, 0x6f, 0xbd, 0x8d, 0xbc, 0xaf, 0x6d, 0xe4, 0x2d,
	0x7c, 0xeb, 0xbc, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xce, 0x98, 0xb9, 0xf1, 0x17, 0x02, 0x00,
	0x00,
}
