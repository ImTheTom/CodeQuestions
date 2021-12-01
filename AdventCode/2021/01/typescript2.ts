import fs from 'fs';
import readline from 'readline';

function add(accumulator: number, a: number) {
  return accumulator + a;
}

async function sumArray(input: Array<number>) {
  const sum = input.reduce(add, 0);
  return sum;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });
  // Note: we use the crlfDelay option to recognize all instances of CR LF
  // ('\r\n') in input.txt as a single line break.

  let previous = null;
  let increased = 0;
  let current: number[] = [];

  for await (const line of rl) {
    const lineValue = Number(line);
    current.push(lineValue);
    if (current.length < 3) {
      continue;
    } else if (current.length === 3) {
      previous = await sumArray(current);
    } else {
      current.shift();
    }

    const newValue = await sumArray(current);

    if (!previous) {
      previous = lineValue;
    } else {
      if (newValue > previous) {
        increased++;
      }
      previous = newValue;
    }
  }

  return increased;
}


async function main() {
  console.log("Running code");

  const increased = await processLineByLine();
  console.log(increased);
}

main();
