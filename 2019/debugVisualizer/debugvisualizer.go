package debugvisualizer

import (
	"encoding/json"
)

// Graph ...
type Graph struct {
	Kind  map[string]bool `json:"kind"`
	Nodes []NodeGraphData `json:"nodes"`
	Edges []EdgeGraphData `json:"edges"`
}

// NodeGraphData ...
type NodeGraphData struct {
	ID    string `json:"id"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
	Shape string `json:"shape,omitempty"`
}

// EdgeGraphData ...
type EdgeGraphData struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Label  string `json:"label,omitempty"`
	ID     string `json:"id"`
	Color  string `json:"color,omitempty"`
	Dashes bool   `json:"dashes,omitempty"`
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		Kind:  map[string]bool{"graph": true},
		Nodes: []NodeGraphData{},
		Edges: []EdgeGraphData{},
	}
}

// ToString for Graph
func (e *Graph) ToString() string {
	rs, _ := json.Marshal(e)
	return string(rs)
}

// D4DataTable ...
type D4DataTable struct {
	Kind map[string]bool `json:"kind"`
	Rows []D4Row         `json:"rows"`
}

// D4Row ...
type D4Row struct {
	SixDigit        string `json:"six digit"`
	TwoDigitsSame   string `json:"tow digits same"`
	NeverDecreasing string `json:"never decreasing"`
}

// NewD4DataTable ...
func NewD4DataTable() *D4DataTable {
	return &D4DataTable{
		Kind: map[string]bool{"table": true},
		Rows: []D4Row{},
	}
}

// ToString for D4DataTable
func (e *D4DataTable) ToString() string {
	rs, _ := json.Marshal(e)
	return string(rs)
}

// MeshData ...
type MeshData struct {
	Type string `json:"type"`
	Mode string `json:"mode,omitempty"`
	X    []int  `json:"x"`
	Y    []int  `json:"y"`
	Z    []int  `json:"z"`
}

// Mesh ...
type Mesh struct {
	Kind map[string]bool `json:"kind"`
	Data []MeshData      `json:"data"`
}

// NewMesh ...
func NewMesh() *Mesh {
	return &Mesh{
		Kind: map[string]bool{"plotly": true},
		Data: []MeshData{},
	}
}

// ToString for Mesh
func (e *Mesh) ToString() string {
	rs, _ := json.Marshal(e)
	return string(rs)
}
