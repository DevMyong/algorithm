def solution(answers):
    paper1 = [1,2,3,4,5]
    paper2 = [2,1,2,3,2,4,2,5]
    paper3 = [3,3,1,1,2,2,4,4,5,5]
    
    score1, score2, score3 = 0, 0, 0
    
    for i in range(len(answers)) :
        if answers[i] == paper1[i%5] :
            score1 += 1
        if answers[i] == paper2[i%8] :
            score2 += 1
        if answers[i] == paper3[i%10] :
            score3 += 1
    
    res = max([score1, score2, score3])
    answer = []
    if res == score1 :
        answer.append(1)
    if res == score2 :
        answer.append(2)
    if res == score3 :
        answer.append(3)
        
    answer.sort()
    return answer


print(solution([1,2,3,4,5]))