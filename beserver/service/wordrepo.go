package service

import (
	"context"
	"strings"
	"sync"

	"github.com/edwardl2004/deltatre-code-test/beserver/proto/wordrepo"
)

type wordRepoService struct {
	repo map[string]int64
	sync.RWMutex
}

func NewWordRepoService() wordrepo.WordRepoServer {
	return &wordRepoService{
		repo: map[string]int64{"hello": 0, "goodbye": 0, "simple": 0, "list": 0, "search": 0, "filter": 0, "yes": 0, "no": 0},
	}
}

func (s *wordRepoService) SearchWord(ctx context.Context, in *wordrepo.SearchWordRequest) (*wordrepo.SearchWordResponse, error) {
	s.RLock()
	word := strings.ToLower(in.Word)
	if _, ok := s.repo[word]; ok {
		s.RUnlock()
		s.Lock()
		defer s.Unlock()
		// Check again if the search word exists, in case the map is updated between s.RUnlock() and s.Lock()
		if _, ok := s.repo[word]; ok {
			s.repo[word]++
			return &wordrepo.SearchWordResponse{Found: true}, nil
		} else {
			return &wordrepo.SearchWordResponse{Found: false}, nil
		}
	}

	s.RUnlock()
	return &wordrepo.SearchWordResponse{Found: false}, nil
}

// UpdateWordList updates the search word list
func (s *wordRepoService) UpdateWordList(ctx context.Context, in *wordrepo.UpdateWordRequest) (*wordrepo.UpdateWordResponse, error) {
	s.Lock()
	defer s.Unlock()

	s.repo = make(map[string]int64)
	for _, word := range in.Words {
		s.repo[strings.ToLower(word)] = 0
	}

	return &wordrepo.UpdateWordResponse{
		Status:  "success",
		Message: "",
	}, nil
}

// GetTopWords returns the top 5 words and the counts they are searched
func (s *wordRepoService) GetTopWords(ctx context.Context, in *wordrepo.GetTopWordRequest) (*wordrepo.GetTopWordResponse, error) {
	s.RLock()
	defer s.RUnlock()

	if len(s.repo) <= 5 {
		list := make([]*wordrepo.TopSearch, len(s.repo))
		index := 0
		for k, v := range s.repo {
			list[index] = &wordrepo.TopSearch{Word: k, Count: v}
			index++
		}

		return &wordrepo.GetTopWordResponse{List: list}, nil
	}

	list := s.getTopNElements(5)
	return &wordrepo.GetTopWordResponse{List: list}, nil
}

func (s *wordRepoService) getTopNElements(n int) []*wordrepo.TopSearch {
	list := make([]*wordrepo.TopSearch, n)
	lastIndex := -1
	newItemAppended := false
	for k, v := range s.repo {
		if lastIndex < n-1 {
			// list not full, append current word to the last of the list
			lastIndex++
			list[lastIndex] = &wordrepo.TopSearch{Word: k, Count: v}
			newItemAppended = true
		} else {
			if v > list[lastIndex].Count {
				// replace the smallest item (ie, the last item) with new one, if new item is larger
				list[lastIndex].Word = k
				list[lastIndex].Count = v
				newItemAppended = true
			}
		}

		// bubble sort the new inserted item to the right position
		if newItemAppended {
			for j := lastIndex; j > 0; j-- {
				if list[j].Count > list[j-1].Count {
					// new item is larger than the previous, swap them
					temp := list[j]
					list[j] = list[j-1]
					list[j-1] = temp
				}
			}
		}
	}
	return list
}
