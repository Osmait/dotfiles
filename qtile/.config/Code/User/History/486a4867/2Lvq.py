from typing import List


def maxSubArray(nums: List[int]) -> int:
    sum_dict = {}
    total = 0

    for i, v in enumerate(nums):
        total += v
        if (total < total + nums[i+1]):
            sum_dict[total] = total
            total = 0

    return sum_dict
