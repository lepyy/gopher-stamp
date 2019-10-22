package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type PlayerSystem struct {
	mouseTracker MouseTracker
	world        *ecs.World
}

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

func (ps *PlayerSystem) New(world *ecs.World) {
	ps.world = world
	fmt.Println("System was added to the Scene")

	ps.mouseTracker.BasicEntity = ecs.NewBasic()
	ps.mouseTracker.MouseComponent = common.MouseComponent{Track: true}
}

// Remove is called whenever an Entity is removed from the World
// in order to remove it from this sytem as well
func (*PlayerSystem) Remove(ecs.BasicEntity) {}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (ps *PlayerSystem) Update(dt float32) {

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&ps.mouseTracker.BasicEntity, &ps.mouseTracker.MouseComponent, nil, nil)
		}
	}

	if engo.Input.Button("addPlayer").JustPressed() {
		// fmt.Println("press F1")

		gopher := Player{BasicEntity: ecs.NewBasic()}
		gopher.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{
				X: ps.mouseTracker.MouseComponent.MouseX,
				Y: ps.mouseTracker.MouseComponent.MouseY,
			},
			Width:  30,
			Height: 64,
		}

		texture, err := common.LoadedSprite("textures/gopher.png")
		if err != nil {
			panic("Unable to load textures: " + err.Error())
		}

		gopher.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{X: 0.1, Y: 0.1},
		}

		for _, system := range ps.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&gopher.BasicEntity, &gopher.RenderComponent, &gopher.SpaceComponent)
			}
		}
	}
}
