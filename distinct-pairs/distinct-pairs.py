def distinctStockPairs(stocksProfit, target):
    seen = set()
    distinct_pairs = set()

    for profit in stocksProfit:
        complement = target - profit
        if complement in seen:
            pair = tuple(sorted([profit, complement]))
            distinct_pairs.add(pair)
        seen.add(profit)

    return len(distinct_pairs)

# Example usage:
stocksProfit = [5, 7, 9, 13, 11, 6, 6, 3, 3]
target = 12
print(distinctStockPairs(stocksProfit, target))  # Output: 3
