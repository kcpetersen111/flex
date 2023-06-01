package main

import "flag"

func main() {

	// go func() {
	// 	// Run the Http server
	// 	server()
	// }()
	var sAddr string
	var port int
	flag.StringVar(&sAddr, "s", "0.0.0.0", "Server Address")
	flag.IntVar(&port, "p", 8080, "Server Port")

	return

}
