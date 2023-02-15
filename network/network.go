package network

import (
	"fmt"
	"log"
	"net"
)

func UDPListen() {
	UDPAddress, err := net.ResolveUDPAddr("udp4", ":20024")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", UDPAddress)

	if err != nil {
		fmt.Println("Error resolving UDP address: ", err)
		return
	}

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		fmt.Println(string(buf[:n]))

	}

}

func UDPSend(message string) {
	UDPAddress, err := net.ResolveUDPAddr("udp4", "255.255.255.255:20024")

	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.DialUDP("udp4", nil, UDPAddress)
	if err != nil {
		fmt.Println("Error creating UDP connection; ", err)
		return
	}

	defer connection.Close()

	buffer := []byte(message)

	_, err = connection.Write(buffer)
	if err != nil {
		fmt.Println("Error sending data over UDP: ", err)
		return
	}
	fmt.Println("String: ", message, ", sucesfully sent")
}
