function containsDuplicate(nums: number[]): boolean {
  const hashset = new Set();
  for (const n in nums) {
    if (hashset.has(n)) {
      console.log(hashset.has(n));
      return true;
    }
    hashset.add(n);
  }
  return false;
}

console.log(containsDuplicate([1, 2, 3, 1]));
