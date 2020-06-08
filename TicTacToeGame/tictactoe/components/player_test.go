package components

import "testing"

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		input1    string
		input2    string
		expected1 string
		expected2 string
	}{
		{"Ameya", OMark, "Ameya", OMark},
		{"Meghan", XMark, "Meghan", XMark},
	}
	for _, test := range tests {
		temp := NewPlayer(test.input1, test.input2)
		if temp.Name != test.expected1 || temp.Mark != test.expected2 {
			t.Error(temp, test.expected1, test.expected2)
		}
	}
}
