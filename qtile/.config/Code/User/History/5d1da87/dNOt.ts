function pivotIndex(nums: number[]): number {
  const Total = nums.reduce((a: number, b: number) => a + b);
  let leftNum = 0;

  for (let i = 0; i < nums.length; i++) {
    if (leftNum === Total - nums[i] - leftNum) return i;
  }
  return -1;
}

console.log(pivotIndex([1, 7, 3, 6, 5, 6]));
