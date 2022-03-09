def find_non_rep(s):
	tmp = {}
	for ch in s:
		tmp[ch] = tmp.get(ch, 0) + 1
	for i, ch in enumerate(s):
		if tmp[ch] == 1:
			return i
	return -1

s = "abracadabra"
print(find_non_rep(s))
