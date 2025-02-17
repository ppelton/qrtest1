package service

import (
	"qrtest1.peltonium.net/api"
)

// Generic service interface
type Service interface {
	GenerateQRCode() ([]byte, error)
	GetItems() ([]api.Item, error)
	SaveItems([]api.Item) error
}
