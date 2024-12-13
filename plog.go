package main

import (
	"fmt"
	"time"
	"net"
	"log"
	"os"
)

var (
	Reset = "\033[0m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Cyan = "\033[36m"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("How to use: ", Cyan, "plog PORT" , Reset)
		fmt.Println("Example: ", Yellow, "plog 8080", Reset)
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

	port := os.Args[1]

	fmt.Println(Yellow + "PORT LOGGER" + Cyan + " - by Conper" + Reset)
	fmt.Println(Blue + "Time format: DAY-MONTH-YEAR / HOUR:MINUTE:SECOND" + Cyan + "      Listening in port: " + port + "\n" + Reset)

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {log.Fatal(err)}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {log.Fatal(err)}
		fmt.Print("Connected received from: ", conn.RemoteAddr())
		go clock()
	}
}

func clock() {
	now := time.Now()
        fmt.Printf("    Time: %02d-%02d-%04d / %02d:%02d:%02d\n", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
	time.Sleep(1 * time.Second)
}
