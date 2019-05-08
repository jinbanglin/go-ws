package go_ws

import (
  "sync"
  "encoding/binary"
  "github.com/jinbanglin/go-ws/bufferpool"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/log"
)

const PacketHeaderLength = 10

type PacketHeader struct {
  PacketLen        uint16 // packet length
  MsgID            uint16 // message id,bind with business logic
  ServerNameLen    uint16
  ServerIDLen      uint16
  ServerAddressLen uint16
  SeqLen           uint16
  ServerName       string // gateway for find server
  ServerID         string // gateway for find server
  ServerAddress    string
  Seq              string // packet sequence
}

var packetPool = &sync.Pool{
  New: func() interface{} {
    return &PacketHeader{}
  },
}

func packetHeaderRelease(header *PacketHeader) {
  header.PacketLen = 0
  header.MsgID = 0
  header.ServerNameLen = 0
  header.ServerIDLen = 0
  header.SeqLen = 0
  header.ServerName = ""
  header.ServerID = ""
  header.Seq = ""
  header.ServerAddress = ""

  packetPool.Put(header)
}

func PackLocalPacket(
  header *PacketHeader,
  payload []byte,
  serverNameLength, serverIDLen, serverAddressLen int,
  serverName, serverID, serverAddress, seq string) (b []byte) {

  header.PacketLen =
    PacketHeaderLength +
      uint16(len(payload)) +
      uint16(len(seq)) +
      uint16(serverNameLength+serverIDLen+serverAddressLen)

  header.Seq = seq
  header.ServerNameLen = uint16(serverNameLength)
  header.ServerIDLen = uint16(serverIDLen)
  header.SeqLen = uint16(len(seq))

  b = make([]byte, header.PacketLen)

  binary.BigEndian.PutUint16(b[0:2], header.PacketLen)
  binary.BigEndian.PutUint16(b[2:4], header.MsgID)
  binary.BigEndian.PutUint16(b[4:6], header.ServerNameLen)
  binary.BigEndian.PutUint16(b[6:8], header.ServerIDLen)
  binary.BigEndian.PutUint16(b[8:10], header.ServerAddressLen)
  binary.BigEndian.PutUint16(b[10:12], header.SeqLen)

  length := PacketHeaderLength
  copy(b[length:length+serverNameLength], helper.String2Byte(serverName))

  length += serverNameLength
  copy(b[length:length+serverIDLen], helper.String2Byte(serverID))

  length += serverIDLen
  copy(b[length:length+serverAddressLen], helper.String2Byte(serverAddress))

  length += serverAddressLen
  copy(b[length:length+len(seq)], helper.String2Byte(seq))

  length += len(seq)
  copy(b[length:header.PacketLen], payload)

  return
}

func ParseRemotePacket(packet []byte) (
  header *PacketHeader, payload *bufferpool.ByteBuffer, packetLength int) {

  header = packetPool.Get().(*PacketHeader)
  header.PacketLen = binary.BigEndian.Uint16(packet[0:2])
  header.MsgID = binary.BigEndian.Uint16(packet[2:4])
  header.ServerNameLen = binary.BigEndian.Uint16(packet[4:6])
  header.ServerIDLen = binary.BigEndian.Uint16(packet[6:8])
  header.ServerAddressLen = binary.BigEndian.Uint16(packet[8:10])
  header.SeqLen = binary.BigEndian.Uint16(packet[10:12])

  log.Debug("--------",len(packet), PacketHeaderLength +
    int(header.ServerNameLen +
      header.ServerIDLen +
      header.ServerAddressLen +
      header.SeqLen))

  if len(packet) < PacketHeaderLength +
    int(header.ServerNameLen +
      header.ServerIDLen +
      header.ServerAddressLen +
      header.SeqLen) {

    log.Error("invalid packet")
    return
  }

  length := PacketHeaderLength
  header.ServerName = helper.Byte2String(packet[length : length+int(header.ServerNameLen)])

  length += int(header.ServerNameLen)
  header.ServerID = helper.Byte2String(packet[length : length+int(header.ServerIDLen)])

  length += int(header.ServerIDLen)
  header.ServerAddress = helper.Byte2String(packet[length : length+int(header.ServerAddressLen)])

  length += int(header.ServerAddressLen)
  header.Seq = helper.Byte2String(packet[length : length+int(header.SeqLen)])

  length += int(header.SeqLen)
  payload = bufferpool.Get()
  payload.Write(packet[length:int(header.PacketLen)])

  packetLength = int(header.PacketLen)

  return
}

func verifyPacket(packetLength, headerPacketLength int) bool {
  return packetLength == headerPacketLength
}
