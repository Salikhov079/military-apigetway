package api

import (
	"github.com/Salikhov079/military/api/handler"
	"github.com/Salikhov079/military/api/middleware"
	_ "github.com/Salikhov079/military/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// @tite Millitary service
// @version 1.0
// @description Millitary service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization

// NewGin sets up a new Gin router with Swagger API endpoints.
func NewGin(h *handler.Handler) *gin.Engine {


	r := gin.Default()
	r.Use(middleware.MiddleWare())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))


	techniques := r.Group("/technique")
	techniques.POST("/create", h.CreateTechnique)
	techniques.GET("/getall", h.GetAllTechniques)
	techniques.GET("/getbyid/:id", h.GetTechnique)
	techniques.PUT("/update/:id", h.UpdateTechnique)
	techniques.DELETE("/delete/:id", h.DeleteTechnique)
	techniques.PUT("/add", h.AddTechnique)
	techniques.PUT("/sub", h.SubTechnique)


	fuel := r.Group("/fuel")
	fuel.POST("/create", h.CreateFuel)
	fuel.GET("/getall", h.GetAllFuels)
	fuel.GET("/getbyid/:id", h.GetFuel)
	fuel.PUT("/update/:id", h.UpdateFuel)
	fuel.DELETE("/delete/:id", h.DeleteFuel)
	fuel.PUT("/add", h.AddFuel)
	fuel.PUT("/sub", h.SubFuel)


	soldiers := r.Group("/soldier")
	soldiers.POST("/create", h.CreateSoldier)
	soldiers.GET("/getall", h.GetAllSoldiers)
	soldiers.GET("/getbyid/:id", h.GetSoldier)
	soldiers.PUT("/update/:id", h.UpdateSoldier)
	soldiers.DELETE("/delete/:id", h.DeleteSoldier)
	soldiers.POST("/usebullet", h.UseBullet)
	soldiers.POST("/usefuel", h.UseFuel)
	
	soldiers.GET("/dashbord", h.Dashbord)
	soldiers.GET("/getallweaponstatistik", h.GetAllWeaponStatistik)
	soldiers.GET("/getallfuelstatistik", h.GetAllFuelStatistik)

	commanders := r.Group("/commander")
	commanders.POST("/create", h.CreateCommander)
	commanders.GET("/getall", h.GetAllCommanders)
	commanders.GET("/getbyid/:id", h.GetCommander)
	commanders.PUT("/update/:id", h.UpdateCommander)
	commanders.DELETE("/delete/:id", h.DeleteCommander)


	departments := r.Group("/department")
	departments.POST("/create", h.CreateDepartment)
	departments.GET("/getall", h.GetAllDepartments)
	departments.GET("/getbyid/:id", h.GetDepartment)
	departments.PUT("/update/:id", h.UpdateDepartment)
	departments.DELETE("/delete/:id", h.DeleteDepartment)

	groups := r.Group("/group")
	groups.POST("/create", h.CreateGroup)
	groups.GET("/getall", h.GetAllGroups)
	groups.GET("/getbyid/:id", h.GetGroup)
	groups.PUT("/update/:id", h.UpdateGroup)
	groups.DELETE("/delete/:id", h.DeleteGroup)

	
	bullets := r.Group("/bullet")
	bullets.POST("/create", h.CreateBullet)
	bullets.GET("/getall", h.GetAllBullets)
	bullets.GET("/getbyid/:id", h.GetBullet)
	bullets.PUT("/update/:id", h.UpdateBullet)
	bullets.DELETE("/delete/:id", h.DeleteBullet)
	bullets.PUT("/add", h.AddBullet)
	bullets.PUT("/sub", h.SubBullet)

	r.POST("/ai/chat", h.CHatAi)
	r.GET("/ai/gethistory/:id", h.GetHistory)

	return r
}
