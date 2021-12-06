import fs from 'fs';
import readline from 'readline';

class fish {
  days = 0;
  constructor(d: number) {
    this.days = d;
  }

  grow(): boolean {
    this.days--;
    if (this.days < 0) {
      this.days = 6;
      return true;
    }
    return false;
  }
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

  let fishArray: fish[] = [];

  for (let i = 0; i < numericalInput.length; i++) {
    let tmp = new fish(numericalInput[i]);
    fishArray.push(tmp);
  }

  let newFishCount = 0;

  let tmpFishArray: fish[] = [];

  let totalFishCount = 0;

  for (let j = fishArray.length-1; j >= 0; j--) {
    let tmp = fishArray.pop();
    if (!tmp) {
      continue;
    }
    tmpFishArray.push(tmp);
    for (let i = 0; i < 64; i++) {
      console.log(tmpFishArray.length);
      for (let k = 0; k < tmpFishArray.length; k++) {
        if (tmpFishArray[k].grow()) {
          newFishCount++;
        }

        while (newFishCount > 0) {
          let tmp = new fish(8);
          tmpFishArray.push(tmp);
          newFishCount--;
        }
      }
    }
    totalFishCount += tmpFishArray.length;
    console.log(totalFishCount);
    tmpFishArray = [];
  }
  return (2*totalFishCount)*3;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
