# (F) Отрезки времени (20 баллов), Полное решение: 20 баллов

def validate_time(time_: str):
    h, m, s = map(int, time_.split(':'))
    if h < 0 or m < 0 or s < 0 or h > 23 or m > 59 or s > 59:
        return 0, False
    return s + m * 60 + h * 3600, True


def validate_timeline(timeline: list[str]):
    start, ok = validate_time(timeline[0])
    if not ok:
        return 
    end, ok = validate_time(timeline[1])
    if not ok or start > end:
        return
    return start, end


def main():
    queries = int(input())
    answers = [None] * queries
    for q in range(queries):
        valid = 1
        amount = int(input())
        data = [None] * amount
        for i in range(amount):
            if valid:
                data[i] = input().split('-')
                data[i] = validate_timeline(data[i])
                if data[i] is None:
                    valid = 0
            else:
                input()
            
        if not valid:
            answers[q] = (('NO', 'YES')[valid])
            continue
 
        line = [i for i in range(86_400)]
        for timeline in data:
            if not valid:
                break
            for sec in range(timeline[0], timeline[1] + 1):
                if line[sec] is None:
                    valid = 0
                    break
                line[sec] = None
        answers[q] = (('NO', 'YES')[valid])
    print('\n'.join(answers))

if __name__ == '__main__':
    main()
