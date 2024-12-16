package main

import (
	"log"

	"hillions.net/hktemphum/dht"
)

func getSensorData() (humidity float64, temperature float64) {
	err := dht.HostInit()
	if err != nil {
		log.Print("HostInit error:", err)
		return 0, 0
	}

	// Change pinName to the appropriate GPIO pin with the sensor signal
	dht, err := dht.NewDHT("GPIO17", dht.Celsius, "")
	if err != nil {
		log.Print("NewDHT error:", err)
		return 0, 0
	}

	// A few different configurations are available for reading data from the sensor
	// ReadRetry works well for a non-time sensitive application
	humidity, temperature, err = dht.ReadRetry(20)
	if err != nil {
		log.Print("Read error:", err)
		return 0, 0
	}
	return humidity, temperature
}
