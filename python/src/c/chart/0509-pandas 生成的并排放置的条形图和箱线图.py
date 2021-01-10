import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

plt.style.use('ggplot')
fig, axes = plt.subplots(nrows=1, ncols=2)
ax1, ax2 = axes.ravel()

indexes = ['客户1', '客户2', '客户3', '客户4', '客户5']
metrics = ['Metric1', 'Metric2', 'Metric3']
data = pd.DataFrame(np.random.rand(5, 3), index=indexes,
                    columns=pd.Index(metrics, name='Metrics'))

data.plot(kind='bar', ax=ax1, alpha=0.75, title='Bar Plot')
plt.setp(ax1.get_xticklabels(), rotation=45, fontsize=10)
plt.setp(ax1.get_yticklabels(), rotation=0, fontsize=10)
ax1.set_xlabel('客户')
ax2.set_ylabel('值')
ax1.xaxis.set_ticks_position('bottom')
ax1.yaxis.set_ticks_position('left')
colors = dict(boxes='DarkBlue', whiskers='Gray', medians='Red', caps='Black')

data.plot(kind='box', color=colors, sym='r.', ax=ax2, title='Box Plot')
plt.setp(ax2.get_xticklabels(), rotation=45, fontsize=10)
plt.setp(ax2.get_yticklabels(), rotation=0, fontsize=10)
ax2.set_xlabel('Metric')
ax2.set_ylabel('Value')
ax2.xaxis.set_ticks_position('bottom')
ax2.yaxis.set_ticks_position('left')
plt.show()
