package consistent

import (
	"../consistent"
	"fmt"
	"log"
)

func Instantiation() {
	servers := consistent.New()
	servers.Add("king-cakes")
	servers.Add("strawberry-habanero")
	servers.Add("cheesy-fries")
	packets := []string{"192.168.124.100/50272/81.209.179.69/80/6", "192.168.124.100/50270/81.209.179.69/80/6", "192.168.124.100/50268/81.209.179.69/80/6", "192.168.124.100/50266/81.209.179.69/80/6", "192.168.124.100/50264/81.209.179.69/80/6"}
	for _, p := range packets {
		server, err := servers.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", p, packets)
	}
}

func AddServ(servers, addserver) { //addserver is a list of server names(String)
	for i := 0; i < len(addserver); i++ {
		servers.Add(addserver[i])
	}
	packets := []string{"192.168.124.100/50272/81.209.179.69/80/6", "192.168.124.100/50270/81.209.179.69/80/6", "192.168.124.100/50268/81.209.179.69/80/6", "192.168.124.100/50266/81.209.179.69/80/6", "192.168.124.100/50264/81.209.179.69/80/6"}
	for _, p := range packets {
		server, err := servers.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", p, packets)
	}
}

func RemoveServ(servers, rmvserver) { //rmvserver is a list of server names(String)
	for i := 0; i < len(rmvserver); i++ {
		var inserver = false
		for _, a := range servers {
			if a == rmvserver[i] {
				inserver = true
			}
		}
		if inserver {
			servers.Remove(rmvserver[i])
		}
	}
	packets := []string{"192.168.124.100/50272/81.209.179.69/80/6", "192.168.124.100/50270/81.209.179.69/80/6", "192.168.124.100/50268/81.209.179.69/80/6", "192.168.124.100/50266/81.209.179.69/80/6", "192.168.124.100/50264/81.209.179.69/80/6"}
	for _, p := range packets {
		server, err := servers.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", p, packets)
	}
}
