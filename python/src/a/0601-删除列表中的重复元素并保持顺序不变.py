def dedupe(items):
    marked = set()
    for item in items:
        if item not in marked:
            yield item
            marked.add(item)


if __name__ == '__main__':
    a = [5, 5, 2, 1, 9, 1, 5, 10]

    # [5, 5, 2, 1, 9, 1, 5, 10]
    print(a)

    # [5, 2, 1, 9, 10]
    print(list(dedupe(a)))
