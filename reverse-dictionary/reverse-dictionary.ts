let asciiDict: { [key: string]: number } = {'A': 65, 'B': 66, 'C': 67, 'D': 68}

let reversedDict: { [key: number]: string } = Object.fromEntries(
    Object.entries(asciiDict).map(([key, value]) => [value, key])
)

console.log(reversedDict)