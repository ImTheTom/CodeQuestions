import { DoQuestion } from "./typescriptFunctions";

describe('General Typescript tests', () => {
  it('Able to do first test', () => {
    const testResult = DoQuestion([-1,0,1,2,-1,-4]);
    console.log(testResult);
    expect(testResult).toBe([[-1,-1,2],[-1,0,1]]);
  })
});