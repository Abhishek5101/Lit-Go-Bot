# Lit-Go-Bot
[![Go Report Card](https://goreportcard.com/badge/github.com/Abhishek5101/Lit-Go-Bot)](https://goreportcard.com/report/github.com/Abhishek5101/Lit-Go-Bot)

## Purpose
Using Go lang with Arduino to convert text to Morse code and transmit it with LEDs

![Project Picture](https://ibb.co/BCLkH4k)

![Int'l Morse Code](https://ibb.co/Qc5jxQX)

## Install
This package has the following requirements:
*  Gobot
    - `go get -d -u gobot.io/x/gobot/...`
*  Firmata
*  Morse code library by alwindoss
    - `go get -u -v github.com/alwindoss/morse`
* An arduino Nano/Uno/Compatible board
* Arduino IDE to upload Firmata sketch

### How to use

* Download the required dependanices above
* This program  assumes your LED is connected to Pin 13
* From your Aruino IDE, upload the already available standard Firmata sketch
* Then run this program