package trie


type Node struct {
    nodes []Node
    val byte
    end bool
}

func found(nodes []Node, c byte) *Node {
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

func (root *Node) Exists(val string) bool {
    var cur *Node = root
    for _, c := range val {
        if n := found(cur.nodes, byte(c)); n != nil {
            cur = n
        } else {
            return false;
        }
    }
    return cur.end
}

func (root *Node) Delete(val string) bool {
    var cur *Node = root
    for _, c := range val {
        if n := found(cur.nodes, byte(c)); n != nil {
            cur = n
        } else {
            return false;
        }
    }

    if cur.end {
        cur.end = false
        return true
    }
    return false
}

func (root *Node) Insert(val string) *Node {
    var cur *Node = root
    for _, c := range val {
        if n := found(cur.nodes, byte(c)); n != nil {
            cur = n
        } else {
            cur.nodes = append(cur.nodes, createNode(byte(c)))
            cur = &cur.nodes[len(cur.nodes)-1]
        }
    }
    cur.end = true
    return cur
}

