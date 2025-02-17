package main

import (
	"qrtest1.peltonium.net/service"
	remoteservice "qrtest1.peltonium.net/service/remote"
	"qrtest1.peltonium.net/types"
	"qrtest1.peltonium.net/web"
)

func main() {
	// I wanted to see how to use interfaces, so I created a generic Service
	// interface and a concrete (but fake) implementation of that interface.
	// Ultimately, there could be a local implementation that stores all the data
	// on a user's phone, or any number of remote implementations (e.g. a
	// version that makes the data accessible via GraphQL). I also specifically
	// used the service.Service type here for type-checking.
	var selectedSvc service.Service = new(remoteservice.FakeService)

	// The /generate endpoint invokes the service's GenerateQRCode method
	web.StartServer([]types.ContentProvider{
		{Name: "generate", Path: "/generate", Provide: selectedSvc.GenerateQRCode},
	})
}
