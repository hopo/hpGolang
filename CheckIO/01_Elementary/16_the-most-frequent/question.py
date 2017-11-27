# https://py.checkio.org/mission/the-most-frequent/

"""
You have a sequence of strings, and you’d like to determine the most frequently occurring string in the sequence.
Input: a list of strings.
Output: a string.
"""

def most_frequent(data):
    """
        determines the most frequently occurring string in the sequence.
    """
    # your code here
    return None

if __name__ == '__main__':
    ex1 = most_frequent([
        'a', 'b', 'c',
        'a', 'b',
        'a'
    ]) # 'a'
    print(ex1)
    ex2 = most_frequent(['a', 'a', 'bi', 'bi', 'bi']) # 'bi'
    print(ex2)
