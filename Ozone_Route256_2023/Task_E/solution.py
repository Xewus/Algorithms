# (E) Отчет (15 баллов)
# Полное решение: 15 баллов

def main():
    for _ in range(int(input())):
        tasks = set()
        amount = int(input())
        data = input().split()
        tasks.add(data[0])

        for day in range(1, amount):
            if data[day] == data[day - 1]:
                continue
            if data[day] in tasks:
                print('NO')
                break
            tasks.add(data[day])
        else:
            print('YES')


if __name__ == '__main__':
    main()
