import fs from 'fs';
import readline from 'readline';

async function getTotal(count: number[]): Promise<number> {
  let max = 0;
  let min = 100000000000000000000000000000;
  for await (const c of count) {
    if (c < min) {
      min = c;
    }
    if (c > max) {
      max = c
    }
  }

  console.log(max);
  console.log(min);
  return max - min;
}

async function finalPoly(pairs: string[], counts: number[]): Promise<string> {
  let tmp = '';
  for (let j = 0; j < pairs.length; j++) {
    for (let i = 0; i < counts[j]; i++) {
      tmp = `${tmp}${pairs[j].substring(0, 1)}`
    }
  }
  tmp = `${tmp}${pairs[pairs.length-1].substring(1)}`
  return tmp;
}

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

  let pairs: string[] = [];
  let counts: number[] = [];

  for (let i = 0; i < template.length - 1 ; i++) {
    let tmp = template.substring(i, i+2);
    let tmpIndx = pairs.indexOf(tmp);
    if (tmpIndx == -1) {
      pairs.push(tmp);
      counts.push(1);
    } else {
      counts[tmpIndx]++;
    }
  }

  for(let i = 0; i < 40; i++) {
    let nextPairs: string[] = [];
    let nextCounts: number[] = [];
    console.log(i);

    for (let j = 0; j < pairs.length; j++) {
      let pair = pairs[j];
      let count = counts[j];

      const ruleIndx = inputs.indexOf(pair);
      const newP = output[ruleIndx];

      let newPair1 = `${pair.substring(0, 1)}${newP}`;
      let newPair2 = `${newP}${pair.substring(1)}`;

      let newPair1Idx = nextPairs.indexOf(newPair1);
      if (newPair1Idx == -1) {
        nextPairs.push(newPair1);
        nextCounts.push(count);
      } else {
        nextCounts[newPair1Idx] += count;
      }

      let newPair2Idx = nextPairs.indexOf(newPair2);
      if (newPair2Idx == -1) {
        nextPairs.push(newPair2);
        nextCounts.push(count);
      } else {
        nextCounts[newPair2Idx] += count;
      }
    }

    pairs = nextPairs;
    counts = nextCounts
  }

  let v: string[] = [];
  let vc: number[] = [];

  for (let i = 0; i < pairs.length; i++) {
    let char = pairs[i].substring(0, 1);
    let cou = counts[i];
    let charIdx = v.indexOf(char);
    if (charIdx == -1) {
      v.push(char);
      vc.push(cou);
    } else {
      vc[charIdx] += cou;
    }
  }

  let tmp = pairs[pairs.length-1].substring(1);
  let tmpIdx = v.indexOf(tmp);
  if (tmpIdx == -1) {
    v.push(tmp);
    vc.push(1);
  } else {
    vc[tmpIdx] += 1;
  }

  console.log(v);

  // const final = await finalPoly(pairs, counts);

  return await getTotal(vc);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
