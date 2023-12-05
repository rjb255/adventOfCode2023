import re


def lineDecoder(codeLine: str, alphanumeric=False):
    dictionary = {str(i): i for i in range(10)}
    if alphanumeric:
        words = {
            "one": 1,
            "two": 2,
            "three": 3,
            "four": 4,
            "five": 5,
            "six": 6,
            "seven": 7,
            "eight": 8,
            "nine": 9,
        }
        dictionary.update(words)

    results = re.findall(f"(?=({'|'.join(dictionary.keys())}))", codeLine)
    return dictionary[results[0]] * 10 + dictionary[results[-1]]


with open("day1input.txt", "r") as f:
    code = f.readlines()

part1 = map(lineDecoder, code)
part2 = map(lambda line: lineDecoder(line, True), code)

print(f"Part 1: {sum(part1)}")
print(f"Part 2: {sum(part2)}")
