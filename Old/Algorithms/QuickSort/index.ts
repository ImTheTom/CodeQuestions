function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

let myArray = new Array();
for (let i = 0; i <= 50; i++) {
	myArray.push(getRandomInt(50));
}

console.log(myArray);

function partition(arr: number[], low: number, high: number): number {
	let pivot = arr[Math.floor((low+high) / 2)];
	let leftIndex = low;
	let rightIndex = high;

	while (leftIndex <= rightIndex) {
		while (arr[leftIndex] < pivot) {
			leftIndex++;
		}

		while (arr[rightIndex] > pivot) {
			rightIndex--;
		}

		if ( leftIndex <= rightIndex) {
			let	tmpR = arr[rightIndex];
			arr[rightIndex] = arr[leftIndex];
			arr[leftIndex] = tmpR;

			leftIndex++;
			rightIndex--;
		}
	}
	return leftIndex;
}

function quicksort(arr: number[], low: number, high: number): number[] {
	if (arr.length > 1) {
		let index = partition(arr, low, high);

		if (low < index - 1) {
			quicksort(arr, low, index - 1);
		}

		if (index < high) {
			quicksort(arr, index, high);
		}
	}
	return arr;
}

myArray = quicksort(myArray, 0, myArray.length - 1);

console.log(myArray);