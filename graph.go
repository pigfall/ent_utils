package ent_utils

import(
		"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)


func LoadGraph(schemaPath string)(*Graph,error){
	g,err := entc.LoadGraph(schemaPath,&gen.Config{})
	if err != nil{
		return nil,err
	}

	return FromGraph(g),nil
}

type Graph struct {
	*gen.Graph
}

func FromGraph(g *gen.Graph)*Graph{
	return &Graph{
		Graph: g,
	}
}

func (this *Graph) GetNodes()[]*Type{
	nodes := make([]*Type,0,len(this.Graph.Nodes))
	for _,v := range this.Graph.Nodes{
		nodes = append(nodes, FromType(v))
	}

	return nodes
}
