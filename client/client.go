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
	go func() {
		for c.active {
			if c.Token.Expiry.Before(time.Now()) {
				req := model.AuthenticateRequest{
					APIKey: c.APIKey,
				}
				resp, err := handler.Authenticate(req)
				if err != nil {
					log.Printf("1-authentication-failed: %s", resp.ReturnValue)
				}
				if resp.Status != "OK" {
					log.Printf("2-authentication-failed: %s", resp.ReturnValue)
				}
				c.Token = model.Token{
					Value:  resp.ReturnValue,
					Expiry: time.Now().Add(10 * time.Minute),
				}
			}
			time.Sleep(9 * time.Minute)
		}
	}()
}

func (c *Client) Stop() {
	c.active = false
}
