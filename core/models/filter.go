package models

import "reflect"

type Filter struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type Filters []Filter

func (f Filters) Validate(structure any) *SystemError {
	val := reflect.ValueOf(structure)
	reqType := val.Type()

	for _, filter := range f {
		if filter.Key == "" {
			return &SystemError{
				Code:    SystemErrorCodeValidation,
				Type:    SystemErrorTypeValidation,
				Level:   SystemErrorLevelError,
				Message: "key is empty",
				Details: struct{}{},
			}
		}

		if filter.Value == nil {
			return &SystemError{
				Code:    SystemErrorCodeValidation,
				Type:    SystemErrorTypeValidation,
				Level:   SystemErrorLevelError,
				Message: "value is empty",
				Details: struct{}{},
			}
		}

		field, ok := reqType.FieldByName(filter.Key)
		if !ok {
			return &SystemError{
				Code:    SystemErrorCodeValidation,
				Type:    SystemErrorTypeValidation,
				Level:   SystemErrorLevelError,
				Message: "field not found: " + filter.Key,
				Details: struct{}{},
			}
		}

		filterVal := reflect.ValueOf(filter.Value)
		if filterVal.Kind() != field.Type.Kind() {
			return &SystemError{
				Code:    SystemErrorCodeValidation,
				Type:    SystemErrorTypeValidation,
				Level:   SystemErrorLevelError,
				Message: "field type mismatch: " + filter.Key,
				Details: struct{}{},
			}
		}
	}

	return nil
}

func (f *Filter) Build() (Filter, *SystemError) {
	if f.Key == "" {
		return Filter{}, &SystemError{
			Code:    SystemErrorCodeValidation,
			Type:    SystemErrorTypeValidation,
			Level:   SystemErrorLevelError,
			Message: "key is empty",
			Details: struct{}{},
		}
	}
	if f.Value == nil {
		return Filter{}, &SystemError{
			Code:    SystemErrorCodeValidation,
			Type:    SystemErrorTypeValidation,
			Level:   SystemErrorLevelError,
			Message: "value is empty",
			Details: struct{}{},
		}
	}
	return *f, nil
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p *Pagination) GetOffset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		return 10
	}
	return p.Limit
}

type SearchQuery struct {
	Filters    Filters    `json:"filters"`
	Pagination Pagination `json:"pagination"`
}

func (sq *SearchQuery) Validate(structure any) *SystemError {
	if err := sq.Filters.Validate(structure); err != nil {
		return err
	}
	return nil
}
