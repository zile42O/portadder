package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("\n\x1b[32m× Please enter the port which you want to add on IP addr...", "\x1b[0m")
	fmt.Println("\n\x1b[93mNote: File name need be 'iplist.txt'", "\x1b[0m")
	fmt.Println("\n\x1b[35mCreated with ♥ by Zile42O", "\x1b[0m")
	start := time.Now()

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(string(sentence))
	}

	LinebyLineScan(string(sentence))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("\n\x1b[32m× Program took:\x1b[0m \x1b[33m", elapsed, "\x1b[0m")
	fmt.Println("\n----------------------------------------------------")
	time.Sleep(5000 * time.Millisecond)
}

func IsIpv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

func LinebyLineScan(port string) {
	file, err := os.Open("./iplist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ipnum := 0

	f, err := os.OpenFile("output.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for scanner.Scan() {
		if err := IsIpv4Net(scanner.Text()); err != false {
			//fmt.Println(scanner.Text())
			result := scanner.Text() + ":" + port
			if _, err := f.WriteString(result); err != nil {
				log.Println(err)
			}
			//fmt.Println(result)
		}
		ipnum++
		//fmt.Println(scanner.Text())

	}
	fmt.Println("\n\x1b[35mTotal IPes: ", ipnum, "\x1b[0m")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
