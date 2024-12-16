package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
)

func main() {

	tempInfo := accessory.Info{
		Name:         "thermometer1",
		Manufacturer: "Manufacturer",
		SerialNumber: "000001",
		Model:        "TempSensor1",
		Firmware:     "0.0.1",
	}

	ts := createThermometer(tempInfo)

	hygrometerInfo := accessory.Info{
		Name:         "hygrometer1",
		Manufacturer: "Manufacturer",
		SerialNumber: "000001",
		Model:        "HumiditySensor1",
		Firmware:     "0.0.1",
	}

	hy := createHygrometer(hygrometerInfo)

	// Store the data in the "./db" directory.
	fs := hap.NewFsStore("./db")

	// Create the hap server.
	server, err := hap.NewServer(fs, ts.A, hy.A)
	if err != nil {
		// stop if an error happens
		log.Panic(err)
	}

	// Enter the pin to supply when adding to HomeKit
	server.Pin = "01010101"

	// Periodically change the temperature characteristic
	go func() {
		for {
			temperature, humidity, err := getTempAndHumidity()

			// If there is no error, set the accessory values in homekit
			if err != nil {
				fmt.Println(err)
			} else {
				ts.TempSensor.CurrentTemperature.SetValue(temperature)
				hy.HumiditySensor.CurrentRelativeHumidity.SetValue(humidity)
			}

			// Change this value to alter the sensor polling rate
			// 5 minutes
			time.Sleep(5 * time.Minute)
		}
	}()

	// Run the server.
	ctx := context.WithoutCancel(context.Background())
	server.ListenAndServe(ctx)
}
