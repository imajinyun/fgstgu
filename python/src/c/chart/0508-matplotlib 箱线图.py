import numpy as np
import matplotlib.pyplot as plt

plt.style.use('ggplot')
N = 500
normal = np.random.normal(loc=0.0, scale=1.0, size=N)
lognormal = np.random.lognormal(mean=0.0, sigma=1.0, size=N)
index = np.random.randint(0, N)
normal_sample = normal[index]
lognormal_sample = lognormal[index]
data = [normal, normal_sample, lognormal, lognormal_sample]

fig = plt.figure()
ax = fig.add_subplot(1, 1, 1)
labels = ['normal', 'normal_sample', 'lognormal', 'lognormal_sample']
ax.boxplot(data, notch=False, sym='.', vert=True,
           whis=1.5, showmeans=True, labels=labels)
ax.xaxis.set_ticks_position('bottom')
ax.yaxis.set_ticks_position('left')
ax.set_title('正态分布和对数正态分布数据的箱线图')
ax.set_xlabel('分布')
ax.set_ylabel('值')
plt.show()
