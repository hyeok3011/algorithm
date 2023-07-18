# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
# untime 57 ms Beats 56.71% Memory 16.3 MB Beats 57.14%

class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(val=-101)
        dummy.next = head
        
        pre_node, current_node = dummy, head

        while (current_node and current_node.next):
            if current_node.val == current_node.next.val:
                while (current_node.next and current_node.val == current_node.next.val):
                    current_node.next = current_node.next.next
                pre_node.next = current_node.next
            else:
                pre_node = current_node
            current_node = current_node.next

        return dummy.next