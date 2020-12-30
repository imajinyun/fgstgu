package http

import (
	pb "kratos-demo/api"
	"kratos-demo/internal/model"
	"net/http"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/kratos-demo")
	{
		g.GET("/start", howToStart)
		g.GET("/hello", helloToHandler)
		g.GET("/param1/:name", pathParam)
		g.GET("/param2/:name/:value/:field", pathParam)
		g.GET("/param3/:name/*action", pathParam)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

func articleToShow(c *bm.Context) {
	k := &model.Article{
		ID:      100,
		Content: "this is a test article",
		Author:  "kratos",
	}
	c.JSON(k, nil)
}

func helloToHandler(c *bm.Context) {
	k := &model.Hello{
		Message: "🎉 Hello World, Hello Go❗️",
	}
	c.JSON(k, nil)
}

func pathParam(c *bm.Context) {
	name, _ := c.Params.Get("name")
	value, _ := c.Params.Get("value")
	field, _ := c.Params.Get("field")
	action, _ := c.Params.Get("action")
	path := c.RoutePath
	c.JSONMap(map[string]interface{}{
		"name":   name,
		"value":  value,
		"field":  field,
		"action": action,
		"path":   path,
	}, nil)
}
