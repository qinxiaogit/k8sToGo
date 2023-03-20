package main

import (
	"fmt"
	"math"
)

func twoSum(nums []int, target int) []int {

	for k, v := range nums {
		curr := target - v

		for j := k + 1; j < len(nums); j++ {
			if curr == nums[j] {
				return []int{k, j}
			}
		}

	}
	return []int{0, 0}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carr int

	var node *ListNode
	var head *ListNode
	for l1 != nil || l2 != nil {
		l1Number, l2Number := 0, 0
		if l1 != nil {
			l1Number = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Number = l2.Val
			l2 = l2.Next
		}
		number := l1Number + l2Number + carr
		carr = 0
		if number > 9 {
			carr = 1
			number = number % 10
		}
		if head == nil {
			head = &ListNode{Val: number, Next: nil}
			node = head
		} else {
			head.Next = &ListNode{Val: number}
			head = head.Next
		}
	}
	if carr > 0 && head != nil {
		head.Next = &ListNode{Val: carr}
	}
	return node
}

func main() {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	res := addTwoNumbers(l1, l2)
	fmt.Println("-------------")
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func getHint(secret string, guess string) string {
	if len(secret) != len(guess) {
		return ""
	}
	bullsCount := 0
	cowsCount := 0

	sgMap := make(map[uint8]int)

	for i, _ := range secret {
		if secret[i] == guess[i] {
			bullsCount++
		} else {
			if sgMap[secret[i]] > 0 {

			} else {
				sgMap[secret[i]]++
			}
		}
	}
}

func isValid(s string) bool {
	sl := len(s)
	if sl == 0 {
		return true
	}
	var nums []byte
	for i := 0; i < sl; i++ {
		switch s[i] {
		case '(':
			nums = append(nums, ')')
		case '{':
			nums = append(nums, '}')
		case '[':
			nums = append(nums, ']')
		default:
			ns := len(nums)
			if nums != nil && (ns == 0 || nums[ns-1] != byte(s[i])) {
				return false
			}
			nums = nums[:ns-1]
		}

	}
	math.Pow()
	return len(nums) == 0
}
