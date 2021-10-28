package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapByAnagrams(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input []string
		Expect map[string][]string
	}{
		{
			Description: "Sample One",
			Input: []string{
				"eat", "tea", "tan", "ate", "nat", "bat", "teae", "teaa",
			},
			Expect: map[string][]string{
				"aet":{"eat","tea", "ate"},
				"ant":{"tan", "nat"},
				"abt":{"bat"},
				"aeet":{"teae"},
				"aaet":{"teaa"},
			},
		},
		{
			Description: "Sample Two",
			Input: []string{
				"abc", "cafe", "face", "cab",  "goo",
			},
			Expect: map[string][]string{
				"abc":{"abc", "cab"},
				"acef":{"cafe", "face"},
				"goo":{"goo"},
			},
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := MapByAnagrams(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}
