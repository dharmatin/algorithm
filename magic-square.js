const magicSquares = [
    [[8, 1, 6], [3, 5, 7], [4, 9, 2]],
    [[6, 1, 8], [7, 5, 3], [2, 9, 4]],
    [[4, 9, 2], [3, 5, 7], [8, 1, 6]],
    [[2, 9, 4], [7, 5, 3], [6, 1, 8]],
    [[8, 3, 4], [1, 5, 9], [6, 7, 2]],
    [[4, 3, 8], [9, 5, 1], [2, 7, 6]],
    [[6, 7, 2], [1, 5, 9], [8, 3, 4]],
    [[2, 7, 6], [9, 5, 1], [4, 3, 8]],
];

const notSoMagicSquare = [[4, 8, 2], [4, 5, 7], [6, 1, 6]];

const evaluate = (s) => {
    const totals = [];
    for (let magicSquare of magicSquares) {
        let total = 0;
        for (let i = 0; i < magicSquare.length; i++) {
            const magicRow = magicSquare[i];
            const sRow = s[i];
            for (let j = 0; j < magicRow.length; j++) {
                const magicValue = magicRow[j];
                const sValue = sRow[j];
                if (magicValue !== sValue) {
                    total += Math.max(magicValue, sValue) -
                        Math.min(magicValue, sValue)
                        console.log("TOTAL", total)
                }
            }
        }
        totals.push(total);
    }
    return Math.min(...totals);
}

console.log(evaluate(notSoMagicSquare));
