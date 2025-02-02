package persistent

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNilEmpty(t *testing.T) {
	var tree *Tree[int, int]
	require.True(t, tree.IsEmpty())
}

func TestDefaultEmpty(t *testing.T) {
	var tree Tree[int, int]
	require.True(t, tree.IsEmpty())
}

func TestNilHeight(t *testing.T) {
	var tree *Tree[int, int]
	require.Equal(t, 0, tree.Height())
}

func TestDefaultHeight(t *testing.T) {
	var tree Tree[int, int]
	require.Equal(t, 0, tree.Height())
}

func TestNilSize(t *testing.T) {
	var tree *Tree[int, int]
	require.Equal(t, 0, tree.Size())
}

func TestDefaultSize(t *testing.T) {
	var tree Tree[int, int]
	require.Equal(t, 0, tree.Size())
}

func TestNilFind(t *testing.T) {
	var tree *Tree[int, int]
	p := tree.Find(2)
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestEmptyFind(t *testing.T) {
	var tree Tree[int, int]
	p := tree.Find(2)
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestNilDelete(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Delete(2)
	require.True(t, tree2.IsEmpty())
}

func TestEmptyDelete(t *testing.T) {
	var tree Tree[int, int]
	tree2 := tree.Delete(2)
	require.True(t, tree2.IsEmpty())
}

func TestNilGLB(t *testing.T) {
	var tree *Tree[int, int]
	p := tree.GreatestLowerBound(2)
	require.True(t, p.IsEmpty())
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
}

func TestEmptyGLB(t *testing.T) {
	var tree Tree[int, int]
	p := tree.GreatestLowerBound(2)
	require.True(t, p.IsEmpty())
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
}

func TestNilLUB(t *testing.T) {
	var tree *Tree[int, int]
	p := tree.LeastUpperBound(2)
	require.True(t, p.IsEmpty())
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
}

func TestEmptyLUB(t *testing.T) {
	var tree Tree[int, int]
	p := tree.LeastUpperBound(2)
	require.True(t, p.IsEmpty())
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
}

func TestNilIter(t *testing.T) {
	var tree *Tree[int, int]
	i := tree.Iter()
	require.False(t, i.Next())
}

func TestEmptyIter(t *testing.T) {
	var tree Tree[int, int]
	i := tree.Iter()
	require.False(t, i.Next())
}

func TestNilIterGTE(t *testing.T) {
	var tree *Tree[int, int]
	i := tree.IterGte(5)
	require.False(t, i.Next())
}

func TestEmptyIterGTE(t *testing.T) {
	var tree Tree[int, int]
	i := tree.IterGte(5)
	require.False(t, i.Next())
}

func TestNilLeft(t *testing.T) {
	var tree *Tree[int, int]
	require.True(t, tree.Left().IsEmpty())
}

func TestEmptyLeft(t *testing.T) {
	var tree Tree[int, int]
	require.True(t, tree.Left().IsEmpty())
}

func TestNilRight(t *testing.T) {
	var tree *Tree[int, int]
	require.True(t, tree.Right().IsEmpty())
}

func TestEmptyRight(t *testing.T) {
	var tree Tree[int, int]
	require.True(t, tree.Right().IsEmpty())
}

func TestNilKey(t *testing.T) {
	var tree *Tree[int, int]
	require.Equal(t, 0, tree.Key())
}

func TestEmptyKey(t *testing.T) {
	var tree Tree[int, int]
	require.Equal(t, 0, tree.Key())
}

func TestNilValue(t *testing.T) {
	var tree *Tree[int, int]
	require.Equal(t, 0, tree.Value())
}

func TestEmptyValue(t *testing.T) {
	var tree Tree[int, int]
	require.Equal(t, 0, tree.Value())
}

func TestNilLeast(t *testing.T) {
	var tree *Tree[int, int]
	p := tree.Least()
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestEmptyLeast(t *testing.T) {
	var tree Tree[int, int]
	p := tree.Least()
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestNilMost(t *testing.T) {
	var tree *Tree[int, int]
	p := tree.Most()
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestEmptyMost(t *testing.T) {
	var tree Tree[int, int]
	p := tree.Most()
	if p == nil {
		require.Fail(t, "wtf?")
	}
	require.Equal(t, 0, p.Key())
	require.Equal(t, 0, p.Value())
	require.True(t, p.IsEmpty())
}

func TestNilUpdate(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Update(2, 3)
	require.NotEqual(t, tree, tree2)
	require.False(t, tree2.IsEmpty())
	require.Equal(t, 1, tree2.Size())
	require.Equal(t, 1, tree2.Height())
	require.Equal(t, 2, tree2.Key())
	require.Equal(t, 3, tree2.Value())
	p := tree2.Find(4)
	if p == nil {
		require.Fail(t, "wtf")
	}
	require.True(t, p.IsEmpty())
	p = tree2.Find(2)
	require.False(t, p.IsEmpty())
	require.Equal(t, 2, p.Key())
	require.Equal(t, 3, p.Value())
	p = tree2.Find(1)
	require.True(t, p.IsEmpty())
	i := tree2.Iter()
	require.True(t, i.Next())
	require.Equal(t, 2, i.Current().Key())
	require.Equal(t, 3, i.Current().Value())
	require.False(t, i.Next())
}

func TestEmptyUpdate(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Update(2, 3)
	require.NotEqual(t, tree, tree2)
	require.False(t, tree2.IsEmpty())
	require.Equal(t, 1, tree2.Size())
	require.Equal(t, 1, tree2.Height())
	require.Equal(t, 2, tree2.Key())
	require.Equal(t, 3, tree2.Value())
	p := tree2.Find(4)
	if p == nil {
		require.Fail(t, "wtf")
	}
	require.True(t, p.IsEmpty())
	p = tree2.Find(2)
	require.False(t, p.IsEmpty())
	require.Equal(t, 2, p.Key())
	require.Equal(t, 3, p.Value())
	p = tree2.Find(1)
	require.True(t, p.IsEmpty())
	i := tree2.Iter()
	require.True(t, i.Next())
	require.Equal(t, 2, i.Current().Key())
	require.Equal(t, 3, i.Current().Value())
	require.False(t, i.Next())
}

func TestUpdateReplace(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Update(2, 3).Update(0, 4).Update(4, 5)
	tree3 := tree2.Update(0, 10)
	require.NotEqual(t, tree, tree2)
	require.NotEqual(t, tree2, tree3)
	p := tree3.Find(0)
	require.False(t, p.IsEmpty())
	require.Equal(t, 0, p.Key())
	require.Equal(t, 10, p.Value())
	require.False(t, tree3.Left().IsEmpty())
	require.False(t, tree3.Right().IsEmpty())
	require.True(t, tree3.Left().Left().IsEmpty())
	require.True(t, tree3.Left().Right().IsEmpty())
	require.True(t, tree3.Right().Left().IsEmpty())
	require.True(t, tree3.Right().Right().IsEmpty())
}

func TestRotateLeft(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Update(1, 1).Update(2, 2).Update(3, 3)
	require.False(t, tree2.IsEmpty())
	require.Equal(t, 2, tree2.Key())
	require.Equal(t, 1, tree2.Left().Key())
	require.Equal(t, 3, tree2.Right().Key())
}

func TestRotateRightLeft(t *testing.T) {
	var tree *Tree[int, int]
	items := []int{10, 9, 15, 13, 16, 14}
	for _, i := range items {
		tree = tree.Update(i, i)
	}
	require.Equal(t, 13, tree.Value())
	require.Equal(t, 10, tree.Left().Value())
	require.Equal(t, 15, tree.Right().Value())
	require.Equal(t, 9, tree.Left().Left().Value())
	require.True(t, tree.Left().Right().IsEmpty())
	require.Equal(t, 14, tree.Right().Left().Value())
	require.Equal(t, 16, tree.Right().Right().Value())
}

func TestRotateRight(t *testing.T) {
	var tree *Tree[int, int]
	tree2 := tree.Update(2, 2).Update(1, 1).Update(0, 0)
	require.Equal(t, 1, tree2.Value())
	require.Equal(t, 0, tree2.Left().Value())
	require.Equal(t, 2, tree2.Right().Value())
}

func TestRotateLeftRight(t *testing.T) {
	var tree *Tree[int, int]
	items := []int{20, 21, 15, 17, 14, 19}
	for _, i := range items {
		tree = tree.Update(i, i)
	}
	require.Equal(t, 17, tree.Value())
	require.Equal(t, 15, tree.Left().Value())
	require.Equal(t, 20, tree.Right().Value())
	require.Equal(t, 14, tree.Left().Left().Value())
	require.True(t, tree.Left().Right().IsEmpty())
	require.Equal(t, 19, tree.Right().Left().Value())
	require.Equal(t, 21, tree.Right().Right().Value())
}

func TestLeastUpperBound(t *testing.T) {
	tree := EmptyTree[int, int]()
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
	}
	p := tree.LeastUpperBound(4)
	require.Equal(t, 4, p.Value())
	p = tree.LeastUpperBound(5)
	require.Equal(t, 6, p.Value())
	p = tree.LeastUpperBound(22)
	require.True(t, p.IsEmpty())
	p = tree.LeastUpperBound(-1)
	require.Equal(t, 0, p.Value())
}

func TestGreatestLowerBound(t *testing.T) {
	tree := EmptyTree[int, int]()
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
	}
	p := tree.GreatestLowerBound(4)
	require.Equal(t, 4, p.Value())
	p = tree.GreatestLowerBound(5)
	require.Equal(t, 4, p.Value())
	p = tree.GreatestLowerBound(22)
	require.Equal(t, 18, p.Value())
	p = tree.GreatestLowerBound(-1)
	require.True(t, p.IsEmpty())
}

func TestIter(t *testing.T) {
	tree := EmptyTree[int, int]()
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
	}
	iter := tree.Iter()
	j := 0
	for iter.Next() {
		require.Equal(t, j, iter.Current().Value())
		j += 2
	}
	require.Equal(t, 20, j)
}

func TestIterGte(t *testing.T) {
	tree := EmptyTree[int, int]()
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
	}
	iter := tree.IterGte(7)
	j := 8
	for iter.Next() {
		require.Equal(t, j, iter.Current().Value())
		j += 2
	}
	require.Equal(t, 20, j)
}

func TestDelete(t *testing.T) {
	tree := EmptyTree[int, int]()
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
	}
	t2 := tree.Delete(7)
	if tree != t2 {
		require.Fail(t, "delete missing changed tree")
	}
	t3 := tree.Delete(4).Delete(18).Delete(0).Delete(-1).Delete(22)
	require.False(t, t3.IsEmpty())
	require.Equal(t, 7, t3.Size())
	for i := 0; i < 20; i += 2 {
		tree = tree.Delete(i)
	}

	require.True(t, tree.IsEmpty())
}

