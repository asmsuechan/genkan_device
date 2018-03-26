package main

import (
	"log"
	"os"
	"reflect"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/raspi"

	"github.com/JustinTulloss/firebase"
)

type History struct {
	Action string `json:",omitempty"`
	RanAt  string `json:",omitempty"`
}

func PushOpenToFirebase(c firebase.Client) {
	PushHistory(&History{Action: "open", RanAt: time.Now().String()}, c)
}

func PushCloseToFirebase(c firebase.Client) {
	PushHistory(&History{Action: "close", RanAt: time.Now().String()}, c)
}

func PushHistory(h *History, c firebase.Client) {
	_, err := c.Push(h, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	deviceID := os.Getenv("GENKAN_DEVICE_ID")

	endpoint := os.Getenv("GENKAN_FIREBASE_ENDPOINT")
	auth := os.Getenv("GENKAN_FIREBASE_AUTH")

	c := firebase.NewClient(endpoint+"/devices/"+deviceID+"/history", auth, nil)

	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	// mqtt.ERROR = log.New(os.Stdout, "", 0)
	uri := "tcp://" + os.Getenv("GENKAN_URI")
	id := "device"
	userName := os.Getenv("GENKAN_USERNAME")
	password := os.Getenv("GENKAN_PASSWORD")
	mqttAdaptor := mqtt.NewAdaptorWithAuth(uri, id, userName, password)

	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "12")

	work := func() {
		mqttAdaptor.On("genkan/devices/"+deviceID, func(msg mqtt.Message) {
			if reflect.DeepEqual(msg.Payload(), []byte("open")) {
				servo.Move(uint8(31))
				PushOpenToFirebase(c)
			} else if reflect.DeepEqual(msg.Payload(), []byte("close")) {
				servo.Move(uint8(13))
				PushCloseToFirebase(c)
			}
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}
