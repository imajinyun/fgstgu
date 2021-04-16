def drop(x: int) -> int:
    digits = []
    y = x
    for _ in range(0, len(str(y))):
        digits.append(y % 10)
        y = y // 10
    digits.sort(reverse=True)
    small, large = 0, 0
    for n in digits:
        large = large * 10 + n
    digits.sort()
    for n in digits:
        small = small * 10 + n
    return large - small


def dropdown(x: int):
    numbers = [x]
    blackhole = []
    for _ in range(500):
        y = drop(x)
        if y in numbers:
            blackhole = numbers[numbers.index(y):]
            break
        x = y
        numbers.append(x)
    minpos = blackhole.index(min(blackhole))
    blackhole = blackhole[minpos:] + blackhole[0:minpos]
    return numbers, blackhole


if __name__ == "__main__":
    result = dropdown(37)
    print(result)
