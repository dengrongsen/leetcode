// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
// L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
//
// Related Topics 栈 递归 链表 双指针 👍 1170 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 拆解到数组里
	tmp := head
	var arr []*ListNode
	for ; head != nil; {
		arr = append(arr, head)
		head = head.Next
	}

	// 重排链表
	head = tmp
	s, d := 1, 0
	for i := 1; s+d <= len(arr); i++ {
		if i%2 == 0 { // 双数
			head.Next = arr[s]
			s++
		} else {
			head.Next = arr[len(arr)-d]
			d++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
