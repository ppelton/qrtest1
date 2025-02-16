package service

import (
	"qrtest1.peltonium.net/api"
)

type Service interface {
	GenerateQRCode() ([]byte, error)
	GetItems() ([]api.Item, error)
	SaveItems([]api.Item) error
}
