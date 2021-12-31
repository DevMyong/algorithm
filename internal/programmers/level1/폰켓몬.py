def solution(nums):
    ordered_nums = list(set(nums))

    return len(ordered_nums[0:int(len(nums) / 2)])