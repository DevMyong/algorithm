def solution(n):
    flag = [True] * n
    sqr_n = int(n ** 0.5)

    for i in range(2, sqr_n+1):
        if flag[i-1]:
            for j in range(i*2, n+1, i):
                flag[j-1] = False

    return flag.count(True)-1