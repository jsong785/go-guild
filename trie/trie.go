package trie

type Node struct {
	parent *Node
	nodes  []Node
	val    byte
	end    bool
}

func (root *Node) Insert(val string) *Node {
	var current *Node = root
	for _, c := range val {
		if found, _ := findByte(current.nodes, byte(c)); found != nil {
			current = found
		} else {
			current.nodes = append(current.nodes, createNode(byte(c), current))
			current = &current.nodes[len(current.nodes)-1]
		}
	}
	current.end = true
	return current
}

func (root *Node) Delete(val string) bool {
	found := root.findNode(val, true)
	if found != nil {
		handleDelete(found)
		return true
	}
	return false
}

func (root *Node) GetNode(val string) *Node {
	return root.findNode(val, true)
}

func (root *Node) Exists(val string) bool {
	return root.GetNode(val) != nil
}

func (root *Node) Length() int {
	count := 0
	traverse(root, "", func(*Node, string) {
		count++
	})
	return count
}

func (root *Node) GetMatches(val string) []string {
	n := root.findNode(val, false)
	if n == nil {
		return nil
	}
	res := make([]string, 0)
	traverse(n, val[:len(val)-1], func(node *Node, val string) {
		if node.end {
			res = append(res, val)
		}
	})
	return res
}

func (root *Node) findNode(val string, needsEnding bool) *Node {
	current := root
	for _, c := range val {
		if found, _ := findByte(current.nodes, byte(c)); found != nil {
			current = found
		} else {
			return nil
		}
	}
	if !needsEnding || current.end {
		return current
	}
	return nil
}

func findByte(nodes []Node, c byte) (*Node, int) {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].val == c {
			return &nodes[i], i
		}
	}
	return nil, -1
}

func createNode(val byte, parent *Node) Node {
	return Node{
		parent: parent,
		val:    val,
		end:    false,
	}
}

func traverse(node *Node, current string, nodeFound func(*Node, string)) {
	nodeFound(node, current+string(node.val))
	for _, n := range node.nodes {
		traverse(&n, current+string(node.val), nodeFound)
	}
}

func delete(n []Node, c byte) []Node {
	if found, i := findByte(n, c); found != nil {
		n[i] = n[len(n)-1]
		n = n[:len(n)-1]
	}
	return n
}

func handleDelete(node *Node) {
	node.end = false
	if len(node.nodes) == 0 && node.parent != nil {
		node.parent.nodes = delete(node.parent.nodes, node.val)
		if len(node.parent.nodes) == 0 {
			handleDelete(node.parent)
		}
	}
}
