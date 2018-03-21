package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Exec update script
		exec.Command("./updater.sh", string(msg.Payload())).Run()
	})
}

func main() {
	uri := os.Getenv("GENKAN_URI")
	userName := os.Getenv("GENKAN_USERNAME")
	password := os.Getenv("GENKAN_PASSWORD")

	cloudmqtturl := "mqtt://" + userName + ":" + password + "@" + uri

	parseduri, err := url.Parse(cloudmqtturl)
	if err != nil {
		log.Fatal(err)
	}
	topic := "genkan/update"

	go listen(parseduri, topic)
	select {}
}
