package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKeyType string

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userID string) {
	ctx = context.WithValue(ctx, userIDKeyType("userIDKey"), userID)
	ch := checkMembership(ctx)
	fmt.Printf("Membership status of userID : %s : %v\n", userID, <-ch)
}

func checkMembership(ctx context.Context) <-chan bool {

	ch := make(chan bool)

	go func() {
		defer close(ch)
		userID := ctx.Value(userIDKeyType("userIDKey")).(string)
		ch <- db[userID]
	}()

	return ch
}
