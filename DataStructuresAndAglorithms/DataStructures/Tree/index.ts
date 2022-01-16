console.log('Tree');

interface Node {
	value: number;
	left?: Node|undefined;
	right?: Node|undefined;
}

class BinaryTree {
	root: Node|undefined;

	constructor(root: Node) {
		this.root = root;
	}

	insert(newNode: Node) {
		this.insertNode(this.root!, newNode);
	}

	insertNode(node: Node, newNode: Node) {
		if (newNode.value < node.value) {
			if (node.left === undefined) {
				node.left = newNode;
			} else {
				this.insertNode(node.left, newNode);
			}
		} else {
			if (node.right === undefined) {
				node.right = newNode;
			} else {
				this.insertNode(node.right, newNode);
			}
		}
	}

	remove(deleteNode: Node) {
		this.root = this.removeNode(deleteNode, this.root);
	}

	removeNode(deleteNode: Node, node?: Node): Node|undefined {
		if (node === undefined) {
			return undefined;
		}

		if (deleteNode.value < node.value) {
			node.left = this.removeNode(deleteNode, node.left);
			return node;
		} else if (deleteNode.value > node.value) {
			node.right = this.removeNode(deleteNode, node.right);
			return node;
		} else {
			if (node.left == undefined && node.right == undefined) {
				node = undefined;
				return node;
			}
			if (node.left == undefined) {
				node = node.right;
				return node;
			} else if (node.right == undefined) {
				node = node.left;
				return node;
			}
        	let minNode = this.findMinNode(node.right);
        	node.value = minNode.value;

        	node.right = this.removeNode(node.right, minNode);
        	return node;
		}
	}

	getRootNode() {
	    return this.root;
	}

	search(node: Node|undefined, data: number): Node|undefined {
    	if(node === undefined)
    	    return undefined;
    	else if(data < node.value)
    	    return this.search(node.left, data);
    	else if(data > node.value)
    	    return this.search(node.right, data);
    	else
    	    return node;
	}

	inorder(node: Node|undefined) {
	    if(node !== undefined) {
	        this.inorder(node.left);
	        console.log(node.value);
	        this.inorder(node.right);
	    }
	}

	findMinNode(node: Node): Node {
    	if(node.left === undefined)
    	    return node;
    	else
        	return this.findMinNode(node.left);
	}
}

var BST = new BinaryTree({
	value: 15
} as Node);

// Inserting nodes to the BinarySearchTree
BST.insert({
	value: 25
} as Node);
BST.insert({
	value: 10
} as Node);
BST.insert({
	value: 7
} as Node);
BST.insert({
	value: 22
} as Node);
BST.insert({
	value: 17
} as Node);
BST.insert({
	value: 13
} as Node);
BST.insert({
	value: 5
} as Node);
BST.insert({
	value: 9
} as Node);
BST.insert({
	value: 27
} as Node);

BST.inorder(BST.getRootNode());

BST.remove({
	value: 5
} as Node);

BST.remove({
	value: 7
} as Node);

BST.inorder(BST.getRootNode());