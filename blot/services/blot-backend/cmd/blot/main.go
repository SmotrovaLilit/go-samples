package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gosamples/blot/internal/gen/blotservice/v1beta1"
	"gosamples/blot/internal/transport/grpc/blotservice"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterBlotServiceServer(grpcServer, blotservice.NewBlotServer())
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("net listen error: %v", err)
	}
	log.Println("Server is running  8081...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer.Serve error: %v", err)
	}

	//player1 := player.NewPlayer("Alice")
	//player2 := player.NewPlayer("Bob")
	//player3 := player.NewPlayer("Charlie")
	//player4 := player.NewPlayer("David")

	//team1 := team.NewTeam("Team 1", player1, player2)
	//team2 := team.NewTeam("Team 2", player3, player4)
	//players := []player.Player{player1, player2, player3, player4}

}
