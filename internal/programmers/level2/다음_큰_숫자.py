def solution(n):
    n_bin = bin(n)[2:]

    idx_one = n_bin.rfind("1")
    idx_zero = n_bin.rfind("0", 0, idx_one)

    if idx_zero == -1:
        answer_left = "10"
    else:
        answer_left = n_bin[:idx_zero] + "10"

    num_one = n_bin[idx_zero + 1:].count("1") - 1
    len_right = len(n_bin) - (idx_zero + 2) - num_one

    answer_right = "0" * len_right + "1" * num_one
    answer = answer_left + answer_right

    return int(answer, 2)