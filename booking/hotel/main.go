package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background();
	ctx, cancel := context.WithTimeout(ctx, time.Second * 4);
	defer cancel();

	// go func() {
	// 	time.Sleep(time.Second * 3);
	// 	cancel();
	// }()

	bookHotel(ctx);
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel is full");
	case <-time.After(time.Second * 5):
		fmt.Println("Room reserved with sucess");
	}
}