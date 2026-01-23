package infrastructure

import (
	"log"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	engine *gin.Engine
	RG 	   *gin.RouterGroup
	Env    *Env
}

func NewAppRouter(env *Env) *AppRouter{
	engine := gin.New()
	engine.GET("/health-check", func(ctx *gin.Context) {
		ctx.Writer.WriteString("high and flying")
	})
	return &AppRouter{
		engine: engine,
		RG:		engine.Group("api") ,	
		Env: 	env,
	}
}

func (r *AppRouter) Start(){
	err := r.engine.Run(r.Env.SERVER_PORT)
	if err!=nil {
		log.Fatal(err.Error())
	}
}