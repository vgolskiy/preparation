# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

import re

def verify(w):
    alphas = 0
    digits = 0
    for ch in w:
        if ch.isalpha():
            alphas += 1
        if ch.isdigit():
            digits += 1
    if not alphas % 2 and digits % 2:
        return True
    return False

def solution(s):
    length = 0
    tmp = s.split()
    for word in tmp:
        if re.findall(r'[^a-zA-Z0-9]', word):
            continue
        if not verify(word):
            continue
        l = len(word)
        if l > length:
            length = l
    if length > 0:
        return length
    return -1
