def is_str_possible(arr, s) -> bool:
	if len(s):
		tmp = {}
		for ch in arr:
			tmp[ch] = tmp.get(ch, 0) + 1
		for ch in s:
			if ch not in tmp or tmp[ch] - 1 < 0:
				return False
			tmp[ch]-=1
	return True

if __name__ == "__main__":
	arr = "abcabc"
	s = "aabcc"
	print(is_str_possible(arr, s))
