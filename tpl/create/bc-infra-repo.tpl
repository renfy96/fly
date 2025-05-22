package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg/impl"
	"{{ .ProjectName }}/internal/repository/query"
	"{{ .ProjectName }}/pkg/log"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	q *query.Query
	logger *log.Logger
}

func New{{ .FileName }}Repository(logger *log.Logger, db *gorm.DB) repository.{{ .FileName }}Repository {
	return repositoryImpl{
    		q: query.Use(db),
    		logger: logger,
    }
}

func (r repositoryImpl) GetById(ctx *gin.Context, id uint) agg.{{ .FileName }} {
	{{ .FileNameTitleLower }} := r.q.{{ .FileName }}
	{{ .FileNameTitleLower }}Mol, _ := {{ .FileNameTitleLower }}.WithContext(ctx).Where({{ .FileNameTitleLower }}.ID.Eq(id)).First()
	return poToAgg({{ .FileNameTitleLower }}Mol)
}

func poToAgg({{ .FileNameTitleLower }} *model.{{ .FileName }}) agg.{{ .FileName }} {
	return impl.New{{ .FileName }}Impl({{ .FileNameTitleLower }})
}