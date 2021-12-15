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

async function extendRightMap(map: number[][]): Promise<number[][]> {
  let extendedMap: number[][] = [];
  for (let j = 0; j < map.length; j++) {
    let tmpRow: number[] = [];
    for (let i = 0; i < 5; i++) {
      for (let k = 0; k < map[j].length; k++) {
        let tmpV = map[j][k];
        tmpV += i;
        if (tmpV > 9) {
          tmpV -= 9;
        }
        tmpRow.push(tmpV);
      }
    }
    extendedMap.push(tmpRow)
  }
  return extendedMap;
}

async function extendDownMap(map: number[][]): Promise<number[][]> {
  let extendedMap: number[][] = JSON.parse(JSON.stringify(map));
  for (let i = 1; i < 5; i++) {
    let tmpMap: number[][] = JSON.parse(JSON.stringify(map));
    for (let j = 0; j < map.length; j++) {
      for (let k = 0; k < map[j].length; k++) {
        let tmpV = map[j][k];
        tmpV += i;
        if (tmpV > 9) {
          tmpV -= 9;
        }
        tmpMap[j][k] = tmpV;
      }
    }
    extendedMap.push(...tmpMap);
  }
  return extendedMap;
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

  let extendRight = await extendRightMap(map);
  let extendDown = await extendDownMap(extendRight);

  let start = '0,0';
  const mapL = extendDown.length;
  let finish = [mapL-1, extendDown[mapL-1].length-1].toString();


  return await findShortestPath(start, finish, extendDown);
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
