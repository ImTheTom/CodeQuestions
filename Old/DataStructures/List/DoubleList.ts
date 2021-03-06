interface Node {
	value: number;
	previous?: Node|undefined;
	next?: Node|undefined;
}

export class DoublelyLinkedList {
	Head: Node;
	End: Node;

	constructor(Head: Node) {
		this.Head = Head;
		this.End = Head;
	}

	addNode(newNode: Node) {
		let previousEnd = this.End;
		this.End.next = newNode;
		this.End = newNode;
		this.End.previous = previousEnd;
	}

	addNodeAfter(newNode: Node, value: number) {
		let tempHead: Node|undefined = this.Head;

		while(tempHead != undefined) {
			if (tempHead.value == value) {
				newNode.next = tempHead.next;
				newNode.previous = tempHead;
				tempHead.next = newNode;
				return true;
			}
			tempHead = tempHead.next;
		}

		return false;
	}

	newHead(newNode: Node) {
		newNode.next = this.Head;
		this.Head.previous = newNode;
		this.Head = newNode;
	}

	exists(value: number): boolean {
		let tempHead: Node|undefined = this.Head;
		while(tempHead != undefined) {
			if (tempHead.value == value) {
				return true;
			}
			tempHead = tempHead.next;
		}
		return false;
	}

	delete(value: number): boolean {
		let tempHead: Node|undefined = this.Head;
		if (tempHead.value == value && tempHead.next) {
			this.Head = tempHead.next;
			this.Head.previous = undefined;
			return true;
		}

		while(tempHead != undefined) {
			if (tempHead.value == value) {
				tempHead.previous!.next = tempHead.next;
				if (tempHead.value == this.End.value) {
					this.End = tempHead.previous!;
				}
				return true;
			}
			tempHead = tempHead.next;
		}

		return false;
	}

	print() {
		let tempHead: Node|undefined = this.Head;
		let index = 0;
		while (tempHead != undefined) {
			console.log(`INDEX - ${index} has a value of ${tempHead.value}`)
			index++;
			tempHead = tempHead.next;
		}
	}

	printBackwards() {
		let tempHead: Node|undefined = this.End;
		let index = 0;
		while (tempHead != undefined) {
			console.log(`INDEX - ${index} has a value of ${tempHead.value}`)
			index++;
			tempHead = tempHead.previous;
		}
	}
}

console.log("Single list");

let myList = new DoublelyLinkedList({
	value: 5,
});

console.log("Initial List is just");

myList.print();

console.log("Adding values");

myList.addNode({
	value: 10,
});

myList.addNode({
	value: 15,
});

myList.print();

console.log("Searching");

console.log(`Finding 10? - ${myList.exists(10)}`);

console.log(`Finding 11? - ${myList.exists(11)}`);

console.log("Adding more values for the deletion");

myList.addNode({
	value: 20,
});

myList.addNode({
	value: 25,
});

console.log("Delete the first value");

myList.delete(5);

myList.print();

console.log("Delete the last value");

myList.delete(25);

myList.print();

console.log("Delete the middle value");

myList.delete(15);

myList.print();

console.log("Delete the non existant value");

console.log(`Result of deletion of non existant ${myList.delete(16)}`);

console.log("Add another value after 10");

myList.addNodeAfter({
	value: 15
}, 10);

myList.print();

console.log("Add a new start value");

myList.newHead({
	value: 5
});

myList.print();

console.log("Transverse backwards");

myList.printBackwards();