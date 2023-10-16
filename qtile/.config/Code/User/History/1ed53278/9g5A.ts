function twoSum(nums: number[], target: number): number[] {
  let map: Map<number, number> = new Map();
  for (let i = 0; i < nums.length; i++) {
    let diff = target - nums[i];
    if (map.has(diff)) {
      return [map.get(diff)!, i];
    }
    map.set(nums[i], i);
  }
  return [];
}

console.log(twoSum([3, 2, 4], 6));
