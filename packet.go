package go_ws

import (
  "sync"
  "encoding/binary"
  "github.com/jinbanglin/go-ws/bufferpool"
)

var PacketHeaderLength = 12

type PacketType = uint16

type PacketHeader struct {
  ServerType PacketType // gateway for find server
  ServerID   PacketType // gateway for find server
  PacketLen  PacketType // packet length
  MsgID      PacketType // message id,bind with business logic
  Seq        PacketType // packet sequence
}

func (p *PacketHeader) GetSeq() uint16 {
  return p.Seq
}

func (p *PacketHeader) GetMsgID() uint16 {
  return p.MsgID
}

func (p *PacketHeader) GetServerType() uint16 {
  return p.ServerType
}

func (p *PacketHeader) GetServerID() uint16 {
  return p.ServerID
}

func (p *PacketHeader) GetPacketLength() uint16 {
  return p.PacketLen
}

var packetPool *sync.Pool

func packetHeaderGet() *PacketHeader {
  v := packetPool.Get()
  if v != nil {
    return v.(*PacketHeader)
  }
  return new(PacketHeader)
}

func packetHeaderRelease(header *PacketHeader) {
  header.MsgID = 0
  header.PacketLen = 0
  header.Seq = 0
  header.ServerID = 0
  header.ServerType = 0
  packetPool.Put(header)
}

func pack(header *PacketHeader, payload []byte) (b []byte) {
  packetLength := uint16(len(payload)) + uint16(PacketHeaderLength)
  b = make([]byte, packetLength)
  binary.BigEndian.PutUint16(b[0:2], header.ServerType)
  binary.BigEndian.PutUint16(b[2:4], header.ServerID)
  binary.BigEndian.PutUint16(b[4:6], packetLength)
  binary.BigEndian.PutUint16(b[6:8], header.MsgID)
  binary.BigEndian.PutUint16(b[8:12], header.Seq)
  return
}

func parsePacket(packet []byte) (
  header *PacketHeader, payload *bufferpool.ByteBuffer, packetLength int) {

  if len(packet) < PacketHeaderLength {
    return
  }

  header = packetHeaderGet()
  header.ServerType = binary.BigEndian.Uint16(packet[0:2])
  header.ServerID = binary.BigEndian.Uint16(packet[2:4])
  header.PacketLen = binary.BigEndian.Uint16(packet[4:6])
  header.MsgID = binary.BigEndian.Uint16(packet[6:8])
  header.Seq = binary.BigEndian.Uint16(packet[8:12])

  payload = bufferpool.Get()
  payload.Write(packet[PacketHeaderLength:header.PacketLen])
  packetLength = int(header.PacketLen)

  return
}

func verifyPacket(packetLength, headerPacketLength int) bool {
  return packetLength == headerPacketLength
}
