def solution(N):
    gap = False
    maximal = 0
    qty = 0
    while N:
        if N%2:
            gap = True
            if qty:
                if qty > maximal:
                    maximal = qty
                qty = 0
        elif gap or (not gap and qty):
            qty += 1
            gap = False
        N = N//2
    return maximal
