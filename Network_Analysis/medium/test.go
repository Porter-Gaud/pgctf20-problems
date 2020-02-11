package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
    "log"
	"net"
	"fmt"
	"strings"
)

type Flags struct {
	NS  bool
	CWR bool
	ECE bool
	URG bool
	ACK bool
	PSH bool
	RST bool
	SYN bool
}

func main() {
	flag := "01100110"
	flags := parseBinToFlags(flag)
	for _, x := range flags {
		createAndSendPacket(&x)
	}
}

func sanitizeTCPFields(packetData []byte, srcPort, dstPort layers.TCPPort) error {
	packet := gopacket.NewPacket(packetData, layers.LayerTypeTCP, gopacket.Default)
	tcpLayerType := packet.Layer(layers.LayerTypeTCP)
	if tcpLayerType == nil {
		return fmt.Errorf("packet has no tcp layer\n")
	}
	tcpLayer, ok := tcpLayerType.(*layers.TCP)
	if !ok {
		return fmt.Errorf("tcp layer is not tcp layer :-/")
	}

	if srcPort != tcpLayer.SrcPort || dstPort != tcpLayer.DstPort {
		return fmt.Errorf("malformed tcp layer: srcport %d dstport %d\n", tcpLayer.SrcPort, tcpLayer.DstPort)
	}

	return nil
}

func parseBinToFlags(bin string) (flags []Flags) {
	bins := strings.Split(bin, " ")
	flag := Flags{}
	for _, val := range bins {
		for i, x := range val {
			if x == 49 {
				switch i {
				case 0:
					flag.NS = true
				case 1:
					flag.CWR = true
				case 2:
					flag.ECE = true
				case 3:
					flag.URG = true
				case 4:
					flag.ACK = true
				case 5:
					flag.PSH = true
				case 6:
					flag.RST = true
				case 7:
					flag.SYN = true
				}

			} else {
				switch i {
				case 0:
					flag.NS = false
				case 1:
					flag.CWR = false
				case 2:
					flag.ECE = false
				case 3:
					flag.URG = false
				case 4:
					flag.ACK = false
				case 5:
					flag.PSH = false
				case 6:
					flag.RST = false
				case 7:
					flag.SYN = false
				}
			}
		}
		flags = append(flags, flag)
	}

	return flags
}

func createAndSendPacket(flags *Flags) {
	defer util.Run()()

	var srcIP, dstIP net.IP
	var srcIPstr string = "127.0.0.1"
	var dstIPstr string = "127.0.0.1"

	// source ip
	srcIP = net.ParseIP(srcIPstr)
	if srcIP == nil {
		log.Printf("non-ip target: %q\n", srcIPstr)
	}
	srcIP = srcIP.To4()
	if srcIP == nil {
		log.Printf("non-ipv4 target: %q\n", srcIPstr)
	}

	// destination ip
	dstIP = net.ParseIP(dstIPstr)
	if dstIP == nil {
		log.Printf("non-ip target: %q\n", dstIPstr)
	}
	dstIP = dstIP.To4()
	if dstIP == nil {
		log.Printf("non-ipv4 target: %q\n", dstIPstr)
	}

	// build tcp/ip packet
	ip := layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}
	srcport := layers.TCPPort(645)
	dstport := layers.TCPPort(22)
	tcp := layers.TCP{
		SrcPort: srcport,
		DstPort: dstport,
		Urgent:  0,
		Seq:     11050,
		Ack:     0,
		ACK:     false,
		SYN:     true,
		FIN:     false,
		RST:     true,
		URG:     false,
		ECE:     true,
		CWR:     false,
		NS:      false,
		PSH:     false,
	}

	payload := gopacket.Payload([]byte("pgctf{this_is_a_basic_payload}"))
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	tcp.SetNetworkLayerForChecksum(&ip)
	err := gopacket.SerializeLayers(buf, opts,
		&ip,
		&tcp,
		payload)
	if err != nil {
		panic(err)
	}
	packetData := buf.Bytes()
	// XXX end of packet creation

	// XXX send packet
	ipConn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		panic(err)
	}

	err = sanitizeTCPFields(packetData, srcport, dstport)
	if err != nil {
		//panic(err)
		log.Printf("malformed packet: %s\n", err)
	}

	dstIPaddr := net.IPAddr{
		IP: dstIP,
	}

	_, err = ipConn.WriteTo(packetData, &dstIPaddr)
	log.Printf("%d", packetData)
	if err != nil {
		panic(err)
	}
	log.Print("packet sent!\n")
}