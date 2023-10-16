function containsDuplicate(nums: number[]): boolean {
  const hashset = new Set();

  for (let i = 0; i < nums.length; i++) {
    if (hashset.has(nums[i])) return true;
    else hashset.add(nums[i]);
  }

  return false;
}

console.log(containsDuplicate([1, 2, 3, 1]));
