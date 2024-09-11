let numberList = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]

let i = 0
let n = numberList.length

while (i < n) {
    if (numberList[i] > 50) {
        numberList.splice(i, 1)
        n--
    } else {
        i++
    }
}

console.log(numberList)