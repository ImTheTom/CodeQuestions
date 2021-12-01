import fs from 'fs';
import readline from 'readline';

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

  for await (const line of rl) {
    const lineValue = Number(line);

    if (!previous) {
      previous = lineValue;
    } else {
      if (lineValue > previous) {
        increased++;
      }
      previous = lineValue;
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
