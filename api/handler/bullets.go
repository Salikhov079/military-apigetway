package handler

import (
	"net/http"
	"strconv"

	pb "github.com/Salikhov079/military/genprotos/militaries"

	"github.com/gin-gonic/gin"
)

// CreateBullet handles the creation of a new Bullet
// @Summary      Create Bullet
// @Description  Create a new bullet
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Security  		BearerAuth
// @Param        BulletReq  body     pb.BulletReq  true  "Bullet Request"
// @Success      200        {string} string        "Create Successful"
// @Failure      401        {string} string        "Error while creating"
// @Router       /bullet/create [post]
func (h *Handler) CreateBullet(ctx *gin.Context) {
	var req pb.BulletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.BulletService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateBullet handles the updating of a Bullet
// @Summary      Update Bullet
// @Description  Update an existing bullet
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Security  		BearerAuth
// @Param        id      path    string   true  "Bullet ID"
// @Param        Bullet  body    pb.Bullet  true  "Bullet"
// @Success      200     {string} string  "Update Successful"
// @Failure      401     {string} string  "Error while updating"
// @Router       /bullet/update/{id} [put]
func (h *Handler) UpdateBullet(ctx *gin.Context) {
	var bullet pb.Bullet
	if err := ctx.ShouldBindJSON(&bullet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bullet.Id = ctx.Param("id")
	_, err := h.BulletService.Update(ctx, &bullet)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteBullet handles the deletion of a Bullet
// @Summary      Delete Bullet
// @Description  Delete an existing bullet
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id      path    string   true  "Bullet ID"
// @Success      200     {string} string  "Delete Successful"
// @Failure      401     {string} string  "Error while deleting"
// @Router       /bullet/delete/{id} [delete]
func (h *Handler) DeleteBullet(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.BulletService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetBullet handles getting a Bullet by ID
// @Summary      Get Bullet
// @Description  Get an existing bullet by ID
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id      path    string     true  "Bullet ID"
// @Success      200     {object} pb.Bullet "Get Successful"
// @Failure      401     {string} string    "Error while getting"
// @Router       /bullet/getbyid/{id} [get]
func (h *Handler) GetBullet(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.BulletService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllBullets handles getting all Bullets
// @Summary      Get All Bullets
// @Description  Get all bullets
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query   pb.BulletReq  true  "Query parameter"
// @Success      200    {object} pb.AllBullets "Get All Successful"
// @Failure      401    {string} string       "Error while getting all"
// @Router       /bullet/getall [get]
func (h *Handler) GetAllBullets(ctx *gin.Context) {
	cl := ctx.Query("caliber")
	qu := ctx.Query("quantity")
	ty := ctx.Query("type")
	c, _ := strconv.ParseFloat(cl, 32)
	q, _ := strconv.Atoi(qu)
	req := pb.BulletReq{Caliber: float32(c), Quantity: int32(q), Type: ty}

	res, err := h.BulletService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// Add handles adding quantity to a Bullet
// @Summary      Add Quantity
// @Description  Add quantity to a Bullet
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Bullet body pb.BulletAddSub true "Bullet data"
// @Success      200    {object} pb.Void "Add Successful"
// @Failure      500    {string} string  "Error while adding quantity"
// @Router       /bullet/add [put]
func (h *Handler) AddBullet(ctx *gin.Context) {
	var Bullet pb.BulletAddSub
	if err := ctx.ShouldBindJSON(&Bullet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.BulletService.Add(ctx, &Bullet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

// Sub handles subtracting quantity from a Bullet
// @Summary      Subtract Quantity
// @Description  Subtract quantity from a Bullet
// @Tags         Bullet
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Bullet body pb.BulletAddSub true "Bullet data"
// @Success      200    {object} pb.Void "Subtract Successful"
// @Failure      500    {string} string  "Error while subtracting quantity"
// @Router       /bullet/sub [put]
func (h *Handler) SubBullet(ctx *gin.Context) {
	var Bullet pb.BulletAddSub
	if err := ctx.ShouldBindJSON(&Bullet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.BulletService.Sub(ctx, &Bullet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}