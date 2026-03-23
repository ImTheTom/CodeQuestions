export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

export function DoQuestion(head: ListNode | null, n: number): ListNode | null {
  let length = 0;
  let tmpHead = head;
  while (tmpHead?.next !== undefined) {
    length++;
    tmpHead = tmpHead.next;
  }

  tmpHead = head;
  let previousNode = head;

  const position = length - n;

  for (let i = 0; i < position; i++) {
    previousNode = tmpHead;
    tmpHead = tmpHead!.next;
  }

  if (position == 0) {
    return tmpHead!.next
  }

  previousNode!.next = tmpHead!.next

  return head;
}
