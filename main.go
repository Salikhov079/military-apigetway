package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/military/api"
	"github.com/Salikhov079/military/api/handler"
	pb "github.com/Salikhov079/military/genprotos/militaries"
	pbs "github.com/Salikhov079/military/genprotos/soldiers"
	ai "github.com/Salikhov079/military/genprotos/ai"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":7070"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	auth, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer auth.Close()
	c := pb.NewBulletServiceClient(UserConn)
	ps := pb.NewFuelServiceClient(UserConn)
	ca := pb.NewTechniqueServiceClient(UserConn)
	el := pbs.NewCommanderServiceClient(auth)
	py := pbs.NewDepartmentServiceClient(auth)
	us:=pbs.NewGroupServiceClient(auth)
	so:=pbs.NewSoldierServiceClient(auth)
	ai:=ai.NewAiServiceClient(auth)

	h := handler.NewHandler(c, ps, ca, el, py, us, so, ai)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
