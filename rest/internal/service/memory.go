package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/SvBrunner/thereandbackagain/internal/domain"
)

type MemoryService struct {
	mu sync.Mutex

	journeyName string
	targetKm    float64

	runs []domain.Run
}

func NewMemoryService(journeyName string, targetKm float64) *MemoryService {
	return &MemoryService{
		journeyName: journeyName,
		targetKm:    targetKm,
		runs:        []domain.Run{},
	}
}

func (s *MemoryService) AddRun(ctx context.Context, distanceKm float64) (domain.Run, error) {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	run := domain.Run{
		ID:         newID(),
		Date:       time.Now().UTC(),
		DistanceKm: distanceKm,
	}
	s.runs = append(s.runs, run)
	return run, nil
}

func (s *MemoryService) ListRuns(ctx context.Context) ([]domain.Run, error) {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]domain.Run, len(s.runs))
	copy(out, s.runs)
	return out, nil
}

func newID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
