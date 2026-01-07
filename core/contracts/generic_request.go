package contracts

type IGenericRequest[TData any] interface {
	Build() TData
}

type GenericRequest[TData any] struct {
	data TData
}

func (g *GenericRequest[TData]) Build() TData {
	return g.data
}

func NewGenericRequest[TData any](data TData) *GenericRequest[TData] {
	return &GenericRequest[TData]{data: data}
}
