package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eclipse/paho.mqtt.golang"
)

func onConnect(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message: %!s(MISSING) from topic: %!s(MISSING)\n", message.Payload(), message.Topic())
}

func main() {
	// MQTT broker address and port
	broker := "tcp://localhost:1883"

	// MQTT client ID
	clientId := "mqtt_client"

	// MQTT topic to subscribe to
	topic := "test"

	// Create MQTT client options
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientId)

	// Set MQTT client callbacks
	options.OnConnect = onConnect
	options.DefaultPublishHandler = onMessageReceived

	// Create MQTT client instance
	client := mqtt.NewClient(options)

	// Connect to MQTT broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to MQTT topic
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Wait for SIGINT or SIGTERM signal to exit
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	// Unsubscribe from MQTT topic
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Disconnect from MQTT broker
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}
