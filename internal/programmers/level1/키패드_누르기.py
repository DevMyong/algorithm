def solution(numbers, hand):
    answer = ''
    left_pos = 10
    right_pos = 12

    for num in numbers:
        if num in [1, 4, 7]:
            left_pos = num
            answer += "L"
        elif num in [3, 6, 9]:
            right_pos = num
            answer += "R"
        else:
            if num == 0:
                num = 11

            abs_left = abs(left_pos - num)
            abs_right = abs(right_pos - num)

            if sum(divmod(abs_left, 3)) < sum(divmod(abs_right, 3)):
                left_pos = num
                answer += "L"
            elif sum(divmod(abs_left, 3)) > sum(divmod(abs_right, 3)):
                right_pos = num
                answer += "R"
            else:
                if hand == "right":
                    right_pos = num
                    answer += "R"
                else:
                    left_pos = num
                    answer += "L"
    return answer