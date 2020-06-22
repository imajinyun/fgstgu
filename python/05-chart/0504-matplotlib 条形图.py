"""
> pip3 install matplotlib
"""
import numpy as np
import matplotlib.pyplot as plt

labels = ['ABC', 'DEF', 'GHI', 'JKL', 'MNO']
amounts = [180, 90, 100, 120, 230]
width = 0.35

fig, ax = plt.subplots()
ax.bar(labels, amounts, width)
ax.set_xlabel('Customer Name')
ax.set_ylabel('Sale Amount')
ax.set_title('Number of Values in Bin')
ax.legend()
plt.show()
