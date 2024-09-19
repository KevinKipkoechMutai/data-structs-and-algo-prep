import math

def drawingBook(n, p):
    front = p // 2
    back = (n - p) / 2

    return min(front, math.ceil(back) if n % 2 == 0 else math.floor(back))