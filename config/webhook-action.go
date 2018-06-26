package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"../pinlike"
)

type webhookConfig struct {
	ObjectKind string `json:"object_kind"`
	Calling    []int  `json:"calling"`
}

type WebhookAction interface {
	Run(pin pinlike.PinLike)
}

func (c *webhookConfig) Run(pin pinlike.PinLike) {
	var duration int
	var what string
	for _, call := range c.Calling {
		if call < 0 {
			what = "Low"
			duration = -call
			pin.Low()
		} else {
			what = "High"
			duration = call
			pin.High()
		}
		log.Println(what + " for " + strconv.Itoa(duration) + "ms")
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}
	pin.Low()
}

func LoadWebhookAction(configName string, action string) (WebhookAction, error) {
	raw, err := ioutil.ReadFile("./webhooks/" + configName + ".json")
	if err != nil {
		return nil, err
	}
	var c []webhookConfig
	json.Unmarshal(raw, &c)

	config := findActionConfig(c, action)

	return config, nil
}

func findActionConfig(configs []webhookConfig, action string) *webhookConfig {
	for _, config := range configs {
		if config.ObjectKind == action {
			return &config
		}
	}
	return nil
}
