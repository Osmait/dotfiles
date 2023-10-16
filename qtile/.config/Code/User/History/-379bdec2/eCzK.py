def runningSum(nums):
    newList = []
    for n in nums:
        if len(newList) == 0:
            newList.append(n)
        else:
            newList.append(n + newList[-1])

    return newList


print(runningSum([1, 2, 3, 4]))
