#!/usr/bin/env python
"""Day 07. Advent of Code."""
import re


def hasChilds(program):
    if program[2] is None:
        return False
    return True


with open('input_test.txt', 'r') as x:
    lines = x.read().splitlines()
    programs = []
    child_programs = []
    for line in lines:
        m = re.search('([a-z]+)\s\((\d+)\)(?:\s->\s)*(.+)*', line)
        if m.group(3) is not None:
            [child_programs.append(y.strip()) for y in m.group(3).split(',')]
        program = [m.group(1), m.group(2), m.group(3)]
        programs.append(program)

    for program in programs:
        if program[0] not in child_programs:
            print("Puzzle 1: " + str(program[0]))

    for program in programs:
        if hasChilds(program):
            program[2] = [y.strip() for y in program[2].split(',')]
            temp = [[int(i[1]) for i in programs if i[0] == j]
                    for j in program[2]]
            print(temp)
            # for child in program[2]:
            #     temp = [int(i[1]) for i in programs if i[0] == child]
            #     print(temp)
                # for i in programs:
                #     if child == i[0]:
                #         program[2] = [i[1]]
                #         print(child, i[1])
