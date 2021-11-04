def nonConstructibleChange(coins):
    if len(coins) > 0:
        coins.sort()
        max_change = 0
        for coin in coins:
            if max_change + 1 < coin:
                break
            max_change += coin
        return max_change + 1
    return 1


if __name__ == '__main__':
    print(nonConstructibleChange([1, 1, 1, 1, 1, 1]))
