package ent_utils

import(
	"entgo.io/ent/entc/gen"
)

type Edge struct{
	*gen.Edge
}

func FromEntEdge(edge *gen.Edge)*Edge{
	return &Edge{
		Edge:edge,
	}
}
