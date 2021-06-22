package service

import (
	"testing"

	"github.com/edwardl2004/deltatre-code-test/beserver/proto/wordrepo"
)

func Test_GetTopNElements(t *testing.T) {
	testcases := []struct {
		name           string
		input          map[string]int64
		expectedOutput []*wordrepo.TopSearch
	}{
		{
			name:  "happy path",
			input: map[string]int64{"hello": 6, "goodbye": 1, "simple": 5, "list": 2, "search": 3, "filter": 0, "yes": 7, "no": 8},
			expectedOutput: []*wordrepo.TopSearch{
				{Word: "no", Count: 8},
				{Word: "yes", Count: 7},
				{Word: "hello", Count: 6},
				{Word: "simple", Count: 5},
				{Word: "search", Count: 3},
			},
		},
		{
			name:  "multiple 0",
			input: map[string]int64{"hello": 0, "goodbye": 1, "simple": 0, "list": 0, "search": 0, "filter": 2, "yes": 0, "no": 0},
			expectedOutput: []*wordrepo.TopSearch{
				{Word: "filter", Count: 2},
				{Word: "goodbye", Count: 1},
				{Word: "hello", Count: 0},
				{Word: "simple", Count: 0},
				{Word: "list", Count: 0},
			},
		},
	}

	for _, test := range testcases {
		s := &wordRepoService{
			repo: test.input,
		}

		output := s.getTopNElements(5)
		if len(output) != 5 {
			t.Fatal("test case ", test.name, ", expect list of 5, got ", len(output))
		}

		for index, item := range output {
			if item.Count > 0 { // sequence of iteration of hashmap is determined
				if item.Word != test.expectedOutput[index].Word || item.Count != test.expectedOutput[index].Count {
					t.Fatalf("test case  %s, item %d, expect word %s, count %d, got word %s, count %d", test.name, index, test.expectedOutput[index].Word, test.expectedOutput[index].Count, item.Word, item.Count)
				}
			}
		}
	}
}
