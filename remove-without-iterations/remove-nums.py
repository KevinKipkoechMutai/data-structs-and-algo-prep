number_list = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]

i = 0

n = len(number_list)

while i < n:
    if number_list[i] > 50:
        del number_list[i]
        n -= 1
    else:
        i += 1

print(number_list)