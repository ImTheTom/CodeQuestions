class Fire {
  xVel: number;
  yVel: number;
  xPos: number[];
  yPos: number[];
  targetX: number[];
  targetY: number[];
  maxYPos: number;

  constructor(xVel: number, yVel: number, targetX: number[], targetY: number[]) {
    this.xVel = xVel;
    this.yVel = yVel;
    this.xPos = [0];
    this.yPos = [0];
    this.targetX = targetX;
    this.targetY = targetY;
    this.maxYPos = 0;
  }

  getMaxY(): number {
    return this.maxYPos;
  }

  calculateMaxY() {
    let max = 0;
    for (let i = 0; i < this.yPos.length; i++) {
      if (this.yPos[i] > max) {
        max = this.yPos[i];
      }
    }
    this.maxYPos = max;
  }

  doStepUntilGoalOrPast(): boolean {
    while (true) {
      if (this.isPastFinal()) {
        return false;
      }
      if (this.isInFinal()) {
        return true;
      }
      this.doStep();
    }
  }

  getCurrentXPos(): number {
    let x = this.xPos[this.xPos.length - 1];
    if (x === undefined) {
      return -1;
    }
    return x;
  }

  getCurrentYPos(): number {
    let y = this.yPos[this.yPos.length - 1];
    if (y === undefined) {
      return -1;
    }
    return y;
  }

  isPastFinal(): boolean {
    let currX = this.getCurrentXPos();
    let currY = this.getCurrentYPos();

    let maxX = Math.max(...this.targetX);
    let minY = Math.min(...this.targetY);

    let pastX = currX > maxX;
    let pastY = currY < minY;

    return pastX || pastY;
  }


  isInFinal(): boolean {
    let currX = this.getCurrentXPos();
    let currY = this.getCurrentYPos();
    let inX = false;
    let inY = false;

    for (let i = 0; i < this.targetX.length; i++) {
      if (this.targetX[i] == currX) {
        inX = true;
        break;
      }
    }

    for (let i = 0; i < this.targetY.length; i++) {
      if (this.targetY[i] == currY) {
        inY = true;
        break;
      }
    }

    return inX && inY;
  }

  doStep() {
    let newX = this.getCurrentXPos();

    newX += this.xVel;
    this.xPos.push(newX);

    if (this.xVel != 0) {
      if (this.xVel > 0) {
        this.xVel -= 1;
      } else {
        this.xVel += 1;
      }
    }

    let newY = this.getCurrentYPos();
    newY += this.yVel;
    this.yPos.push(newY);

    this.yVel -= 1;
  }
}

function createArray(start: number, end: number, step: number): number[] {
  let arr: number[] = [];
  for (let i = start; i != end; i += step) {
    arr.push(i);
  }
  arr.push(end);
  return arr;
}

function processStuff(): number {
  const targetX = createArray(241, 275, 1);
  const targetY = createArray(-75, -49, 1);

  let potentialX = createArray(-275, 275, 1);
  let potentialY = createArray(-275, 275, 1);

  let total = 0;

  for (let i = 0; i < potentialX.length; i++) {
    for (let j = 0; j < potentialY.length; j++) {
      let f = new Fire(potentialX[i], potentialY[j], targetX, targetY);
      if (f.doStepUntilGoalOrPast()) {
        total += 1;
      }
    }
  }

  return total;
}

function main() {
  console.log("Running code");

  const result = processStuff();
  console.log(result);
}

main();
