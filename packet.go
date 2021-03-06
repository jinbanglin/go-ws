package ws

import (
  "sync"
  "encoding/binary"
  "github.com/jinbanglin/go-ws/bufferpool"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/log"
  "errors"
)

const PacketHeaderLength = 12

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
  client *Client) (b []byte) {

  header.SeqLen = uint16(len(client.getLogTraceID()))
  header.PacketLen = PacketHeaderLength +
    header.SeqLen +
    uint16(len(payload)) + uint16(client.ServerNameLen +
    client.ServerIDLen +
    client.ServerAddressLen)

  header.Seq = client.getLogTraceID()
  header.ServerNameLen = uint16(client.ServerNameLen)
  header.ServerIDLen = uint16(client.ServerIDLen)
  header.ServerAddressLen = uint16(client.ServerAddressLen)

  b = make([]byte, header.PacketLen)
  binary.BigEndian.PutUint16(b[0:2], header.PacketLen)
  binary.BigEndian.PutUint16(b[2:4], header.MsgID)
  binary.BigEndian.PutUint16(b[4:6], header.ServerNameLen)
  binary.BigEndian.PutUint16(b[6:8], header.ServerIDLen)
  binary.BigEndian.PutUint16(b[8:10], header.ServerAddressLen)
  binary.BigEndian.PutUint16(b[10:12], header.SeqLen)

  length := PacketHeaderLength
  copy(b[length:length+client.ServerNameLen], helper.String2Byte(client.ServerName))

  length += client.ServerNameLen
  copy(b[length:length+client.ServerIDLen], helper.String2Byte(client.ServerID))

  length += client.ServerIDLen
  copy(b[length:length+client.ServerAddressLen], helper.String2Byte(client.ServerAddress))

  length += client.ServerAddressLen
  copy(b[length:length+int(header.SeqLen)], helper.String2Byte(client.getLogTraceID()))

  length += int(header.SeqLen)
  copy(b[length:header.PacketLen], payload)
  return
}

func ParseRemotePacket(packet []byte) (
  header *PacketHeader, payload *bufferpool.ByteBuffer, err error) {

  header = packetPool.Get().(*PacketHeader)
  header.PacketLen = binary.BigEndian.Uint16(packet[0:2])
  header.MsgID = binary.BigEndian.Uint16(packet[2:4])
  header.ServerNameLen = binary.BigEndian.Uint16(packet[4:6])
  header.ServerIDLen = binary.BigEndian.Uint16(packet[6:8])
  header.ServerAddressLen = binary.BigEndian.Uint16(packet[8:10])
  header.SeqLen = binary.BigEndian.Uint16(packet[10:12])

  if len(packet) < PacketHeaderLength +
    int(header.ServerNameLen +
      header.ServerIDLen +
      header.ServerAddressLen +
      header.SeqLen) {

    log.Error("invalid packet")
    return header, nil, errors.New("invalid packet")
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

  return
}

func verifyPacket(packetLength, headerPacketLength int) bool {
  return packetLength == headerPacketLength
}
