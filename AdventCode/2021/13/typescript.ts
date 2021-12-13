import fs from 'fs';
import readline from 'readline';

class Paper {
  points: number[][];
  maxX: number;
  maxY: number;
  total: number;

  constructor(hashes: number[][], x: number, y:number) {
    this.total = 0;
    this.maxX = x+1;
    this.maxY = y+1;
    this.points = Array.from(Array(this.maxY), _ => Array(this.maxX).fill(0));
    for (let i = 0; i < hashes.length; i++) {
      const h = hashes[i];
      this.points[h[1]][h[0]] = 1;
    }
  }

  foldUp(row: number) {
    for (let i = 0; i < row; i++) {
      let from = row + i + 1;
      let to = row - i - 1;
      for (let j = 0; j < this.points[from].length; j++) {
        let value = this.points[from][j];
        if (value == 1) {
          this.points[to][j] = 1;
        }
      }
    }
    for (let i = 0; i <= row; i++) {
      this.points.pop();
    }
    this.maxY = row;
  }

  foldLeft(column: number) {
    for (let i = 0; i < column; i++) {
      let from = column + i + 1;
      let to = column - i - 1;
      for (let j = 0; j < this.maxY; j++) {
        let value = this.points[j][from];
        if (value == 1) {
          this.points[j][to] = 1;
        }
      }
    }

    for(let i = 0; i <= column; i++) {
      for(let j = 0; j < this.maxY; j++) {
        this.points[j].pop();
      }
    }
    this.maxX = column;
  }

  getTotal(): number {
    for (let i = 0; i < this.maxY; i++) {
      for (let j = 0; j < this.maxX; j++) {
        if (this.points[i][j] == 1) {
          this.total++;
        }
      }
    }
    return this.total;
  }
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let hashes: number[][] = [];
  let XFolds: number[] = [];
  let YFolds: number[] = [];

  let recievingHashes = true;

  let maxX = 0;
  let maxY = 0;

  for await (const line of rl) {
    if (line == '') {
      recievingHashes = false;
    }
    if (recievingHashes) {
      const hash = line.split(',').map(Number);
      if (hash[0] > maxX) {
        maxX = hash[0];
      }
      if (hash[1] > maxY) {
        maxY = hash[1];
      }
      hashes.push(hash);
    } else {
      const folds = line.split('=')
      if (folds[0] == 'fold along y') {
        YFolds.push(parseInt(folds[1]));
      } else if (folds[0] == 'fold along x') {
        XFolds.push(parseInt(folds[1]));
      }
    }
  }

  let paper = new Paper(hashes, maxX, maxY);
  paper.foldLeft(XFolds[0])

  return paper.getTotal();
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
