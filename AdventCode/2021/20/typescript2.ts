import fs from 'fs';
import readline from 'readline';
// const algorithm = '#.#.#.#.#......#.#.#.#.##..#.##.##..#..##...#.#.#.#...##.##.##.###....#..#...#.#..###.#...#..##.#.###..#..####.###...#.#.#..##..##.##..##..###..#....#.#....#####.#...###...#.#....###...#..##.##..#..#.##..###..#.##.###..#.####...#.##.....#.###...#.##.##.#.#######...#.###..##..##..#.#.#.#####...#....#.....##.#.#...##.######....#..#......#.#.#.#.##...######.#.#####..#####..#.#.#.#.###.#.#....#..##..#..#.#.#..##....##..#.#.......##...#..####.####.#.#..#.###..#...#......###...#...#.##.#.####..#.#....###.####..#.'
// const algorithm = '..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#'
const algorithm = '#..###.##....#.#.#...#.#.#...##...####......##.##..###...#.####..#..#..#####..#.##.....#..#.###.##...#.#.....#...##.##.##...#####.#.#.#.##.###.#.##..#.##.##.#..#...####.#.#.....#..#.....###.#..#.#.#.#...#.###..#.###..##.#..#...##...####.#.........###..#.##.#..#.#...##.#.#.##.####.###....#####..###...##..#####..###..##..#.#.#..###.##.###..#.#######.####..#....###.##...#.####..#.#######...###...##.##.###...##..#.....#.###....#..#.#..###.#...#######.#...##..#.#..##.#...##.#..##.##..#...#.#.##.####........#..#.'
const algorithmSplit = algorithm.split('');

async function createArrayAtPoint(i: number, j: number, map: string[][]): Promise<string[][]> {
  let imgCond: string[][] = [];
  let top:string[] = [];
  let mid:string[] = [];
  let bot:string[] = [];

  if (i != 0 && j != 0) {
    top.push(map[i-1][j-1])
  } else {
    top.push('.')
  }

  if (i != 0) {
    top.push(map[i-1][j])
  } else {
    top.push('.')
  }

  if (j != map[0].length-1 && i != 0) {
    top.push(map[i-1][j+1])
  } else {
    top.push('.')
  }

  if (j != 0) {
    mid.push(map[i][j-1])
  } else {
    mid.push('.')
  }
  mid.push(map[i][j])
  if (j != map[0].length-1) {
    mid.push(map[i][j+1])
  } else {
    mid.push('.')
  }

  if (i != map.length-1 && j != 0) {
    bot.push(map[i+1][j-1])
  } else {
    bot.push('.')
  }
  if (i != map.length-1) {
    bot.push(map[i+1][j])
  } else {
    bot.push('.')
  }
  if (i != map.length-1 && j != map[0].length-1) {
    bot.push(map[i+1][j+1])
  } else {
    bot.push('.')
  }

  imgCond.push(top);
  imgCond.push(mid);
  imgCond.push(bot);

  return imgCond;
}

async function turnIntoBinary(map: string[][]): Promise<number> {
  let bin = ""
  for (let i = 0; i < map.length; i++) {
    for (let j = 0; j < map[0].length; j++) {
      bin += map[i][j] === "#" ? "1" : "0"
    }
  }
  return parseInt(bin, 2);
}

async function findActualValue(numb: number): Promise<string> {
  return algorithmSplit[numb]
}

async function calculate(img: string[][], bounds: number): Promise<number> {
  let total = 0
  for (let i = 50; i < bounds-50; i++) {
    for (let j = 50; j < bounds-50; j++) {
      total += img[i][j] === "#" ? 1 : 0
    }
  }
  return total
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let img: string[][] = [];

  for await (let line of rl) {
    img.push(line.split(''));
  }

  const bounds = img.length + 4 * 50;

  let currentImg = Array.from({length: bounds}, () => Array(bounds).fill(0));

  for (let i = 0; i < img.length; i++) {
    for (let j = 0; j < img[i].length; j++) {
      if(img[i][j] == "#"){
        currentImg[i+(2*50)][j+(2*50)] = "#"
      }
    }
  }

  for (let i = 0; i < 50; i++) {
    const currentNewImage = Array.from(
      { length: bounds },
      () => Array(bounds).fill(0),
    );
    for (let y = 1; y < bounds - 1; y++) {
      for (let x = 1; x < bounds - 1; x++) {
        const arr = await createArrayAtPoint(y, x, currentImg);
        const num = await turnIntoBinary(arr);
        let ac = await findActualValue(num);

        currentNewImage[y][x] = ac;
      }
    }
    currentImg = currentNewImage;
  }
  console.log(currentImg);

  return await calculate(currentImg, bounds);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
