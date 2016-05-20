package main

import (
	"golang.org/x/oauth2"
	/*"./msg"
	"github.com/golang/protobuf/proto"
	"log"*/
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

var conf = &oauth2.Config{
	ClientID:     "7_5az7pj935owsss8kgokcco84wc8osk0g0gksow0ow4s4ocwwgc",
	ClientSecret: "49p1ynqfy7c4sw84gwoogwwsk8cocg8ow8gc8o80c0ws448cs4",
	Scopes:       []string{"trading"},
	//RedirectURL: "https://sandbox-id.ctrader.com",
	RedirectURL: "http://localhost:8080/auth",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://sandbox-connect.spotware.com/oauth/v2/auth",
		TokenURL: "https://sandbox-connect.spotware.com/oauth/v2/token",
	},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/login", login)
	router.GET("/auth", auth)
	router.Run(":" + port)
}

func login(c *gin.Context) {
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	c.Redirect(http.StatusFound, conf.AuthCodeURL("state", oauth2.AccessTypeOffline))
}

func auth(c *gin.Context) {
	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	code := c.Query("code");
	//var token oauth2.Token = nil;
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}
	//client := conf.Client(oauth2.NoContext, token)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"token" : token.AccessToken,
	})
}
