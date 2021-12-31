def solution(record):
    answer = []
    users = dict()

    map_action = {"Leave": "님이 나갔습니다.", "Enter": "님이 들어왔습니다."}
    for r in record:
        r = r.split()
        uid, action, name = r[1], r[0], ""

        if action != "Leave":
            name = r[2]

        if action in ["Enter", "Change"]:
            users[uid] = name

    for r in record:
        r = r.split()
        uid, action = r[1], r[0]
        name = users[uid]

        if action != "Change":
            answer.append(f'{users[uid]}{map_action[action]}')

    return answer