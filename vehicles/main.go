package main

import (
	"github.com/paulnasdaq/fms-v2/vehicles/db"
	"github.com/paulnasdaq/fms-v2/vehicles/service"
	"github.com/paulnasdaq/fms-v2/vehicles/transport"
	"log"
)

func main() {
	r, err := db.NewRepository()
	if err != nil {
		panic(err)
	}

	s := service.NewVehiclesService(r)
	server := transport.NewGRPCServer(s)

	log.Fatal(server.Listen(3000))
}
