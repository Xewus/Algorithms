# (H) Валидация карты (25 баллов)
# Полное решение: 25 баллов

def create_mapa(n: int, m: int) -> list:
    mapa = [[['', 0] for _ in range(m)] for i in range(n)]
    for row in range(n):
        data = input()
        for col, v in enumerate(data):
            if v == '.':
                mapa[row][col] = ['.', 1]
            else:
                mapa[row][col] = [v, 0]
    return mapa


def worm(mapa: list, was: set, row=0, col=0, color='') -> list:
    if row >= len(mapa) or col >= len(mapa[0]) or col < 0 or mapa[row][col][1] == 1:
        return mapa

    if mapa[row][col][0] != color:
        return mapa
    else:
        mapa[row][col][1] = 1

    if col < len(mapa[0]) - 2:
        mapa = worm(mapa, was, row, col + 2, mapa[row][col][0])
    if col > 1:
        mapa = worm(mapa, was, row, col - 2, mapa[row][col][0])

    if row < len(mapa):
        mapa = worm(mapa, was, row + 1, col + 1, mapa[row][col][0])
        mapa = worm(mapa, was, row + 1, col - 1, mapa[row][col][0])
    if row > 0:
        mapa = worm(mapa, was, row - 1, col + 1, mapa[row][col][0])
        mapa = worm(mapa, was, row - 1, col - 1, mapa[row][col][0])
    return mapa


def chek_mapa(mapa: list) -> str:
    for row in range(len(mapa)):
        for col in mapa[row]:
            if col[1] == 0:
                return "NO"
    return 'YES'


def solution():
    n, m = map(int, input().split())
    mapa = create_mapa(n, m)
    was = set()    

    for row in range(n):
        not_even = row & 1
        for col in range(0, m, 2):
            if not_even:
                col += 1
                if col >= m:
                    break
            if mapa[row][col][1] == 1 or mapa[row][col][0] in was:
                continue
            was.add(mapa[row][col][0])
            mapa = worm(mapa, was, row, col, mapa[row][col][0])

    print(chek_mapa(mapa))

def main():
    for _ in range(int(input())):
        solution()


if __name__ == '__main__':
    main()
