
import collections

sample_list = [10, 20, 60, 30, 20, 40, 30, 60, 70, 80]
# exist = {}
# duplicates = []

# for x in sample_list:
#     if x not in exist:
#         exist[x] = 1
#     else:
#         duplicates.append(x)

# print(duplicates)

duplicates = []

for item, count in collections.Counter(sample_list).items():
    if count > 1:
        duplicates.append(item)

print(duplicates)