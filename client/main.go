package main

import (
	"context"
	"log"

	v1 "github.com/atom-apps/dictionary/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// dial a grpc client connect to localhost:9800
	conn, err := grpc.DialContext(context.Background(), "localhost:61996", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := v1.NewDictionaryServiceClient(conn)
	out, err := client.GetItems(context.Background(), &v1.GetItemsRequest{
		Language: "zhch",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out.Items)
}
