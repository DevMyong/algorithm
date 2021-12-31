def solution(s):
    m = {"zero": '0', "one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

    word = ""
    answer = ""
    idx = 0

    for i in range(len(s)):
        if s[i] >= 'A':
            # 문자일 때
            word = s[idx:i + 1]
        else:
            # 숫자일 때
            answer += s[i]
            idx = i + 1

        if word in m:
            # 단어가 완성됐을 때
            answer += m.get(word)
            idx = i + 1
            word = ""

    return int(answer)