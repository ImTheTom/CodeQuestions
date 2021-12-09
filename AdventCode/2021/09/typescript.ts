import fs from 'fs';
import readline from 'readline';

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let lines: number[][] = [];

  for await (const line of rl) {
    lines.push(line.split('').map(Number));
  }

  let totalRisk = 0;

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
        console.log(compareValues);
        console.log(currentValue);
        totalRisk += (currentValue+1)
      }
    }
  }

  return totalRisk;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
