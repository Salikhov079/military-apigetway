package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/militaries"

	"github.com/gin-gonic/gin"
)


// CreateTechnique handles the creation of a new Technique
// @Summary      Create Technique
// @Description  Create a new technique entry
// @Tags         Technique
// @Accept       json
// @Produce      json
// @Param        TechniqueReq  body     pb.TechniqueReq  true  "Technique Request"
// @Success      200           {string} string           "Create Successful"
// @Failure      400           {string} string           "Error while creating"
// @Router       /technique/create [post]
func (h *Handler) CreateTechnique(ctx *gin.Context) {
	var req pb.TechniqueReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.TechniqueService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateTechnique handles the updating of a Technique
// @Summary      Update Technique
// @Description  Update an existing technique entry
// @Tags         Technique
// @Accept       json
// @Produce      json
// @Param        id         path     string       true  "Technique ID"
// @Param        Technique  body     pb.Technique true  "Technique"
// @Success      200        {string} string       "Update Successful"
// @Failure      400        {string} string       "Error while updating"
// @Router       /technique/update/{id} [put]
func (h *Handler) UpdateTechnique(ctx *gin.Context) {
	var technique pb.Technique
	if err := ctx.ShouldBindJSON(&technique); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	technique.Id = ctx.Param("id")
	_, err := h.TechniqueService.Update(ctx, &technique)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteTechnique handles the deletion of a Technique
// @Summary      Delete Technique
// @Description  Delete an existing technique entry
// @Tags         Technique
// @Accept       json
// @Produce      json
// @Param        id    path     string    true  "Technique ID"
// @Success      200   {string} string    "Delete Successful"
// @Failure      400   {string} string    "Error while deleting"
// @Router       /technique/delete/{id} [delete]
func (h *Handler) DeleteTechnique(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.TechniqueService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetTechnique handles getting a Technique by ID
// @Summary      Get Technique
// @Description  Get an existing technique entry by ID
// @Tags         Technique
// @Accept       json
// @Produce      json
// @Param        id    path     string       true  "Technique ID"
// @Success      200   {object} pb.Technique "Get Successful"
// @Failure      400   {string} string       "Error while getting"
// @Router       /technique/get/{id} [get]
func (h *Handler) GetTechnique(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.TechniqueService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllTechniques handles getting all Techniques
// @Summary      Get All Techniques
// @Description  Get all technique entries
// @Tags         Technique
// @Accept       json
// @Produce      json
// @Param        query  query    pb.TechniqueReq true  "Query parameter"
// @Success      200    {object} pb.AllTechnique "Get All Successful"
// @Failure      400    {string} string          "Error while getting all"
// @Router       /technique/getall [get]
func (h *Handler) GetAllTechniques(ctx *gin.Context) {
	var req pb.TechniqueReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.TechniqueService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}