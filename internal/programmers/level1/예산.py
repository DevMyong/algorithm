def solution(department, budget):
    department.sort()
    answer = 0
    for d in department:
        if budget >= d:
            budget = budget - d
            answer += 1
    return answer