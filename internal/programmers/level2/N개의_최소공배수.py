def solution(nums):
    def gcd(a, b):
        n = 0
        a, b = max(a, b), min(a, b)

        while b != 0:
            n = a % b
            a = b
            b = n

        return a

    answer = nums[0]
    for num in nums:
        answer = num * answer // gcd(num, answer)

    return answer