package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"time"
)

// тестовая длительность материала.
// в реальной программе сначало нужно узнать длительность материала коммандой ffprobe
const full_duration = 300.000000

func main() {

	PORT := ":" + "9090"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

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
		// Считать 256 байт до символа 'd'
		 s1, err := bufio.NewReaderSize(c,256).ReadBytes('d')

		//fmt.Println(s1)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Ищем строку с микросекундами
		reg1 := regexp.MustCompile(`out_time_ms=.*`)
		if err != nil {
			log.Fatal(err)
		}

		// Удаляем все кроме букв
		reg2 := regexp.MustCompile(`[0-9]+`)
		if err != nil {
			log.Fatal(err)
		}

		// Накладываем два фильтра регулярных выражений
		t1 := reg1.Find([]byte(string(s1)))
		t2 := reg2.Find([]byte(string(t1)))

		real_time, err := strconv.ParseFloat(string(t2), 64)

		// Вычисление процента завершения
		complete := int(((real_time / 1000000) * 100)/full_duration)

		fmt.Println(complete)

	}
	c.Close()
}



