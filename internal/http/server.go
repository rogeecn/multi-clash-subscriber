package http

import (
	"fmt"
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/conf"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

var cacheData []byte

func Serve() error {
	var err error
	if cacheData, err = conf.GetYaml(); err != nil {
		log.Fatal(err)
	}

	go func() {
		ticker := time.NewTicker(time.Minute * 10)
		for range ticker.C {
			if cacheData, err = conf.GetYaml(); err != nil {
				log.Println("err: ", err)
				continue
			}
		}
	}()

	// start gin server on config port
	service := gin.Default()
	service.Use(checkToken(config.C.App.Token))

	service.GET("/rules", func(c *gin.Context) {
		if len(cacheData) == 0 {
			c.String(http.StatusBadGateway, "no data")
			return
		}

		fileName := "MultiClash"
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Accept-Length", fmt.Sprintf("%d", len(cacheData)))
		c.Header("profile-update-interval", "1")
		c.Data(http.StatusOK, "application/text/plain", cacheData)
	})

	service.GET("/tunnels", func(c *gin.Context) {
		tpl, err := template.New("tunnels").Parse(tpl)
		if err != nil {
			log.Println("err: ", err)
			c.String(http.StatusBadGateway, err.Error())
			return
		}

		_ = tpl.Execute(c.Writer, config.C.Subscribes)
	})

	return service.Run(fmt.Sprintf(":%d", config.C.App.Port))
}

func checkToken(t string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := c.Query("token"); token != t {
			log.Println("need token: ", t, ", got: ", token)
			c.String(http.StatusUnauthorized, "token is invalid")
			c.Abort()
			return
		}

		c.Next()
	}
}
