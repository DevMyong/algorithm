def solution(nums):
    answer = 0

    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            for k in range(j + 1, len(nums)):
                sum_nums = nums[i] + nums[j] + nums[k]

                for p in range(2, sum_nums // 2 + 1):
                    if sum_nums % p == 0:
                        break
                else:
                    answer += 1

    return answer