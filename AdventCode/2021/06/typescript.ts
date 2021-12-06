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

  for (let i = 0; i < 80; i++) {
    for (let j = 0; j < fishArray.length; j++) {
      if (fishArray[j].grow()) {
        newFishCount++;
      }
    }

    while (newFishCount > 0) {
      let tmp = new fish(8);
      fishArray.push(tmp);
      newFishCount--;
    }
  }

  return fishArray.length;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
