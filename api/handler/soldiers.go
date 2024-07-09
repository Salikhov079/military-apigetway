package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/gin-gonic/gin"
)


// CreateSoldier handles the creation of a new Soldier
// @Summary      Create Soldier
// @Description  Create a new soldier
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        SoldierReq  body     pb.CreateSoldier  true  "Soldier Request"
// @Success      200         {string} string         "Create Successful"
// @Failure      401         {string} string         "Error while creating"
// @Router       /soldier/create [post]
func (h *Handler) CreateSoldier(ctx *gin.Context) {
	var req pb.SoldierReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.SoldierService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateSoldier handles the updating of a Soldier
// @Summary      Update Soldier
// @Description  Update an existing soldier
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id       path     string     true  "Soldier ID"
// @Param        Soldier  body     pb.Soldier true  "Soldier"
// @Success      200      {string} string     "Update Successful"
// @Failure      401      {string} string     "Error while updating"
// @Router       /soldier/update/{id} [put]
func (h *Handler) UpdateSoldier(ctx *gin.Context) {
	var soldier pb.Soldier
	if err := ctx.ShouldBindJSON(&soldier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	soldier.Id = ctx.Param("id")
	_, err := h.SoldierService.Update(ctx, &soldier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteSoldier handles the deletion of a Soldier
// @Summary      Delete Soldier
// @Description  Delete an existing soldier
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id     path     string   true  "Soldier ID"
// @Success      200    {string} string  "Delete Successful"
// @Failure      401    {string} string  "Error while deleting"
// @Router       /soldier/delete/{id} [delete]
func (h *Handler) DeleteSoldier(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.SoldierService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetSoldier handles getting a Soldier by ID
// @Summary      Get Soldier
// @Description  Get an existing soldier by ID
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id     path     string     true  "Soldier ID"
// @Success      200    {object} pb.Soldier "Get Successful"
// @Failure      401    {string} string     "Error while getting"
// @Router       /soldier/get/{id} [get]
func (h *Handler) GetSoldier(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.SoldierService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllSoldiers handles getting all Soldiers
// @Summary      Get All Soldiers
// @Description  Get all soldiers
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query    pb.GetAllSoldierFilter  true  "Query parameter"
// @Success      200    {object} pb.AllSoldiers "Get All Successful"
// @Failure      401    {string} string         "Error while getting all"
// @Router       /soldier/getall [get]
func (h *Handler) GetAllSoldiers(ctx *gin.Context) {
	var req pb.SoldierReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.SoldierService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UseBullet handles the use of bullets by a soldier
// @Summary      Use Bullet
// @Description  Record the use of bullets by a soldier
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        UseB  body     pb.UseB  true  "Use Bullet"
// @Success      200   {string} string   "Use Bullet Successful"
// @Failure      401   {string} string   "Error while using bullet"
// @Router       /soldier/usebullet [post]
func (h *Handler) UseBullet(ctx *gin.Context) {
	var req pb.UseB
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.SoldierService.UseBullet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Use Bullet Successful")
}

// UseFuel handles the use of fuel by a soldier
// @Summary      Use Fuel
// @Description  Record the use of fuel by a soldier
// @Tags         Soldier
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        UseF  body     pb.UseF  true  "Use Fuel"
// @Success      200   {string} string   "Use Fuel Successful"
// @Failure      401   {string} string   "Error while using fuel"
// @Router       /soldier/usefuel [post]
func (h *Handler) UseFuel(ctx *gin.Context) {
	var req pb.UseF
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.SoldierService.UseFuel(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Use Fuel Successful")
}



// Dashbord handles getting all Dashbord
// @Summary      Get All Dashbord
// @Description  Get all Dashbord
// @Tags         Dashbord
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query    pb.SoldierReq  true  "Query parameter"
// @Success      200    {object} pb.AllSoldiers "Get All Successful"
// @Failure      401    {string} string         "Error while getting all"
// @Router       /soldier/dashbord [get]
func (h *Handler) Dashbord(ctx *gin.Context) {
	var req pb.SoldierReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.SoldierService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
