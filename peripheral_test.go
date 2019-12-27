package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPeripheral_ParseTemperature_Test1(t *testing.T) {
	// arrange
	var s = "33011b480e10177000d4020427016500"
	b, _ := hex.DecodeString(s)
	p := new(Peripheral)
	x := fmt.Sprintf("%02X%02X", b[8], b[9])

	// act
	parseTemperature(p, &x)

	// assert
	if p.Temperature != 21.2 {
		t.Fail()
	}
}

func TestPeripheral_ParseTemperature_Test2(t *testing.T) {
	// arrange
	var s = "33011b480e1017700154020427016500"
	b, _ := hex.DecodeString(s)
	p := new(Peripheral)
	x := fmt.Sprintf("%02X%02X", b[8], b[9])

	// act
	parseTemperature(p, &x)

	// assert
	if p.Temperature != 34 {
		t.Fail()
	}
}

func TestPeripheral_ParseBattery_Test1(t *testing.T) {
	// arrange
	var s = "33011b480e10177000d4020427016500"
	b, _ := hex.DecodeString(s)
	p := new(Peripheral)
	x := fmt.Sprintf("%02X", b[3])

	// act
	parseBattery(p, &x)

	// assert
	if p.BatteryLevel != 72 {
		t.Fail()
	}
}

func TestPeripheral_ParseBattery_Test2(t *testing.T) {
	// arrange
	var s = "33011b590e10177000d4020427016500"
	b, _ := hex.DecodeString(s)
	p := new(Peripheral)
	x := fmt.Sprintf("%02X", b[3])

	// act
	parseBattery(p, &x)

	// assert
	if p.BatteryLevel != 89 {
		t.Fail()
	}
}
