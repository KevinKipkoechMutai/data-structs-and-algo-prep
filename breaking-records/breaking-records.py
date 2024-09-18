def breakingRecords(scores):
    max_count = 0
    min_count = 0

    minimum = scores[0]
    maximum = scores[0]

    for i in range(len(scores)):
        if scores[i] < minimum:
            minimum = scores[i]
            min_count += 1
        elif scores[i] > maximum:
            maximum = scores[i]
            max_count += 1
    
    return [max_count, min_count]