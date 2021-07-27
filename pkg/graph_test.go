package pkg_test

import (
	"strconv"
	"testing"

	"github.com/tonygilkerson/algo1/pkg"
)

func Test1(t *testing.T) {
	want := "mygraph"

	var graph pkg.Graph
	graph.Name = "mygraph"
	got := graph.Name

	if want != got {
		t.Errorf("Set graph name, want: %s, got: %s\n", want, got)
	}
}

func Test2(t *testing.T) {

	graph := pkg.NewGraph("My Graph")

	// Add a few nodes
	graph.AddNode("first")

	for i := 0; i < 9; i++ {
		id := "node-" + strconv.Itoa(i)
		graph.ConnectNewNode("first", id, "edge")
	}

	graph.ConnectNewNode("a", "b", "edge a-b")
	graph.ConnectNewNode("b", "c", "edge b-c")
	graph.ConnectNewNode("c", "d", "edge c-d")

	want := 10
	got := len(graph.Nodes)
	if want != got {
		t.Errorf("Check graph size, want %d, got %d\n", want, got)
	}

}

func Test3(t *testing.T) {
	graph := pkg.NewGraph("My Graph")

	// Add a few nodes
	graph.AddNode("a")

	// Add edges so that a <--> b <--> c <--> d
	graph.ConnectNewNode("a", "b", "edge a-b")
	graph.ConnectNewNode("b", "c", "edge b-c")
	graph.ConnectNewNode("c", "d", "edge c-d")

	// Start at a, traverse graph and make sure you end up at d
	// This is a simplified case where each node has at most 2 edge
	// and one of them points to the node we came from so we want to ignore it
	// when walking the graph

	// Walk a->b->c->d
	nodeA := graph.GetNode("a")
	nodeB := nodeA.Edges[0].TargetOf["a"] // A only has one edge
	nodeC := nodeB.Edges[1].TargetOf["b"] // 0 - goes back, 1 - goes forward
	nodeD := nodeC.Edges[1].TargetOf["c"]

	// Walk it back
	nodeC = nodeD.Edges[0].TargetOf["d"] // D only has one edge
	nodeB = nodeC.Edges[0].TargetOf["c"] // 0 - goes back, 1 - goes forward
	nodeA = nodeB.Edges[0].TargetOf["b"]

	// If all that worked then the nodeID of nodeA will be "a"
	want := "a"
	got := nodeA.ID

	if want != got {
		t.Errorf("Walk graph, want %s, got %s\n", want, got)
	}

}

func Test4(t *testing.T) {
	graph := pkg.NewGraph("My Graph")

	// Add a few nodes
	graph.AddNode("a")

	// Add edges so that a <--> b <--> c <--> d
	graph.ConnectNewNode("a", "b", "edge a-b")
	graph.ConnectNewNode("b", "c", "edge b-c")
	graph.ConnectNewNode("c", "d", "edge c-d")

	startNode := graph.GetNode("a")
	path := new([]*pkg.Node)

	graph.Walk(startNode, path)

}

func Test5(t *testing.T) {
	graph := pkg.NewGraph("My Graph")

	// Add a few nodes
	graph.AddNode("a")

	// Add edges so that a <--> b <--> c <--> d
	graph.ConnectNewNode("a", "b", "edge a-b")
	graph.ConnectNewNode("b", "c", "edge b-c")
	graph.ConnectNewNode("c", "d", "edge c-d")

	want := true
	got := graph.IsEulerian()

	if want != got {
		t.Errorf("Check if graph is Eulerian, want %t, got %t\n", want, got)
	}
}

func Test6(t *testing.T) {
	graph := pkg.NewGraph("The Seven Bridges of Königsberg")

	// Add a few nodes
	graph.AddNode("a")

	// Add edges so that it is the The Seven Bridges of Königsberg
	graph.ConnectNewNode("a", "b", "edge1 a-b")
	graph.AddEdge("a", "b", "edge2 a-b")

	graph.ConnectNewNode("b", "c", "edge3 b-c")
	graph.AddEdge("b", "c", "edge4 b-c")

	graph.ConnectNewNode("b", "d", "edge5 b-d")
	graph.AddEdge("a", "d", "edge6 a-d")
	graph.AddEdge("c", "d", "edge7 c-d")

	want := false
	got := graph.IsEulerian()

	if want != got {
		t.Errorf("Check The Seven Bridges of Königsberg is not Eulerian, want %t, got %t\n", want, got)
	}
}
