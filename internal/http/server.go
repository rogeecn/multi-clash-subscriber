package http

import (
	"fmt"
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/source"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve(config *config.Config) error {
	// start gin server on config port
	service := gin.Default()

	service.GET("/rules", func(c *gin.Context) {
		for _, s := range config.Subscribes {
			log.Println("downloading source ", s.URL)

			if err := source.New(s).Parse(); err != nil {
				log.Println("err: ", err)
				c.String(http.StatusBadGateway, "text/html;charset=UTF-8", err.Error())
				continue
			}
		}

		clash, err := config.Generate()
		if err != nil {
			log.Println("err: ", err)
			c.String(http.StatusBadGateway, "text/html;charset=UTF-8", err.Error())
			return
		}
		b, err := clash.Bytes()
		if err != nil {
			log.Println("err: ", err)
			c.String(http.StatusBadGateway, "text/html;charset=UTF-8", err.Error())
			return
		}

		fileName := "FullClash"
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Accept-Length", fmt.Sprintf("%d", len(b)))
		c.Data(http.StatusOK, "application/text/plain", b)
	})

	service.GET("/tunnels", func(c *gin.Context) {

	})

	return service.Run(fmt.Sprintf(":%d", config.App.Port))
}
