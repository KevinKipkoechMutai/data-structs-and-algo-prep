function maxValue(n, index, maxSum) {
    // Helper function to calculate the total sum
    const calculateSum = (mid) => {
        let leftLen = index;
        let rightLen = n - index - 1;
        
        // Calculate sum for the left side
        let leftSum;
        if (mid > leftLen) {
            leftSum = ((mid - 1) + (mid - leftLen)) * leftLen / 2;
        } else {
            leftSum = ((mid - 1) + 1) * mid / 2 + (leftLen - mid + 1);
        }

        // Calculate sum for the right side
        let rightSum;
        if (mid > rightLen) {
            rightSum = ((mid - 1) + (mid - rightLen)) * rightLen / 2;
        } else {
            rightSum = ((mid - 1) + 1) * mid / 2 + (rightLen - mid + 1);
        }

        return leftSum + rightSum + mid;
    };

    let low = 1, high = maxSum;
    while (low < high) {
        let mid = Math.floor((low + high + 1) / 2);
        if (calculateSum(mid) <= maxSum) {
            low = mid;
        } else {
            high = mid - 1;
        }
    }

    return low;
}

// Example usage:
let n = 3;
let index = 1;
let maxSum = 6;
console.log(maxValue(n, index, maxSum));  // Output: 2