func TestMarshalJson(t *testing.T) {
	tree := EmptyTree[int, int]()
	expected := make(map[int]int)
	for i := 0; i < 20; i += 2 {
		tree = tree.Update(i, i)
		expected[i] = i
	}
	serialized, err := json.Marshal(tree)
	require.NoError(t, err)
	var actual map[int]int
	err = json.Unmarshal(serialized, &actual)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestUnmarshalJson(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 20; i += 2 {
		m[i] = i
	}
	serialized, err := json.Marshal(m)
	require.NoError(t, err)
	var tree *Tree[int, int]
	err = json.Unmarshal(serialized, &tree)
	require.NoError(t, err)
	require.Equal(t, 10, tree.Size())
	iter := tree.Iter()
	i := 0
	for iter.Next() {
		require.Equal(t, i, iter.Current().Key())
		i += 2
	}
	require.Equal(t, 20, i)
}

func TestUnMarshalJsonStringKey(t *testing.T) {
	expected := map[string]int{
		"Hello": 2,
		"World": 4,
	}
	serialized, err := json.Marshal(expected)
	require.NoError(t, err)

	var x *Tree[string, int]
	err = json.Unmarshal(serialized, &x)
	require.NoError(t, err)
	require.Equal(t, 2, x.Find("Hello").Value())
	require.Equal(t, 4, x.Find("World").Value())
}
