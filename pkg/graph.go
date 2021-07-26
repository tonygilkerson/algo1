package pkg

import "fmt"

type Graph struct {
	Name  string
	Nodes map[string]*Node
}

type Node struct {
	ID    string
	Edges []*Edge
}

type Edge struct {
	ID    string
	Taken bool
	TargetOf map[string]*Node
}

//
// NewGraph - The Grap constructor
//
func NewGraph(name string) *Graph {
	graph := new(Graph)
	graph.Name = name
	graph.Nodes = make(map[string]*Node)

	return graph
}

//
// AddNode - Add a node to the graph
//
func (graph *Graph) AddNode(ID string) {
	node := newNode(ID)
	graph.Nodes[ID] = node
}

//
// AddEdge - Add an edge to the graph
//
func (graph *Graph) AddEdge(edgeID string, leftNodeID string, rightNodeID string) {

	// get from and to nodes from graph
	leftNode := graph.Nodes[leftNodeID]
	rightNode := graph.Nodes[rightNodeID]

	// create edge such that left -> right and right -> left
	edge := newEdge(edgeID, leftNode, rightNode)

	leftNode.Edges = append(leftNode.Edges, edge)
	rightNode.Edges = append(rightNode.Edges, edge)
}

//
// GetNode - Get node by ID
//
func (graph *Graph) GetNode(nodeID string) *Node {
	return graph.Nodes[nodeID]
}

//
// newNode - The Node constructor
//
func newNode(ID string) *Node {
	node := new(Node)
	node.ID = ID
	node.Edges = make([]*Edge, 0)

	return node
}

//
// Edge constructor
//
func newEdge(edgeID string, leftNode, rightNode *Node) *Edge {
	edge := new(Edge)
	edge.ID = edgeID
	edge.Taken = false
	edge.TargetOf = make(map[string]*Node)

	edge.TargetOf[leftNode.ID] = rightNode
	edge.TargetOf[rightNode.ID] = leftNode

	return edge
}

//
// Walk - walk the graph
//  node - walk starting from this node
//  path - the path taken
//  edge - the edge you came from
func (graph *Graph) Walk(node *Node, path *[]*Node) {

	*path = append(*path, node)

	for _, edge := range node.Edges {

		if !edge.Taken {
			edge.Taken = true
			// fmt.Printf("%s", node.ID)
			fmt.Printf("%s -(%s)-> %s\n", node.ID, edge.ID, edge.TargetOf[node.ID].ID)
			graph.Walk(edge.TargetOf[node.ID], path)
		}

	}
	// fmt.Printf("%s\n", node.ID)
	// fmt.Printf("\n")
}
