def solution(places):
    answer = []

    for place in places:
        answer.append(check_room(place))

    return answer


def check_room(place):
    for y in range(5):
        for x in range(5):
            if place[y][x] == "X":
                continue

            if place[y][x] == "P":
                for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                    if y + dy < 0 or y + dy >= 5 or x + dx < 0 or x + dx >= 5:
                        continue

                    if place[y + dy][x + dx] == "P":
                        return 0

            if place[y][x] == "O":
                p_cnt = 0
                for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                    if y + dy < 0 or y + dy >= 5 or x + dx < 0 or x + dx >= 5:
                        continue

                    if place[y + dy][x + dx] == "P":
                        p_cnt += 1
                        if p_cnt == 2:
                            return 0
    return 1