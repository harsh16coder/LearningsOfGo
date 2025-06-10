package main

import "fmt"

type SpecialPostion struct {
	Position
}

func (s *SpecialPostion) MoveSpecial(x, y float64) {
	s.x += x * x
	s.y += y * y
}

type Player struct {
	*Position
}
type Enemy struct {
	*SpecialPostion
}
type Position struct {
	x float64
	y float64
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}
func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPostion: &SpecialPostion{},
	}
}
func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func main() {
	player := NewPlayer()
	player.Move(1.2, 3.5)
	fmt.Println("Player current position: ", player.x, player.y)
	player.Teleport(5.4, 7.8)
	fmt.Println("Player current position: ", player.x, player.y)
	player.Move(1.2, 3.5)
	fmt.Println("Player current position: ", player.x, player.y)
	//Enemy
	enemy := NewEnemy()
	enemy.Move(2.3, 5.4)
	enemy.MoveSpecial(4.5, 6.3)
	fmt.Println("New enemy position: ", enemy.x, enemy.y)
	fmt.Println("Player current position: ", player.x, player.y)

}
