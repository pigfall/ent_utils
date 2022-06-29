package ent_utils

import(
		"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)


func LoadGraph(schemaPath string)(*gen.Graph,error){
	return entc.LoadGraph(schemaPath,nil)
}
