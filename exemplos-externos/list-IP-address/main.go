package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	fmt.Println("getting local IP Address...")
	for _, i := range ifaces {

		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			fmt.Println(addr.Network())
			fmt.Println(addr.String())
		}
	}
}
