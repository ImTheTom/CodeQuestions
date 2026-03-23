export function DoQuestion(nums: number[]): number[][] {
  let sortedNums = nums.sort();
  let totalLength = nums.length;
  const totalArray: number[][] = [];

  if (totalLength < 3) {
    return totalArray;
  }

  for (let  i = 0; i < totalLength - 1; i++) {
    for (let j = i+1; j < totalLength; j++) {
      const first = nums[i];
      const second = nums[j];
      const third = -(second + first);
      const thirdLoc = nums.indexOf(third)

      if (thirdLoc !== -1 && thirdLoc != j && thirdLoc != i) {
        if (!alreadyExists(totalArray, [first, second, third])) {
          totalArray.push([first, second, third].sort())
        }
      }
    }
  }

  return totalArray
}

function alreadyExists(total: number[][], newNums: number[]): boolean {
  let sortedNewNums = newNums.sort();
  for (let i = 0; i < total.length; i++) {
    let potential = total[i];
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