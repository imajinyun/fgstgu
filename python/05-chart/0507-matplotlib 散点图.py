import numpy as np
import matplotlib.pyplot as plt

plt.style.use('ggplot')
x = np.arange(start=1., stop=15., step=1.)
linear = x + 5. * np.random.randn(14)
quadratic = x**2 + 10. * np.random.randn(14)
fn_linear = np.poly1d(np.polyfit(x, linear, deg=1))
fn_quadratic = np.poly1d(np.polyfit(x, quadratic, deg=2))
fig = plt.figure()
ax = fig.add_subplot(1, 1, 1)
ax.plot(x, linear, 'bo', x, quadratic, 'go', x, fn_linear(
    x), 'b-', x, fn_quadratic(x), 'g-', linewidth=2.)
ax.xaxis.set_ticks_position('bottom')
ax.yaxis.set_ticks_position('left')
ax.set_title('散点图回归线')
plt.xlabel('x')
plt.ylabel('f(x)')
plt.xlim((min(x) - 1., max(x) + 1.))
plt.ylim((min(quadratic) - 10., max(quadratic) + 10.))
plt.show()
