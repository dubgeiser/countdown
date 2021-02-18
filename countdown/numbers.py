"""
Countdown (Cijfers&Letters)
Number game
https://www.datagenetics.com/blog/august32014/index.html
"""
import random

from countdown.utils import remove_random_element


def generate(big_count):
    """ Generate the Numbers problem
    return a tuple where the first element is the target and the second a list
    of numbers to make that target.
    """
    big_numbers = [25, 50, 75, 100]
    small_numbers = [1 , 1 , 2 , 2 , 3 , 3 , 4 , 4 , 5 , 5 , 6 , 6 , 7 , 7 , 8 , 8 , 9 , 9 , 10 , 10]
    total_amount_of_numbers = 6
    small_count = total_amount_of_numbers - big_count
    if big_count + small_count != total_amount_of_numbers:
        raise Exception(f"total amount of numbers must be {total_amount_of_numbers}")
    selection = []
    for i in range(0, big_count):
        selection.append(remove_random_element(big_numbers))
    for i in range(0, small_count):
        selection.append(remove_random_element(small_numbers))
    target = random.choice(range(101, 1000))
    return (target, selection)
