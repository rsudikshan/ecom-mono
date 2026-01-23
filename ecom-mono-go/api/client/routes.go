package client

// import "github.com/gin-gonic/gin"

// type ClientRoutes struct {
// 	RG 			  *gin.RouterGroup
// 	clientHandler ClientHandler
// }

// func NewClientRoutes(rg *gin.RouterGroup, clientHandler ClientHandler) *ClientRoutes{
// 	return &ClientRoutes{
// 		RG: rg.Group("client-portal"),
// 		clientHandler: clientHandler,
// 	}
// }

// func (r *ClientRoutes) Setup() {
// 	r.RG.POST("register", r.clientHandler.RegisterUser)
// }