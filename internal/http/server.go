package http

import (
	"fmt"
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/source"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var tpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tunnels</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">
</head>
<body>
    <table class="table table-striped"
        <tr>
            <th>Name</th>
            <th>Traffic</th>
            <th>Expires</th>
        </tr>
        {{range .}}
        <tr>
            <th>{{.Name}}</th>
			<td>
				<div class="progress" role="progressbar">
					<div class="progress-bar" style="width: {{ .UserInfo.Progress }}%"></div>
				</div>
			</td>
            <td>{{.UserInfo.ExpireAt}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

var cacheData []byte

func Serve(config *config.Config) error {
	var err error
	if cacheData, err = getYaml(config); err != nil {
		log.Fatal(err)
	}

	go func() {
		ticker := time.NewTicker(time.Hour)
		for range ticker.C {
			if cacheData, err = getYaml(config); err != nil {
				log.Println("err: ", err)
				continue
			}
		}
	}()

	// start gin server on config port
	service := gin.Default()
	service.Use(checkToken(config.App.Token))

	service.GET("/rules", func(c *gin.Context) {
		if len(cacheData) == 0 {
			c.String(http.StatusBadGateway, "no data")
			return
		}

		fileName := "FullClash"
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Accept-Length", fmt.Sprintf("%d", len(cacheData)))
		c.Data(http.StatusOK, "application/text/plain", cacheData)
	})

	service.GET("/tunnels", func(c *gin.Context) {
		tpl, err := template.New("tunnels").Parse(tpl)
		if err != nil {
			log.Println("err: ", err)
			c.String(http.StatusBadGateway, err.Error())
			return
		}

		_ = tpl.Execute(c.Writer, config.Subscribes)
	})

	return service.Run(fmt.Sprintf(":%d", config.App.Port))
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

func getYaml(config *config.Config) ([]byte, error) {
	var eg errgroup.Group
	for _, s := range config.Subscribes {
		s := s
		eg.Go(func() error {
			log.Println("downloading source ", s.URL)
			if err := source.New(s).Parse(); err != nil {
				log.Println("err: ", err)
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	clash, err := config.Generate()
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	return clash.Bytes()
}
