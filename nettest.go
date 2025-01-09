package main

import (
	"fmt"
	"net"
	"time"
	// "regexp"
	"strconv"

)

var Blue = "\033[34m"
var Reset = "\033[0m"
var Yellow = "\033[33m"

func title() {
	asciiArt := `
                          _    _            _             
   __ _  ___   _ __   ___| |_ | |_ ___  ___| |_ ___ _ __  
  / _' |/ _ \ | '_ \ / _ \ __|| __/ _ \/ __| __/ _ \ '__| 
 | (_| | (_) || | | |  __/ |_ | ||  __/\__ \ ||  __/ |    
  \__, |\___(_)_| |_|\___|\__(_)__\___||___/\__\___|_|    
  |___/                                                   
`
	fmt.Println(Yellow + asciiArt + Reset)
}


func main() {

	title()
	// re := regexp.MustCompile(`:.*$`)
	var addr string
	// var ipaddr string

	fmt.Println("Please give address and port e.g. 8.8.8.8:53 : ")
	fmt.Scanf("%s", &addr)

	for i := 50; i < 55; i++ {

		var ipaddr = addr + ":" + strconv.Itoa(i)
		fmt.Println(ipaddr)
		timeout := 20 * time.Second
		conn, err := net.DialTimeout("tcp", ipaddr, timeout)
		fmt.Println(conn)
		if err != nil {
			fmt.Println(strconv.Itoa(i) + " Is Closed.")
			// fmt.Println("Connection Failed :(")
			// fmt.Println(err)
		} else {
			fmt.Println(strconv.Itoa(i) + " Is Open.")
		}

		// defer conn.Close()

	// addrfixed := re.ReplaceAllString(addr, "")


	//fmt.Println(conn.RemoteAddr())
	//fmt.Println(conn.LocalAddr())
	// fmt.Println(net.LookupAddr(addrfixed))
	//fmt.Println()

	}
}

