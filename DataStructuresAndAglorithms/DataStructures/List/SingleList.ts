interface Node {
	value: number;
	next?: Node|undefined;
}

export class SinglelyLinkedList {
	Head: Node;
	End: Node;

	constructor(Head: Node) {
		this.Head = Head;
		this.End = Head;
	}

	addNode(newNode: Node) {
		this.End.next = newNode;
		this.End = newNode;
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
		let previous: Node|undefined = undefined;
		if (tempHead.value == value && tempHead.next) {
			this.Head = tempHead.next;
			return true;
		}

		while(tempHead != undefined) {
			if (tempHead.value == value) {
				if (previous && tempHead.next) {
					previous.next = tempHead.next;
				} else if (previous && previous?.next) {
					previous.next = undefined;
				}
				return true;
			}
			previous = tempHead;
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
}

console.log("Single list");

let myList = new SinglelyLinkedList({
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