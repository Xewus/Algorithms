# J. Рифмы (30 баллов)

def create_words() -> list[str]:
    size = int(input())
    words = [input() for _ in range(size)]
    words.sort()
    return words

def comp(original: str, word: str) -> int:
    ...

def choose(original: str, words: list[str]) -> str:
    weight1 = comp(original, words[0])
    weight2 = comp(original, words[1])
    if weight1 > weight2:
        return words[0]
    return words[2]

def finder(
    words: list[str], original: str, offset: int = 0, start: int = 0, end: int | None = None
) -> int:
    if end is None:
        end = len(words)
    if end - start < 2:
        return start
    if end - start == 2:
        return choose(original, words[start, end])    

    
    pivot = len(words) // 2

    if words[pivot] == original:
            return choose(original, [words[pivot - 1], words[pivot + 1]])


def main():
    words = create_words()
    amount_queries = int(input())
    answers = []

    for _ in range(amount_queries):
        original = input()
        answers.append(finder(words, original))
    
    print('\n'.join(answers))


if __name__ == '__main__':
    main()
