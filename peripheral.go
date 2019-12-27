package main

import (
	"fmt"
	"log"
	"strconv"
)

type Peripheral struct {
	Temperature float64
	Humidity float64
	BatteryLevel int64
}

// ParseData takes a byte [] and separates out the pieces from the bluetooth
// temperature sensor readings
func (p *Peripheral) ParseData(b []byte)  {
	if len(b) < 13 {
		return
	}

	t := fmt.Sprintf("%02X%02X", b[8], b[9])
	h := fmt.Sprintf("%02X%02X", b[10], b[11])
	bt := fmt.Sprintf("%02X", b[3])

	parseTemperature(p, &t)
	parseHumidity(p, &h)
	parseBattery(p, &bt)
}

// parseTemperature takes a peripheral and a hex string parses it
// out to create the temperature
func parseTemperature(p *Peripheral, temp *string) {
	i, err := strconv.ParseUint(*temp, 16, 64)

	if err != nil {
		log.Println("Error parsing temp: ", err)
		return
	}
	p.Temperature = float64(i) / 10.0
}

// parseHumidity takes a peripheral and a hex string parses it
// out to create the humidity
func parseHumidity(p *Peripheral, humidity *string) {
	i, err := strconv.ParseUint(*humidity, 16, 64)

	if err != nil {
		log.Println("Error parsing humidity: ", err)
		return
	}
	p.Humidity = float64(i) / 10.0
}

// parseBattery takes a peripheral and a hex string parses it
// out to create the battery
func parseBattery(p *Peripheral, batt *string) {
	b, err := strconv.ParseInt(*batt, 16, 64)

	if err != nil {
		log.Println("Error parsing battery: ", err)
		return
	}

	p.BatteryLevel = b
}