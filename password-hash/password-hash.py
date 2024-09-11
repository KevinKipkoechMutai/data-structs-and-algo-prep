def authEvents(events):
    p = 131
    M = 1000000007

    password_hash = 0

    def hash_string(s):
        hash_val = 0
        for c in s:
            hash_val = (hash_val * p + ord(c)) % M
        return hash_val

    def authorize(x, current_hash):
        if x == current_hash:
            return 1
        for i in range(128):  # ASCII range
            new_hash = (current_hash * p + i) % M
            if new_hash == x:
                return 1
        return 0

    result = []
    for command, param in events:
        if command == "setPassword":
            password_hash = hash_string(param)
        elif command == "authorize":
            x = int(param)
            result.append(authorize(x, password_hash))

    return result

# Example usage:
events = [
    ["setPassword", "cAr1"],
    ["authorize", "223691457"],
    ["authorize", "303580761"]
]
print(authEvents(events))  # Expected output: [1, 1]