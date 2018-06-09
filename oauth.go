package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func oauthLogin(c *gin.Context) {
	url := oauthClient.AuthCodeURL("", oauth2.AccessTypeOnline)
	c.Redirect(302, url)
}

func oauthCallback(c *gin.Context) {
	session := sessions.Default(c)
	code := c.Query("code")
	var err error

	defer func() {
		if err != nil {
			c.JSON(500, gin.H{"ok": false, "err": err.Error()})
		}
	}()

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
		return
	}

	var uid int
	_uid := token.Extra("user_id")
	if _uid == nil {
		err = fmt.Errorf("uid not found")
		return
	}
	switch v := _uid.(type) {
	case string:
		err = fmt.Errorf("uid is string: %+v", v)
		return
	case int:
		uid = v
	case float64:
		uid = int(v)
	default:
		err = fmt.Errorf("uid invalid: %+v", v)
		return
	}

	err = storeToken(uid, token)
	if err != nil {
		return
	}

	session.Set("uid", uid)
	session.Save()

	data, err := json.Marshal(token)
	if err != nil {
		return
	}
	c.Redirect(302, getQRCode(base64.StdEncoding.EncodeToString(data)))

}
