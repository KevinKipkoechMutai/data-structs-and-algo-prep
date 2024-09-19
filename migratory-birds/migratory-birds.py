
def migratoryBirds(arr):
    frequency = {}

    for num in arr:
        # if num in frequency:
        #     frequency[num] += 1
        # else:
        #     frequency[num] = 1
        frequency[num] = frequency.get(num, 0) + 1
    
    return min(frequency, key= lambda x: (-frequency[x], x))