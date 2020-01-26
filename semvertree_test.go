package semvertree_test

import (
	"fmt"
	"testing"

	semvertree "github.com/bgokden/semver-tree"
)

func TesSemverTreeBasic(t *testing.T) {
	st := semvertree.New()
	st.AddVersion("v1.0.1", map[string]string{"rule": "test 3"})
	st.AddVersion("v2.0.0", map[string]string{"rule": "test 2"})
	st.AddVersion("v1.0.0", map[string]string{"rule": "test 1"})
	st.AddVersion("v2.0.1", map[string]string{"rule": "test 2"})
	st.AddVersion("v2.0.3", map[string]string{"rule": "test 2"})
	st.Rebuild()
	fmt.Printf("----\n")
	st.Walk(func(prev, value *semvertree.Value, degreeOfDifference int) {
		if prev == nil {
			fmt.Printf("root: %v\n", value.Version)
			return
		}
		fmt.Printf("prev: %v next: %v degreeOfDifference: %v data: %v\n", prev.Version, value.Version, degreeOfDifference, value.Data)
	})
	fmt.Printf("----\n")
}

func TestSemverTreeBasic2(t *testing.T) {
	st := semvertree.New()
	st.AddVersion("v1.0.1", map[string]string{"rule": "test 3"})
	st.AddVersion("v2.0.0", map[string]string{"rule": "test 2"})
	st.AddVersion("v1.0.0", map[string]string{"rule": "test 1"})
	st.AddVersion("v2.0.1", map[string]string{"rule": "test 2"})
	st.AddVersion("v2.0.3", map[string]string{"rule": "test 2"})
	st.AddVersion("v1.0.3", map[string]string{"rule": "test 4"})
	st.AddVersion("v1.1.0", map[string]string{"rule": "test 4"})
	st.AddVersion("v1.1.2", map[string]string{"rule": "test 5"})
	st.Rebuild()
	fmt.Printf("----\n")
	st.Walk(func(prev, value *semvertree.Value, degreeOfDifference int) {
		if prev == nil {
			fmt.Printf("root: %v\n", value.Version)
			return
		}
		fmt.Printf("prev: %v next: %v degreeOfDifference: %v data: %v\n", prev.Version, value.Version, degreeOfDifference, value.Data)
	})
	fmt.Printf("----\n")
}

func TestSemverTreeBasic3(t *testing.T) {
	st := semvertree.New()
	st.AddVersion("v1.0.1", map[string]string{"rule": "test 3"})
	st.AddVersion("v2.0.0", map[string]string{"rule": "test 2"})
	st.AddVersion("v1.0.0", map[string]string{"rule": "test 1"})
	st.AddVersion("v2.0.1", map[string]string{"rule": "test 2"})
	st.AddVersion("v2.0.3", map[string]string{"rule": "test 2"})
	st.AddVersion("v1.0.3", map[string]string{"rule": "test 4"})
	st.AddVersion("v1.1.0", map[string]string{"rule": "test 4"})
	st.AddVersion("v1.1.2", map[string]string{"rule": "test 5"})
	st.AddVersion("v3.0.1", map[string]string{"rule": "test 5"})
	st.Rebuild()
	fmt.Printf("----\n")
	st.Walk(func(prev, value *semvertree.Value, degreeOfDifference int) {
		if prev == nil {
			fmt.Printf("root: %v\n", value.Version)
			return
		}
		fmt.Printf("prev: %v next: %v degreeOfDifference: %v data: %v\n", prev.Version, value.Version, degreeOfDifference, value.Data)
	})
	fmt.Printf("----\n")
}
