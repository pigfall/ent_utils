package ent_utils

import (
	"entgo.io/ent/entc/gen"
)

type Type struct {
	*gen.Type
}

func FromType(tpe *gen.Type) *Type {
	return &Type{
		Type: tpe,
	}
}

func (this *Type) Name() string {
	return this.Type.Name
}

func (this *Type) AllFldsIncludePK() []*gen.Field {
	flds := make([]*gen.Field, len(this.Fields)+1)
	flds[0] = this.Type.ID
	copy(flds[1:], this.Type.Fields)
	return flds
}

func (this *Type) AllFldsExlucdePK() []*gen.Field {
	return this.Fields
}
