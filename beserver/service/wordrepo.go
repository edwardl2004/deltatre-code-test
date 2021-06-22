package service

import (
	"context"

	"github.com/edwardl2004/deltatre-code-test/beserver/proto/wordrepo"
)

type wordRepoService struct {
	repo map[string]int64
}

func NewWordRepoService() wordrepo.WordRepoServer {
	return &wordRepoService{
		repo: map[string]int64{},
	}
}

func (s *wordRepoService) SearchWord(ctx context.Context, in *wordrepo.SearchWordRequest) (*wordrepo.SearchWordResponse, error) {
	if _, ok := s.repo[in.Word]; ok {
		s.repo[in.Word]++
		return &wordrepo.SearchWordResponse{Found: true}, nil
	}

	return &wordrepo.SearchWordResponse{Found: false}, nil
}

// UpdateWordList updates the search word list
func (s *wordRepoService) UpdateWordList(ctx context.Context, in *wordrepo.UpdateWordRequest) (*wordrepo.UpdateWordResponse, error) {
	for _, word := range in.Words {
		if _, ok := s.repo[word]; !ok {
			s.repo[word] = 0
		}
	}

	return &wordrepo.UpdateWordResponse{
		Status:  "success",
		Message: "",
	}, nil
}

// GetTopWords returns the top 5 words and the counts they are searched
func (s *wordRepoService) GetTopWords(ctx context.Context, in *wordrepo.GetTopWordRequest) (*wordrepo.GetTopWordResponse, error) {
}
