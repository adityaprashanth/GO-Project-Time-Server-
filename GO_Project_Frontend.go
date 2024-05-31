package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	fmt.Println("Enter the timezones separated by commas (IND,USA,UK,AUS,JPN,RSA,SA):")

	reader := bufio.NewReader(os.Stdin)
	timezones, _ := reader.ReadString('\n')
	timezones = strings.TrimSpace(timezones)

	if timezones == "" {
		fmt.Println("No timezones provided. Exiting...")
		return
	}

	timezoneList := strings.Split(timezones, ",")

	for _, tz := range timezoneList {
		go func(tz string) {
			response := sendRequest(tz)
			fmt.Printf("Response from backend for timezone %s: %s\n", tz, response)
		}(strings.TrimSpace(tz))
	}

	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}

func sendRequest(timezone string) string {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error connecting to backend:", err)
		return ""
	}
	defer conn.Close()

	_, err = conn.Write([]byte(timezone))
	if err != nil {
		fmt.Println("Error sending request to backend:", err)
		return ""
	}

	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error reading response from backend:", err)
		return ""
	}

	response := string(data[:n])
	return response
}
