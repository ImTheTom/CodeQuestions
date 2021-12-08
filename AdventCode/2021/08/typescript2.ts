import fs from 'fs';
import readline from 'readline';

const uniqueValues = [2,4,3,7];

async function processInputs(inputs: string[][], outputs: string[][]): Promise<number> {
  let total = 0;
  for (let i = 0; i < inputs.length; i++) {
    for (let j = 0; j < outputs[i].length; j++) {
      if (uniqueValues.includes(outputs[i][j].length)) {
        total++;
      }
    }
  }
  return total;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  var lines: string[] = [];

  for await (const line of rl) {
    lines.push(line);
  }

  var inputs: string[][] = [];
  var outputs: string[][] = [];

  for (const line of lines) {
    var splitValue = line.split('|');
    var input = splitValue[0].split(' ').filter(Boolean);
    inputs.push(input);

    var output = splitValue[1].split(' ').filter(Boolean);
    outputs.push(output);
  }

  const result = await processInputs(inputs, outputs);

  return result;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
