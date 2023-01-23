# (C) Парное программирование (10 баллов), Полное решение: 10 баллов.

def solution():
    devs = int(input())
    skills = list(map(int, input().split()))

    if devs == 2:
        print(1, 2, '\n')
        return

    first = 0
    second = 1

    for _ in range(devs // 2 - 1):
        min_div = abs(skills[first] - skills[second])

        for i in range(first + 2, devs):
            if skills[i] is None:
                continue

            if abs(skills[first] - skills[i]) < min_div:
                min_div = abs(skills[first] - skills[i])
                second = i

            if min_div == 0:
                break
    
        skills[first] = skills[second] = None
        print(first + 1, second + 1)

        first += 1
        while first < devs:
            if not skills[first] is None:
                break
            first += 1
        
        second = first + 1
        while first < devs:
            if not skills[second] is None:
                break
            second += 1

    print(first + 1, second + 1, '\n')


def main():
    for _ in range(int(input())):
        solution()


if __name__ == '__main__':
    main()
