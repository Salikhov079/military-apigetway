package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/ai"

	"github.com/gin-gonic/gin"
)

// CHat handles the creation of a new Bullet
// @Summary      CHAT
// @Description  CHat with AI
// @Tags         AI
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        BulletReq  body     pb.AiCHat  true  "Bullet Request"
// @Success      200        {string} pb.AiCHat       
// @Failure      401        {string} string        "Error while creating"
// @Router       /ai/chat [post]
func (h *Handler) CHatAi(ctx *gin.Context) {
	var req pb.AiCHat
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Ai.CHat(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}


// CHat handles the creation of a new Bullet
// @Summary      GetHistory
// @Description  CHat with AI
// @Tags         AI
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id      path    string     true  "User ID"
// @Success      200        {string} pb.GetHistoryResponse       
// @Failure      401        {string} string        "Error while creating"
// @Router       /ai/gethistory/{id} [get]
func (h *Handler) GetHistory(ctx *gin.Context) {
	var req pb.GetHistoryRequest
	req.Id=ctx.Param("id")
	res, err := h.Ai.GetHistory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}