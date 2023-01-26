# (I) Планировщик задач (30 баллов)
# Полное решение: 30 баллов

from heapq import heapify, heappop, heappush
 
 
def check_available(queue_p: list, current_time: int, processors: list) -> None:
    if queue_p:
        while True:
            try:
                if current_time >= queue_p[0][0]:
                    heappush(processors, heappop(queue_p)[1])
                else:
                    break
            except IndexError:
                break
 
 
def calc_power_consumption(processors: list, tasks: list) -> int:
    tasks = list(reversed(tasks))
    heapify(processors)
    queued = []
    consumption = 0
    current_time = tasks[-1][0]
    while tasks:
        check_available(queued, current_time, processors)
        upcoming_task = tasks.pop()
        if processors:
            using_processor = heappop(processors)
            consumption += upcoming_task[1] * using_processor
            heappush(queued, (upcoming_task[1] + current_time, using_processor))
 
        current_time = tasks[-1][0] if tasks else current_time
 
    return consumption
 
 
def get_task_data_from_txt() -> tuple[list[int], list[tuple[int]], int]:
    _, queries = input().split()
    proc = list(map(lambda x: int(x), input().split()))
    tasks = [tuple(map(lambda x: int(x), input().split())) for _ in range(int(queries))]
 
    return proc, tasks
 
def main() -> None:
    proc, t = get_task_data_from_txt()
    print(calc_power_consumption(proc, t))
 
 
if __name__ == '__main__':
    main()