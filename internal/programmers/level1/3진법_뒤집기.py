def solution(n):
    reversed_string = ''

    while n > 0:
        n, mod = divmod(n, 3)
        reversed_string += str(mod)

    return int(reversed_string, 3)