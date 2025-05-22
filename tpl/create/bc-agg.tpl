package agg

import "{{ .ProjectName }}/internal/repository/model"

type {{ .FileName }} interface {
	Value() *model.{{ .FileName }}
	IsExist() bool
}
