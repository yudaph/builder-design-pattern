package builder

import (
	"fmt"
	"strings"
)

type order int

type operators int

const (
	Asc order = iota
	Desc
)

func (o order) String() string {
	return []string{"asc", "desc"}[o]
}

const (
	Equals operators = iota
	NotEquals
	GreaterThan
	GreaterThanEquals
	LessThan
	LessThanEquals
)

func (o operators) String() string {
	return []string{"=", "<>", ">", ">=", "<", "<="}[o]
}

type selectBuilder struct {
	selectField []string
	from        string
	where       []string
	order       []string
}

func NewSelectBuilder(selectField []string, from string) *selectBuilder {
	return &selectBuilder{selectField: selectField, from: from}
}

func (s *selectBuilder) AddSelectedField(fields ...string) *selectBuilder {
	for _, field := range fields {
		if field != "" {
			s.selectField = append(s.selectField, field)
		}
	}
	return s
}

func (s *selectBuilder) Where(field string, operation operators, value interface{}) *selectBuilder {
	switch value.(type) {
	case int, int8, int16, int32, int64, string, float32, float64, bool:
		s.where = append(s.where, fmt.Sprintf("%s %s %#v", field, operation, value))
	default:
	}
	return s
}

func (s *selectBuilder) Order(field string, order order) *selectBuilder {
	s.order = append(s.order, fmt.Sprintf("%s %s", field, order))
	return s
}

func (s *selectBuilder) Build() string {
	result := "SELECT "
	result += strings.Join(s.selectField, ",")
	result += " FROM " + s.from
	if len(s.from) > 0 {
		result += " WHERE "
		result += strings.Join(s.where, " AND ")
	}
	if len(s.order) > 0 {
		result += " ORDER BY "
		result += strings.Join(s.order, ",")
	}
	return result
}
