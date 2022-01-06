export function DoQuestion(nums: number[]): number[][] {
  let totalLength = nums.length;
  const totalArray: number[][] = [];

  for (let  i = 0; i < totalLength - 2; i++) {
    for (let j = i+1; j < totalLength - 1; j++) {
      for (let k = i+2; k < totalLength; k++) {
        if (i == j || j == k || i == k) {
          continue;
        }
        const first = nums[i];
        const second = nums[j];
        const third = nums[k];
        if ((first + second + third) === 0) {
          if (!alreadyExists(totalArray, [first, second, third])) {
            totalArray.push([first, second, third])
          }
        }
      }
    }
  }
  return totalArray;
}

function alreadyExists(total: number[][], newNums: number[]): boolean {
  let sortedNewNums = newNums.sort();
  for (let i = 0; i < total.length; i++) {
    let potential = total[i].sort();
    if (isSameArray(potential, sortedNewNums)) {
      return true;
    }
  }
  return false;
}

function isSameArray(first: number[], second: number[]): boolean {
  for (var i = 0; i < first.length; ++i) {
    if (first[i] !== second[i]) return false;
  }
  return true;
}
