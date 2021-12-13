import fs from 'fs';
import readline from 'readline';

// Credit tymscar for inspiration

class Graph {
  nodes: Node[];
  paths: Node[][];
  total: number;

  constructor() {
    this.nodes = [];
    this.paths = [];
    this.total = 0;
  }

  CreateNode(nodeName: string): Node {
    let n = new Node(nodeName, nodeName == nodeName.toLowerCase());
    this.AddNode(n);
    return n;
  }

  GetNodeOrCreateNode(nodeName: string): Node {
    let find = this.GetNode(nodeName);
    if (find !== null) {
      return find;
    }
    return this.CreateNode(nodeName);
  }

  GetNode(nodeName: string): Node|null {
    for (const n of this.nodes) {
      if (n.getName() === nodeName) {
        return n;
      }
    }
    return null;
  }

  AddNode(node: Node) {
    this.nodes.push(node);
  }

  AddNewInput(from: string, to: string) {
    let fromNode = this.GetNodeOrCreateNode(from);
    let toNode = this.GetNodeOrCreateNode(to);

    fromNode.AddEdge(toNode);
    toNode.AddEdge(fromNode);
  }

  FindPaths(start: string, goal: string) {
    const queue = [[start]];
    while (queue.length > 0) {
      let curr = queue.pop();
      if (!curr) {
        return;
      }

      let currLastPlace = curr[curr.length-1];
      if (!currLastPlace) {
        return;
      }

      if (currLastPlace == goal) {
        this.total++;
        continue;
      }

      const actualNode = this.GetNode(currLastPlace);
      if (!actualNode) {
        return;
      }

      for (let i = 0; i < actualNode.getEdges().length; i++) {
        let tmp = actualNode.getEdges()[i];
        if (tmp.getName() == 'start') {
          continue;
        }
        if (tmp.isSmall() && curr.includes(tmp.getName())) {
          continue;
        }
        queue.push([...curr, tmp.getName()])
      }
    }
  }

  GetTotal(): number {
    return this.total;
  }
}

class Node {
  name: string;
  small: boolean;
  edges: Node[];

  constructor(name: string, isSmall: boolean) {
    this.name = name;
    this.small = isSmall;
    this.edges = [];
  }

  AddEdge(node: Node) {
    this.edges.push(node);
  }

  isSmall(): Boolean {
    return this.small;
  }

  getName(): string {
    return this.name;
  }

  getEdges(): Node[] {
    return this.edges;
  }
}

async function processLineByLine() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });

  let graph = new Graph();

  for await (const line of rl) {
    let inputs = line.split('-');
    graph.AddNewInput(inputs[0], inputs[1]);
  }

  graph.FindPaths('start', 'end');

  return graph.GetTotal();
}

async function main() {
  console.log("Running code");

  const result = await processLineByLine();
  console.log(result);
}

main();
