import fs from 'fs';
import readline from 'readline';

async function buildFullBinary(input: Array<number>, total: number, mostCommon: boolean): Promise<string> {
  let result: string = '';
  for (let i = 0; i < input.length; i++) {
    const numberOfOnes = input[i];
    if (numberOfOnes > (total/2)) {
      result = mostCommon ? result.concat('1') : result.concat('0');
    } else {
      result = mostCommon ? result.concat('0') : result.concat('1');
    }
  }
  return result;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let numberOfOnes = new Array(12);
  let firstRun = true;
  let total = 0;

  for await (const line of rl) {

    for (let i = 0; i < line.length; i++) {
      if (firstRun) {
        numberOfOnes[i] = 0;
      }
      if (line[i] === '1') {
        numberOfOnes[i]++;
      }
    }

    firstRun = false;
    total++;
  }

  const gammaB = await buildFullBinary(numberOfOnes, total, true);
  const epsilonB = await buildFullBinary(numberOfOnes, total, false);

  const gamma = parseInt(gammaB, 2)
  const epsilon = parseInt(epsilonB, 2)

  console.log(gammaB);
  console.log(epsilonB);

  console.log(gamma);
  console.log(epsilon);

  return gamma * epsilon;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
