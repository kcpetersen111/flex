package main

import (
	"net"
	"os"
	"bufio"
	"fmt"
	"log"
)

func main() {

	ln, err := net.Listen("tcp",":8888")
	if err != nil {
		log.Fatalf("Failed to start up the server, ",err)
	}

	

	conn, err := ln.Accept()
	if err != nil{
		fmt.Println("connection failed")
	}
	defer conn.Close()

	log.Printf("Connection accepted ")

	// fmt.Fprintf(conn, "test")

	fmt.Fprintf(conn, "What video would you like to see?  ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	movie := scanner.Text()

	file,err := os.Open(movie)
	if err != nil {
		fmt.Fprintf(conn, "Error in opening the movie, %v\n",err)
	}
	defer file.Close()
	
	bytearr := make([]byte, 4086)

	

	for {
		length, err := file.Read(bytearr)
		if err != nil {
			break
		}
		for i:=0; i<length; i++{
			fmt.Fprintf(conn, "%c",bytearr[i])
		}
	}


	
}