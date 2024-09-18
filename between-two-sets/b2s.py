# import math

# # Function to compute the Least Common Multiple (LCM) of two numbers
# def lcm(x, y):
#     return abs(x * y) // math.gcd(x, y)

# # Function to compute the Greatest Common Divisor (GCD) of an array
# def gcd_array(arr):
#     return math.gcd(arr[0], arr[1]) if len(arr) == 2 else math.gcd(arr[0], gcd_array(arr[1:]))

# # Function to compute the LCM of an array
# def lcm_array(arr):
#     result = arr[0]
#     for num in arr[1:]:
#         result = lcm(result, num)
#     return result

# # Function to determine how many numbers are between two arrays
# def getTotalX(a, b):
#     # Step 1: Find the LCM of array a
#     lcm_a = lcm_array(a)

#     # Step 2: Find the GCD of array b
#     gcd_b = gcd_array(b)

#     # Step 3: Find the numbers between the two arrays
#     count = 0
#     multiple = lcm_a
#     while multiple <= gcd_b:
#         if gcd_b % multiple == 0:
#             count += 1
#         multiple += lcm_a

#     return count

# # Sample input
# a = [2, 4]
# b = [16, 32, 96]
# getTotalX(a, b)

def getTotalX(a, b):
    res = 0

    for i in range(max(a), min(b) + 1):
        res += (all([i % x == 0 for x in a]) and all([y % i == 0 for y in b]))
    return res
