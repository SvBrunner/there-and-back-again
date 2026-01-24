package service

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/SvBrunner/there-and-back-again/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteService struct {
	mu sync.Mutex
	db *gorm.DB
}

type runModel struct {
	ID            int64 `gorm:"primaryKey;autoIncrement"`
	DistanceKm    float64
	TimeInMinutes int32
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

func (runModel) TableName() string { return "run" }

type targetModel struct {
	ID         int64 `gorm:"primaryKey;autoIncrement"`
	DistanceKm float64
	Name       string
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (targetModel) TableName() string { return "target" }

func NewSqliteService(dsn string) (*SqliteService, error) {
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}
	// Recommended for SQLite
	sqlDB.SetMaxOpenConns(1)

	if err := gdb.AutoMigrate(&runModel{}, &targetModel{}); err != nil {
		return nil, err
	}
	return &SqliteService{db: gdb}, nil
}
func (s *SqliteService) AddRun(ctx context.Context, distanceKm float64, timeInMinutes int32) (domain.Run, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m := runModel{
		DistanceKm:    distanceKm,
		TimeInMinutes: timeInMinutes,
	}
	if err := s.db.WithContext(ctx).Create(&m).Error; err != nil {
		return domain.Run{}, err
	}
	return domain.Run{
		ID:            strconv.FormatInt(m.ID, 10),
		DistanceInKm:  m.DistanceKm,
		TimeInMinutes: m.TimeInMinutes,
		Timestamp:     m.CreatedAt,
	}, nil
}

func (s *SqliteService) ListRuns(ctx context.Context) ([]domain.Run, error) {

	var models []runModel
	if err := s.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	runs := make([]domain.Run, len(models))
	for i, m := range models {
		runs[i] = domain.Run{
			ID:            strconv.FormatInt(m.ID, 10),
			DistanceInKm:  m.DistanceKm,
			TimeInMinutes: m.TimeInMinutes,
			Timestamp:     m.CreatedAt,
		}
	}
	return runs, nil
}

func (s *SqliteService) AddTarget(ctx context.Context, distanceKm float64, name string) (domain.Target, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m := targetModel{
		DistanceKm: distanceKm,
		Name:       name,
	}
	if err := s.db.WithContext(ctx).Create(&m).Error; err != nil {
		return domain.Target{}, err
	}
	return domain.Target{
		ID:           strconv.FormatInt(m.ID, 10),
		DistanceInKm: m.DistanceKm,
		Name:         m.Name,
	}, nil
}

func (s *SqliteService) ListTargets(ctx context.Context) ([]domain.Target, error) {
	var models []targetModel
	if err := s.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	runs := make([]domain.Target, len(models))
	for i, m := range models {
		runs[i] = domain.Target{
			ID:           strconv.FormatInt(m.ID, 10),
			DistanceInKm: m.DistanceKm,
			Name:         m.Name,
		}
	}
	return runs, nil
}
