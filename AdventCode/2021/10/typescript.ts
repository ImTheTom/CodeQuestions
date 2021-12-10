import fs from 'fs';
import readline from 'readline';

let openers = ['(', '[', '{', '<'];
let closers = [')', ']', '}', '>'];
let values = [3, 57, 1197, 25137];
let totalInvalid = [0, 0, 0, 0];

async function calculateInvalidScore(): Promise<number> {
  let total = 0;
  console.log(totalInvalid);
  for(let i = 0; i < totalInvalid.length; i++) {
    total += (totalInvalid[i] * values[i]);
  }
  return total;
}

async function determineFailure(line: string) {
  let previousOpeners: string[] = []
  for (let i = 0; i < line.length; i++) {
    if (openers.indexOf(line[i]) !== -1) {
      previousOpeners.push(line[i]);
    } else {
      const previousOpener = previousOpeners.pop();
      const closerIndex = closers.indexOf(line[i]);
      const opener = openers[closerIndex];

      if (opener !== previousOpener) {
        totalInvalid[closerIndex]++;
        return;
      }
    }
  }
  return;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let lines: string[] = []

  for await (const line of rl) {
    lines.push(line);
  }

  for await (const line of lines) {
    await determineFailure(line);
  }

  return await calculateInvalidScore();
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
