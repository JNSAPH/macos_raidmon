package core

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)

func SetupSystray() {
	systray.Run(onReady, onExit)
}

func onExit() {
	fmt.Println("Exiting app...")
	os.Exit(0)
}
