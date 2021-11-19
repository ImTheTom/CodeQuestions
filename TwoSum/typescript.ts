function twoSum(nums: number[], target: number): number[] {
	const numLength = nums.length;
	for (let i = 0; i < numLength-1; i++) {
		const current = nums[i];

		for (let j = i + 1; j < numLength; j++) {
			const otherValue = nums[j];
			if (current + otherValue === target) {
				return [i, j];
			}
		}
	}

	throw "Failed to find target value";
};

const nums = [-1,-2,-3,-4,-5];
const target = -8;
const expected = [2, 4];

const returnedValue = twoSum(nums, target);

console.log(`Got ${returnedValue} Expected ${expected}`);
