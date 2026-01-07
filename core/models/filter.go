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
