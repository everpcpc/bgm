package main

import (
	"encoding/json"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

func getUIDKey(uid int) string {
	return "token_" + strconv.Itoa(uid)
}

func storeToken(uid int, token *oauth2.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return redisClient.Set(getUIDKey(uid), data, token.Expiry.Sub(time.Now())).Err()
}

func getToken(uid int) (token *oauth2.Token, err error) {
	ret, err := redisClient.Get(getUIDKey(uid)).Bytes()
	if err != nil {
		return
	}
	token = &oauth2.Token{}
	err = json.Unmarshal(ret, token)
	return
}
