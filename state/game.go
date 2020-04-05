package state

import (
	"math"
	object "sockets/object"
	"sockets/validation"
)

//PlayerList initializes a map to store players

//GameState t
type GameState struct {
	Players     map[int]*object.Player
	Projectiles map[int]*object.Projectile
}

//ProjectileList initializes a map to store players

//AddPlayer adds a player to the map table
func AddPlayer(player *object.Player, state *GameState) {
	uniqueID := validation.ValidatePlayerID(100, state.Players)

	state.Players[uniqueID] = player

}

//UpdatePlayer update the players state in the game
func UpdatePlayer(ID int, state *GameState, newInput *object.UserInput) {
	player := state.Players[ID]
	player.IsShooting = newInput.IsShooting
}

//RemoveProjectile player from List
func RemoveProjectile(ID int, state *GameState) {
	delete(state.Projectiles, ID)
}

//RemovePlayer player from List
func RemovePlayer(ID int, state *GameState) {
	delete(state.Players, ID)
}

//AddProjectile adds a projectile to state
func AddProjectile(projectile *object.Projectile, state *GameState) {
	uniqueID := validation.ValidateProjectileID(5000, state.Projectiles)

	state.Projectiles[uniqueID] = projectile

	// angleX, angleY := setAngle(startCoord, targetCoord)
	// projectiles[uniqueID] = &object.Projectile{
	// 	Angle: &object.Angle{
	// 		X: (angleX),
	// 		Y: (angleY),
	// 	},
	// 	Position: startCoord,
	// 	ID:       float32(uniqueID),
	// }

}

//sets the angle for the projectile path
func setAngle(startCoord, targetCoord *object.Position) (float64, float64) {

	initX, initY := startCoord.X, startCoord.Y
	targetX, targetY := targetCoord.X, targetCoord.Y

	deltaX, deltaY := float64(targetX-initX), float64(targetY-initY)
	angle := math.Atan2(deltaY, deltaX)

	x, y := math.Cos(angle), math.Sin(angle)
	return x, y
}

//makes sure id is unique
