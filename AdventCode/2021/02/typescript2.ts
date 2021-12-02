import fs from 'fs';
import readline from 'readline';

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let horizontal = 0;
  let vertical = 0;
  let aim = 0;

  for await (const line of rl) {
    const instructionArray = line.split(" ");
    const instruction = instructionArray[0];
    const units = parseInt(instructionArray[1]);

    if (instruction === 'forward') {
      horizontal += units;
      vertical += aim * units;
    } else if (instruction === 'down') {
      aim += units;
    } else if (instruction === 'up') {
      aim -= units;
    }
  }

  return horizontal * vertical;
}

async function main() {
  console.log("Running code");

  const increased = await processLineByLine();
  console.log(increased);
}

main();
