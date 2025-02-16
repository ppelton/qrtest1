package main

import (
	"qrtest1.peltonium.net/service"
	remoteservice "qrtest1.peltonium.net/service/remote"
	"qrtest1.peltonium.net/types"
	"qrtest1.peltonium.net/web"
)

func main() {
	var selectedSvc service.Service = new(remoteservice.FakeService)

	web.StartServer([]types.ContentProvider{
		{Name: "generate", Path: "/generate", Provide: selectedSvc.GenerateQRCode},
	})
}
