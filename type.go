package ent_utils

import(
	"entgo.io/ent/entc/gen"
)


type Type struct{
	*gen.Type
}

func FromType(tpe *gen.Type) *Type{
	return &Type{
		Type:tpe,
	}
}

func (this *Type) Name() string{
	return this.Type.Name
}
