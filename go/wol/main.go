package main

import (
	"log"
	"net"

	"github.com/mdlayher/wol"
)

const (
	mac = "00:00:5e:00:53:01"
)

func main() {
	mac, err := net.ParseMAC(mac)
	if err != nil {
		log.Fatal("Get MAC failed: ", err)
	}

	wake, err := wol.NewClient()
	if err != nil {
		log.Fatal("Create New woi client failed: ", err)
	}
	defer wake.Close()

	wake.Wake("192.168.31.51/24", mac)
}
