package main

import "testing"

// Функція, яку ми тестуємо
func plus(x, y int) int {
	return x + y
}

func TestPlus(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"додаємо позитивні числа", 2, 3, 5},
		{"додаємо нуль", 0, 5, 5},
		{"додаємо від'ємне число", -2, 3, 1},
		{"додаємо два від'ємні числа", -3, -4, -7},
		{"додаємо нулі", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := plus(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("plus(%d, %d) = %d; want %d", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}