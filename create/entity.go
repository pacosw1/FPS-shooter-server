package entity

import (
	"math"
	object "sockets/object"
	"sockets/validation"
)

//PlayerList initializes a map to store players
func PlayerList() map[int]*object.Player {
	return make(map[int]*object.Player, 100)
}

//ProjectileList initializes a map to store players
func ProjectileList() map[int]*object.Projectile {
	return make(map[int]*object.Projectile, 500)
}

//AddPlayer adds a player to the map table
func AddPlayer(x, y int, players map[int]*object.Player) {
	uniqueID := validation.ValidatePlayerID(100, players)

	players[uniqueID] = &object.Player{
		Health: 100,
		ID:     uniqueID,
		Position: &object.Position{
			X: x,
			Y: y,
		},
		SequenceID: 0,
	}
}

//RemoveProjectile player from List
func RemoveProjectile(ID int, projectiles map[int]*object.Projectile) {
	delete(projectiles, ID)
}

//RemovePlayer player from List
func RemovePlayer(ID int, players map[int]*object.Player) {
	delete(players, ID)
}

//AddProjectile adds a projectile to state
func AddProjectile(startCoord *object.Position, targetCoord *object.Position, projectiles map[int]*object.Projectile) {

	angleX, angleY := setAngle(startCoord, targetCoord)
	uniqueID := validation.ValidateProjectileID(5000, projectiles)
	projectiles[uniqueID] = &object.Projectile{
		Angle: &object.Angle{
			X: (angleX),
			Y: (angleY),
		},
		Position: startCoord,
		ID:       float32(uniqueID),
	}

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
