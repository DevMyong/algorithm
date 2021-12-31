def solution(absolutes, signs):
    answer = 0
    for absolute, sign in zip(absolutes, signs):
        if not sign:
            answer += absolute * -1
        else:
            answer += absolute
    return answer