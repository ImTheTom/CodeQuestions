import fs from 'fs';
import readline from 'readline';

async function count(input: number[]): Promise<number> {
  let total = 0;
  for (let i = 0; i < input.length; i++) {
    total += input[i];
  }
  return total;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  var input;

  for await (const line of rl) {
    input = line;
  }

  if (input === undefined) {
    return 0;
  }

  var numericalInput = input.split(',').map(Number);

  let fishArray: number[] = new Array(9).fill(0);

  for (let i = 0; i < numericalInput.length; i++) {
    fishArray[numericalInput[i]]++;
  }

  console.log(fishArray);

  let shiftDown = 0;
  let shiftDown2 = 0;

  let totalResetFishes = 0;

  for (let i = 0; i<256; i++) {
    totalResetFishes = fishArray[0];

    shiftDown = 0;
    shiftDown2 = 0;

    for (let j = fishArray.length-1; j >= 0; j--){
      if(j == fishArray.length-1) {
        shiftDown = fishArray[j];
      } else {
        shiftDown2 = fishArray[j];
        fishArray[j] = shiftDown;
        shiftDown = shiftDown2;
      }
    }

    fishArray[6] += totalResetFishes;
    fishArray[8] = totalResetFishes;

    console.log(fishArray);
  }

  return await count(fishArray);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
