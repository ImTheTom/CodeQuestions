const openers = ["(", "{", "["];
const closers = [")", "}", "]"];

interface info {
  opener: boolean,
  index: number,
}

export function DoQuestion(s: string): boolean {
  const openers = [];

  for (let i = 0; i < s.length; i++) {
    const current = createInfo(s[i]);
    if (current.opener) {
      openers.push(current.index)
    } else {
      const last = openers.pop();
      if (last !== current.index) {
        return false;
      }
    }
  }

  if (openers.length > 0) {
    return false;
  }

  return true;
}

function createInfo(s: string): info {
  for (let i = 0; i < openers.length; i++) {
    if (s == openers[i]) {
      return {
        opener: true,
        index: i,
      }
    } else if (s == closers[i]) {
      return {
        opener: false,
        index: i
      }
    }
  }
  return {
    opener: false,
    index: -1
  }
}
