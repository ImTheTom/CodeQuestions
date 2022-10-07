export class Stack {
	q: Array<any>;

	constructor() {
		this.q = new Array();
	}

	enqueue(element: any) {
		this.q.push(element);
	}

	dequeue(): any {
		return this.q.pop();
	}

	exists(searchValue: any): boolean {
		for (const v of this.q) {
			if (v == searchValue) {
				return true;
			}
		}
		return false;
	}

	print() {
		for (const v of this.q) {
			console.log(v);
		}
	}
}

let stack = new Stack();

console.log("Add and print out queue")

stack.enqueue(5);
stack.enqueue(6);
stack.enqueue(7);
stack.enqueue(8);

stack.print();

console.log("Dequeue")
console.log(stack.dequeue());

console.log("print out current queue")
stack.print();
