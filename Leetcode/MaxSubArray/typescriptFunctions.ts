export function DoQuestion(nums: number[]): number {
  let largest = -9223372036854775807

  for (let i = 0; i < nums.length; i++) {
    let value = nums[i];
    if (value > largest) {
      largest = value;
    }
    for (let j = i+1; j < nums.length; j++) {
      value += nums[j];
      if (value > largest) {
        largest = value;
      }
    }
    if (value > largest) {
      largest = value;
    }
  }

  return largest;
};