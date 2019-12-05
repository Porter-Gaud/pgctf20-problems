package main

import (
	//"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"log"
	"net"
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
	flag := "01100110 01101100 01100001 01100111 01111011 01110100 01100101 01110011 01110100 01111101"
	flags := parseBinToFlags(flag)
	for _, x := range flags {
		createAndSendPacket(&x)
	}
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

	srcport := layers.TCPPort(666)
	dstport := layers.TCPPort(1234)
	tcp := layers.TCP{
		SrcPort: srcport,
		DstPort: dstport,
		Window:  1505,
		Urgent:  0,
		Seq:     11050,
		Ack:     0,
		ACK:     flags.ACK,
		SYN:     flags.SYN,
		FIN:     false,
		RST:     flags.RST,
		URG:     flags.URG,
		ECE:     flags.ECE,
		CWR:     flags.CWR,
		NS:      flags.NS,
		PSH:     flags.PSH,
	}

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	tcp.SetNetworkLayerForChecksum(&ip)

	ipHeaderBuf := gopacket.NewSerializeBuffer()
	err := ip.SerializeTo(ipHeaderBuf, opts)
	if err != nil {
		panic(err)
	}
	ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
	if err != nil {
		panic(err)
	}

	tcpPayloadBuf := gopacket.NewSerializeBuffer()
	payload := gopacket.Payload([]byte("garbage"))
	err = gopacket.SerializeLayers(tcpPayloadBuf, opts, &tcp, payload)
	if err != nil {
		panic(err)
	}
	// XXX end of packet creation

	// XXX send packet
	var packetConn net.PacketConn
	var rawConn *ipv4.RawConn
	packetConn, err = net.ListenPacket("ip4:tcp", "127.0.0.1")
	if err != nil {
		panic(err)
	}
	rawConn, err = ipv4.NewRawConn(packetConn)
	if err != nil {
		panic(err)
	}

	err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil)
	log.Printf("packet of length %d sent!\n", (len(tcpPayloadBuf.Bytes()) + len(ipHeaderBuf.Bytes())))

}
