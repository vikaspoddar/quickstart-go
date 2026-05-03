package contexts

import (
	"context"
	"fmt"
)

type contextKey string

const currentContextKey contextKey = "key"

func doSomething(ctx context.Context) {
	fmt.Printf("do something with the key's value: %s\n", ctx.Value("key"))

	anotherCtx := context.WithValue(ctx, currentContextKey, "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("do something with the key's value: %s\n", ctx.Value("key"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("do another with the key's value: %s\n", ctx.Value("key"))
}
func ContextDemo() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, currentContextKey, "value")
	doSomething(ctx)
}
