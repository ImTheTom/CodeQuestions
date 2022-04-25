import fs from 'fs';

export function DoQuestion(input: string[]): number {
  let total = 0;
  for (let i = 0; i < input.length; i++) {
    total += parseInt(input[i]);
  }
  return total;
}

export function ReadInFile(inputFile: string): string[] {
  return fs.readFileSync(inputFile).toString().split("\n");
}