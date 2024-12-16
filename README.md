# HomeKit Temperature and Humidity Sensor
Building a HomeKit integration for the DHT22 Temperature/Humidity sensor and a Raspberry Pi. Connects over Wifi - no bridge required.

# Description
TODO

# References
[go-dht](https://github.com/MichaelS11/go-dht)

[go-rpio](https://github.com/stianeikeland/go-rpio)

[hap](https://github.com/brutella/hap)

# Configuration
TODO

# Parts List
100k ohm resistor

2N2222 transistor

DHT22 temperature and humidity sensor

Raspberry Pi Zero 2 W

Blank PCB

# Wiring Diagram
TODO

# Build
Building for Raspberry Pi on an M2 MacBook Pro:

`env GOOS=linux GOARCH=arm go build`

