var displayDuplicates = function (nums) {
    // create array to store duplicates
    var duplicates = [];
    var seen = [];
    // loop through numbers
    // add new number to the duplicates array
    // if it's already there, skip it
    for (var i = 0; i < nums.length; i++) {
        if (seen.includes(nums[i]) && !duplicates.includes(nums[i])) {
            duplicates.push(nums[i]);
        }
        else {
            seen.push(nums[i]);
        }
    }
    //return the length of the duplicates array
    return duplicates.length;
};
console.log(displayDuplicates([10, 20, 60, 30, 20, 40, 30, 60, 70, 80]));
