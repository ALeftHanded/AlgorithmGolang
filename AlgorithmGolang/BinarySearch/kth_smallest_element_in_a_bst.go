package BinarySearch

// * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

// * 树中的节点数为 n 。
// * 1 <= k <= n <= 10^4
// * 0 <= Node.val <= 10^4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	return 0
}
