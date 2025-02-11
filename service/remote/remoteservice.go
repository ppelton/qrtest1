package remoteservice

import (
	qrcode "github.com/skip2/go-qrcode"

	"log"
	"net"
)

type RemoteService struct{}

func (l RemoteService) GenerateQRCode() ([]byte, error) {
	return qrcode.Encode("http://"+getOutboundIP().String()+":8080", qrcode.Medium, 128)
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
