package main

import (
	"fmt"
	"log"
	strconv "strconv"
)

type Peripheral struct {
	Temperature float64
	Humidity float64
	BatteryLevel int
}

func (p *Peripheral) ParseData(b []byte)  {
	if len(b) < 12 {
		return
	}

	t := fmt.Sprintf("%02X%02X", b[8], b[9])
	i, err := strconv.ParseUint(t, 16, 64)
	if err != nil {
		log.Println("Error parsing temp: ", err)
		return
	}
	p.BatteryLevel = int(b[3])
	p.Temperature = float64(i) / 10.0
	// p.Humidity = float32((b[10] + b[11]) / 10.0)
}

func parseTemperature(p *Peripheral, temp *string) {
	i, err := strconv.ParseUint(*temp, 16, 64)
	if err != nil {
		log.Println("Error parsing temp: ", err)
		return
	}
	p.Temperature = float64(i) / 10.0
}