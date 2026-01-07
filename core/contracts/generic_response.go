package contracts

type IGenericResponse[TResult any] interface {
	Data() TResult
}

type GenericResponse[TResult any] struct {
	data TResult
}

func (g *GenericResponse[TResult]) Data() TResult {
	return g.data
}
