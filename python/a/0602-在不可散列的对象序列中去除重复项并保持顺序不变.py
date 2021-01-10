def dedupe(items, key=None):
    marked = set()
    for item in items:
        val = item if key is None else key(item)
        if val not in marked:
            yield item
            marked.add(val)


if __name__ == "__main__":
    a = [
        {'x': 2, 'y': 3},
        {'x': 1, 'y': 4},
        {'x': 2, 'y': 3},
        {'x': 2, 'y': 3},
        {'x': 9, 'y': 8},
    ]

    # [{'x': 2, 'y': 3}, {'x': 1, 'y': 4}, {'x': 2, 'y': 3}, {'x': 2, 'y': 3}, {'x': 9, 'y': 8}]
    print(a)

    # [{'x': 2, 'y': 3}, {'x': 1, 'y': 4}, {'x': 9, 'y': 8}
    print(list(dedupe(a, key=lambda a: (a['x'], a['y']))))
