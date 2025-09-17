// Package poll provides a repository for polls
package poll

import (
	"sync"

	"pollparlor/internal/domain"
)

// MemoryRepo is a repository for polls in memory
type MemoryRepo struct {
	mu    sync.RWMutex
	polls []domain.Poll
}

//
// var _ domain.PollRepository = (*MemoryRepo)(nil)
//
// // NewMemoryRepo creates a new repository for polls in memory
// func NewMemoryRepo(seed []domain.Poll) *MemoryRepo {
// 	return &MemoryRepo{polls: seed}
// }
//
// // List returns all polls in the repository
// func (r *MemoryRepo) List(limit, skip int64) ([]domain.Poll, error) {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()
//
// 	if limit == 0 {
// 		limit = int64(len(r.polls))
// 	}
// 	if skip < 0 {
// 		skip = 0
// 	}
//
// 	out := make([]domain.Poll, 0, limit)
// 	for i := skip; i < int64(len(r.polls)) && i < skip+limit; i++ {
// 		out = append(out, r.polls[i])
// 	}
// 	return out, nil
// }
//
// // GetByID returns a poll by its ID
// func (r *MemoryRepo) GetByID(id bson.ObjectID) (*domain.Poll, error) {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()
// 	for _, p := range r.polls {
// 		if p.ID == id {
// 			cp := p
// 			return &cp, nil
// 		}
// 	}
// 	return nil, errors.New("not found")
// }
//
// // Create adds a poll to the repository
// func (r *MemoryRepo) Create(p domain.Poll, cands []domain.Candidate) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	r.polls = append(r.polls, p)
// 	return nil
// }
