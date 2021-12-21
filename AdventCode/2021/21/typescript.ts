let dieCount = 1;
let totalRolls = 0;

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
  let p2 = new Player(6);
  while (true) {
    if (p2.getScore() < 1000) {
      let tmpRolls = generateRolls();
      p1.doRoll(tmpRolls.reduce((a,b) => a + b, 0));
    } else {
      return totalRolls * p1.getScore();
    }
    if (p1.getScore() < 1000) {
      let tmpRolls = generateRolls();
      p2.doRoll(tmpRolls.reduce((a,b) => a + b, 0));
    } else {
      return totalRolls * p2.getScore();
    }
  }
}

function main() {
  console.log("Running code");

  const result = processFunction();
  console.log(result);
}

main();
