package handler

import (
	"net/http"
	"strconv"

	pb "github.com/Salikhov079/military/genprotos/militaries"

	"github.com/gin-gonic/gin"
)

// CreateFuel handles the creation of a new Fuel
// @Summary      Create Fuel
// @Description  Create a new fuel entry
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        FuelReq  body     pb.FuelReq  true  "Fuel Request"
// @Success      200      {string} string      "Create Successful"
// @Failure      400      {string} string      "Error while creating"
// @Router       /fuel/create [post]
func (h *Handler) CreateFuel(ctx *gin.Context) {
	var req pb.FuelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.FuelService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateFuel handles the updating of a Fuel
// @Summary      Update Fuel
// @Description  Update an existing fuel entry
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id    path     string    true  "Fuel ID"
// @Param        Fuel  body     pb.Fuel   true  "Fuel"
// @Success      200   {string} string    "Update Successful"
// @Failure      400   {string} string    "Error while updating"
// @Router       /fuel/update/{id} [put]
func (h *Handler) UpdateFuel(ctx *gin.Context) {
	var fuel pb.Fuel
	if err := ctx.ShouldBindJSON(&fuel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fuel.Id = ctx.Param("id")
	_, err := h.FuelService.Update(ctx, &fuel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteFuel handles the deletion of a Fuel
// @Summary      Delete Fuel
// @Description  Delete an existing fuel entry
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id    path     string    true  "Fuel ID"
// @Success      200   {string} string    "Delete Successful"
// @Failure      400   {string} string    "Error while deleting"
// @Router       /fuel/delete/{id} [delete]
func (h *Handler) DeleteFuel(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.FuelService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetFuel handles getting a Fuel by ID
// @Summary      Get Fuel
// @Description  Get an existing fuel entry by ID
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id    path     string    true  "Fuel ID"
// @Success      200   {object} pb.Fuel   "Get Successful"
// @Failure      400   {string} string    "Error while getting"
// @Router       /fuel/getbyid/{id} [get]
func (h *Handler) GetFuel(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.FuelService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllFuels handles getting all Fuels
// @Summary      Get All Fuels
// @Description  Get all fuel entries
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query    pb.FuelReq  true  "Query parameter"
// @Success      200    {object} pb.AllFuels "Get All Successful"
// @Failure      400    {string} string      "Error while getting all"
// @Router       /fuel/getall [get]
func (h *Handler) GetAllFuels(ctx *gin.Context) {
	qu := ctx.Query("quantity")
	ty := ctx.Query("type")
	q, _ := strconv.Atoi(qu)
	req := pb.FuelReq{Quantity: int32(q), Type: ty}
	res, err := h.FuelService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// Add handles adding quantity to a Fuel
// @Summary      Add Quantity
// @Description  Add quantity to a Fuel
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Fuel body pb.FuelAddSub true "Fuel data"
// @Success      200    {object} pb.Void "Add Successful"
// @Failure      500    {string} string  "Error while adding quantity"
// @Router       /fuel/add [put]
func (h *Handler) AddFuel(ctx *gin.Context) {
	var Fuel pb.FuelAddSub
	if err := ctx.ShouldBindJSON(&Fuel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.FuelService.Add(ctx, &Fuel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

// Sub handles subtracting quantity from a Fuel
// @Summary      Subtract Quantity
// @Description  Subtract quantity from a Fuel
// @Tags         Fuel
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Fuel body pb.FuelAddSub true "Fuel data"
// @Success      200    {object} pb.Void "Subtract Successful"
// @Failure      500    {string} string  "Error while subtracting quantity"
// @Router       /fuel/sub [put]
func (h *Handler) SubFuel(ctx *gin.Context) {
	var Fuel pb.FuelAddSub
	if err := ctx.ShouldBindJSON(&Fuel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.FuelService.Sub(ctx, &Fuel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}