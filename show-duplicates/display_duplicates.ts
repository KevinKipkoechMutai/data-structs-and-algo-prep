

const displayDuplicates = (nums: number[]): number => {
    // create array to store duplicates
    let duplicates: number[] = []
    let seen: number[] = []

    // loop through numbers
    // add new number to the duplicates array
    // if it's already there, skip it
    for (let i = 0; i < nums.length; i++) {
        if (seen.includes(nums[i]) && !duplicates.includes(nums[i])) {
            duplicates.push(nums[i])
        } else {
            seen.push(nums[i])
        }
    }

    //return the length of the duplicates array
    return duplicates.length
}

console.log(displayDuplicates([10, 20, 60, 30, 20, 40, 30, 60, 70, 80]))