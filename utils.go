package main

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

const (
	// ref: https://developer.google.com/chart/infographics/docs/qr_codes
	qrURL = "https://chart.googleapis.com/chart?"
)

func getUIDKey(uid int) string {
	return "token_" + strconv.Itoa(uid)
}

func storeToken(uid int, token *oauth2.Token) error {
	if redisClient == nil {
		return nil
	}
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return redisClient.Set(getUIDKey(uid), data, token.Expiry.Sub(time.Now())).Err()
}

func getToken(uid int) (token *oauth2.Token, err error) {
	if redisClient == nil {
		return
	}
	ret, err := redisClient.Get(getUIDKey(uid)).Bytes()
	if err != nil {
		return
	}
	token = &oauth2.Token{}
	err = json.Unmarshal(ret, token)
	return
}

func getQRCode(data string) string {
	v := url.Values{}
	v.Add("chl", data)
	v.Add("chs", "400x400")
	v.Add("cht", "qr")
	return qrURL + v.Encode()
}
