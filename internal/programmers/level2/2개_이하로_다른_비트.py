def solution(numbers):
    answer = []

    for num in numbers:
        if num % 2 == 0:
            answer.append(num + 1)
            continue

        bin_num = bin(num)[2:]
        idx_lsb0 = bin_num.rfind('0')
        if idx_lsb0 == -1:
            answer.append(int("10" + bin_num[1:], 2))
            continue

        bin_num = bin_num[:idx_lsb0] + "10" + bin_num[idx_lsb0 + 2:]
        answer.append(int(bin_num, 2))

    return answer