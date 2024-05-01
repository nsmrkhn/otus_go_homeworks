package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type countMap struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	words := strings.Fields(s)
	wordsCounter := make(map[string]int)

	for _, word := range words {
		wordsCounter[word]++
	}

	sortedMap := make([]countMap, 0, len(wordsCounter))
	for k, v := range wordsCounter {
		sortedMap = append(sortedMap, countMap{k, v})
	}

	sort.Slice(sortedMap, func(i, j int) bool {
		if sortedMap[i].Count == sortedMap[j].Count {
			return sortedMap[i].Word < sortedMap[j].Word
		}
		return sortedMap[i].Count > sortedMap[j].Count
	})

	if len(sortedMap) > 10 {
		sortedMap = sortedMap[:10]
	}

	result := make([]string, 0, len(sortedMap))
	for _, Key := range sortedMap {
		result = append(result, Key.Word)
	}

	return result
}
