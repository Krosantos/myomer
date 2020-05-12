package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/krosantos/myomer/v2/elements"
)

func postTest(c *gin.Context) {
	tm := elements.TileMap{Tiles: make(map[int]map[int]*elements.Tile)}
	t0 := elements.Tile{
		X:       0,
		Y:       0,
		Z:       1,
		Unit:    nil,
		Corpse:  nil,
		Terrain: "grass",
		Map:     &tm,
	}
	t1 := elements.Tile{
		X:       0,
		Y:       1,
		Z:       1,
		Unit:    nil,
		Corpse:  nil,
		Terrain: "horses",
		Map:     &tm,
	}
	t2 := elements.Tile{
		X:       1,
		Y:       1,
		Z:       1,
		Unit:    nil,
		Corpse:  nil,
		Terrain: "water",
		Map:     &tm,
	}
	t3 := elements.Tile{
		X:       1,
		Y:       2,
		Z:       1,
		Unit:    nil,
		Corpse:  nil,
		Terrain: "horses",
		Map:     &tm,
	}
	tm.Set(&t0)
	tm.Set(&t1)
	tm.Set(&t2)
	tm.Set(&t3)

	unitJSON := `{
		"name": "Knight",
		"cost": 500,
		"color": "green",
		"strength": 3,
		"health": 4,
		"speed": 2,
		"moxie": 33,
		"attackRange": 1,
		"attackType": "basic",
		"moveType": "basic",
		"onAttack": [],
		"onDie": [],
		"onKill": [],
		"onMove": [
			"grassy"
		],
		"onRoundEnd": [],
		"onStrike": [],
		"onStruck": [],
		"activeAbilities": []
	}`
	unit := elements.BuildUnit(unitJSON, 0)
	unit.Tile = &t0
	t0.Unit = &unit
	canMoveT1 := unit.MoveIsValid(&t1)
	canMoveT2 := unit.MoveIsValid(&t2)
	canMoveT3 := unit.MoveIsValid(&t3)
	fmt.Println(canMoveT1, canMoveT2, canMoveT3)

	unit.Move(&t3)
	fmt.Println(t0.Terrain, t3.Terrain)

	c.Status(200)
}
