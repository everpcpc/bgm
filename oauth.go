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
	c.Redirect(302, url)
}

func oauthCallback(c *gin.Context) {
	code := c.Query("code")

	var httpClient *http.Client
	if proxyURL.String() != "" {
		httpClient = &http.Client{
			Timeout:   10 * time.Second,
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		}
	} else {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	token, err := oauthClient.Exchange(ctx, code)
	if err != nil {
		c.JSON(200, gin.H{
			"ok":      false,
			"err_msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":    true,
		"token": token,
		"uid":   token.Extra("user_id"),
	})
}
