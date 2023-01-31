package client

import (
	"log"
	"time"

	"github.com/plaenkler/booklooker/handler"
	"github.com/plaenkler/booklooker/model"
)

type Client struct {
	APIKey string
	Token  model.Token
	active bool
}

func (c *Client) Start() {
	c.active = true
	done := make(chan bool)
	go func() {
		for c.active {
			req := model.AuthenticateRequest{
				APIKey: c.APIKey,
			}
			resp, err := handler.Authenticate(req)
			if err != nil || resp.Status != "OK" {
				log.Printf("token-renewal-failed: %s", resp.ReturnValue)
			}
			c.Token = model.Token{
				Value:  resp.ReturnValue,
				Expiry: time.Now().Add(10 * time.Minute),
			}
			done <- true
			time.Sleep(9 * time.Minute)
		}
	}()
	<-done
}

func (c *Client) Stop() {
	c.active = false
}
