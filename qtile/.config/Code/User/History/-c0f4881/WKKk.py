from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        l1_val = 0
        l2_val = 0
        while l1:
            l1_val = l1_val * 10 + l1.val
            l1 = l1.next

        while l2:
            l2_val = l2_val * 10 + l2.val
            l2 = l2.next

        sum_val = l1_val + l2_val
        sum_val_str = str(sum_val)
        sum_val_list = list(sum_val_str)
        sum_val_list.reverse()

        sum_val_list = [int(i) for i in sum_val_list]
        sum_val_list = [ListNode(i) for i in sum_val_list]
        return sum_val_list


if __name__ == "__main__":
    l1 = ListNode(2, ListNode(4, ListNode(3)))
    l2 = ListNode(5, ListNode(6, ListNode(4)))
    s = Solution()
    result = s.addTwoNumbers(l1, l2)
    print(result)
