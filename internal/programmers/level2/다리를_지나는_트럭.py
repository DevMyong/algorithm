def solution(bridge_length, weight, truck_weights):
    from collections import deque

    wait = deque(truck_weights)
    crossing = deque()
    time_que = deque()

    count = 0
    crossing.append(wait.popleft())
    time_que.append(count)

    while crossing:
        # 건넌 시간 확인
        if time_que and count - time_que[0] >= bridge_length:
            crossing.popleft()
            time_que.popleft()

        # 도로가 비어있는지 확인
        if not crossing and wait:
            crossing.append(wait.popleft())
            time_que.append(count)
        # 무게 초과하는지 확인
        elif wait and sum(crossing) + wait[0] <= weight:
            crossing.append(wait.popleft())
            time_que.append(count)
        count += 1

    return count