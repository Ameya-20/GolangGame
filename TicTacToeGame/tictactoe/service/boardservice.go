package service

import (
	"fmt"
	"github.com/Ameya-20/GolangGame/components"

)

type BoardService struct {
	*components.Board
}

func NewBoardService(board *components.Board) *BoardService {
	return &BoardService{board}
}

func (b *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	err := b.BoardCells[position].SetMark(player.Mark)
	return err
}

func (b *BoardService) PrintBoard() string {

	tempStr := ""
	for i := uint8(0); i < (b.Size * b.Size); i++ {
		if i%b.Size == 0 {
			tempStr = tempStr + fmt.Sprint("\n\t")
		}
		if i >= b.Size*(b.Size-1) {
			if i%b.Size == (b.Size - 1) {
				tempStr += fmt.Sprintf("     %s     ", b.BoardCells[i].GetMark())
			} else {
				tempStr += fmt.Sprintf("     %s     |", b.BoardCells[i].GetMark())
			}
		} else {
			if i%b.Size == (b.Size - 1) {
				tempStr += fmt.Sprintf("_____%s_____", b.BoardCells[i].GetMark())
			} else {
				tempStr += fmt.Sprintf("_____%s_____|", b.BoardCells[i].GetMark())
			}
		}
	}
	tempStr += fmt.Sprint("\n")
	return tempStr

	// for i := uint8(0); i < (b.Size * b.Size); i++ {
	// 	if i%b.Board.Size == 0 {
	// 		tempStr = tempStr + fmt.Sprint("\n\t")
	// 	}
	// 	tempStr = tempStr + fmt.Sprint(b.BoardCells[i].GetMark(), " ")
	// }
	// return tempStr
}

func (b *BoardService) CheckBoardIsFull() bool {
	for _, cell := range b.BoardCells {
		if cell.GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
