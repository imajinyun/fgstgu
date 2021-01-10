import os
from _collections import deque


def search(lines: str, pattern: str, history=5)->None:
    dq = deque(maxlen=history)
    for line in lines:
        if pattern in line:
            yield line, dq


"""
deque([], maxlen=5)
java and python

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
deque([], maxlen=5)
python

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
deque([1, 2, 3], maxlen=3)
deque([2, 3, 4], maxlen=3)
"""
if __name__ == "__main__":
    basedir = os.path.dirname(__file__)
    filename = os.path.join(basedir, '', '0528.txt')
    with open(filename, 'r') as lines:
        for line, dq in search(lines, 'python', 5):
            print(dq)
            for element in dq:
                if element is not None:
                    print('Element: ' + element)
            print(line)
            print('+' * 80)

q = deque(maxlen=3)
q.append(1)
q.append(2)
q.append(3)
print(q)
q.append(4)
print(q)
