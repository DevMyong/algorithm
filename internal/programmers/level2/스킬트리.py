def solution(sequence, skill_trees):
    answer = 0

    for skill_tree in skill_trees:
        s = ""
        
        for skill in skill_tree:
            if skill in sequence:
                s += skill

        for i in range(len(s)):
            if s[i] != sequence[i]:
                break
        else:
            answer += 1

    return answer