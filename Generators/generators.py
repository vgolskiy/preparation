# reading big files
def csv_reader(file):
    for row in open(file, "r"):
        yield row


# infinite sequence
def infinite_sequence():
    n = 0
    while True:
        yield n
        n += 1


# is palindrome
def is_palindrome(n):
    if not n//10:
        return False
    tmp = n
    reverse = 0
    while tmp:
        reverse = (reverse * 10) + tmp % 10
        tmp //= 10
    if n == reverse:
        return True
    return False
