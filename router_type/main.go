package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get\n")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post\n")
	})

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete\n")
	})

	//[GIN-debug] GET    /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] POST   /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] PUT    /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] PATCH  /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] HEAD   /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] OPTIONS /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] DELETE /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] CONNECT /any                      --> main.main.func4 (3 handlers)
	//[GIN-debug] TRACE  /any                      --> main.main.func4 (3 handlers)

	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any request\n")
	})
	r.Run()
}
