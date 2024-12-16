package main

import (
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// Turn the sensor on, read the data, and then turn it back off to conserve power
func getTempAndHumidity() (temperature float64, humidity float64, err error) {
	err = rpio.Open()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer rpio.Close()

	// GPIO4 is where the transistor switch is connected, which controls power to the sensor
	controlPin := rpio.Pin(4)
	controlPin.Output()

	// Ensure the switch is off to start
	controlPin.Low()

	// GPIO17 is where sensor data will be read
	inputPin := rpio.Pin(17)
	inputPin.Input()

	// There are a number of pauses in this application which may not be necessary
	time.Sleep(time.Second)

TempLoop:

	// Reading the sensor can return errors, retry until a valid value is returned
	for x := 0; x < 10; x++ {

		// Turn on the sensor with the control pin
		controlPin.High()

		// Wait for the sensor to fully turn on before reading data
		time.Sleep(3 * time.Second)

		// Use the DHT library to actually read the data from the sensor
		humidity, temperature = getSensorData()

		// End the loop if a valid value is returned
		// Humidity should never be 0
		if humidity != 0 {
			break TempLoop
		}

		// Turn off the sensor
		controlPin.Low()

		time.Sleep(1 * time.Second)
	}

	// End with the sensor off
	controlPin.Low()

	return temperature, humidity, nil
}
