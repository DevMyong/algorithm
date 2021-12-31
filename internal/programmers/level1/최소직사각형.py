def solution(sizes):
    for i in range(len(sizes)):
        if sizes[i][0] < sizes[i][1]:
            sizes[i][1], sizes[i][0] = sizes[i][0], sizes[i][1]

    max_row = max(sizes, key=lambda x: x[0])[0]
    max_col = max(sizes, key=lambda x: x[1])[1]

    return max_row * max_col