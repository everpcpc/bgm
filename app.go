package main

import (
	"context"
	"net/url"

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

	r := gin.New()

	auth := r.Group("/oauth")
	{
		auth.GET("/callback", oauthCallback)
		auth.GET("/url", oauthURL)
		r.Run(viper.GetString("listen"))
	}
}
