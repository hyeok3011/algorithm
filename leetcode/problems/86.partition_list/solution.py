# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        start_pre_node = None
        current = head 
        while current != None:
            if current.val >= x:
                if start_pre_node == None:
                    start_pre_node = pre
                pre = current
                current = current.next
            else:
                if start_pre_node != None:
                    dummy_next = current.next
                    current.next = start_pre_node.next
                    start_pre_node.next = current
                    start_pre_node = current
                    current = dummy_next
                else:
                    pre = current
                    current = pre.next
        
        return head




