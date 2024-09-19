
def solve(s, d, m):
    count = 0

    for i in range((len(s) - m) + 1):
        if sum(s[i:i+m]) == d:
            count += 1
    
    return count


print(solve([2,1,2,3,4,1,2], 3, 2))