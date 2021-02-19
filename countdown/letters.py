""" Letters
Could not immediately find the distribution/frequency for the letters in Dutch,
but it seems that Countdown uses the same distribution as Scrabble:

For Dutch:
E ×18, N ×10, A ×6, O ×6, I ×4
D ×5, R ×5, S ×5, T ×5
G ×3, K ×3, L ×3, M ×3, B ×2, P ×2
U ×3, F ×2, H ×2, J ×2, V ×2, Z ×2
C ×2, W ×2
X ×1, Y ×1
Q ×1
"""
from countdown.utils import remove_random_element


# These need to be module level, since we're gonna pick from them in multiple
# HTTP requests.
# TODO This needs to be reset every game!!!
vowels = [i for i in 'aaaaaaeeeeeeeeeeeeeeeeeeiiiioooooouuuy']
consonants = [i for i in 'bbccdddddffggghhjjkkklllmmmnnnnnnnnnnppqrrrrrssssstttttvvwwxzz']
total_amount_of_letters = 9


def pick(vowel_or_consonant, game):
    if vowel_or_consonant.lower() not in ['vowel', 'consonant']:
        raise Exception('Only vowels or consonants can be picked')
    if len(game['selection']) == 9:
        return
    game['selection'].append(remove_random_element(game[vowel_or_consonant]))


def new_game():
    game = {
            'vowel' : vowels,
            'consonant' : consonants,
            'selection' : []
            }
    return game
