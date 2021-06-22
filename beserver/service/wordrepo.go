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
}

// UpdateWordList updates the search word list
func (s *wordRepoService) UpdateWordList(ctx context.Context, in *wordrepo.UpdateWordRequest) (*wordrepo.UpdateWordResponse, error) {
}

// GetTopWords returns the top 5 words and the counts they are searched
func (s *wordRepoService) GetTopWords(ctx context.Context, in *wordrepo.GetTopWordRequest) (*wordrepo.GetTopWordResponse, error) {
}