package application

import (
    "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"
    "{{ .ProjectName }}/pkg/log"
)

type QueryService struct {
    {{ .FileNameTitleLower }}Repo repository.{{ .FileName }}Repository
    logger *log.Logger
}

func NewQueryService(logger *log.Logger, repo repository.{{ .FileName }}Repository) *QueryService {
    return &QueryService{ {{ .FileNameTitleLower }}Repo: repo, logger: logger}
}

