package go_ws

import (
  "sync"
  "encoding/binary"
  "github.com/jinbanglin/go-ws/bufferpool"
  "github.com/jinbanglin/helper"
  "time"
)

var PacketHeaderLength = 16

type PacketHeader struct {
  ServerType uint16  // gateway for find server
  ServerID   uint16  // gateway for find server
  PacketLen  uint16  // packet length
  MsgID      uint16  // message id,bind with business logic
  Seq        uint64 // packet sequence
}

func (p *PacketHeader) GetSeq() uint64 {
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

func PackLocalPacket(header *PacketHeader, payload []byte, seq uint64) (b []byte) {
  packetLength := uint16(len(payload)) + uint16(PacketHeaderLength)
  b = make([]byte, packetLength)
  binary.BigEndian.PutUint16(b[0:2], header.ServerType)
  binary.BigEndian.PutUint16(b[2:4], header.ServerID)
  binary.BigEndian.PutUint16(b[4:6], packetLength)
  binary.BigEndian.PutUint16(b[6:8], header.MsgID)
  binary.BigEndian.PutUint64(b[8:16], seq)
  return
}

func MakeSeq() uint64 {
  return uint64(time.Now().Unix()*100000) + helper.RandUInt64(100000, 999999)
}

func ParseRemotePacket(packet []byte) (
  header *PacketHeader, payload *bufferpool.ByteBuffer, packetLength int) {

  if len(packet) < PacketHeaderLength {
    return
  }

  header = packetHeaderGet()
  header.ServerType = binary.BigEndian.Uint16(packet[0:2])
  header.ServerID = binary.BigEndian.Uint16(packet[2:4])
  header.PacketLen = binary.BigEndian.Uint16(packet[4:6])
  header.MsgID = binary.BigEndian.Uint16(packet[6:8])
  header.Seq = binary.BigEndian.Uint64(packet[8:16])

  payload = bufferpool.Get()
  payload.Write(packet[PacketHeaderLength:int(header.PacketLen)])
  packetLength = int(header.PacketLen)

  return
}

func verifyPacket(packetLength, headerPacketLength int) bool {
  return packetLength == headerPacketLength
}
