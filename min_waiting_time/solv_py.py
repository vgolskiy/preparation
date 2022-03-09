def min_wait_time(arr) -> int:
	arr.sort()
	res = 0
	n = len(arr)
	for i, q in enumerate(arr):
		res += q * (n - (i + 1))
	return res

if __name__ == "__main__":
	arr = [3, 2, 1, 2, 6]
	print(min_wait_time(arr))