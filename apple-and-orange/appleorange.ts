const countApplesAndOranges = (s: number, t: number, a: number, b:number, apples: number[], oranges: number[]) => {
    let acount = 0
    let bcount = 0

    for (let i = 0; i < apples.length; i++) {
        let temp = a + apples[i]
        if (temp >= s && temp <= t) {
            acount++
        }
    }

    for (let i = 0; i < oranges.length; i++) {
        let temp =  + oranges[i]
        if (temp >= s && temp <= t) {
            bcount++
        }
    }

    console.log(acount)
    console.log(bcount)
}