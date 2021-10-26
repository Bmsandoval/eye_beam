package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextBigger(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input int
		Expect int
	}{
		{
			Description: "Happy Path",
			Input: 5,
			Expect: 5,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := Do(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result int
func benchmarkDo(i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Do(i)
	}
	result = r
}

func BenchmarkDo10x3(b *testing.B)   { benchmarkDo(1000, b) }
func BenchmarkDo10x6(b *testing.B)   { benchmarkDo(1000000, b) }
func BenchmarkDo10x9(b *testing.B)   { benchmarkDo(1000000000, b) }
func BenchmarkDo10x12(b *testing.B)  { benchmarkDo(1000000000000, b) }
func BenchmarkDo10x15(b *testing.B)  { benchmarkDo(1000000000000000, b) }
func BenchmarkDo10x16(b *testing.B)  { benchmarkDo(1000000000000000000, b) }
