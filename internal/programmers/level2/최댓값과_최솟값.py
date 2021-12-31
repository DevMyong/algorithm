def solution(s):
    r = [int(ch) for ch in s.split()]
    return f'{min(r)} {max(r)}'