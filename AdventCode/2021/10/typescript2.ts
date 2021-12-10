import fs from 'fs';
import readline from 'readline';

let openers = ['(', '[', '{', '<'];
let closers = [')', ']', '}', '>'];
let values = [1, 2, 3, 4];

async function calculateInvalidScore(requiredClosures: string[]): Promise<number> {
  let total = 0;
  for(let i = 0; i < requiredClosures.length; i++) {
    total *= 5;
    const valueIndex = closers.indexOf(requiredClosures[i]);
    total += values[valueIndex];
  }
  return total;
}

async function findMissingClosers(line: string): Promise<number> {
  let previousOpeners: string[] = []
  for (let i = 0; i < line.length; i++) {
    if (openers.indexOf(line[i]) !== -1) {
      previousOpeners.push(line[i]);
    } else {
      const previousOpener = previousOpeners.pop();
      const closerIndex = closers.indexOf(line[i]);
      const opener = openers[closerIndex];

      if (opener !== previousOpener) {
        return 0;
      }
    }
  }
  let requiredClosures: string[] = [];
  while (previousOpeners.length > 0) {
    let missing = previousOpeners.pop();
    if (!missing) {
      break;
    }
    const openerIndex = openers.indexOf(missing);
    const closer = closers[openerIndex];
    requiredClosures.push(closer);
  }
  return await calculateInvalidScore(requiredClosures);
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

  let values: number[] = []

  for await (const line of lines) {
    const v = await findMissingClosers(line);
    if ( v !== 0) {
      values.push(v);
    }
  }

  values.sort(function(a, b) {
    return a - b;
  });

  return values[Math.floor(values.length/2)];
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
