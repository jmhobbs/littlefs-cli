package tree

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node_Print(t *testing.T) {
	root := &Node{
		"root",
		[]*Node{
			{"hello.txt", nil},
			{"home", []*Node{
				{"desktop", []*Node{{"avatar.jpg", nil}}},
				{"todo.txt", nil},
			}},
			{"notes", nil},
			{
				"subdir", []*Node{
					{"a-file", nil},
					{"another-file", nil},
					{"deeper", []*Node{{"deepest", nil}}},
					{"subfile.txt", nil},
				},
			},
		},
	}

	var buf bytes.Buffer
	root.Print(&buf)

	expected := `root
├── hello.txt
├── home
│   ├── desktop
│   │   └── avatar.jpg
│   └── todo.txt
├── notes
└── subdir
    ├── a-file
    ├── another-file
    ├── deeper
    │   └── deepest
    └── subfile.txt
`

	if buf.String() != expected {
		fmt.Println("Expected:", expected)
		fmt.Println("Got:", buf.String())
		t.Fail()
	}
}

func Test_Node_Print_One(t *testing.T) {
	root := &Node{"root", nil}
	var buf bytes.Buffer
	root.Print(&buf)
	expected := "root\n"
	if buf.String() != expected {
		fmt.Println("Expected:", expected)
		fmt.Println("Got:", buf.String())
		t.Fail()
	}
}

func Test_Node_Insert(t *testing.T) {
	root := &Node{"root", []*Node{}}
	root.Insert("hello.txt")
	root.Insert("home")
	root.Insert("home/todo.txt")
	root.Insert("home/todone.txt")
	root.Insert("home/zeeper")
	root.Insert("home/zeeper/deepest.txt")
	root.Insert("notes")
	root.Insert("allo")

	/*
	   root
	   ├── allo
	   ├── hello.txt
	   ├── home
	   │   ├── todo.txt
	   │   ├── todone.txt
	   │   └── zeeper
	   │       └── deepest.txt
	   └── notes
	*/

	expected := &Node{
		"root", []*Node{
			{"allo", []*Node{}},
			{"hello.txt", []*Node{}},
			{"home", []*Node{
				{"todo.txt", []*Node{}},
				{"todone.txt", []*Node{}},
				{"zeeper", []*Node{
					{"deepest.txt", []*Node{}},
				}},
			}},
			{"notes", []*Node{}},
		},
	}

	assert.Equal(t, expected, root)
}
