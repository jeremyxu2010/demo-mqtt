/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(true)
	opts.SetUsername("testuser")
	opts.SetPassword("testpassword")


	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	qos := byte(0)
	//qos := byte(1)
	//qos := byte(2)

	if token := c.Subscribe("go-mqtt/sample", qos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", qos, false, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)

	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}