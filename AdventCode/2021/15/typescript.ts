import fs from 'fs';
import readline from 'readline';

const inf = 1000000;

async function removeFromOpenSet(openSet: string[], current: string): Promise<string[]> {
  for (let i = 0; i < openSet.length; i++) {
    if (openSet[i] == current) {
      openSet.splice(i, 1);
    }
  }
  return openSet
}

async function addToOpenSet(openSet: string[], current: string): Promise<string[]> {
  for (let i = 0; i < openSet.length; i++) {
    const value = openSet[i];
    if (value == current) {
      return openSet
    }
  }
  openSet.push(current);
  return openSet
}

async function findNeighbours(currentString: string, map: number[][]): Promise<string[]> {
  let current = await stringToArrayOfNumbers(currentString);
  let neighbours: string[] = [];
  if (current[0] != 0) {
    neighbours.push(await arrayOfNumbersToString([current[0] - 1, current[1]]))
  }

  if (current[1] != 0) {
    neighbours.push(await arrayOfNumbersToString([current[0], current[1] - 1]))
  }

  if (current[0] != map.length-1) {
    neighbours.push(await arrayOfNumbersToString([current[0] + 1, current[1]]))
  }

  if (current[1] != map[2].length) {
    neighbours.push(await arrayOfNumbersToString([current[0], current[1] + 1]))
  }

  return neighbours
}

async function stringToArrayOfNumbers(info: string): Promise<number[]> {
	return info.split(',').map(Number);
}

async function arrayOfNumbersToString(info: number[]): Promise<string> {
	return info.join(',');
}

async function fullPath(cameFrom: Map<string, string>, current: string, start: string): Promise<string[]> {
  let path: string[] = [];
  let last: string = current;
  while (last != start) {
    path.unshift(current);
    current = cameFrom.get(current) ?? '';
    last = current;
  }
  return path;
}

async function calculateCost(path: string[], map: number[][]): Promise<number> {
  let total = 0;
  for (let i = 0; i < path.length; i++) {
    let tmpV = await stringToArrayOfNumbers(path[i]);
    total += map[tmpV[0]] [tmpV[1]]
  }
  return total;
}

async function findShortestPath(start: string, goal: string, map: number[][]): Promise<number> {
  let openSet = [start];
  let cameFrom = new Map();

  // route from goal to n Default infinity
  let gScore = new Map();
  gScore.set(start, 0);

  while (openSet.length > 0) {
    let current = openSet.shift();
    if (!current) {
      return -2;
    }

    if (current == goal) {
      const path = await fullPath(cameFrom, current, start);
      return await calculateCost(path, map);
    }

    openSet = await removeFromOpenSet(openSet, current);

    const neighbours = await findNeighbours(current, map);

    for await (const n of neighbours) {
      const nArray = await stringToArrayOfNumbers(n);
      const currGScore = gScore.get(current) ?? inf;
      const tmpgScore = currGScore + map[nArray[0]][nArray[1]];

      const nGScore = gScore.get(n) ?? inf;

      if (tmpgScore < nGScore) {
        cameFrom.set(n, current);
        gScore.set(n, tmpgScore);

        openSet = await addToOpenSet(openSet, n);
      }
    }
  }
  return -1;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let map: number[][] = [];

  for await (const line of rl) {
    map.push(line.split('').map(Number));
  }

  let start = '0,0';
  const mapL = map.length;
  let finish = [mapL-1, map[mapL-1].length-1].toString();

  return await findShortestPath(start, finish, map);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
