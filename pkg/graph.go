package pkg

import (
	"errors"
	"fmt"
)

type Graph struct {
	Name  string
	Nodes map[string]*Node
}

type Node struct {
	ID    string
	Edges []*Edge
}

type Edge struct {
	ID       string
	Taken    bool
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
// AddNode - Add a node to the graph.
//           Can only add the first node, afte that you must
//           connect to an existing node
//
func (graph *Graph) AddNode(ID string) error {

	if len(graph.Nodes) > 0 {
		return errors.New("can not add a node if the graph has existing nodes. Try connecting to an existing node")
	}

	node := newNode(ID)
	graph.Nodes[ID] = node
	return nil
}

//
// ConnectNewNode - Connect a new node to an existing node in the graph
//
func (graph *Graph) ConnectNewNode(existingNodeID, newNodeID, edgeID string) error {

	// find existing node
	existingNode, ok := graph.Nodes[existingNodeID]
	if !ok {
		return errors.New("node not found for existingNodeID")
	}

	// create new node
	newNode := newNode(newNodeID)
	graph.Nodes[newNodeID] = newNode

	// create edge such that existing -> new and new -> existing
	edge := newEdge(edgeID, existingNode, newNode)

	existingNode.Edges = append(existingNode.Edges, edge)
	newNode.Edges = append(newNode.Edges, edge)

	return nil
}

//
// AddEdge - Add an edge to two existing nodes
//
func (graph *Graph) AddEdge(xNodeID, yNodeID, edgeID string) error {

	// find existing nodes
	xNode, ok := graph.Nodes[xNodeID]
	if !ok {
		return errors.New("xNodeID not found")
	}

	yNode, ok := graph.Nodes[yNodeID]
	if !ok {
		return errors.New("yNodeID not found")
	}


	// create edge such that existing -> new and new -> existing
	edge := newEdge(edgeID, xNode, yNode)

	xNode.Edges = append(xNode.Edges, edge)
	yNode.Edges = append(yNode.Edges, edge)

	return nil
}

//
// 

//
// GetNode - Get node by ID
//
func (graph *Graph) GetNode(nodeID string) *Node {
	return graph.Nodes[nodeID]
}

//
// GetDegree
//
func (graph *Graph) GetGegree(node *Node) int {

	return len(node.Edges)

}

//
// IsEulerian - Check to see if the graph is Eulerian, or partly Eulerian
//              i.e. all nodes have even number of edges
//
func (graph *Graph) IsEulerian() bool {

	if len(graph.Nodes) < 2 {
		// must have at least two nodes
		return false
	}

	var oddCount int = 0
	for _, node := range graph.Nodes {

		if len(node.Edges)%2 != 0 {
			oddCount++
		}
	}

	// all nodes must have even degree
	// or exactly two nodes have an odd degree
	if oddCount == 0 || oddCount == 2 {
		return true
	} else {
		return false
	}
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
