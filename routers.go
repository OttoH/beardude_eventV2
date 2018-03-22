package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"qpet-engine/dao"
	"qpet-engine/handlers/event"
	"qpet-engine/handlers/racer"
)

func SetupRouter(r *gin.Engine, dao *dao.DAO, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// pass dao instance as dependency injection
	depInjEventHandler := &eventHandlers.DepInj{Dao: dao}
	depInjRacerHandler := &racerHandlers.DepInj{Dao: dao}

	// login
	r.POST("/login", authMiddleware.LoginHandler)

	//Racer
	// signup
	r.POST("/signup", depInjRacerHandler.CreateRacer)

	authRoute := r.Group("/racer")
	authRoute.Use(authMiddleware.MiddlewareFunc())
	{
		// test use
		authRoute.GET("/hello", helloHandler)
		authRoute.GET("/refresh_token", authMiddleware.RefreshHandler)

		// racer/get
		authRoute.POST("/get", depInjRacerHandler.GetRacer)
		// racer/update
		authRoute.PUT("/update", depInjRacerHandler.UpdateRacer)
		// racer/delete
		authRoute.DELETE("/delete", depInjRacerHandler.RemoveRacer)
	}

	// Events
	r.POST("/event/create", depInjEventHandler.CreateEvent)

	return r
}

// test use
func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}
