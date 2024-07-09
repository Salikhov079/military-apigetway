package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/military/api"
	"github.com/Salikhov079/military/api/handler"
	ai "github.com/Salikhov079/military/genprotos/ai"
	pb "github.com/Salikhov079/military/genprotos/militaries"
	pbs "github.com/Salikhov079/military/genprotos/soldiers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mil, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer mil.Close()

	sol, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":7070"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer sol.Close()

	a, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8086"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer a.Close()

	c := pb.NewBulletServiceClient(mil)
	ps := pb.NewFuelServiceClient(mil)
	ca := pb.NewTechniqueServiceClient(mil)
	el := pbs.NewCommanderServiceClient(sol)
	py := pbs.NewDepartmentServiceClient(sol)
	us := pbs.NewGroupServiceClient(sol)
	so := pbs.NewSoldierServiceClient(sol)
	ai := ai.NewAiServiceClient(a)

	h := handler.NewHandler(c, ps, ca, el, py, us, so, ai)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
