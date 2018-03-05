package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func oauthURL(c *gin.Context) {
	url := oauthClient.AuthCodeURL("", oauth2.AccessTypeOnline)
	c.JSON(200, gin.H{
		"url": url,
	})
}

func oauthCallback(c *gin.Context) {
	code := c.Query("code")

	if proxyURL.String() != "" {
		proxyedClient := &http.Client{
			Timeout:   10 * time.Second,
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, proxyedClient)
	}

	token, err := oauthClient.Exchange(ctx, code)
	if err != nil {
		c.JSON(200, gin.H{
			"ok":      false,
			"err_msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"OK":    true,
		"code":  code,
		"token": token,
	})
}