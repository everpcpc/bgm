package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	ctx         = context.Background()
	oauthClient *oauth2.Config
	proxyURL    *url.URL
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/bgm")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	proxyURL, err = url.Parse(viper.GetString("http_proxy"))
	if err != nil {
		panic(err)
	}
	oauthClient = &oauth2.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		RedirectURL:  viper.GetString("domain") + "/oauth/callback",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://bgm.tv/oauth/access_token",
			AuthURL:  "https://bgm.tv/oauth/authorize",
		},
	}

	r := gin.Default()

	auth := r.Group("/oauth")
	{
		auth.GET("/callback", authCallback)
		auth.GET("/url", authURL)
	}

	r.Run(viper.GetString("listen"))
}

func authURL(c *gin.Context) {
	url := oauthClient.AuthCodeURL("", oauth2.AccessTypeOnline)
	c.JSON(200, gin.H{
		"url": url,
	})
}

func authCallback(c *gin.Context) {
	code := c.Query("code")
	proxyedClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, proxyedClient)
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
