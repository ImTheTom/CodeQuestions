import fs from 'fs';
import readline from 'readline';

async function process(input: number[]): Promise<number> {
  const max = Math.max(...input);
  const min = Math.min(...input);
  let totalCost = 9223372036854775807;
  let currentCost = 0;
  for (let i = min; i <= max; i++) {
    for (let j = 0; j < input.length; j++) {
      currentCost += Math.abs(i - input[j]);
    }
    if (currentCost < totalCost) {
      totalCost = currentCost;
    }
    currentCost = 0;
  }
  return totalCost;
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

  return await process(numericalInput);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
