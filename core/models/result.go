package models

import "fmt"

type DynamicResult struct {
	data any
}

func NewDynamicResult(data any) *DynamicResult {
	return &DynamicResult{data: data}
}

// Función genérica standalone
func Convert[T any](r *DynamicResult) (T, error) {
	var zero T

	if r == nil || r.data == nil {
		return zero, fmt.Errorf("result is nil")
	}

	value, ok := r.data.(T)
	if !ok {
		return zero, fmt.Errorf("cannot convert %T to %T", r.data, zero)
	}

	return value, nil
}

// Alternativa con Must para cuando estés seguro del tipo
func MustConvert[T any](r *DynamicResult) T {
	value, err := Convert[T](r)
	if err != nil {
		panic(err)
	}
	return value
}

// Uso en tests o código de ejemplo
func Example() {
	result := NewDynamicResult("hello")

	// Conversión segura
	str, err := Convert[string](result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Success:", str)

	// Conversión con panic si falla (usar solo cuando estés seguro)
	str2 := MustConvert[string](result)
	fmt.Println("Must:", str2)
}
