arr = [1,2,3,4,5,6]
s = 5


def find_sum_memb(arr, s):
	tmp = {}
	for n in arr:
		if n in tmp:
			return [n, tmp[n]]
		else:
			tmp[s-n]=n
	return []

print(find_sum_memb(arr, s))
