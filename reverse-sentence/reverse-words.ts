
const reverseWords = (sentence: string): string => {
    // split sentence into individual words
    const splitWords: string[] = sentence.split(" ")
    let reversedWords: string[] = []

    //iterate over split words and reverse each of them
    for (let i = 0; i < splitWords.length; i++) {
        const reversedWord: string = splitWords[i].split("").reverse().join("")
        reversedWords.push(reversedWord)
    }

    return reversedWords.join(" ")
}

console.log(reverseWords("My Name is Kevin"))