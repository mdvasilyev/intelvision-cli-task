package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Port struct {
	Type  string
	Value int
}

func Read(ports []Port, pos int) (int, error) {
	if pos >= len(ports) {
		return -1, errors.New("no such port")
	}
	if ports[pos].Type != "IN" {
		return -1, errors.New("port must be IN")
	}
	return ports[pos].Value, nil
}

func Write(ports []Port, pos int, value int) error {
	if pos >= len(ports) {
		return errors.New("no such port")
	}
	if ports[pos].Type != "OUT" {
		return errors.New("port must be OUT")
	}
	ports[pos].Value = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter number of IN ports")
	scanner.Scan()
	inNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Only integers are accepted")
	}

	fmt.Println("Enter number of OUT ports")
	scanner.Scan()
	outNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Only integers are accepted")
	}

	inPorts := make([]Port, inNumber)
	for i := 0; i < inNumber; i++ {
		inPorts[i] = Port{Type: "IN", Value: -1}
	}

	outPorts := make([]Port, outNumber)
	for i := 0; i < outNumber; i++ {
		outPorts[i] = Port{Type: "OUT", Value: -1}
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

			val, err := Read(inPorts, portNumber)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Printf("Value of port %d = %d\n", portNumber, val)
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

			valToWrite, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("wrong number parameter")
				continue
			}

			err = Write(outPorts, portNumber, valToWrite)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Printf("Wrote value %d to port %d\n", valToWrite, portNumber)
		}
	}
}
