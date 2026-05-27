package tree

type TreeData interface {
	GetName() string
}

// TreeNode represents a custom data type that contains data along with it's dependent TreeNodes
type TreeNode struct {
	Data       TreeData
	Dependents []*TreeNode
}

// GetAllNodes returns level order traversal of tree starting from current node
func (t *TreeNode) GetAllNodes() []*TreeNode { _ = "STUB: not implemented"; return nil }

func (t *TreeNode) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TreeNode) AddDependent(depNode *TreeNode) *TreeNode { _ = "STUB: not implemented"; return nil }

// NewTreeNode creates an instance of TreeNode
func NewTreeNode(data TreeData) *TreeNode { _ = "STUB: not implemented"; return nil }
