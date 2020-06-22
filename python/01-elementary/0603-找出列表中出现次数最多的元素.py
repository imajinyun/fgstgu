from collections import Counter

words = [
    'a', 'b', 'a', 'd', 'e', 'b', 'a', 'b', 'c', 'd', 'b', 'f',
    'c', 'd', 'e', 'f', 'a', 'b', 'e', 'f', 'a', 'b', 'c', 'e',
    'f', 'a', 'a', 'd', 'e', 'a', 'a', 'c', 'b', 'b', 'a', 'a',
]
counts = Counter(words)
top_three = counts.most_common(3)

# [('a', 11), ('b', 8), ('e', 5)]
print(top_three)
