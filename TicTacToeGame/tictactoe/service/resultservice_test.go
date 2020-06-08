package service

import (
	"testing"
	"github.com/Ameya-20/GolangGame/components"

)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		input    *ResultService
		inp1     string
		inp2     uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, 3, false},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 4, true},
	}

	for _, test := range tests {
		if test.input.checkRows(test.inp1, test.inp2) != test.expected {
			t.Error("Row check failed!!")
		}
	}
}

func TestCheckColumns(t *testing.T) {
	tests := []struct {
		input    *ResultService
		inp1     string
		inp2     uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, 3, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 5, false},
	}

	for _, test := range tests {
		if test.input.checkColumns(test.inp1, test.inp2) != test.expected {
			t.Error("Column check failed!!")
		}
	}
}

func TestCheckDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		inp1     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			BoardCells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		if test.input.checkDiagonal(test.inp1) != test.expected {
			t.Error(test, "Diagonal check failed!!")
		}
	}
}
