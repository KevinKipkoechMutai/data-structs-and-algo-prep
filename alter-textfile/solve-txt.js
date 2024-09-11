import * as fs from 'fs'

fs.readFile('sample.txt', 'utf-8', (err, data) => {
    if (err) {
        console.error(err)
        return
    }
    const modifiedData = data.replace(/\n/g, ' ')
    console.log(modifiedData)
})