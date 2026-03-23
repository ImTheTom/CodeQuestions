export function DoQuestion(strs: string[]): string[][] {
  const final: string[][] = [];

  strs.sort((a,b) => {
    return a.length > b.length ? 1 : -1;
  });

  let lastLength = 0;

  for (const val of strs) {
    if (lastLength !== val.length) {
      final.push([val]);
      lastLength = val.length;
      continue;
    }

    const sortedVal = val.split('').sort().join('');
    let added = false;
    lastLength = val.length;


    for (let i = 0; i < final.length; i++) {
      const finalArrValue = final[i][0];
      if (finalArrValue.length != sortedVal.length) {
        continue
      }

      const sortedFinalValue = finalArrValue.split('').sort().join('');
      if (sortedFinalValue == sortedVal) {
        final[i].push(val);
        added = true;
        break;
      }
    }

    if (!added) {
      final.push([val])
    }
  }
  return final;
}
