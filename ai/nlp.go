package ai

import (
	"regexp"
	"sort"
	"strings"
)

// A tiny, zero-dependency TF-IDF style keyword extractor
type NLP struct {
	stopWords map[string]bool
}

func NewNLP() *NLP {
	words := []string{
		"the", "be", "to", "of", "and", "a", "in", "that", "have", "i", "it", "for", "not", "on", "with", "he", "as", "you", "do", "at",
		"this", "but", "his", "by", "from", "they", "we", "say", "her", "she", "or", "an", "will", "my", "one", "all", "would", "there",
		"their", "what", "so", "up", "out", "if", "about", "who", "get", "which", "go", "me", "when", "make", "can", "like", "time", "no",
		"just", "him", "know", "take", "people", "into", "year", "your", "good", "some", "could", "them", "see", "other", "than", "then",
		"now", "look", "only", "come", "its", "over", "think", "also", "back", "after", "use", "two", "how", "our", "work", "first", "well",
		"way", "even", "new", "want", "because", "any", "these", "give", "day", "most", "us", "is", "are", "was", "were", "has", "had", "been",
	}

	stopMap := make(map[string]bool)
	for _, w := range words {
		stopMap[w] = true
	}
	return &NLP{stopWords: stopMap}
}

type wordFreq struct {
	Word  string
	Count int
}

// ExtractKeywords finds the most meaningful words in a text
func (n *NLP) ExtractKeywords(text string, topN int) []string {
	// Simple regex to match words (letters only)
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(text, -1)

	freq := make(map[string]int)
	for _, w := range words {
		lower := strings.ToLower(w)
		if len(lower) < 4 {
			continue // skip very short words
		}
		if n.stopWords[lower] {
			continue
		}
		freq[lower]++
	}

	var list []wordFreq
	for k, v := range freq {
		list = append(list, wordFreq{k, v})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	var result []string
	for i := 0; i < len(list) && i < topN; i++ {
		result = append(result, list[i].Word)
	}

	return result
}
