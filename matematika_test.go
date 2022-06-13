package matematika

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCekGanjilGenap(t *testing.T) {
	tests := []struct {
		name     string
		expectes string
		param    []int
	}{
		{
			name:     "01",
			expectes: "Ganjil, Genap, Ganjil, Genap, Ganjil",
			param:    []int{1, 2, 3, 4, 5},
		},
		{
			name:     "02",
			expectes: "Genap, Ganjil, Genap, Ganjil, Genap",
			param:    []int{6, 7, 8, 9, 10},
		},
		{
			name:     "03",
			expectes: "Ganjil, Genap, Ganjil, Genap, Ganjil",
			param:    []int{11, 12, 13, 14, 15},
		},
	}
	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CekGanjilGenap(test.param...)
			assert.Equal(test.expectes, result)
		})
	}
}

func BenchmarkCekGanjilGenap(b *testing.B) {
	tests := []struct {
		name  string
		param []int
	}{
		{
			name:  "01",
			param: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "02",
			param: []int{6, 7, 8, 9, 10},
		},
		{
			name:  "03",
			param: []int{11, 12, 13, 14, 15},
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CekGanjilGenap(test.param...)
			}
		})
	}
}
