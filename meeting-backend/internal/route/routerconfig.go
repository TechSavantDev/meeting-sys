package route

import "github.com/gin-gonic/gin"

var Routers []Router = []Router{
	{
		groupPath: "/api",
		routes: []Router{
			{
				method: GET,
				handler: func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "Hello, World!",
					})
				},
				path: "/hello",
			},
		},
	},
}
