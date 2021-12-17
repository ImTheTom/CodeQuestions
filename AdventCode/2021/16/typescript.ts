import fs from 'fs';
import readline from 'readline';

let totalPackets: Packet[] = [];

interface Packet {
	version: number,
	type: number,
	value: number
}

async function addLiteral(info: string[], version: number, type: number): Promise<string> {
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
		value: parseInt(total, 2)
	} as Packet);

	return info.join('')
}

async function addOperatorTypeZero(info: string[], version: number, type: number): Promise<string> {
  const subPacketLength = parseInt(info.splice(0, 15).join(''), 2);
  let content = info.splice(0, subPacketLength).join('');

  while (content.length > 0) {
    content = await parseBinary(content);
  }

	totalPackets.push({
		version,
		type,
		value: 0
	} as Packet);

  return info.join('');
}

async function addOperatorTypeOne(info: string[], version: number, type: number): Promise<string> {
  const subPacketCount = parseInt(info.splice(0, 11).join(''), 2);
  let content = info.join('');

  for (let i = 0; i < subPacketCount; i++) {
    let remaining = await parseBinary(content);
    content = remaining;
  }

	totalPackets.push({
		version,
		type,
		value: 0
	} as Packet);

  return content;
}

async function parseBinary(input: string): Promise<string> {
	let info = input.split('');
	const version = parseInt(info.splice(0, 3).join(''), 2);
	const type = parseInt(info.splice(0, 3).join(''), 2);

	if (type == 4) {
		let remaining = await addLiteral(info, version, type);
		return remaining;
	} else {
		const lengthOperatorType = info.splice(0, 1)[0];
    if (lengthOperatorType == '0') {
      let remaining = await addOperatorTypeZero(info, version, type);
      return remaining;
    } else {
      let remaining = await addOperatorTypeOne(info, version, type);
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

async function calculate(): Promise<number> {
  let total = 0;
  for (let i = 0; i < totalPackets.length; i++) {
    total += totalPackets[i].version;
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
