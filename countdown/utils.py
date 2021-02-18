import random


def remove_random_element(l):
    """ Remove random element from a list and return it """
    choice = random.choice(l)
    l.remove(choice)
    return choice
