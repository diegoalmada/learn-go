package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "xyz")
	bookingHotel(ctx)
}

// Por convencao , o contexto é o primeiro parametro.
func bookingHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Print(token)
}
