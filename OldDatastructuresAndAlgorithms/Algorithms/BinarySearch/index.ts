// Works on sorted arrays

let myArray = new Array();
for (let i = 0; i <= 100; i++) {
	myArray.push(i);
}

function BinarySearch(arr: number[], target: number): number {
	let L = 0;
	let R = arr.length - 1;
	while (L <= R) {
		let m = Math.floor((L+R) / 2);
		if (arr[m] < target) {
			L = m + 1;
		} else if (arr[m] > target) {
			R = m - 1;
		} else {
			return m;
		}
	}

	return -1;
}

console.log(myArray);

console.log(BinarySearch(myArray, 5));
console.log(BinarySearch(myArray, -1));
console.log(BinarySearch(myArray, 70));
console.log(BinarySearch(myArray, 700));