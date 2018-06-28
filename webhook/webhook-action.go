package webhook

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"../pinlike"
)

type webhookConfigAction struct {
	ObjectKind string `json:"object_kind"`
	Calling    []int  `json:"calling"`
}
type webhookConfig struct {
	SecretToken string                `json:"secret_token"`
	Actions     []webhookConfigAction `json:"actions"`
}

type actionImpl struct {
	calling     []int
	secretToken string
}

type WebhookAction interface {
	Run(pin pinlike.PinLike)
	IsAuthorized(req *http.Request) bool
}

func (a *actionImpl) IsAuthorized(req *http.Request) bool {
	token := req.Header.Get("X-Gitlab-Token")
	return a.secretToken == "" || a.secretToken == token
}

func (c *actionImpl) Run(pin pinlike.PinLike) {
	var duration int
	var what string
	for _, call := range c.calling {
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
	var c webhookConfig
	json.Unmarshal(raw, &c)

	config := findActionConfig(c, action)
	if config == nil {
		return nil, nil
	}

	return config, nil
}

func findActionConfig(conf webhookConfig, action string) *actionImpl {
	for _, config := range conf.Actions {
		if config.ObjectKind == action {
			return &actionImpl{config.Calling, conf.SecretToken}
		}
	}
	return nil
}
