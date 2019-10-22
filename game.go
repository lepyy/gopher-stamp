package main

import (
	"mygame/systems"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/gopher.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(updater engo.Updater) {
	// common.SetBackground(color.White)
	engo.Input.RegisterButton("addPlayer", engo.KeyF1)

	world, _ := updater.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.PlayerSystem{})

	kbs := common.NewKeyboardScroller(
		200,
		engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis,
	)

	world.AddSystem(kbs)
	world.AddSystem(&common.MouseSystem{})

	// world.AddSystem(&common.EdgeScroller{ScrollSpeed: 400, EdgeMargin: 20})
	// world.AddSystem(&common.MouseZoomer{ZoomSpeed: -0.125})
}

func main() {
	opts := engo.RunOptions{
		Title:          "gopher stamp",
		Width:          640,
		Height:         480,
		StandardInputs: true,
	}
	engo.Run(opts, &myScene{})
}
