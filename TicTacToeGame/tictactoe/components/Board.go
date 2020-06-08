package components

type Board struct {
	Size       uint8
	BoardCells []*Cell
}

func CreateBoard(size uint8) *Board {
	tempcells := make([]*Cell, size*size)
	for i := 0; i < int(size*size); i++ {
		tempcells[i] = NewCell()
	}
	return &Board{
		Size:       size,
		BoardCells: tempcells,
	}
}

// func SetBoardCell(index uint8, mark string) {

// }
