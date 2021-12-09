import fs from 'fs';
import readline from 'readline';

function arrayAlreadyHasArray(arr: number[][], subarr: number[]){
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

async function exploreBasin(point: number[], lines: number[][]): Promise<number> {
  let queue: number[][] = [];
  let explored: number[][] = [];
  queue.push(point);

  let total = 0;

  while (queue.some( function (a) { return a.length })) {
    let current = queue.pop();
    console.log(current);
    if (!current) {
      break;
    }

    let x = current[0] ?? 0;
    let y = current[1] ?? 0;
    let value = lines[x][y];

    if(value == 9) {
      continue;
    }

    total++;

    if (x != 0) {
      let tmp = [x-1, y];
      if (!arrayAlreadyHasArray(queue, tmp) && !(arrayAlreadyHasArray(explored, tmp))) {
        queue.push(tmp);
      }
    }
    if (y != 0) {
      let tmp = [x, y-1];
      if (!arrayAlreadyHasArray(queue, tmp) && !(arrayAlreadyHasArray(explored, tmp))) {
        queue.push(tmp);
      }
    }
    if (y != lines[x].length-1) {
      let tmp = [x, y+1];
      if (!arrayAlreadyHasArray(queue, tmp) && !(arrayAlreadyHasArray(explored, tmp))) {
        queue.push(tmp);
      }
    }
    if (x != lines.length-1) {
      let tmp = [x+1, y];
      if (!arrayAlreadyHasArray(queue, tmp) && !(arrayAlreadyHasArray(explored, tmp))) {
        queue.push(tmp);
      }
    }

    explored.push(current);
  }

  return total;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let lines: number[][] = [];

   let lowpoints: number[][] = [];

  for await (const line of rl) {
    lines.push(line.split('').map(Number));
  }

  for (let i = 0; i < lines.length; i++) {
    for (let j = 0; j < lines[i].length; j++) {
      let currentValue = lines[i][j];
      let compareValues: number[] = [];
      if (j != 0) {
        compareValues.push(lines[i][j-1]);
      }
      if (i != 0) {
        compareValues.push(lines[i-1][j]);
      }

      if (j != lines[i].length-1) {
        compareValues.push(lines[i][j+1]);
      }

      if (i != lines.length-1) {
        compareValues.push(lines[i+1][j]);
      }

      let low = true;
      for await (const otherValue of compareValues) {
        if (otherValue <= currentValue) {
          low = false;
        }
      }

      if (low) {
        lowpoints.push([i,j])
      }
    }
  }

  let basinDepths: number[] = [];

  for await (const lp of lowpoints) {
    basinDepths.push(await exploreBasin(lp, lines));
  }

  basinDepths.sort(function(a, b) {
    return a - b;
  });

  let lastThree = basinDepths.slice(basinDepths.length-3, basinDepths.length)

  console.log(lastThree);

  return lastThree.reduce((a,b) => a*b, 1);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
