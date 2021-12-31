def solution(participant, completion):
    participant.sort()
    completion.sort()

    for i in range(0, len(participant)):
        if i == len(participant)-1 or participant[i] != completion[i]:
            return participant[i]