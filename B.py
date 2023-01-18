# B. Сумма к оплате (10 баллов)

def sol() -> int:
    d = {}
    for i in map(int, input().split()):
        if i in d:
            d[i] += 1
            continue
        d[i] = 1
    return sum((v - v // 3 ) * k for k, v in d.items())

  
def main():
    for i in range(int(input())):
        input()
        print(sol())
 

if __name__ == '__main__':
    main()
