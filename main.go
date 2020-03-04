package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"github.com/alwindoss/morse"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Hello from Arduino"))
	if err != nil {
		return
	}
	// morseTest := "-.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... ."
	lettersAndSpaces := strings.ReplaceAll(string(morseCode), " - / -", "")
	lettersOnly := strings.ReplaceAll(lettersAndSpaces, " ", "")
	fmt.Printf("Morse Translation: %s\n", string(lettersOnly))
	writeFile("morse.txt", string(lettersOnly))

	firmataAdaptor := firmata.NewAdaptor("/dev/cu.usbmodem142401")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		for a := 0; a < len(lettersOnly); a++ {
			if string(lettersOnly[a]) == "." {
				gobot.Every(1*time.Second, func() {
					led.Toggle()
				})
			} else {
				gobot.Every(3*time.Second, func() {
					led.Toggle()
				})
			}
		    }
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)
	robot.Start()

}

func writeFile(name string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}
