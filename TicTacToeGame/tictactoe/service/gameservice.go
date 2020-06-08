package service

import (
	"github.com/Ameya-20/GolangGame/components"
	"fmt"
)

type GameService struct {
	*ResultService
	players [2]*components.Player
	turn    int
}

//var turn int = 0

func NewGameService(rs *ResultService, players [2]*components.Player) *GameService {
	return &GameService{rs, players, 0}

}

func (gs *GameService) Play(pos uint8) (Result, error) {
	if pos < 0 || pos >= gs.Size*gs.Size {
		return Result{"", false, false}, fmt.Errorf("Please enter the number in range (0,%d)", gs.Size*gs.Size-1)
	}
	var res Result
	if gs.turn%2 == 0 {
		err := gs.PutMarkInPosition(gs.players[0], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = gs.GetResult(gs.players[0], pos)
	} else if gs.turn%2 == 1 {
		err := gs.PutMarkInPosition(gs.players[1], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = gs.GetResult(gs.players[1], pos)
	}
	gs.turn++
	return res, nil
}
