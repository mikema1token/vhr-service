package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vhr-service/service"
)

func main() {
	//engine := gin.Default()
	////registerRouter(engine)
	////GinAsciiJSONFeature(engine)
	////GinHtmlTemplate(engine)
	////GinHttpPush(engine)
	////GinJsonP(engine)
	////GinFormBind(engine)
	////GinQuery(engine)
	////SecureJson(engine)
	////GinXmlYaml(engine)
	//GinReader(engine)
	//engine.Run()
	ShutdownGrace()
}

func registerRouter(engine *gin.Engine) {
	engine.POST("login", service.UserLogin)
}

func GinAsciiJSONFeature(engine *gin.Engine) {
	engine.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语音",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	engine.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{"data": "<b>this is b</b>"})
	})
}
func GinHtmlTemplate(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
		})
	})
}

func GinHttpPush(engine *gin.Engine) {
	var html = template.Must(template.New("https").Parse(`
			<html>
			<head>
			  <title>Https Test</title>
			  <script src="/assets/app.js"></script>
			</head>
			<body>
			  <h1 style="color:red;">Welcome, Ginner!</h1>
			</body>
			</html>
			`))
	engine.Static("/assets", "./assets")
	engine.SetHTMLTemplate(html)
	engine.GET("/", func(context *gin.Context) {
		if pusher := context.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		context.HTML(http.StatusOK, "https", gin.H{"status": "success"})
	})
}

func GinJsonP(r *gin.Engine) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, data)
	})
}

func GinFormBind(r *gin.Engine) {
	type LoginForm struct {
		User     string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	r.POST("/login2", func(c *gin.Context) {
		var f LoginForm
		if err := c.ShouldBindWith(&f, binding.Form); err == nil {
			if f.User == "user" && f.Password == "password" {
				c.JSON(http.StatusOK, gin.H{"code": "ok"})
			} else {
				c.JSON(401, gin.H{"code": "bad"})
			}
		}
	})
	r.POST("/login3", func(c *gin.Context) {
		user := c.PostForm("user")
		password := c.PostForm("password")
		if user == "user" && password == "password" {
			c.JSON(200, gin.H{"code": "ok"})
		} else {
			c.JSON(401, gin.H{"code": "bad"})
		}
	})
}

//func GinQuery(r *gin.Engine) {
//	r.GET("/query", func(c *gin.Context) {
//		value := c.Query("id")
//		c.JSON(200, gin.H{"id": value})
//	})
//}

//func SecureJson(engine *gin.Engine) {
//	engine.GET("/secure", func(c *gin.Context) {
//		c.SecureJSON(200, []string{"tony", "angle"})
//	})
//}

func GinXmlYaml(r *gin.Engine) {
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"data": "ok"})
	})
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"data": "bad"})
	})
}

func GinReader(r *gin.Engine) {
	r.GET("/reader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != 200 {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		length := response.ContentLength
		cttType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, length, cttType, reader, extraHeaders)
	})
}

func ShutdownGrace() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(200, "Welcome Gin Server")
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
	defer cf()
	err := srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
