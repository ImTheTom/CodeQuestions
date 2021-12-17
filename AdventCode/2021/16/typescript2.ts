import fs from 'fs';
import readline from 'readline';

let totalPackets: Packet[] = [];

let currID = 0;

interface Packet {
	version: number,
	type: number,
	value: number,
	parent?: number,
	children?: number,
}

async function getID(): Promise<number> {
	currID += 1;
	return currID;
}

async function addLiteral(info: string[], version: number, type: number, parent?: number): Promise<string> {
	let total = '';

	while (true) {
		let sub = info.splice(0, 5).join('').padEnd(5, '0');
		total += sub.substring(1);
		if (sub[0] == '0') {
			break;
		}
	}

	totalPackets.push({
		version,
		type,
		value: parseInt(total, 2),
		parent,
		children: undefined,
	} as Packet);

	return info.join('')
}

async function addOperatorTypeZero(info: string[], version: number, type: number, parent?: number): Promise<string> {
  const subPacketLength = parseInt(info.splice(0, 15).join(''), 2);
  let content = info.splice(0, subPacketLength).join('');

  let id = await getID();

  while (content.length > 0) {
    content = await parseBinary(content, id);
  }

	totalPackets.push({
		version,
		type,
		value: 0,
		parent: parent,
		children: id,
	} as Packet);

  return info.join('');
}

async function addOperatorTypeOne(info: string[], version: number, type: number, parent?: number): Promise<string> {
  const subPacketCount = parseInt(info.splice(0, 11).join(''), 2);
  let content = info.join('');

  let id = await getID();

  for (let i = 0; i < subPacketCount; i++) {
    let remaining = await parseBinary(content, id);
    content = remaining;
  }

	totalPackets.push({
		version,
		type,
		value: 0,
		parent: parent,
		children: id,
	} as Packet);

  return content;
}

async function parseBinary(input: string, related?: number): Promise<string> {
	let info = input.split('');
	const version = parseInt(info.splice(0, 3).join(''), 2);
	const type = parseInt(info.splice(0, 3).join(''), 2);

	if (type == 4) {
		let remaining = await addLiteral(info, version, type, related);
		return remaining;
	} else {
		const lengthOperatorType = info.splice(0, 1)[0];
    if (lengthOperatorType == '0') {
      let remaining = await addOperatorTypeZero(info, version, type, related);
      return remaining;
    } else {
      let remaining = await addOperatorTypeOne(info, version, type, related);
      return remaining;
    }
	}
}

function toBinary(input: string): string {
	return input
		.split('')
		.map(char => parseInt(char, 16).toString(2).padStart(4, '0'))
		.join('')
		.replace(/0+$/, '');
}

async function findChildrenAndRemove(p: Packet): Promise<Packet[]> {
  let childrenOfP: Packet[] = [];
  for await(const newP of totalPackets) {
    if (newP.parent == p.children) {
      childrenOfP.push(newP);
      totalPackets = totalPackets.filter(news => newP !== news);
    }
  }
  return childrenOfP;
}

async function determineValue(p: Packet): Promise<number> {
  let total = 0;
  switch(p.type) {
    case 0:
      let childrenCaseZero = await findChildrenAndRemove(p);
      for (let i = 0; i < childrenCaseZero.length; i++) {
        total += await determineValue(childrenCaseZero[i])
      }
      return total;
    case 1:
      let childrenCaseOne = await findChildrenAndRemove(p);
      total = 1;
      for (let i = 0; i < childrenCaseOne.length; i++) {
        total *= await determineValue(childrenCaseOne[i])
      }
      return total;
    case 2:
      let childrenCaseTwo = await findChildrenAndRemove(p);
      let minimum = 100000000000000;
      for (let i = 0; i < childrenCaseTwo.length; i++) {
        let tmp = await determineValue(childrenCaseTwo[i]);
        if (tmp < minimum) {
          minimum = tmp;
        }
      }
      return minimum;
    case 3:
      let childrenCaseThree = await findChildrenAndRemove(p);
      let max = -1;
      for (let i = 0; i < childrenCaseThree.length; i++) {
        let v = await determineValue(childrenCaseThree[i])
        if (v > max) {
          max = v;
        }
      }
      return max;
    case 4:
      return p.value;
    case 5:
      let childrenCaseFive = await findChildrenAndRemove(p);
      let valuePOne = await determineValue(childrenCaseFive[0]);
      let valuePTwo = await determineValue(childrenCaseFive[1]);
      return valuePOne > valuePTwo ? 1 : 0;
    case 6:
      let childrenCaseSix = await findChildrenAndRemove(p);
      let valuePOneCaseSix = await determineValue(childrenCaseSix[0]);
      let valuePTwoCaseSix = await determineValue(childrenCaseSix[1]);
      return valuePOneCaseSix < valuePTwoCaseSix ? 1 : 0;
    case 7:
      let childrenCaseSeven = await findChildrenAndRemove(p);
      let valuePOneCaseSeven = await determineValue(childrenCaseSeven[0]);
      let valuePTwoCaseSeven = await determineValue(childrenCaseSeven[1]);
      return valuePOneCaseSeven == valuePTwoCaseSeven ? 1 : 0;
  }
  return 1;
}

async function calculate(): Promise<number> {
  let total = 0;
  while (totalPackets.length >= 0) {
	  let current = totalPackets.pop();
    if (!current) {
      return total;
    }
    total += await determineValue(current);
  }
  return total;
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let input: string = '';
  for await (const line of rl) {
    input = line;
  }
  let binaryInput = toBinary(input);

  await parseBinary(binaryInput);

  console.log(totalPackets);

  return await calculate();
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
