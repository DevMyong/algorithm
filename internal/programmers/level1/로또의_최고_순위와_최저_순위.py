def solution(lottos, win_nums):
    n_deleted = lottos.count(0)
    deleted_num = 0

    for win_num in win_nums:
        if win_num in lottos:
            deleted_num += 1

    high_score = 7 - (deleted_num + n_deleted) if 7 - (deleted_num + n_deleted) <= 6 else 6
    low_score = 7 - deleted_num if 7 - deleted_num <= 6 else 6

    return [high_score, low_score]