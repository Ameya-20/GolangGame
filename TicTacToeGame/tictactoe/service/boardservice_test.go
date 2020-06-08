package service

import (
	"errors"
	"testing"
	"github.com/Ameya-20/GolangGame/components"
)

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		index    uint8
		player   *components.Player
		expected error
	}{
		{&BoardService{components.CreateBoard(3)}, 0, &components.Player{Name: "Ameya", Mark: components.OMark}, nil},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		}, 0, &components.Player{Name: "Meghan", Mark: components.XMark}, errors.New("Cell marked already!!")},
	}
	for _, test := range tests {
		actual := test.input1.PutMarkInPosition(test.player, test.index)
		if actual != nil {
			if actual.Error() != test.expected.Error() {
				t.Error(actual, test.expected)
			}
		}
	}
}

func TestPrintBoard(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		expected string
	}{
		{&BoardService{components.CreateBoard(3)}, "\n\t_____-_____|_____-_____|_____-_____\n\n\t_____-_____|_____-_____|_____-_____\n\n\t     -     |     -     |     -     \n"},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		}, "\n\t_____O_____|_____X_____\n\t     -     |     X     \n"},
	}
	for _, test := range tests {
		actual := test.input1.PrintBoard()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	tests := []struct {
		input1   *BoardService
		expected bool
	}{
		{&BoardService{components.CreateBoard(3)}, false},
		{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
		},
		}, true},
	}
	for _, test := range tests {
		actual := test.input1.CheckBoardIsFull()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}
