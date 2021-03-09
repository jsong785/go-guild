package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testfound(t *testing.T) {
	assert.Nil(t, findByte(createNodes("a"), byte('z')))
	assert.Nil(t, findByte(createNodes("a"), byte('z')))
	assert.Nil(t, findByte(createNodes("abc"), byte('z')))

	assert.NotNil(t, findByte(createNodes("a"), byte('a')))
	assert.NotNil(t, findByte(createNodes("abc"), byte('a')))
	assert.NotNil(t, findByte(createNodes("abc"), byte('c')))
}

func TestInsert(t *testing.T) {
	var root Node
	assert.Nil(t, root.nodes)

	lNode := root.Insert("val")
	assert.NotNil(t, lNode)
	assert.Equal(t, lNode.val, byte('l'))
	assert.True(t, lNode.end)
	assert.Nil(t, lNode.nodes)
	assert.NotNil(t, root.nodes)

	eNode := root.Insert("value")
	assert.NotNil(t, eNode)
	assert.Equal(t, eNode.val, byte('e'))
	assert.True(t, eNode.end)
	assert.Nil(t, eNode.nodes)
	assert.NotNil(t, root.nodes)

	assert.NotNil(t, lNode.nodes)

	// ensure lNode goes to eNode
	assert.Equal(t, lNode.val, byte('l'))
	assert.True(t, lNode.end)
	assert.Equal(t, eNode, &lNode.nodes[0].nodes[0])
}

func TestExists(t *testing.T) {
	var root Node
	assert.False(t, root.Exists("val"))

	root.Insert("value")
	assert.False(t, root.Exists("val"))
	assert.False(t, root.Exists("valu"))
	assert.True(t, root.Exists("value"))

	root.Insert("val")
	assert.True(t, root.Exists("val"))
	assert.False(t, root.Exists("valu"))
	assert.True(t, root.Exists("value"))
}

func TestDeleteLazy(t *testing.T) {
	var root Node
	root.Insert("vaccinate")
	root.Insert("val")
	root.Insert("value")
	assert.True(t, root.Exists("vaccinate"))
	assert.True(t, root.Exists("val"))
	assert.True(t, root.Exists("value"))

	root.Delete("what")
	assert.True(t, root.Exists("vaccinate"))
	assert.True(t, root.Exists("val"))
	assert.True(t, root.Exists("value"))
	root.Delete("val")
	assert.True(t, root.Exists("vaccinate"))
	assert.False(t, root.Exists("val"))
	assert.True(t, root.Exists("value"))

	root.Delete("v")
	assert.True(t, root.Exists("vaccinate"))

	root.Delete("va")
	assert.True(t, root.Exists("vaccinate"))

	root.Delete("vac")
	assert.True(t, root.Exists("vaccinate"))

	root.Delete("vacc")
	assert.True(t, root.Exists("vaccinate"))

	root.Delete("vaccinat")
	assert.True(t, root.Exists("vaccinate"))

	root.Delete("vaccinate")
	assert.False(t, root.Exists("vaccinate"))

	root.Delete("value")
	assert.False(t, root.Exists("value"))
}

func TestDeleteLazyThenInsert(t *testing.T) {
	var root Node
	root.Insert("val")
	assert.True(t, root.Exists("val"))

	root.Delete("val")
	assert.False(t, root.Exists("val"))

	root.Insert("val")
	assert.True(t, root.Exists("val"))

	root.Delete("val")
	assert.False(t, root.Exists("val"))

	root.Insert("value")
	assert.False(t, root.Exists("val"))
	assert.True(t, root.Exists("value"))
}

func TestRudimentaryForeign(t *testing.T) {
	var root Node
	root.Insert("こん")
	root.Insert("こんにちは")
	assert.True(t, root.Exists("こん"))
	assert.True(t, root.Exists("こんにちは"))

	assert.True(t, root.Delete("こん"))
	assert.True(t, root.Delete("こんにちは"))
	assert.False(t, root.Exists("こんにち"))
}

func TestGetMatches(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		var root Node
		root.Insert("value")
		root.Insert("value_bargain")

		res := root.GetMatches("v")
		assert.Equal(t, []string{"value", "value_bargain"}, res)

		res = root.GetMatches("abc")
		assert.Nil(t, res)
	})

	t.Run("complex", func(t *testing.T) {
		var root Node
		root.Insert("abc")
		root.Insert("abacus")
		root.Insert("dinner")
		root.Insert("dictionary")

		res := root.GetMatches("a")
		assert.Equal(t, []string{"abc", "abacus"}, res)

		res = root.GetMatches("abc")
		assert.Equal(t, []string{"abc"}, res)

		res = root.GetMatches("di")
		assert.Equal(t, []string{"dinner", "dictionary"}, res)

		res = root.GetMatches("din")
		assert.Equal(t, []string{"dinner"}, res)
	})
}

func createNodes(val string) []Node {
	nodes := make([]Node, 0)
	for i, c := range val {
		nodes = append(nodes, Node{val: byte(c), end: i == len(val)-1})
	}
	return nodes
}
