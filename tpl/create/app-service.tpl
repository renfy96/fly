package service

import (
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/application"
	"{{ .ProjectName }}/pkg/log"
)

type Service struct {
	query *application.QueryService
	command *application.CommandService
	logger *log.Logger
}

func NewService(logger *log.Logger, query *application.QueryService, command *application.CommandService) *Service {
	return &Service{query: query, command: command, logger: logger}
}