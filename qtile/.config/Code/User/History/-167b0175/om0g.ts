function RunningSum(nums: number[]): number[] {
  let sum = 0;
  return nums.map((num) => (sum += num));
}

console.log(RunningSum([1, 2, 3, 4]));
