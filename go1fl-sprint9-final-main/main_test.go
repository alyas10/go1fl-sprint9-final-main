package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    []int
		expectedValue int
	}{
		{
			name:          "Пустой слайс",
			inputValue:    []int{},
			expectedValue: 0,
		},
		{
			name:          "Слайс с одним элементом",
			inputValue:    []int{100},
			expectedValue: 100,
		},
		{
			name:          "Одинаковые элементы",
			inputValue:    []int{1, 1, 1, 1, 1},
			expectedValue: 1,
		},
		{
			name:          "Отрицательные числа",
			inputValue:    []int{-100, -10, -256},
			expectedValue: -10,
		},
		{
			name:          "Обычный слайс",
			inputValue:    []int{1, 2, 3, 4, 5},
			expectedValue: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Используем assert.Equal, так как типы совпадают (int)
			assert.Equal(t, test.expectedValue, maximum(test.inputValue))
		})
	}
}

func TestGenerateRandomElements(t *testing.T) {
	t.Run("Размер 0", func(t *testing.T) {
		result := generateRandomElements(0)
		assert.Nil(t, result, "Для размера 0 должен возвращаться nil")
	})

	t.Run("Отрицательный размер", func(t *testing.T) {
		result := generateRandomElements(-5)
		assert.Nil(t, result, "Для отрицательного размера должен возвращаться nil")
	})

	t.Run("Корректный размер", func(t *testing.T) {
		size := 100
		result := generateRandomElements(size)
		assert.NotNil(t, result, "Слайс не должен быть nil")
		assert.Len(t, result, size, "Длина слайса должна соответствовать запрошенному размеру")

		for _, val := range result {
			assert.Greater(t, val, 0, "Все сгенерированные числа должны быть строго больше 0")
		}
	})
}
