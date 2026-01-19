package service

import (
	"context"

	"github.com/SvBrunner/thereandbackagain/internal/domain"
)

type Service interface {
	AddRun(ctx context.Context, distanceKm float64) (domain.Run, error)
	ListRuns(ctx context.Context) ([]domain.Run, error)
}<
