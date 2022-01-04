def solution(s):
    flag21, flag12 = False, False

    i = 0
    while i < len(s) - 1:
        if not flag21 and s[i:i + 2] == "21":
            flag21 = True
            i += 1
        elif not flag12 and s[i:i + 2] == "12":
            flag12 = True
            i += 1

        i += 1
        if flag21 & flag12:
            return "Yes"
    return "No"


def solution2(s):
    if "12" not in s or "21" not in s:
        return "No"

    if s.index("12") < s.index("21") and "21" in s[s.index("12") + 2:]:
        return "Yes"
    elif s.index("12") > s.index("21") and "12" in s[s.index("21") + 2:]:
        return "Yes"
    return "No"


if __name__ == "__main__":
    s = input()
    print(solution(s))
    print(solution2(s))
