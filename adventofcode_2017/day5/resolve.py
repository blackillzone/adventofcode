#!/usr/bin/env python
"""Day 05. Advent of Code."""


def isValidIndex(list, index):
    """Function to test if out of index."""
    try:
        list[index]
        return True
    except IndexError:
        return False


with open('input.txt', 'r') as x:
    numbers = x.read().splitlines()
    numbers = [int(a) for a in numbers]
    i = 0
    count = 0

    while isValidIndex(numbers, i):
        numbers[i] += 1
        i += numbers[i] - 1
        count += 1

    print('Puzzle 1: ' + str(count))


with open('input.txt', 'r') as x:
    numbers = x.read().splitlines()
    numbers = [int(a) for a in numbers]
    i = 0
    count = 0

    while isValidIndex(numbers, i):
        if numbers[i] >= 3:
            numbers[i] -= 1
            i += numbers[i] + 1
        else:
            numbers[i] += 1
            i += numbers[i] - 1
        count += 1

    print('Puzzle 2: ' + str(count))
