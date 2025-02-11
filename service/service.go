package service

type Service interface {
	GenerateQRCode() ([]byte, error)
}
