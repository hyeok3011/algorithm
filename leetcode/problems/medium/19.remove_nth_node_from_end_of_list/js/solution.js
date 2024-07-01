/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
// two pointer 방식으로 풀면 될듯 하다. 레퍼런스이긴 하나 array에 모두 담아두는것은 비효율적
var removeNthFromEnd = function(head, n) {
    let first = head;
    let second = head;
    let i = 0;
    while (first) {
        if (i > n) {
            second = second.next;
        }
        
        first = first.next;
        i++;
    }

    if (n == i) {
        head = head.next;
    } else {
        second.next = second.next.next;
    }

    return head
};

