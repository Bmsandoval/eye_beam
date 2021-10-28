package main

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


// GroupByAnagrams groups by anagrams an array of strings into arrays of strings
func GroupByAnagrams(receivedItems []string) [][]string {
	mappedItems := MapByAnagrams(receivedItems)
	// Flatten map's values
	var groupedItems [][]string
	for _,mappedItem := range mappedItems {
		groupedItems = append(groupedItems, mappedItem)
	}
	return groupedItems
}


// MapByAnagrams maps by anagrams an array of strings into arrays of strings
func MapByAnagrams(receivedItems []string) map[string][]string {
	// Map received items into arrays of anagrams
	mappedItems := map[string][]string{}
	for _,receivedItem := range receivedItems {
		processedItem := SortByRunes(receivedItem)
		mappedItems[processedItem] = append(mappedItems[processedItem], receivedItem)
	}
	return mappedItems
}


// SortByRunes sort a string by its runes
func SortByRunes(item string) string {
	itemsRunes := []rune(item)
	sort.Sort(byRunes(itemsRunes))
	return string(itemsRunes)
}


// byRune sort array of runes
type byRunes []rune
func (s byRunes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s byRunes) Len() int { return len(s) }