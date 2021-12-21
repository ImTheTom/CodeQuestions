let dieCount = 1;
let totalRolls = 0;
let p1TotalWins = 0;
let p2TotalWins = 0;
let games: Game[] = [];

function generateRolls(): number[] {
  let rolls: number[] = [];
  for (let i = 0; i < 3; i++) {
    totalRolls++;
    rolls.push(dieCount);
    dieCount++;

    if (dieCount > 100) {
      dieCount = 1;
    }
  }

  return rolls;
}

class Game {
  p1: Player;
  p2: Player;

  constructor(p1: Player, p2: Player) {
    this.p1 = p1;
    this.p2 = p2;
  }

  playTillWinner() {
    while(true) {
      if (this.p2.getScore() < 21) {
        let tmpRolls = generateRolls();
        for (let i = 0; i < tmpRolls.length; i++) {
          this.p1.doRoll(tmpRolls[i]);
          let tmG = new Game(this.p1, this.p2);
          games.push(tmG);
        }
      } else {
        p2TotalWins += 1;
        return;
      }
      if (this.p1.getScore() < 21) {
        let tmpRolls = generateRolls();
        for (let i = 0; i < tmpRolls.length; i++) {
          this.p2.doRoll(tmpRolls[i]);
          let tmG = new Game(this.p1, this.p2);
          games.push(tmG);
        }
        return;
      } else {
        p1TotalWins += 1;
        return;
      }
    }
  }
}

class Player {
  position: number;
  score: number;

  constructor(p: number) {
    this.position = p;
    this.score = 0;
  }

  getPos(): number {
    return this.position;
  }

  getScore(): number {
    return this.score;
  }

  doRoll(num: number) {
    let tmpNewPos = this.position + num;
    let newScore = tmpNewPos % 10;
    if (newScore === 0) {
      newScore = 10;
    }
    this.position = newScore;
    this.score += newScore;
  }
}

function processFunction(): number {
  let p1 = new Player(4);
  let p2 = new Player(8); // my input = 6
  let g = new Game(p1, p2);
  games.push(g);

  while (games.length > 0) {
    console.log(games.length);
    let tmpG = games.shift();
    if (tmpG === undefined) {
      return -1;
    }
    tmpG.playTillWinner();
  }

  console.log(p1TotalWins);
  console.log(p2TotalWins);

  return p1TotalWins;
}

function main() {
  console.log("Running code");

  const result = processFunction();
  console.log(result);
}

main();
