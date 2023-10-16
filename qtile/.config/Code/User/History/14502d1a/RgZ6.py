

from typing import List


def pivotIndex(nums: List[int]) -> int:
    total = sum(nums)
    leftSum = 0
    for i, num in enumerate(nums):
        if leftSum == total - leftSum - num:
            return i
        leftSum += num
    return -1


print(pivotIndex([1, 7, 3, 6, 5, 6]))
