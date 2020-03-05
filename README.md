# Lit-Go-Bot
[![Go Report Card](https://goreportcard.com/badge/github.com/Abhishek5101/Lit-Go-Bot)](https://goreportcard.com/report/github.com/Abhishek5101/Lit-Go-Bot)

## Purpose
Using Go lang with Arduino to convert text to Morse code and transmit it with LEDs

![Project Picture](https://raw.githubusercontent.com/Abhishek5101/Lit-Go-Bot/master/project_picture.jpg)

![Int'l Morse Code](https://raw.githubusercontent.com/Abhishek5101/Lit-Go-Bot/master/colored_morse_code.jpg)

### Presentation slides
[Morse Code](https://docs.google.com/presentation/d/19pMEI79wN8EOLtRxAUrVsAwAV1u26ckRt4PgVCkrl90/edit?usp=sharing)

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

## 📝 License

By contributing, you agree that your contributions will be licensed under its MIT License.