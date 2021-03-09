package trie

type Node struct {
	nodes []Node
	val   byte
	end   bool
}

func (root *Node) Insert(val string) *Node {
	var current *Node = root
	for _, c := range val {
		if found := findByte(current.nodes, byte(c)); found != nil {
			current = found
		} else {
			current.nodes = append(current.nodes, createNode(byte(c)))
			current = &current.nodes[len(current.nodes)-1]
		}
	}
	current.end = true
	return current
}

func (root *Node) Delete(val string) bool {
	found := root.findNode(val, true)
	if found != nil {
		found.end = false
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
		if found := findByte(current.nodes, byte(c)); found != nil {
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

func findByte(nodes []Node, c byte) *Node {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].val == c {
			return &nodes[i]
		}
	}
	return nil
}

func createNode(val byte) Node {
	return Node{
		val: val,
		end: false,
	}
}

func traverse(node *Node, current string, nodeFound func(*Node, string)) {
	nodeFound(node, current+string(node.val))
	for _, n := range node.nodes {
		traverse(&n, current+string(node.val), nodeFound)
	}
}
