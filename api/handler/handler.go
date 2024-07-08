package handler

import (
	pb "github.com/Salikhov079/military/genprotos/militaries"
	pbs "github.com/Salikhov079/military/genprotos/soldiers"
	ai  "github.com/Salikhov079/military/genprotos/ai"
)


type Handler struct {
	BulletService pb.BulletServiceClient
	FuelService pb.FuelServiceClient
	TechniqueService pb.TechniqueServiceClient
	CommanderService pbs.CommanderServiceClient
	DepartmentService pbs.DepartmentServiceClient
	GroupService pbs.GroupServiceClient
	SoldierService pbs.SoldierServiceClient
	Ai  ai.AiServiceClient


}

func NewHandler(bu pb.BulletServiceClient, fu pb.FuelServiceClient,
	te pb.TechniqueServiceClient, co pbs.CommanderServiceClient, de pbs.DepartmentServiceClient, 
	ge pbs.GroupServiceClient, so pbs.SoldierServiceClient, ai ai.AiServiceClient) *Handler {
	return &Handler{bu, fu, te, co, de, ge, so, ai}
}
