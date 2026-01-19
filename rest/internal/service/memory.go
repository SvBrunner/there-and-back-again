package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/SvBrunner/there-and-back-again/internal/domain"
)

type MemoryService struct {
	mu sync.Mutex

	runs    []domain.Run
	targets []domain.Target
}

func NewMemoryService() *MemoryService {
	return &MemoryService{
		runs:    []domain.Run{},
		targets: []domain.Target{},
	}
}

func (s *MemoryService) AddRun(ctx context.Context, distanceKm float64, timeinminutes int32) (domain.Run, error) {
	select {
	case <-ctx.Done():
		return domain.Run{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	run := domain.Run{
		ID:            newID(),
		Timestamp:     time.Now().UTC(),
		DistanceInKm:  distanceKm,
		TimeInMinutes: timeinminutes,
	}
	s.runs = append(s.runs, run)
	return run, nil
}

func (s *MemoryService) ListRuns(ctx context.Context) ([]domain.Run, error) {
	select {
	case <-ctx.Done():
		return []domain.Run{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]domain.Run, len(s.runs))
	copy(out, s.runs)
	return out, nil
}

func (s *MemoryService) AddTarget(ctx context.Context, distanceKm float64, name string) (domain.Target, error) {
	select {
	case <-ctx.Done():
		return domain.Target{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	target := domain.Target{
		ID:           newID(),
		DistanceInKm: distanceKm,
		Name:         name,
	}
	s.targets = append(s.targets, target)
	return target, nil
}

func (s *MemoryService) ListTargets(ctx context.Context) ([]domain.Target, error) {
	select {
	case <-ctx.Done():
		return []domain.Target{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]domain.Target, len(s.targets))
	copy(out, s.targets)
	return out, nil
}

func newID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
