"""
[seaborn: statistical data visualization](http://seaborn.pydata.org/index.html)
"""
import seaborn as sns
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from pylab import savefig

sns.set(color_codes=True)

# 直方图
x = np.random.normal(size=100)
sns.distplot(x, bins=20, kde=False, rug=True, label='Hisgogram w/o Density')
plt.xlabel('Value')
plt.ylabel('Frequency')
plt.title('Histogram of a Random Sample from a Normal Distribution')
plt.legend()

# 带有回归线直线的散点图与单变量直方图
mean, cov = [5, 10], [(1, .5), (.5, 1)]
data = np.random.multivariate_normal(mean, cov, 200)
df = pd.DataFrame(data, columns=['x', 'y'])
sns.jointplot(x='x', y='y', data=df, kind='reg').set_axis_labels('x', 'y')
plt.suptitle('Joint Plot of Two Variables with Bivariate and Univariate Graphs')

# 成对变量之间的落散点图与单变量直方图
iris = sns.load_dataset('iris')
sns.pairplot(iris)

# 按照某几个变量生成的箱线图
tips = sns.load_dataset('tips')
sns.factorplot(x='time', y='total_bill', hue='smoker',
               col='day', data=tips, kind='box', size=4, aspect=.5)

# 带有 bootstrap 置信区间的线性回归模型
sns.lmplot(x='total_bill', y='tip', data=tips)

# 带有 bootstrap 置信区间的逻辑斯蒂回归模型
tips['big_tip'] = tuple((tips.tip / tips.total_bill) > .15)
sns.lmplot(x='total_bill', y='big_tip', data=tips).set_axis_labels(
    'Total Bill', 'Big Tip')
plt.title('Logistic Regression of Big Tip vs. Total Bill')
plt.show()
