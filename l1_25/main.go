package main

import (
	"context"
	"fmt"
	"time"
)

// custom time sleep

func Sleep(ms int) {
	// можно так же <-time.After

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(ms))
	defer cancel()
	<-ctx.Done()

}

func main() {

	fmt.Println("strating")
	Sleep(2000)
	fmt.Println("done")

}
