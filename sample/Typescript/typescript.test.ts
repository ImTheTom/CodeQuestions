import { DoQuestion } from "./typescriptFunctions";

describe('General Typescript tests', () => {
  it('Able to do first test', () => {
    const testResult = DoQuestion(1, 1);
    expect(testResult).toBe(2);
  })
});