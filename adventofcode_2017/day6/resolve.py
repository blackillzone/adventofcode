#!/usr/bin/env python
"""Day 05. Advent of Code."""

with open('input.txt', 'r') as x:
    banks = x.read().split(' ')
    banks = [int(a) for a in banks]
    ref = []

    while banks not in ref:
        ref.append(banks[:])
        m = max(banks)
        i = banks.index(m)
        banks[i] = 0
        while m:
            i = (i+1) % len(banks)
            banks[i] += 1
            m -= 1

    print('Puzzle 1: ' + str(len(ref)))
    print('Puzzle 2: ' + str(len(ref)-ref.index(banks)))
