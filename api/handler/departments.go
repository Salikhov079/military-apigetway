package handler

import (
	"net/http"
	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/gin-gonic/gin"
)



// CreateDepartment handles the creation of a new Department
// @Summary      Create Department
// @Description  Create a new department
// @Tags         Department
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        Department  body     pb.CreateDeportment  true  "Department"
// @Success      200         {string} string         "Create Successful"
// @Failure      401         {string} string         "Error while creating"
// @Router       /department/create [post]
func (h *Handler) CreateDepartment(ctx *gin.Context) {
	var dept pb.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.DepartmentService.Create(ctx, &dept)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Create Successful")
}

// UpdateDepartment handles the updating of a Department
// @Summary      Update Department
// @Description  Update an existing department
// @Tags         Department
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id         path     string         true  "Department ID"
// @Param        Department body     pb.Department  true  "Department"
// @Success      200        {string} string         "Update Successful"
// @Failure      401        {string} string         "Error while updating"
// @Router       /department/update/{id} [put]
func (h *Handler) UpdateDepartment(ctx *gin.Context) {
	var dept pb.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dept.Id = ctx.Param("id")
	_, err := h.DepartmentService.Update(ctx, &dept)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update Successful")
}

// DeleteDepartment handles the deletion of a Department
// @Summary      Delete Department
// @Description  Delete an existing department
// @Tags         Department
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id     path     string   true  "Department ID"
// @Success      200    {string} string  "Delete Successful"
// @Failure      401    {string} string  "Error while deleting"
// @Router       /department/delete/{id} [delete]
func (h *Handler) DeleteDepartment(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.DepartmentService.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Delete Successful")
}

// GetDepartment handles getting a Department by ID
// @Summary      Get Department
// @Description  Get an existing department by ID
// @Tags         Department
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        id     path     string      true  "Department ID"
// @Success      200    {object} pb.Department "Get Successful"
// @Failure      401    {string} string       "Error while getting"
// @Router       /department/get/{id} [get]
func (h *Handler) GetDepartment(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.DepartmentService.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// GetAllDepartments handles getting all Departments
// @Summary      Get All Departments
// @Description  Get all departments
// @Tags         Department
// @Accept       json
// @Produce      json
// @Security  		BearerAuth
// @Param        query  query    pb.GetAllDepartmentFilter  true  "Query parameter"
// @Success      200    {object} pb.AllDepartments "Get All Successful"
// @Failure      401    {string} string           "Error while getting all"
// @Router       /department/getall [get]
func (h *Handler) GetAllDepartments(ctx *gin.Context) {
	name := ctx.Query("name")
	req := pb.Department{Name: name}
	res, err := h.DepartmentService.GetAll(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
