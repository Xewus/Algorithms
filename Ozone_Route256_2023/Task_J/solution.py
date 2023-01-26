# (J) Рифмы (30 баллов)
# Полное решение: 30 баллов

def create_words() -> list[str]:
    size = int(input())
    words = [input()[::-1] for _ in range(size)]
    words.sort()
    return words


def choose(words: list[str], request: str) -> str:
    rate = []
    for word in words:
        lim = min(len(request), len(word))
        accord = 0
        for idx in range(lim):
            if request[idx] != word[idx]:
                break
            accord += 1
        if accord == len(request):
            if accord != len(word):
                return word
            continue
        rate.append((accord, word))
    return max(rate, key=lambda x: x[0])[1]


def finder(
    words: list[str], request: str, start: int, end: int
) -> int:
    dist = end - start
    if dist == 1:
        return start

    pivot = dist // 2 + start
    if words[pivot] == request:
        return pivot
    if words[pivot] > request:
        return finder(words, request, start, pivot)
    return finder(words, request, pivot, end)


def get_word(words: list[str], request: str) -> str:
    if len(words) < 3:
        return choose(words, request)

    index = finder(words, request, 0, len(words))
    if index == 0:
        return choose(words[:2], request)
    return choose(words[index - 1: index + 2], request)

 
def main():
    words = create_words()
    amount_queries = int(input())
    answers = []
    cache = {}

    for i in range(amount_queries):
        request = input()[::-1]
        answer = cache.get(request)
        if not answer is None:
            answers.append(answer)
            continue
        answer = get_word(words, request)
        answers.append(answer[::-1])
        cache[request] = answer[::-1]

    print('\n'.join(answers))


if __name__ == '__main__':
    main()
