# (B) Сумма к оплате (10 баллов), Полное решение: 10 баллов.

def solution() -> int:
    basket = {}
    for product in map(int, input().split()):
        if product in basket:
            basket[product] += 1
            continue
        basket[product] = 1
    return sum(
        (amount - amount // 3) * prod for prod, amount in basket.items()
    )


def main():
    for _ in range(int(input())):
        input()
        print(solution())


if __name__ == '__main__':
    main()
