package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/raff/goble"
)

var redisService *RedisService

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	verbose := flag.Bool("verbose", false, "dump all events")
	dups := flag.Bool("allow-duplicates", true, "allow duplicates when scanning")
	flag.Parse()
	_ = New("127.0.0.1:6379")
	var quit chan bool

	ble := goble.New()
	ble.SetVerbose(*verbose)

	if *verbose {
		ble.On(goble.ALL, func(ev goble.Event) (done bool) {
			log.Println("Event", ev)
			return
		})
	}

	ble.On("stateChange", func(ev goble.Event) (done bool) {
		if *verbose {
			fmt.Println("stateChange", ev.State)
		}
		if ev.State == "poweredOn" {
			ble.StartScanning(nil, *dups)
		} else {
			ble.StopScanning()
			done = true
			quit <- true
		}

		return
	})

	ble.On("discover", func(ev goble.Event) (done bool) {
		if  ev.Peripheral.Advertisement.LocalName != "Freezer" {
			return
		}

		p := Peripheral{}
		p.ParseData(ev.Peripheral.Advertisement.ManufacturerData)
		return
	})

	if *verbose {
		log.Println("Init...")
	}

	ble.Init()

	<-quit
}