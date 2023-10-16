from typing import List


def maxSubArray(nums: List[int]) -> int:

    total = 0

    for i, v in enumerate(nums):
        total += v
        if (len(nums) > i+1):
            if (total < total + nums[i+1]):
                pass

        print(total + nums[i+1])
    return


nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
print(maxSubArray(nums))
