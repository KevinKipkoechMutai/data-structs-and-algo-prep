function distinctStockPairs(stocksProfit, target) {
    const seen = new Set()
    const distinctPairs = new Set()

    for (let profit of stocksProfit) {
        const complement = target - profit

        if (seen.has(complement)) {
            const pair = [Math.min(profit, complement), Math.max(profit, complement)]
            distinctPairs.add(pair.toString())
        }
        seen.add(profit)
    }
    return distinctPairs.size
}

console.log(distinctStockPairs([5,7,9,13,11,6,6,3,3], 12))