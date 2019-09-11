package semvertree

import (
	"github.com/blang/semver"
)

// Value is an element in tree
type Value struct {
	Version string
	Data    interface{}
}

// VersionTree is a wrapper struct for all functions
type VersionTree struct {
	Root *Tree
}

// A Tree is a binary tree with values.
type Tree struct {
	Left  *Tree
	Value *Value
	Right *Tree
}

// insert triggers a recursive insert method
func insert(t *Tree, value *Value) *Tree {
	if t == nil {
		return &Tree{
			Left:  nil,
			Value: value,
			Right: nil,
		}
	}
	diff := CompareVersions(value.Version, t.Value.Version)
	degreeDiff := DegreeOfDifference(value.Version, t.Value.Version)
	degreeDiffRight := 0
	if t.Right != nil && t.Right.Value != nil {
		degreeDiffRight = DegreeOfDifference(t.Value.Version, t.Right.Value.Version)
	}
	if diff == 0 {
		return t
	}
	if diff < 0 || degreeDiffRight > degreeDiff {
		t.Left = insert(t.Left, value)
		return t
	}

	t.Right = insert(t.Right, value)
	return t
}

// New returns just an empty instance
func New() *VersionTree {
	return &VersionTree{}
}

// Add allow insterting a data
func (b *VersionTree) Add(value *Value) {
	if b.Root == nil {
		b.Root = &Tree{nil, value, nil}
	} else {
		insert(b.Root, value)
	}
}

// AddVersion allow insterting a data
func (b *VersionTree) AddVersion(version string, data interface{}) {
	b.Add(
		&Value{
			Version: version,
			Data:    data,
		},
	)
}

// Walk lets you iterate over whole data
func (b *VersionTree) Walk(step func(prev, next *Value, degreeOfDifference int)) {
	if b.Root == nil {
		return
	}
	walk(b.Root, nil, step)
}

// walk is recursive function needed for Walk
func walk(t *Tree, prev *Value, step func(prev, next *Value, degreeOfDifference int)) {
	if t == nil {
		return
	}
	var prevVersion string
	if prev != nil {
		prevVersion = prev.Version
	} else {
		prevVersion = t.Value.Version
	}
	walk(t.Left, t.Value, step)
	step(prev, t.Value, DegreeOfDifference(prevVersion, t.Value.Version))
	walk(t.Right, t.Value, step)
}

func CompareVersions(a, b string) int {
	v, _ := semver.ParseTolerant(a)
	o, _ := semver.ParseTolerant(b)
	return v.Compare(o)
}

func DegreeOfDifference(a, b string) int {
	v, _ := semver.ParseTolerant(a)
	o, _ := semver.ParseTolerant(b)

	if v.Major != o.Major {
		return 1
	}
	if v.Minor != o.Minor {
		return 2
	}
	if v.Patch != o.Patch {
		return 3
	}

	// Quick comparison if a version has no prerelease versions
	if len(v.Pre) != len(o.Pre) {
		return 4
	}

	i := 0
	for ; i < len(v.Pre) && i < len(o.Pre); i++ {
		if comp := v.Pre[i].Compare(o.Pre[i]); comp == 0 {
			continue
		} else {
			return 4
		}
	}
	return 0
}

// Rebuild with sorted versions
func (b *VersionTree) Rebuild() {
	versions := make([]*Value, 0)
	b.Walk(func(prev, value *Value, degreeOfDifference int) {
		versions = append(versions, value)
	})
	b.Root = nil
	for _, value := range versions {
		b.Add(value)
	}
}
