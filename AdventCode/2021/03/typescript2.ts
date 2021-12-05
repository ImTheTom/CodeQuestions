import fs from 'fs';

function findNextValue(input: Array<string>, column: number, findingMax: boolean): string {
  let numberOfOnes = 0;
  let numberOfZeros = 0;

  for (let i = 0; i < input.length; i++) {
    if (input[i][column] === '1') {
      numberOfOnes++;
    } else {
      numberOfZeros++;
    }
  }

  if (findingMax) {
    if (numberOfOnes > numberOfZeros) {
      return '1';
    } else if (numberOfZeros > numberOfOnes) {
      return '0';
    } else {
      return '1';
    }
  }

  if (numberOfOnes > numberOfZeros) {
    return '0';
  } else if (numberOfZeros > numberOfOnes) {
    return '1';
  } else {
    return '0';
  }
}

function processInput() {
  const fileStream = fs.readFileSync('input.txt').toString('utf-8');

  let oxygenInputs = fileStream.split("\n");
  let co2Inputs = fileStream.split("\n");

  const byteLength = oxygenInputs[0].length;

  for (let i = 0; i < byteLength && oxygenInputs.length > 1; i++) {
    const mostCommonValue = findNextValue(oxygenInputs, i, true);
    oxygenInputs = oxygenInputs.filter(x => x[i] == mostCommonValue);
  }

  for (let i = 0; i < byteLength && co2Inputs.length > 1; i++) {
    const leastCommonValue = findNextValue(co2Inputs, i, false);
    co2Inputs = co2Inputs.filter(x => x[i] == leastCommonValue);
  }

  const oxygen = parseInt(oxygenInputs[0], 2)
  const c02 = parseInt(co2Inputs[0], 2)

  return oxygen*c02;
}

function main() {
  console.log("Running code");

  const result = processInput();
  console.log(result);
}

main();
