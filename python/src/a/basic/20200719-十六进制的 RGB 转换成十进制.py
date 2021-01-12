"""
颜色代码转换为 RGB
"""
color = 0xD200C1
r, g, b = (color >> 16) & 0xFF, (color >> 8) & 0xFF, (color & 0XFF)
print(r, g, b)


"""
RGB 转换为颜色代码
"""
r, g, b = 100, 56, 200
color = (r << 16) | (g << 8) | b
print(color, hex(color))
