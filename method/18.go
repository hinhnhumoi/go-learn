package main

import "fmt"

type IpAddr [4]byte

func (addr IpAddr) String() string {
	return fmt.Sprintf("address is %v %v %v %v", addr[0], addr[1], addr[2], addr[3])
}

func func18() {
	fmt.Println("func18: ")
	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}