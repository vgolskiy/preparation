def solution(message, K):
    words = message.split()
    qty = 0
    res = ""
    for w in words:
        n = len(w)
        if qty + n + 1 <= K:
            res += w + " "
            qty += n + 1
        elif qty + n <= K:
            res += w
            qty += n
        else:
            break
    if len(res) and res[-1] == " ":
        return res[:len(res)-1]
    return res

print(solution("test or not test", 7))
