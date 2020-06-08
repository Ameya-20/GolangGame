package service

import (
	"github.com/Ameya-20/GolangGame/components"
)

type ResultService struct {
	*BoardService
}
type Result struct {
	CurrResult string
	Win        bool
	Draw       bool
}

func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (rs *ResultService) checkRows(mark string, pos uint8) bool {
	ret := true
	i := (pos - pos%rs.Size)
	//fmt.Printf("checking row %d", i/rs.Size+1)
	for j := i; j < (i + rs.Size); j++ {
		if rs.BoardCells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (rs *ResultService) checkColumns(mark string, pos uint8) bool {
	ret := true
	//fmt.Printf("checking column %d", pos%rs.Size+1)
	for j := pos % rs.Size; j < (rs.Size * rs.Size); j = j + rs.Size {
		if rs.BoardCells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}
func (rs *ResultService) checkDiagonal(mark string) bool {
	ret := true
	for i := uint8(0); i < rs.Size; i++ {
		if rs.BoardCells[rs.Size*i+i].GetMark() != mark {
			ret = false
		}
	}
	if ret {
		return ret
	}
	ret = true
	for i := uint8(0); i < rs.Size; i++ {
		if rs.BoardCells[(rs.Size*i)+(rs.Size-1-i)].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (rs *ResultService) PrintResult(resultStat string) string {
	retString := ""

	for i := 0; i < 3; i++ {
		retString += "\n\t"
		for j := 0; j < len(resultStat)+6; j++ {
			if i == 0 || i == 2 {
				retString += "-"
			}
		}
		if i == 1 {
			retString += "|  " + resultStat + "  |"
		}
	}
	retString += "\n"
	return retString

}

func (rs *ResultService) GetResult(player *components.Player, pos uint8) Result {

	if rs.checkRows(player.Mark, pos) || rs.checkColumns(player.Mark, pos) || rs.checkDiagonal(player.Mark) {
		return Result{rs.PrintResult(player.Name + " Won"), true, false}
	} else if rs.CheckBoardIsFull() {
		return Result{rs.PrintResult("Draw"), false, true}
	}
	return Result{"", false, false}
}
