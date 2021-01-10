from ggplot import *

print(mtcars.head())

plt1 = ggplot(aex(x='mpg'), data=mtcars) + \
    geom_histogram(fill='darkblue', binwidth=2) + xlim(10, 35) + \
    ylim(0, 10) + xlab('MPG') + ylab('Frequency') + \
    ggtitle('Histogram of MPG') + theme_matplotlib()
print(plt1)
print(meat.head())

plt2 = ggplot(aes(x='date', y='beef'), data=meat) + \
    geom_line(color='purple', size=1.5, alpha=0.75) + \
    stat_smooth(color='blue', size=2.0, span=0.15) + \
    xlab('年') + ylab('Head') + \
    ggtitle('Beef') + theme_seaborn()
print(plt2)
print(diamonds.head())

plt3 = ggplot(diamonds, aes(x='carat', y='price', colour='cut')) + \
    geom_point(alpha=0.5) + \
    scale_color_gradient(low='#05D9F6', high='#5011D1') + \
    xlim(0, 6) + ylim(0, 20000) + \
    xlab('克拉') + ylab('价格') + \
    ggtitle('钻石') + \
    + theme_gray()
print(plt3)
