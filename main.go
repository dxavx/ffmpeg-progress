package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
)

// Test duration of the material.
// In a real program you first need to know the duration of the material with the ffprobe command
const fullDuration = 300.000000

func main() {

	PORT := ":" + "9090"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		// Read 256 bytes to the 'd' character
		s1, err := bufio.NewReaderSize(c, 256).ReadBytes('d')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Looking for a string with microseconds
		reg1 := regexp.MustCompile(`out_time_ms=.*`)
		if err != nil {
			log.Fatal(err)
		}

		// Delete all but the letters
		reg2 := regexp.MustCompile(`[0-9]+`)
		if err != nil {
			log.Fatal(err)
		}

		// Consistent use of regular expressions
		t1 := reg1.Find(s1)
		t2 := reg2.Find(t1)

		realTime, err := strconv.ParseFloat(string(t2), 64)

		// Calculating the completion percentage
		complete := int(((realTime / 1000000) * 100) / fullDuration)

		fmt.Println(complete)

	}
	c.Close()
}
