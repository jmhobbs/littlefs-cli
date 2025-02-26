package tree

import (
	"cmp"
	"fmt"
	"io"
	"path/filepath"
	"slices"
	"strings"
)

type Node struct {
	Name     string
	Children []*Node
}

func cmpNodes(a, b *Node) int {
	return cmp.Compare(a.Name, b.Name)
}

func NewNode(name string) *Node {
	return &Node{name, []*Node{}}
}

func (f *Node) Insert(path string) {
	// TODO: there _must_ be a stdlib to do this
	split := strings.Split(path, string([]byte{filepath.Separator}))
	if len(split) <= 1 {
		f.Children = append(f.Children, NewNode(path))
		slices.SortFunc(f.Children, cmpNodes)
		return
	}

	for _, child := range f.Children {
		if split[0] == child.Name {
			child.Insert(filepath.Join(split[1:]...))
		}
	}
}

// todo: optional color escapes
func (root *Node) Print(out io.Writer) {
	root.print(out, []bool{})
}

func (root *Node) print(out io.Writer, depths []bool) {
	// special case for root node, print it's name
	// all other depths have their name printed by
	// their parent
	if len(depths) == 0 {
		fmt.Fprintln(out, filepath.Base(root.Name))
	}

	// if we have no children nodes, we have nothing left to do
	if len(root.Children) == 0 {
		return
	}

	// the final node gets treated slightly differently
	// so we track the index for testing
	finalNodeIndex := len(root.Children) - 1

	for i, node := range root.Children {
		// print prefix space for tree depth
		for _, blank := range depths {
			if blank {
				fmt.Fprint(out, "    ")
			} else {
				fmt.Fprint(out, "│   ")
			}
		}
		// print the name of this node
		isFinal := i == finalNodeIndex
		if isFinal {
			fmt.Fprintln(out, "└──", node.Name)
		} else {
			fmt.Fprintln(out, "├──", node.Name)
		}
		// and descend into it
		node.print(out, append(depths, isFinal))
	}
}
