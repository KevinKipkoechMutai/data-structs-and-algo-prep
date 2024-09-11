
const reverseWords = (sentence: string) => {
    const sentenceList = sentence.split("")
    for (let i = 0; i < sentenceList.length; i++) {
        sentenceList[i].split('').reverse().join('')
    }

    return sentenceList.join(" ")
}