package application

import (
    "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"
    "{{ .ProjectName }}/pkg/log"
)

type CommandService struct {
	{{ .FileNameTitleLower }}Repo repository.{{ .FileName }}Repository
	logger *log.Logger
}

func NewCommandService(logger *log.Logger, repo repository.{{ .FileName }}Repository) *CommandService {
	return &CommandService{ {{ .FileNameTitleLower }}Repo: repo, logger: logger}
}

