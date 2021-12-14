import fs from 'fs';
import readline from 'readline';

async function getMax (str: string): Promise<number> {
 var max = 0;
  str.split('').forEach(function(char: string){
    if(str.split(char).length > max) {
        max = str.split(char).length;
     }
  });
  return max;
};

async function getMin (str: string): Promise<number> {
 var min = 1000;
  str.split('').forEach(function(char: string){
    if(str.split(char).length < min) {
        min = str.split(char).length;
     }
  });
  return min;
};

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let template = '';
  let inputs: string[] = [];
  let output: string[] = [];

  for await (const line of rl) {
    if (template == '') {
      template = line;
    } else if (line == '') {
      continue
    } else {
      let spl = line.split(' -> ');
      inputs.push(spl[0]);
      output.push(spl[1]);
    }
  }

  for (let i = 0; i < 10; i++) {
    for (let j = 0; j < template.length-1; j++) {
      let tmp = template.substring(j, j+2);
      for (let k = 0; k < inputs.length; k++) {
        if (inputs[k] == tmp) {
          template = [template.slice(0, j+1), output[k], template.slice(j+1)].join('');
          break;
        }
      }
      j++;
    }
  }

  return await getMax(template) - await getMin(template);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
