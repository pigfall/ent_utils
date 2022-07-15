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

func (this *Type) AllFldsIncludePK()[]*gen.Field{
	flds := make([]*gen.Field,len(this.Fields)+1)
	copy(flds,this.Type.Fields)
	flds[len(flds)-1] = this.Type.ID
	return flds
}

func (this *Type) AllFldsExlucdePK()[]*gen.Field{
	return this.Fields
}
