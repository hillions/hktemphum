package main

import (
	"log"
	"net/http"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

// Added a new kind of accessory "Hygrometer" (Humidity Sensor) since the hap package doesn't yet include this
type Hygrometer struct {
	*accessory.A
	HumiditySensor *service.HumiditySensor
}

func NewHumiditySensor(info accessory.Info) *Hygrometer {
	a := Hygrometer{}
	a.A = accessory.New(info, accessory.TypeThermostat)
	a.HumiditySensor = service.NewHumiditySensor()
	a.AddS(a.HumiditySensor.S)

	return &a
}

func createHygrometer(info accessory.Info) *Hygrometer {

	// Create the hygrometer accessory
	hy := NewHumiditySensor(info)
	hy.HumiditySensor.CurrentRelativeHumidity.OnValueUpdate(func(new float64, old float64, r *http.Request) {
		log.Printf("New humidity %%%v, old value %%%v", new, old)
	})

	return hy
}

func createThermometer(info accessory.Info) *accessory.Thermometer {

	// Create the thermometer accessory
	ts := accessory.NewTemperatureSensor(info)
	// This is the range of the DHT22 sensor
	ts.TempSensor.CurrentTemperature.MinVal = -40.0
	ts.TempSensor.CurrentTemperature.MaxVal = 125.0

	ts.TempSensor.CurrentTemperature.OnValueUpdate(func(new float64, old float64, r *http.Request) {
		log.Printf("New temperature %vC, old value %vC", new, old)
	})

	return ts
}
