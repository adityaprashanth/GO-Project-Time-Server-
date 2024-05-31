package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connected with:", conn.RemoteAddr())

	for {
		fmt.Println("Waiting to receive data...")
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}
		Data := string(data[:n])
		fmt.Println("\nReceived data:", Data)

		if strings.TrimSpace(Data) == "quit" || strings.TrimSpace(Data) == "QUIT" {
			fmt.Println("Client requested to quit. Closing connection...")
			break
		}

		go func(data string) {
			timezones := strings.Split(data, ",")
			var responses []string
			for _, tz := range timezones {
				tz = strings.TrimSpace(tz)
				var response string
				switch tz {
				case "IND":
					response = IND()
				case "USA":
					response = USA()
				case "UK":
					response = UK()
				case "AUS":
					response = AUS()
				case "JPN":
					response = JPN()
				case "RSA":
					response = RSA()
				case "SA":
					response = SA()
				default:
					response = "Timezone not found"
				}
				responses = append(responses, fmt.Sprintf("%s: %s", tz, response))
			}

			finalResponse := strings.Join(responses, ", ")
			_, err := conn.Write([]byte(finalResponse))
			if err != nil {
				fmt.Println("Error sending response:", err)
				return
			}
		}(Data)
	}
}

func IND() string {
	now := time.Now()
	return now.Format("15:04:05")
}

func USA() string {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func UK() string {
	loc, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func AUS() string {
	loc, _ := time.LoadLocation("Australia/Sydney")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func JPN() string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func RSA() string {
	loc, _ := time.LoadLocation("Africa/Maputo")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func SA() string {
	loc, _ := time.LoadLocation("Chile/Continental")
	now := time.Now().In(loc)
	return now.Format("15:04:05")
}

func main() {
	PORT := "8080"

	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on port", PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleClient(conn)
	}
}
