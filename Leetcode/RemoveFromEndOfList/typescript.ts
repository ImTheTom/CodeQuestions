import { DoQuestion, ListNode } from "./typescriptFunctions";

function main() {
  console.log("Running code");
  const fifth = new ListNode(5);
  const fourth = new ListNode(4, fifth);
  const third = new ListNode(3, fourth);
  const second = new ListNode(2, null);
  const first = new ListNode(1, null);
  let result = DoQuestion(first, 1);
  while (result?.next !== undefined) {
    console.log(`At ${result.val}`);
    result = result.next;
  }
}

main();
