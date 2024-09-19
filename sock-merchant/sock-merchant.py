def findPairs(arr):
    color_frequency = {}
    count = 0

    for num in arr:
        color_frequency[num] = color_frequency.get(num, 0) + 1
    
    for value in color_frequency.values():
        if value % 2 != 0:
            count += (value - 1) / 2
        else:
            count += value / 2