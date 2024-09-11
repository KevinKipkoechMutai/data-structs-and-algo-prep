def maxValue(n, index, maxSum):
    # Binary search to find the maximum value at index 'k'
    def calculateSum(mid):
        # Compute the sum of the array with 'mid' at the given index
        left_len = index
        right_len = n - index - 1
        
        # Sum of the left part of the array
        if mid > left_len:
            left_sum = (mid - 1 + mid - left_len) * left_len // 2
        else:
            left_sum = (mid - 1 + 1) * mid // 2 + (left_len - mid + 1)
        
        # Sum of the right part of the array
        if mid > right_len:
            right_sum = (mid - 1 + mid - right_len) * right_len // 2
        else:
            right_sum = (mid - 1 + 1) * mid // 2 + (right_len - mid + 1)
        
        return left_sum + right_sum + mid
    
    low, high = 1, maxSum
    while low < high:
        mid = (low + high + 1) // 2
        if calculateSum(mid) <= maxSum:
            low = mid
        else:
            high = mid - 1
    
    return low

# Example usage:
n = 3
index = 1
maxSum = 6
print(maxValue(n, index, maxSum))  # Output: 2
