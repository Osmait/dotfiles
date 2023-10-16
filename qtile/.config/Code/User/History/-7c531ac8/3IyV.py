class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# class Solution(object):
#     def mergeTwoLists(self, list1, list2):
#         """
#         :type list1: Optional[ListNode]
#         :type list2: Optional[ListNode]
#         :rtype: Optional[ListNode]
#         """
#         if list1 is None:
#             return list2
#         if list2 is None:
#             return list1
#         if list1.val < list2.val:
#             list1.next = self.mergeTwoLists(list1.next, list2)
#             return list1
#         else:
#             list2.next = self.mergeTwoLists(list1, list2.next)
#             return list2


def mergeTwoLists(l1: ListNode, l2: ListNode) -> ListNode:
    # create a dummy node to start the merged list
    dummy = ListNode(0)
    # create a pointer to traverse the merged list
    ptr = dummy

    # traverse both lists and compare values of each node
    while l1 and l2:
        if l1.val < l2.val:
            ptr.next = l1
            l1 = l1.next
        else:
            ptr.next = l2
            l2 = l2.next
        ptr = ptr.next

    # if one of the lists is exhausted, append the rest of the other list
    if l1:
        ptr.next = l1
    else:
        ptr.next = l2

    # return the head of the merged list (after the dummy node)
    return dummy.next


lis1 = ListNode(1)
lis1.next = ListNode(2)
lis1.next.next = ListNode(4)
lis2 = ListNode(1)
lis2.next = ListNode(3)
lis2.next.next = ListNode(4)
print(mergeTwoLists(lis1, lis2).val)
