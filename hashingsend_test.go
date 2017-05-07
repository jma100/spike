package main

import (
	"fmt"
	"log"
	"stathat.com/c/consistent"
)

func main() {
	servers := createNew()
	packets := []string{"192.168.124.100/50272/81.209.179.69/80/6", "192.168.124.100/50270/81.209.179.69/80/6", "192.168.124.100/50268/81.209.179.69/80/6", "192.168.124.100/50266/81.209.179.69/80/6", "192.168.124.100/50264/81.209.179.69/80/6"}
	for _, p := range packets {
		server, err := servers.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", p, server)
	
}
	//AddServ()
	//RemoveServ()

}

func createNew() {
	servers := consistent.New()
	servers.Add("king-cakes")
	servers.Add("strawberry-habanero")
	servers.Add("cheesy-fries")
	return servers
}

func AddServ(servers consistent, []addserver)([]updatedservers) { //addserver is a list of server names(String)
	for i := 0; i < len(addserver); i++ {
		servers.Add(addserver[i])
	}
	return servers
}

func RemoveServ(servers consistent, []rmvserver)([]updatedservers) { //rmvserver is a list of server names(String)
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
	return servers
}
