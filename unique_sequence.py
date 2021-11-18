def clear_sequence(n):
    i = 0
    last = input().strip()
    res = [last]
    while i < n - 1:
        num = input().strip()
        i += 1
        if num != last:
            res.append(num)
            last = num
    print("\n".join(res))


if __name__ == "__main__":
    n = int(input().strip())
    clear_sequence(n)
