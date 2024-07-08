package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/gin-gonic/gin"
)

// CreateGroup handles the creation of a new Group
// @Summary      Create Group
// @Description  Create a new group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        GroupReq  body     pb.GroupReq  true  "Group Request"
// @Success      200       {string} string       "Create Successful"
// @Failure      401       {string} string       "Error while creating"
// @Router       /group/create [post]
func (h *Handler) CreateGroup(ctx *gin.Context) {
	var req pb.GroupReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.GroupService.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateGroup handles the updating of a Group
// @Summary      Update Group
// @Description  Update an existing group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        id       path     string     true  "Group ID"
// @Param        Group    body     pb.Group   true  "Group"
// @Success      200      {string} string     "Update Successful"
// @Failure      401      {string} string     "Error while updating"
// @Router       /group/update/{id} [put]
func (h *Handler) UpdateGroup(ctx *gin.Context) {
	var group pb.Group
	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group.Id = ctx.Param("id")
	_, err := h.GroupService.Update(ctx, &group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteGroup handles the deletion of a Group
// @Summary      Delete Group
// @Description  Delete an existing group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        id     path     string   true  "Group ID"
// @Success      200    {string} string  "Delete Successful"
// @Failure      401    {string} string  "Error while deleting"
// @Router       /group/delete/{id} [delete]
func (h *Handler) DeleteGroup(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.GroupService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetGroup handles getting a Group by ID
// @Summary      Get Group
// @Description  Get an existing group by ID
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        id     path     string     true  "Group ID"
// @Success      200    {object} pb.Group  "Get Successful"
// @Failure      401    {string} string    "Error while getting"
// @Router       /group/get/{id} [get]
func (h *Handler) GetGroup(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.GroupService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllGroups handles getting all Groups
// @Summary      Get All Groups
// @Description  Get all groups
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        query  query    pb.GetAllDepartmentFilter  true  "Query parameter"
// @Success      200    {object} pb.AllGroups "Get All Successful"
// @Failure      401    {string} string       "Error while getting all"
// @Router       /group/getall [get]
func (h *Handler) GetAllGroups(ctx *gin.Context) {
	var req pb.GroupReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.GroupService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
