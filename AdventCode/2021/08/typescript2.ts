import fs from 'fs';
import readline from 'readline';

const uniqueValues = [2,4,3,7];

async function test(string: string, substring: string) {
    var letters = [...string];
    return [...substring].every(x => {
        var index = letters.indexOf(x);
        if (~index) {
            letters.splice(index, 1);
            return true;
        }
    });
}


async function processInputs(inputs: string[][], outputs: string[][]): Promise<number> {
  let total = 0;
  for (let i = 0; i < inputs.length; i++) {
    // 1, 4, 7, 8
    let currentInput = inputs[i];
    let currentOutput = outputs[i];
    let values: string[] = new Array(10).fill('');
    let fives: string[] = [];
    let sixes: string[] = [];

    let currentInputTotal = 0;

    for (let j = 0; j < currentInput.length; j++) {
      if (currentInput[j].length == 2) {
        values[1] = currentInput[j]
      } else if (currentInput[j].length == 4) {
        values[4] = currentInput[j]
      } else if (currentInput[j].length == 3) {
        values[7] = currentInput[j]
      } else if (currentInput[j].length == 7) {
        values[8] = currentInput[j]
      } else if (currentInput[j].length == 5) {
        fives.push(currentInput[j])
      } else if (currentInput[j].length == 6) {
        sixes.push(currentInput[j])
      }
    }

    // 6
    for await (const six of sixes) {
      if (!await test(six, values[1])) {
        values[6] = six
        sixes = sixes.filter(function(e) { return e !== six })
      }
    }

    // 3
    for await (const five of fives) {
      if (await test(five, values[1])) {
        values[3] = five
        fives = fives.filter(function(e) { return e !== five })
      }
    }

    // 9
    for await (const six of sixes) {
      if (await test(six, values[3])) {
        values[9] = six
        sixes = sixes.filter(function(e) { return e !== six })
      }
    }

    // 0
    values[0] = sixes.pop() ?? '';

    // 5
    for await (const five of fives) {
      if (await test(values[6], five)) {
        values[5] = five
        fives = fives.filter(function(e) { return e !== five })
      }
    }

    // 2
    values[2] = fives.pop() ?? '';

    console.log(values);

    for (let m = 0; m < currentOutput.length; m++) {
      currentInputTotal = currentInputTotal * 10;
      currentInputTotal += values.indexOf(currentOutput[m]);
    }

    total += currentInputTotal;
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

  for await (const line of lines) {
    var splitValue = line.split('|');
    var input = splitValue[0].split(' ').filter(Boolean).sort();
    for (let i = 0; i < input.length; i++) {
      input[i] = input[i].split('').sort().join('')
    }
    inputs.push(input);

    var output = splitValue[1].split(' ').filter(Boolean);
    for (let i = 0; i < output.length; i++) {
      output[i] = output[i].split('').sort().join('')
    }
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
