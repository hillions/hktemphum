# HomeKit Temperature and Humidity Sensor
Building a HomeKit integration for the DHT22 Temperature/Humidity sensor and a Raspberry Pi. Connects over Wifi - no bridge required.

## Description
There aren't many reliable temperature/humidity sensors available for HomeKit which do not require their own bridge. This project details how to build one based on a Raspberry Pi and the DHT22 sensor which is commonly available. Since the Raspberry Pi is powered via DC, one could use either a wall AC/DC converter (recommended) or battery to power the sensor.

A simple transistor switch is used to power down the sensor when not in use to conserve power. The system first turns on the sensor and then waits to take a reading. Once a clean reading is taken, the sensor is powered off for an interval.

## References
This project would not be possible without the prior work done in these Go packages:

[go-dht](https://github.com/MichaelS11/go-dht) - reading the DHT22 sensor

[go-rpio](https://github.com/stianeikeland/go-rpio) - basic handling of GPIO pins

[hap](https://github.com/brutella/hap) - allowing for the development of HomeKit accessories

## Configuration
TODO

## Parts List
100k ohm resistor

2N2222 transistor

DHT22 temperature and humidity sensor

Raspberry Pi Zero 2 W

Blank PCB / Wire / Soldering Kit

## Wiring Diagram
https://pinout.xyz
TODO

## Build
Building for Raspberry Pi on an M2 MacBook Pro:

`env GOOS=linux GOARCH=arm go build`

## Linux Service
TODO

