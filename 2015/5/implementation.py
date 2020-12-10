FORBIDDEN_PHRASES= {
    'ab': True,
    'cd': True,
    'pq': True,
    'xy': True,
}

VOWELS = {
    'a': True,
    'e': True,
    'i': True,
    'o': True,
    'u': True,
}

def is_nice(input_string):
    last_letter = ''
    seen_vowels = {
        'a': 0,
        'e': 0,
        'i': 0,
        'o': 0,
        'u': 0,
    }
    repeat_seen = False
    for letter in input_string:
        if last_letter + letter in FORBIDDEN_PHRASES:
            return False
        if letter in VOWELS:
            seen_vowels[letter] = seen_vowels[letter] + 1
        if last_letter == letter:
            repeat_seen = True
        last_letter = letter
    return sum(list(seen_vowels.values())) > 2 and repeat_seen

test_table = [
    ['ugknbfddgicrmopn', True],
    ['aaa', True],
    ['jchzalrnumimnmhp', False],
    ['haegwjzuvuyypxyu', False],
    ['dvszwmarrgswjxmb', False],
]

def tests():
    for tt in test_table:
        result = is_nice(tt[0])
        if result != tt[1]:
            raise AssertionError('{} expected {} got {}'.format(tt[0], tt[1], result))

if __name__ == '__main__':
    tests()
    with open('input.txt') as infile:
        input_strings = infile.readlines()
    nice_strings = [i for i in input_strings if is_nice(i)]
    print(len(nice_strings))
