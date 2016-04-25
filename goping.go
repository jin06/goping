package goping

import "net"

type ICMPMessage interface {
	Marshal()(b []byte)
	Unmarshal(b []byte)
}

type ICMPHeader struct {
	Type int
	Code int
	CheckSum string
}

func (header *ICMPHeader) SetCheckSum() (err error ) {
	return
}

type PingMessage struct {
	ICMPHeader
	Identifier int
	SequenceNumber int
}

func (msg *PingMessage) Marshal()(b []byte){
	b = []byte{}
	b = append(b,byte(msg.Type),byte(msg.Code))
	return
}

func (msg *PingMessage) Unmarshal(b []byte) {
	return
}

func newPingMessage()(msg *ICMPMessage){
	msg = &PingMessage{
		//ICMPHeader.Type:8,
		//ICMPHeader.Code:0,
		Identifier:10779,
		SequenceNumber:1,
	}
	return
}

func main() {
	//c, err := net.Dial("ip4:icmp", "127.0.0.1")
	//if err != nil {
	//	return err
	//}
	//defer c.Close()
	//c.Write()
}
