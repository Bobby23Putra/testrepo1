package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/goburrow/modbus"
)

//pertama itu serial berapa
//slave id
//register
//quantity

func main() {
	serial := os.Args[1]
	slave := os.Args[2]
	register := os.Args[3]
	length := os.Args[4]

	slave1, _ := strconv.Atoi(slave)
	client := modbus.NewRTUClientHandler(serial)
	client.SlaveId = byte(slave1)
	client.BaudRate = 9600
	client.DataBits = 8
	client.Parity = "N"
	client.StopBits = 1
	client.Timeout = 3 * time.Second

	err := client.Connect()
	if err != nil {
		log.Fatalf("RTU Connect: %v", err)
	}
	defer client.Close()
	cli := modbus.NewClient(client)
	register1, _ := strconv.ParseUint(register, 0, 16)
	lenght1, _ := strconv.ParseUint(length, 0, 16)
	// println(serial, slave, register, length, slave1, register1, lenght1)
	result1, err1 := cli.ReadHoldingRegisters(uint16(register1), uint16(lenght1))
	print(result1, err1)
}
