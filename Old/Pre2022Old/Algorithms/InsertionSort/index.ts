// Is typically O(n^2)

function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}

let myArray = new Array();
for (let i = 0; i <= 50; i++) {
	myArray.push(getRandomInt(50));
}

console.log(myArray);

function InsertionSort(arr: number[]): number[] {
	let i = 1;
	while (i < arr.length) {
		let j = i;
		while (j > 0 && arr[j - 1] > arr[j]) {
			let	tmpA = arr[j];
			arr[j] = arr[j-1];
			arr[j - 1] = tmpA;
			j = j - 1;
		}
		i = i + 1;
	}
	return arr;
}

myArray = InsertionSort(myArray);

console.log(myArray);