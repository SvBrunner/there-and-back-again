package service

import (
	"context"
	"errors"

	"github.com/SvBrunner/there-and-back-again/internal/domain"
)

var ErrNotFound = errors.New("not found")

type Service interface {
	AddRun(ctx context.Context, distanceKm float64, timeInMinutes int32) (domain.Run, error)
	ListRuns(ctx context.Context) ([]domain.Run, error)
	AddTarget(ctx context.Context, distanceKm float64, name string) (domain.Target, error)
	ListTargets(ctx context.Context) ([]domain.Target, error)
}
