from itertools import permutations


def solution(numbers):
    permuted_nums = set()

    for i in range(1, len(numbers) + 1):
        for permuted_num in set(permutations(list(numbers), i)):
            num = int(''.join(list(permuted_num)))
            permuted_nums.add(num)

    permuted_nums -= set(range(0, 2))

    maximum = max(permuted_nums)
    for i in range(2, int(maximum ** 0.5)):
        permuted_nums -= set(range(i * 2, maximum + 1, i))

    return len(permuted_nums)