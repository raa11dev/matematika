package matematika

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCekGanjilGenap(t *testing.T) {
	tests := []struct {
		name     string
		expectes string
		param    int
	}{
		{
			name:     "01",
			expectes: "Ganjil",
			param:    1,
		},
		{
			name:     "02",
			expectes: "Genap",
			param:    2,
		},
		{
			name:     "03",
			expectes: "Ganjil",
			param:    3,
		},
	}
	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(test.expectes, CekGanjilGenap(test.param))
		})
	}
}

func BenchmarkCekGanjilGenap(b *testing.B) {
	tests := []struct {
		name  string
		param int
	}{
		{
			name:  "01",
			param: 1,
		},
		{
			name:  "02",
			param: 2,
		},
		{
			name:  "03",
			param: 3,
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CekGanjilGenap(test.param)
			}
		})
	}
}
