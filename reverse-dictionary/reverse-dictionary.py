ascii_dict={ 'A': 65, 'B': 66, 'C': 67, 'D': 68 }

reversed_dict = { value: key for key, value in ascii_dict.items() }

print(reversed_dict)