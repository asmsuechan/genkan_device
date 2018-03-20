package main

import (
	"fmt"
	"os"
	"reflect"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	// mqtt.ERROR = log.New(os.Stdout, "", 0)
	uri := os.Getenv("GENKAN_URI")
	id := "device"
	userName := os.Getenv("GENKAN_USERNAME")
	password := os.Getenv("GENKAN_PASSWORD")
	fmt.Println(uri, userName, password)
	mqttAdaptor := mqtt.NewAdaptorWithAuth(uri, id, userName, password)

	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "12")

	work := func() {
		servo.Move(uint8(27))

		mqttAdaptor.On("test", func(msg mqtt.Message) {
			if reflect.DeepEqual(msg.Payload(), []byte("open")) {
				servo.Move(uint8(31))
			} else if reflect.DeepEqual(msg.Payload(), []byte("close")) {
				servo.Move(uint8(13))
			}
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}
