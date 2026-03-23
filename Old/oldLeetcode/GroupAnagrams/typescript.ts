import { DoQuestion } from "./typescriptFunctions";

function main() {
  console.log("Running code");
  const result = DoQuestion(["eat","tea","tan","ate","nat","bat"]);
  for (let i = 0; i < result.length; i++) {
    console.log(`Got ${result[i]}`)
  }
}

main();
