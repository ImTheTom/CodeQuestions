import fs from 'fs';
import readline from 'readline';

async function unmarkedValue(input: string[], gameBoard: string[]) {
  let total = 0;
  for (let i = 0; i < gameBoard.length; i++) {
    if (input.indexOf(gameBoard[i]) == -1) {
      total += parseInt(gameBoard[i])
    }
  }
  console.log(total);
  return total;
}

async function calculateWinnerValue(input: string[], gameBoard: string[], winningValue: number) {
  const umarkedV = await unmarkedValue(input, gameBoard);
  return umarkedV * winningValue;
}

async function sumArr(input: string[]) {
  return parseInt(input[0]) + parseInt(input[1]) + parseInt(input[2]) + parseInt(input[3]) + parseInt(input[4]);
}

async function checker(arr: string[], target: string[]) {
  return target.every(v => arr.includes(v));
}

async function checkWinnersHorizontally(input: string[], gameBoard: string[], row: number) {
  let gameRow = [gameBoard[(row*5)+0], gameBoard[(row*5)+1], gameBoard[(row*5)+2], gameBoard[(row*5)+3], gameBoard[(row*5)+4]];
  if (await checker(input, gameRow)) {
    return await sumArr(gameRow)
  }
  return 0;
}

async function checkWinnersVertically(input: string[], gameBoard: string[], column: number) {
  let gameColumn = [gameBoard[column+(5*0)], gameBoard[column+(5*1)], gameBoard[column+(5*2)], gameBoard[column+(5*3)], gameBoard[column+(5*4)]];
  if (await checker(input, gameColumn)) {
    return await sumArr(gameColumn)
  }
  return 0;
}

async function checkForWinners(input: string[], gameBoards: string[][]) {
  for (let i = 0; i < gameBoards.length; i++) {
    for (let j = 0; j < 5; j++) {
      let result = await checkWinnersHorizontally(input, gameBoards[i], j);
      if (result != 0) {
        if (gameBoards.length > 1) {
          gameBoards.splice(i, 1);
          break;
        } else {
          return await calculateWinnerValue(input, gameBoards[i], parseInt(input[input.length - 1]));
        }
      }
      result = await checkWinnersVertically(input, gameBoards[i], j);
      if (result != 0) {
        if (gameBoards.length > 1) {
          gameBoards.splice(i, 1);
          break;
        } else {
          return await calculateWinnerValue(input, gameBoards[i], parseInt(input[input.length - 1]));
        }
      }
    }
  }
  return 0;
}

async function actuallyProcess(input: string[], gameBoards: string[][]) {
  let availableInput: string[] = [];
  for (let i = 0; i < input.length; i++) {
    availableInput.push(input[i]);
    const winnerValue =  await checkForWinners(availableInput, gameBoards)
    if (winnerValue !== 0) {
      return winnerValue;
    }
  }

  return 0;
}

async function process() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let inputString = '';
  let gameBoard = '';
  let gameBoards: Array<string[]> = [];

  for await (const line of rl) {
    if (inputString == '' && line != '') {
      inputString = line;
      continue;
    }

    if (line == '' && gameBoard != '') {
      const tmpBoard = gameBoard.split(' ');
      gameBoards.push(tmpBoard.filter(x => x != ''));
      gameBoard = '';
    } else {
      if (gameBoard == '') {
        gameBoard = line;
      } else {
        gameBoard = gameBoard + ' ' + line;
      }
    }
  }

  const bingoInput = inputString.split(',')

  const result = await actuallyProcess(bingoInput, gameBoards);

  return result;
}

async function main() {
  console.log("Running code");

  const result = await process();
  console.log(result);
}

main();
