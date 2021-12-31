def solution(k, dungeons):
    from itertools import permutations

    max_cnt = 0
    for seq in list(permutations(dungeons)):
        cnt = 0
        point = k
        for need_point, use_point in seq:
            if need_point > point:
                break

            point -= use_point
            cnt += 1

        max_cnt = max(max_cnt, cnt)

    return max_cnt