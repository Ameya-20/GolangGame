package components

import (
	"errors"
	"testing"
)

func TestNewCell(t *testing.T) {
	test := NewCell()
	expected := NoMark
	if test.Mark != expected {
		t.Error(test.Mark, expected)
	}
}

func TestSetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		input2   string
		expected error
	}{
		{&Cell{Mark: NoMark}, XMark, nil},
		{&Cell{Mark: XMark}, OMark, errors.New("Cell is already marked")},
		{&Cell{Mark: OMark}, NoMark, errors.New("Cell is already marked")},
	}
	for _, test := range tests {
		actual := test.input1.SetMark(test.input2)
		if actual != nil {
			if actual.Error() != test.expected.Error() {
				t.Error(actual, test.expected)
			}
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		expected string
	}{
		{&Cell{Mark: NoMark}, NoMark},
		{&Cell{Mark: OMark}, OMark},
		{&Cell{Mark: XMark}, XMark},
	}
	for _, test := range tests {
		actual := test.input1.GetMark()
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}

}
