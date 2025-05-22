package impl

import (
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
	"{{ .ProjectName }}/internal/repository/model"
)

type {{ .FileNameTitleLower }}Impl struct {
	value *model.{{ .FileName }}
}

func New{{ .FileName }}Impl(v *model.{{ .FileName }}) agg.{{ .FileName }} {
	return &{{ .FileNameTitleLower }}Impl{
		value: v,
	}
}

func (a {{ .FileNameTitleLower }}Impl) IsExist() bool {
	if a.value == nil {
		return false
	}
	return a.value.ID > 0
}

func (a {{ .FileNameTitleLower }}Impl) Value() *model.{{ .FileName }} {
	return a.value
}

