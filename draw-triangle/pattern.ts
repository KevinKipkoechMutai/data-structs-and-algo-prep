let rows = 5
let x = 0

for (let i = rows; i > 0; i--) {
    x++
    console.log((x.toString() + " ").repeat(i))
}