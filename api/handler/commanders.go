package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/gin-gonic/gin"
)



// CreateCommander handles the creation of a new Commander
// @Summary      Create Commander
// @Description  Create a new commander
// @Tags         Commander
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        CommanderReq  body     pb.CreateCommand  true  "Commander Request"
// @Success      200           {string} string           "Create Successful"
// @Failure      401           {string} string           "Error while creating"
// @Router       /commander/create [post]
func (h *Handler) CreateCommander(ctx *gin.Context) {
	var req pb.CommanderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.CommanderService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateCommander handles the updating of a Commander
// @Summary      Update Commander
// @Description  Update an existing commander
// @Tags         Commander
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id          path     string           true  "Commander ID"
// @Param        Commander   body     pb.Commander     true  "Commander"
// @Success      200         {string} string           "Update Successful"
// @Failure      401         {string} string           "Error while updating"
// @Router       /commander/update/{id} [put]
func (h *Handler) UpdateCommander(ctx *gin.Context) {
	var commander pb.Commander
	if err := ctx.ShouldBindJSON(&commander); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commander.Id = ctx.Param("id")
	_, err := h.CommanderService.Update(ctx, &commander)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteCommander handles the deletion of a Commander
// @Summary      Delete Commander
// @Description  Delete an existing commander
// @Tags         Commander
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id       path     string   true  "Commander ID"
// @Success      200      {string} string  "Delete Successful"
// @Failure      401      {string} string  "Error while deleting"
// @Router       /commander/delete/{id} [delete]
func (h *Handler) DeleteCommander(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.CommanderService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetCommander handles getting a Commander by ID
// @Summary      Get Commander
// @Description  Get an existing commander by ID
// @Tags         Commander
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id       path     string      true  "Commander ID"
// @Success      200      {object} pb.Commander "Get Successful"
// @Failure      401      {string} string       "Error while getting"
// @Router       /commander/get/{id} [get]
func (h *Handler) GetCommander(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.CommanderService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllCommanders handles getting all Commanders
// @Summary      Get All Commanders
// @Description  Get all commanders
// @Tags         Commander
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query    pb.GetAllFilter  true  "Query parameter"
// @Success      200    {object} pb.AllCommanders "Get All Successful"
// @Failure      401    {string} string           "Error while getting all"
// @Router       /commander/getall [get]
func (h *Handler) GetAllCommanders(ctx *gin.Context) {
	var req pb.CommanderReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CommanderService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
