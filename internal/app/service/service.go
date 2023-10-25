package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int {
	startTime := time.Now()
	endTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	timeToEnd := endTime.Sub(startTime)

	return int(timeToEnd.Hours()) / 24
}
