def get_2_num_sum(arr, target):
    buf = {}
    for num in arr:
        if num in buf:
            return [num, buf[num]]
        else:
            buf[target - num] = num
    return []


if __name__ == '__main__':
    print(get_2_num_sum([2, 9, -1, 3, 10, 11], 8))

