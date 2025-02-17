package remoteservice

import (
	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"

	"log"
	"net"

	"qrtest1.peltonium.net/api"
)

type FakeService struct{}

// Generates a QR code representing the local IP address and listening port
func (s FakeService) GenerateQRCode() ([]byte, error) {
	return qrcode.Encode("http://"+getOutboundIP().String()+":8080", qrcode.Medium, 128)
}

// Returns a fake list of items
func (s FakeService) GetItems() ([]api.Item, error) {
	var items = []api.Item{
		{Id: uuid.NewString(), Name: "test", Description: "desc"},
	}

	return items, nil
}

func (s FakeService) SaveItems([]api.Item) error {
	return nil
}

// Private functions

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
