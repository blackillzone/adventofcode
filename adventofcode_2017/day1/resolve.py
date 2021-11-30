#!/usr/bin/env python
"""Day 01. Advent of Code."""

input_data = ''
with open('input.txt', 'r') as keyfile:
    input_data = list(map(int, keyfile.read().strip()))


def puzzle01(input_data):
    """Function to resolve first puzzle of advent of code."""
    result = 0
    ref_list = input_data[1:] + input_data[:1]

    for i, j in zip(input_data, ref_list):
        if i == j:
            result += i
    return result


def puzzle02(input_data):
    """Function to resolve second puzzle of advent of code."""
    result = 0
    half = int(len(input_data)/2)
    ref_list = input_data[half:] + input_data[:half]

    for i, j in zip(input_data, ref_list):
        if i == j:
            result += i
    return result


print ("Puzzle 01: " + str(puzzle01(input_data)))
print ("Puzzle 02: " + str(puzzle02(input_data)))
