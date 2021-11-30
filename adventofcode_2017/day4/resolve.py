#!/usr/bin/env python
"""Day 04. Advent of Code."""


def passphrase_duplicate(passphrase):
    """Function to compare set lenght vs list lenght."""
    if len(passphrase.split()) != len(set(passphrase.split())):
        return False
    return True


def passphrase_anagram(passphrase):
    """Function to test anagrams in passphrase."""
    import itertools
    for a, b in itertools.combinations(passphrase.split(), 2):
        if sorted(a) == sorted(b):
            return False
    return True


with open('input.txt', 'r') as x:
    lines = x.read().splitlines()
    total_pz1 = [passphrase_duplicate(i) for i in lines]
    total_pz2 = [passphrase_anagram(i) for i in lines]
    print('Puzzle 1: ' + str(sum(total_pz1)))
    print('Puzzle 2: ' + str(sum(total_pz2)))
