package tree

import (
	"errors"
)

// ErrCyclicDependencyEncountered is triggered a tree has a cyclic dependency
var ErrCyclicDependencyEncountered = errors.New("a cycle dependency encountered in the tree")

// MultiRootTree - represents a data type which has multiple independent root nodes
// all root nodes have their independent tree based on depdencies of TreeNode.
// it also maintains a map of nodes for faster lookups and managing node data.
type MultiRootTree struct {
	rootNodes []string
	dataMap   map[string]*TreeNode
}

func (t *MultiRootTree) GetRootNodes() []*TreeNode { _ = "STUB: not implemented"; return nil }

// MarkRoot marks a node as root
func (t *MultiRootTree) MarkRoot(node *TreeNode) { _ = "STUB: not implemented"; return }

func (t *MultiRootTree) AddNode(node *TreeNode) { _ = "STUB: not implemented"; return }

func (t *MultiRootTree) AddNodeIfNotExist(node *TreeNode) { _ = "STUB: not implemented"; return }

func (t *MultiRootTree) GetNodeByName(dagName string) (*TreeNode, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// get sorted nodeNames from dataMap
func (t *MultiRootTree) getSortedNodeNames() []string { _ = "STUB: not implemented"; return nil }

// ValidateCyclic - detects if there are any cycles in the tree
func (t *MultiRootTree) ValidateCyclic() ([]string, error) {
	_ = "STUB: not implemented"
	// runs a DFS on a given tree using visitor pattern
	return nil, nil
}

func prettifyPaths(paths []string) string { _ = "STUB: not implemented"; return "" }

// NewMultiRootTree returns an instance of multi root dag tree
func NewMultiRootTree() *MultiRootTree { _ = "STUB: not implemented"; return nil }
