package main

import (
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
        "github.com/alwindoss/morse"
        "strings"
		"fmt"
		"io/ioutil"
)

func main() {
        firmataAdaptor := firmata.NewAdaptor("/dev/cu.usbmodem141401")
        led := gpio.NewLedDriver(firmataAdaptor, "13")

        work := func() {
                gobot.Every(5*time.Second, func() {
                        led.Toggle()
                })
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led},
                work,
        )

        robot.Start()
        h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Convert this to Morse"))
	if err != nil {
		return
	}
	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	fmt.Printf("Abhishek: %s\n", string(morseCode))
	writeFile("morse.txt", string(morseCode))

}

func writeFile(name string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}
