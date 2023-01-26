# (D) Электронная таблица (10 баллов)
# Полное решение: 10 баллов

def main():
    for _ in range(int(input())):
        input()  # empty string in StdIn
        n, _ = map(int, input().split())
        mtrx = [list(map(int, input().split())) for i in range(n)]
        input()
        cols = tuple(map(int, input().split()))
        for col in cols:
            col -= 1
            mtrx.sort(key=lambda row: row[col])

        for row in mtrx:
            print(' '.join(str(i) for i in row))

if __name__ == '__main__':
    main()
