import { DoQuestion, ReadInFile } from "./typescriptFunctions";

function main() {
  console.log("Running code");
  const inputArray = ReadInFile("input.txt");
  const result = DoQuestion(inputArray);
  console.log(`Got ${result}`)
}

main();
