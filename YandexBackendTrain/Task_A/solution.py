# Python 3.9 (PyPy 7.3.11)	
# OK 283ms 41.90Mb

def inp() -> list:
    input()
    return list(map(int, input().split()))

def solution(arr: list) -> int:
    if sorted(arr) != arr:
        return -1
    return arr[-1] - arr[0]

    return arr[p] - arr[0]

print(solution(inp()))
