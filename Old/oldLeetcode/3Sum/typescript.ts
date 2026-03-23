import { DoQuestion } from "./typescriptFunctions";

function main() {
  console.log("Running code");
  const result = DoQuestion([-1,0,1,2,-1,-4]);
  console.log(`Got ${result}`)
}

main();
