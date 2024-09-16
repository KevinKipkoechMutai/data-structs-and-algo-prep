
const kangaroo = (x1: number, v1: number, x2: number, v2: number) => {
    if (x1 < x2 && v1 < v2) {
        return "NO"
    }
    else {
        if (v1 !== v2 && (x2 - x1)%(v2 - v1) === 0) {
            return "YES"
        } else {
            return "NO"
        }
    }
}