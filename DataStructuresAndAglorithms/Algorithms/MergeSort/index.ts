function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

let myArray = new Array();
for (let i = 0; i <= 50; i++) {
	myArray.push(getRandomInt(50));
}

console.log(myArray);

function sortMerge(arrayA: number[], arrayB: number[]): number[] {
	let merged: number[] = new Array();
	while(arrayA.length > 0 && arrayB.length > 0) {
		if (arrayA[0] <= arrayB[0]) {
			merged.push(arrayA.shift()!);
		} else {
			merged.push(arrayB.shift()!);
		}
	}

	while (arrayA.length > 0) {
		merged.push(arrayA.shift()!);
	}

	while (arrayB.length > 0) {
		merged.push(arrayB.shift()!);
	}

	return merged;
}

function mergeSort(arr: number[]): number[] {
	if (arr.length <= 1) {
		return arr;
	}

	let middle = Math.ceil(arr.length / 2);
	let leftArr = arr.slice(0, middle);
	let rightArr = arr.slice(leftArr.length, arr.length);

	return sortMerge(mergeSort(leftArr), mergeSort(rightArr));
}

myArray = mergeSort(myArray);

console.log(myArray);