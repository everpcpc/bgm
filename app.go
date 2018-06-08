package main

import (
	"context"
	"log"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	ctx         = context.Background()
	proxyURL    *url.URL
	oauthClient *oauth2.Config
	redisClient *redis.Client
	clientID    string
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/bgm")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config error: %+v", err)
	}
	proxyURL, err = url.Parse(viper.GetString("http_proxy"))
	if err != nil {
		log.Fatalf("parse http_proxy error: %+v", err)
	}
	clientID = viper.GetString("client_id")
	oauthClient = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: viper.GetString("client_secret"),
		RedirectURL:  viper.GetString("domain") + "/oauth/callback",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://bgm.tv/oauth/access_token",
			AuthURL:  "https://bgm.tv/oauth/authorize",
		},
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis"),
		DB:   0,
	})

	r := gin.New()

	sessionStore := cookie.NewStore([]byte(viper.GetString("secret")))
	r.Use(sessions.Sessions("mysession", sessionStore))

	r.GET("", index)

	auth := r.Group("/oauth")
	{
		auth.GET("/login", oauthLogin)
		auth.GET("/callback", oauthCallback)
	}

	r.Run(viper.GetString("listen"))
}

func index(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")
	if uid == nil {
		c.Redirect(302, "/oauth/login")
		return
	}
	_, err := getToken(uid.(int))
	if err != nil {
		if err == redis.Nil {
			c.Redirect(302, "/oauth/login")
			return
		}
		c.JSON(500, gin.H{"ok": false, "err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ok": true, "uid": uid})
}
