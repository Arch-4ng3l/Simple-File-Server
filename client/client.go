package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

func getOs() uint8{
	OS := runtime.GOOS
	if OS == "linux" {
		return 1
	} 
	return 0
}

func connect() net.Conn {
	var conn net.Conn = nil

	for (conn == nil) {
		conn, _ = net.DialTimeout("tcp", "localhost:3333", time.Second * 5) 
	}

	return conn

}

func sendFile(fileName string, conn net.Conn) bool {
	

	file, err := os.Open(fileName)
		
	if err != nil {
		fmt.Println(color.Bold + color.Red + "Coulnt Find File" + color.Reset)
		return false
	}
	defer file.Close()
	file.Seek(0, 0)

	fileReader := bufio.NewReader(file)

	_, err = io.Copy(conn, fileReader)
	if err != nil {
		return false
	}
	return true

}


func getInput() string {

	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" => ")

	input, _ := reader.ReadString('\n')

	input = strings.ReplaceAll(input, "\n", "")

	return input

}


func printHelp() {
	
	exec.Command("clear")
	fmt.Println("#########################################################################")
	fmt.Println("#" + color.Bold + color.Green + "  ls   " + color.Reset + "- Prints all Files in the Current Directory")
	fmt.Println("#" + color.Bold + color.Green + "  pwd  " + color.Reset + "- Prints the Current Directory" ) 
	fmt.Println("#" + color.Bold + color.Green + "  cd   " + color.Reset + "- Change Current Directory")
	fmt.Println("#" + color.Bold + color.Green + "  send " + color.Reset + "- Send File to the Server")
	fmt.Println("#########################################################################")
}

func printLs(OS uint8) {
	var cmd = &exec.Cmd {}
	if OS == 1 {
		cmd = exec.Command("ls")
	} else {
		cmd = exec.Command("dir")
	}

	output, _ := cmd.Output()
	fmt.Printf("%s", output)
}

func printPwd(OS uint8) {
	var cmd = &exec.Cmd {}
	if OS == 1 {
		cmd = exec.Command("pwd")
	} else {
		cmd = exec.Command("echo %cd%")
	}
	output, _ := cmd.Output()
	fmt.Printf("%s", output)
}

func send(conn net.Conn) {
	fmt.Print(color.Bold + color.Green + "FileName"+ color.Reset)
	fileName := getInput()
	res := sendFile(fileName, conn)	

	if res {
		fmt.Println("Closing Connection")
		conn.Close()
	}
}

func chdir() {
	fmt.Println(color.Bold + color.Green + "Directory" + color.Reset)
	dir := getInput()
	os.Chdir(dir)
}
func menu(conn net.Conn, OS uint8) {

	input := getInput()
	switch input {
		case "help":
			printHelp()
		case "ls":
			printLs(OS)
		case "pwd":
			printPwd(OS)
		case "cd": 
			chdir()
		case "send":
			send(conn)
			return 
		default: 
			fmt.Println("No Valid Command")

	}

}

func main() {
	
	conn := connect()	

	defer conn.Close()

	fmt.Println(color.InGreen("[+] Connected"))

	fmt.Println(color.Bold + color.Green + " help " + color.Reset + "to see all commands")
	for {
		menu(conn, getOs())
	}

}
