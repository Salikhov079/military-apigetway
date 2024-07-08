package api

import (
	"github.com/Salikhov079/military/api/handler"
	_ "github.com/Salikhov079/military/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewGin sets up a new Gin router with Swagger API endpoints.
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Swagger API documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// Techniques API
	techniques := r.Group("/technique")
	techniques.POST("/create", h.CreateTechnique)
	techniques.GET("/getall", h.GetAllTechniques)
	techniques.GET("/getbyid/:id", h.GetTechnique)
	techniques.PUT("/update/:id", h.UpdateTechnique)
	techniques.DELETE("/delete/:id", h.DeleteTechnique)

	// Fuel API
	fuel := r.Group("/fuel")
	fuel.POST("/create", h.CreateFuel)
	fuel.GET("/getall", h.GetAllFuels)
	fuel.GET("/getbyid/:id", h.GetFuel)
	fuel.PUT("/update/:id", h.UpdateFuel)
	fuel.DELETE("/delete/:id", h.DeleteFuel)

	// Soldiers API
	soldiers := r.Group("/soldier")
	soldiers.POST("/create", h.CreateSoldier)
	soldiers.GET("/getall", h.GetAllSoldiers)
	soldiers.GET("/getbyid/:id", h.GetSoldier)
	soldiers.PUT("/update/:id", h.UpdateSoldier)
	soldiers.DELETE("/delete/:id", h.DeleteSoldier)
	soldiers.POST("/usebullet", h.UseBullet)
	soldiers.POST("/usefuel", h.UseFuel)
	
	soldiers.GET("/dashbord", h.Dashbord)
	// Commanders API
	commanders := r.Group("/commander")
	commanders.POST("/create", h.CreateCommander)
	commanders.GET("/getall", h.GetAllCommanders)
	commanders.GET("/getbyid/:id", h.GetCommander)
	commanders.PUT("/update/:id", h.UpdateCommander)
	commanders.DELETE("/delete/:id", h.DeleteCommander)

	// Departments API
	departments := r.Group("/department")
	departments.POST("/create", h.CreateDepartment)
	departments.GET("/getall", h.GetAllDepartments)
	departments.GET("/getbyid/:id", h.GetDepartment)
	departments.PUT("/update/:id", h.UpdateDepartment)
	departments.DELETE("/delete/:id", h.DeleteDepartment)

	// Groups API
	groups := r.Group("/group")
	groups.POST("/create", h.CreateGroup)
	groups.GET("/getall", h.GetAllGroups)
	groups.GET("/getbyid/:id", h.GetGroup)
	groups.PUT("/update/:id", h.UpdateGroup)
	groups.DELETE("/delete/:id", h.DeleteGroup)

	// Bullets API
	bullets := r.Group("/bullet")
	bullets.POST("/create", h.CreateBullet)
	bullets.GET("/getall", h.GetAllBullets)
	bullets.GET("/getbyid/:id", h.GetBullet)
	bullets.PUT("/update/:id", h.UpdateBullet)
	bullets.DELETE("/delete/:id", h.DeleteBullet)

	r.POST("/ai/chat", h.CHatAi)
	r.GET("/ai/gethistory/:id", h.GetHistory)

	return r
}
