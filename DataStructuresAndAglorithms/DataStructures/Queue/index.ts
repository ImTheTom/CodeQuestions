export class Queue {
	q: Array<any>;

	constructor() {
		this.q = new Array();
	}

	enqueue(element: any) {
		this.q.push(element);
	}

	dequeue(): any {
		return this.q.shift();
	}

	print() {
		for (const v of this.q) {
			console.log(v);
		}
	}
}

let q = new Queue();

console.log("Add and print out queue")

q.enqueue(5);
q.enqueue(6);
q.enqueue(7);
q.enqueue(8);

q.print();

console.log("Dequeue")
console.log(q.dequeue());

console.log("print out current queue")
q.print();
