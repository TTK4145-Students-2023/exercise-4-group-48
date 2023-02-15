package main

import (
	"Project/network"
	"Project/timer"
	"fmt"
	"os"
	"os/exec"
)

func listen(){
	network.UDPListen()

	// Make errorhandling if we dont get a heartbeat
}

func send(){
	timer.A()
	network.UDPSend()
}

func main() {
	file, err := os.Create("backupFile.txt")

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	timer.A(file)
	//exec.Command("gnome-terminal", "--", "go", "run", "timer.go").Run()

	network.UDPListen()





	listen() // backup

	exec.Command("gnome-terminal", "--", "go", "run", "main.go").Run()
	
	counting()  // sending

	exec.Command("exit").Run()


}
