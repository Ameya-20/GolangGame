package components

import (
	"errors"
)

type Cell struct {
	Mark string
}

const (
	NoMark = "-"
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{Mark: NoMark}
}

func (c *Cell) SetMark(mark string) error {
	if c.Mark == NoMark {
		c.Mark = mark
		return nil
	}
	return errors.New("cell is already marked")
}

func (c *Cell) GetMark() string {
	return c.Mark
}
