function authEvents(events) {
    const p = 131
    const M = Math.pow(10, 9) + 7

    const hashString = (s) => {
        let hash = 0
        for (let i = 0; i < s.length; i++) {
            hash = (hash * p + s.charCodeAt(i)) % M
        }
        return hash
    }

    const authorize = (x, currentHash) => {
        if (x === currentHash) return 1

        for (let i = 0; i < 128; i++) {
            let newHash = (currentHash * p + 1) % M
            if (newHash === x) return 1
        }
        return 0
    }

    const result = []
    for (let [command, param] of events) {
        if (command === "setPassword") {
            passwordHash = hashString(param)
        } else if (command === "authorize") {
            const x = parseInt(param)
            result.push(authorize(x, passwordHash))
        }
    }

    return result
}