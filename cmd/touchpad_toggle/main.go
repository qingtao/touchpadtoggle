package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/qingtao/touchpadtoggle/touchpad"
)

func toggleStateOfTouchpad(ctx context.Context, id string) (exit int) {
	if id == "" {
		var err error
		id, err = touchpad.FindTouchpad(ctx)
		if err != nil {
			log.Printf("find touchpad error: %s", err)
			exit = 1
			return
		}
	}
	if err := touchpad.ToggleStateOfTouchpad(ctx, id); err != nil {
		log.Printf("touchpad toggle state error: %s ", err)
		exit = 2
	}
	return
}

var (
	debug      = flag.Bool("debug", false, "启用debug")
	touchpadId = flag.String("id", "", "指定touchpad的id而不是使用自动查询")
)

func main() {
	flag.Parse()
	touchpad.SetDebug(*debug)
	os.Exit(toggleStateOfTouchpad(context.Background(), *touchpadId))
}
