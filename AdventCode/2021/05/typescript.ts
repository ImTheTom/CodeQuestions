import fs from 'fs';
import readline from 'readline';

async function calculate(field: number[][]): Promise<number> {
  let total = 0;
  for (let i =0; i < field.length; i++) {
    for (let j =0; j < field[i].length; j++) {
      if(field[i][j] > 1) {
        total++;
      }
    }
  }
  return total;
}

async function runMovesOnField(moves: string[][], field: number[][]): Promise<number[][]> {
  for (let i =0; i < moves.length; i++) {
    let move = [parseInt(moves[i][0]), parseInt(moves[i][1]),parseInt(moves[i][2]),parseInt(moves[i][3])];
    if (move[0] === move[2]) {
      if (move[1] < move[3]) {
        while(move[1] <= move[3]) {
          field[move[0]][move[1]]++;
          move[1]++;
        }
      } else {
        while(move[3] <= move[1]) {
          field[move[0]][move[3]]++;
          move[3]++;
        }
      }
    }else if (move[1] === move[3]) {
      if (move[0] < move[2]) {
        while(move[0] <= move[2]) {
          field[move[0]][move[1]]++;
          move[0]++;
        }
      } else {
        while(move[2] <= move[0]) {
          field[move[2]][move[1]]++;
          move[2]++;
        }
      }
    }
  }
  return field;
}

async function zeros(dimensions: Array<number>): Promise<number[][]> {
  var array: number[][] = [];

  for (var i = 0; i < dimensions[0]; ++i) {
    array.push(Array(dimensions[1]).fill(0));
  }

  return array;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let moves: string[][] = [];

  for await (const line of rl) {
    const firstSplit = line.split(' ');
    const secondSplit = firstSplit[0].split(',');
    const thirdSplit = firstSplit[2].split(',');
    const full = [...secondSplit, ...thirdSplit];
    moves.push(full);
  }

  const field = await zeros([1000, 1000]);

  const finalField = await runMovesOnField(moves, field);

  const score = await calculate(finalField);

  return score;
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
