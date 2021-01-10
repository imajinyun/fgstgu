"""
解决中文汉字方块问题：

1. 下载 [SimHei](http://www.fontpalace.com/font-download/SimHei/) 字体；
2. 复制字体文件到 matplotlib 字体目录下的 ttf 文件夹中；

```bash
# 查看 matplotlib 字体目录
import matplotlib
matplotlib.matplotlib_fname()
```

3. 修改配置文件；

```bash
$ vim matplotlibrc
font.family         : sans-serif
font.sans-serif     : SimHei, DejaVu Sans, Bitstream Vera Sans, Lucida Grande, Verdana, Geneva, Lucid, Arial, Helvetica, Avant Garde, sans-serif
axes.unicode_minus  : False
```

4. 删除相关缓存；

```bash
$ rm -rf ~/.matplotlib/*
```

5. 终端进入 Python 命令行解释器；

```bash
from matplotlib.font_manager import _rebuild
_rebuild()
```
"""
# -*- coding: utf-8 -*-
import numpy as np
import matplotlib.pyplot as plt
import matplotlib

# 随机生成（10000）服从正态分布的数据
data = np.random.randn(10000)

a, b, sigma = 100, 130, 15
x1 = a + sigma * np.random.randn(10000)
x2 = b + sigma * np.random.randn(10000)
fig = plt.figure()
fig.suptitle('Histograms', fontsize=14, fontweight='bold')
ax = fig.add_subplot(1, 1, 1)

"""
绘制直方图：

data: 必选参数，绘图数据
bins: 直方图的长条形数目，可选项，默认为 10
normed: 是否将得到的直方图向量归一化，可选项，默认为 0，代表不归一化，显示频数。normed=1，表示归一化，显示频率
facecolor: 长条形的颜色
edgecolor: 长条形边框的颜色
alpha: 透明度
"""
ax.hist(x1, bins=50, normed=0, facecolor="green",
        edgecolor="black", alpha=0.7)
ax.hist(x2, bins=50, normed=0, facecolor="blue",
        edgecolor="black", alpha=0.5)
ax.xaxis.set_ticks_position('bottom')
ax.yaxis.set_ticks_position('left')
plt.xlabel('区间')
plt.ylabel("频数/频率")
plt.title("频数/频率分布直方图")
plt.show()
