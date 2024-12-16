# HomeKit Temperature and Humidity Sensor
Building a HomeKit integration for the DHT22 Temperature/Humidity sensor and a Raspberry Pi. Connects over Wifi - no bridge required.

## Description
TODO

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

Blank PCB

## Wiring Diagram
TODO

## Build
Building for Raspberry Pi on an M2 MacBook Pro:

`env GOOS=linux GOARCH=arm go build`

