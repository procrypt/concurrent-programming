package main

import (
	"context"
"fmt"
	)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func userID(ctx context.Context) string {
	return ctx.Value(ctxUserID).(string)
}

func authToken(ctx context.Context) string  {
	return ctx.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleRequest(ctx)
}

func HandleRequest(ctx context.Context)  {
	fmt.Printf("Handling Response for %v (auth: %v)", userID(ctx), authToken(ctx))
}

func main() {
	ProcessRequest("Abhishek", "7827")
}