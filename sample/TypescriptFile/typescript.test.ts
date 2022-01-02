import { DoQuestion, ReadInFile } from "./typescriptFunctions";

describe('General Typescript tests', () => {
  it('Able to do first test', () => {
    const input = ReadInFile("input.txt");
    const testResult = DoQuestion(input);
    expect(testResult).toBe(2);
  })
});