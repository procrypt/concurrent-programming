package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("Abhishek", "1234")
}

func ProcessRequest(userID, authToken string)  {
	ctx := context.WithValue(context.Background(), "userID",userID)
	ctx  = context.WithValue(ctx,"authToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context)  {
	fmt.Printf("Handling response for %v (%v)",
		ctx.Value("userID"),
		ctx.Value("authToken"),
		)
}