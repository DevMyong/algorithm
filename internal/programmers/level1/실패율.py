def solution(N, stages):
    result = {}
    tried_person = len(stages)

    for stage in range(1, N+1):
        failed = stages.count(stage)
        if tried_person > 0:
            result[stage] = failed / tried_person
        else:
            result[stage] = 0
        tried_person -= failed
    return sorted(result, key=lambda x: result[x], reverse=True)