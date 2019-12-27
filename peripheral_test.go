package main

import "testing"

func TestPeripheral_ParseData(t *testing.T) {
	type fields struct {
		Temperature  float32
		Humidity     float32
		BatteryLevel int
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Peripheral{
				Temperature:  tt.fields.Temperature,
				Humidity:     tt.fields.Humidity,
				BatteryLevel: tt.fields.BatteryLevel,
			}
		})
	}
}