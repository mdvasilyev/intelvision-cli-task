package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

type PortType int

const (
	IN PortType = iota
	OUT
)

type Port struct {
	Type        PortType
	Number      int
	Value       int
	Transaction chan int
}

func (p *Port) Read(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-p.Transaction:
			fmt.Printf("Reading from IN port number %d value %d\n", p.Number, rand.Intn(2))
		}
	}
}

func (p *Port) Write(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case tr := <-p.Transaction:
			p.Value = 42
			fmt.Printf("Writing to OUT port number %d value %d. Transaction %d\n", p.Number, p.Value, tr)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter number of IN ports")
	scanner.Scan()
	inPortsNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Unable to parse a number of IN ports")
	}

	fmt.Println("Enter number of OUT ports")
	scanner.Scan()
	outPortsNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Unable to parse a number of OUT ports")
	}

	var wg sync.WaitGroup
	wg.Add(inPortsNum + outPortsNum)

	inPorts := make([]Port, 0, inPortsNum)
	outPorts := make([]Port, 0, outPortsNum)

	for i := 0; i < inPortsNum; i++ {
		port := Port{Type: IN, Number: i, Value: 0, Transaction: make(chan int)}
		inPorts = append(inPorts, port)
		go port.Read(&wg)
	}

	for i := 0; i < outPortsNum; i++ {
		port := Port{Type: OUT, Number: i, Value: 0, Transaction: make(chan int)}
		outPorts = append(outPorts, port)
		go port.Write(&wg)
	}

	for scanner.Scan() {
		s := scanner.Text()
		args := strings.Split(s, " ")

		if args[0] == "READ" {
			if len(args) != 2 {
				fmt.Println("Usage: READ <n>")
				continue
			}

			portNumber, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("wrong number parameter")
				continue
			}

			inPorts[portNumber].Transaction <- 0
		} else if args[0] == "WRITE" {
			if len(args) != 3 {
				fmt.Println("Usage: WRITE <n> <n>")
				continue
			}

			portNumber, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("wrong number parameter")
				continue
			}

			tr, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("wrong number parameter")
				continue
			}

			outPorts[portNumber].Transaction <- tr
		}
	}

	wg.Wait()
}
