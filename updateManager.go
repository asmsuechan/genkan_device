package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
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

// Usage:
// ./updateManager [version]
// example
// $ ./updateManager 1.0.0
func main() {
	uri := os.Getenv("GENKAN_URI")
	userName := os.Getenv("GENKAN_USERNAME")
	password := os.Getenv("GENKAN_PASSWORD")

	cloudmqtturl := "mqtt://" + userName + ":" + password + "@" + uri

	parseduri, err := url.Parse(cloudmqtturl)
	if err != nil {
		// TODO: Send error log to a service (not decided)
		log.Fatal(err)
	}

	client := connect("pub", parseduri)

	flag.Parse()
	args := flag.Args()

	topic := "genkan/update"
	token := client.Publish(topic, 0, false, args[0])
	token.Wait()
}
