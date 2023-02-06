package client

import (
	"fmt"
	"time"

	"github.com/plaenkler/booklooker/handler"
	"github.com/plaenkler/booklooker/model"
)

type Client struct {
	APIKey string
	Token  model.Token
	active bool
}

func (c *Client) Start() error {
	c.active = true
	done := make(chan error)
	go func() {
		for c.active {
			req := model.AuthenticateRequest{
				APIKey: c.APIKey,
			}
			resp, err := handler.Authenticate(req)
			if err != nil {
				done <- fmt.Errorf("client-error-1: %s", err)
				return
			}
			if resp.Status != "OK" {
				done <- fmt.Errorf("client-error-2: %s", resp.ReturnValue)
				return
			}
			c.Token = model.Token{
				Value:  resp.ReturnValue,
				Expiry: time.Now().Add(10 * time.Minute),
			}
			done <- nil
			time.Sleep(9 * time.Minute)
		}
	}()
	return <-done
}

func (c *Client) Stop() {
	c.active = false
}
