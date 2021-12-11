import fs from 'fs';
import readline from 'readline';

class Octopuses {
  octopuses: number[][];
  totalFlashes: number;

  constructor(octopuses: number[][]) {
    this.octopuses = octopuses;
    this.totalFlashes = 0;
  }

  getTotalFlashes(): number {
    return this.totalFlashes;
  }

  arrayAlreadyHasArray(arr: number[][], subarr: number[]): boolean {
      for(var i = 0; i<arr.length; i++){
          let checker = false
          for(var j = 0; j<arr[i].length; j++){
              if(arr[i][j] === subarr[j]){
                  checker = true
              } else {
                  checker = false
                  break;
              }
          }
          if (checker){
              return true
          }
      }
      return false
  }

  grow() {
    for (let i = 0; i < this.octopuses.length; i++) {
      for (let j = 0; j < this.octopuses[i].length; j++) {
        this.octopuses[i][j]++;
      }
    }
  }

  flash() {
    let alreadyFlashed: number[][] = [];
    for (let i = 0; i < this.octopuses.length; i++) {
      for (let j = 0; j < this.octopuses[i].length; j++) {
        if (this.octopuses[i][j] > 9 && !this.arrayAlreadyHasArray(alreadyFlashed, [i,j])) {
          this.totalFlashes++;
          alreadyFlashed.push([i, j]);
          alreadyFlashed = this.growAdjacent(i, j, alreadyFlashed);
        }
      }
    }

    while (alreadyFlashed.some( function (a) { return a.length })) {
      let current = alreadyFlashed.pop();
      if (!current) {
        continue;
      }
      this.octopuses[current[0]][current[1]] = 0;
    }
  }

  growAdjacent(i: number, j: number, alreadyFlashed: number[][]): number[][] {
    const maxI = this.octopuses.length;
    const maxJ = this.octopuses[0].length;

    if (i !== 0) {
      alreadyFlashed = this.growAndAddToFlash(i-1, j, alreadyFlashed)
    }

    if (j !== 0) {
      alreadyFlashed = this.growAndAddToFlash(i, j-1, alreadyFlashed)
    }

    if (i !== 0 && j !== 0) {
      alreadyFlashed = this.growAndAddToFlash(i-1, j-1, alreadyFlashed)
    }

    if (i !== maxI-1) {
      alreadyFlashed = this.growAndAddToFlash(i+1, j, alreadyFlashed)
    }

    if (j !== maxJ-1) {
      alreadyFlashed = this.growAndAddToFlash(i, j+1, alreadyFlashed)
    }

    if (i !== maxI-1 && j !== maxJ-1) {
      alreadyFlashed = this.growAndAddToFlash(i+1, j+1, alreadyFlashed)
    }

    if (i !== 0 && j !== maxJ-1) {
      alreadyFlashed = this.growAndAddToFlash(i-1, j+1, alreadyFlashed)
    }

    if (i !== maxI-1 && j !== 0) {
      alreadyFlashed = this.growAndAddToFlash(i+1, j-1, alreadyFlashed)
    }

    return alreadyFlashed;
  }

  growAndAddToFlash(i: number, j: number, alreadyFlashed: number[][]): number[][] {
    this.octopuses[i][j]++;
    if (this.octopuses[i][j] > 9 && !this.arrayAlreadyHasArray(alreadyFlashed, [i, j])) {
      this.totalFlashes++;
      alreadyFlashed.push([i, j]);
      this.growAdjacent(i, j, alreadyFlashed);
    }

    return alreadyFlashed;
  }
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let octopuses: number[][] = [];

  for await (const line of rl) {
    const lineSplit = line.split('').map(Number);
    octopuses.push(lineSplit);
  }

  const octo = new Octopuses(octopuses);

  console.log(octo);

  for ( let i = 0; i < 100; i++) {
    octo.grow();
    octo.flash();
  }

  console.log(octo);

  return octo.getTotalFlashes();
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
