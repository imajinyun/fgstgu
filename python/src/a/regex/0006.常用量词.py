# coding=utf-8

import re


def solution() -> None:
    """
    * == {0,}  -> 可能出现，也可能不出现，出现次数没有上限；
    + == {1,}  -> 至少出现 1 次，出现次数没有上限；
    ? == {0,1} -> 至多出现 1 次，也可能不出现；
    """
    items = [
        # re.search(r'^travell?er$', 'traveller') is not None,  # True
        # re.search(r'^travell?er$', 'traveler') is not None,  # True
        # re.search(r'^<[^>]+>$', '<body>') is not None,  # True
        # re.search(r'^<[^>]+>$', '</table>') is not None,  # True
        # re.search(r'^<[^>]+>$', '<>') is not None,  # False
        # re.search(r'^\"[^\"]*\"$', "\"some\"") is not None,  # True
        # re.search(r'^\"[^\"]*\"$', "\"\"") is not None,  # True

        # 匹配 Open Tag
        re.search(r'^<[^/][^>]*>$', '<b>') is not None,  # True
        re.search(r'^<[^/][^>]*>$', '<article>') is not None,  # True
        re.search(r'^<[^/][^>]*>$', '<>') is not None,  # False

        # 匹配 Close Tag
        re.search(r'^</[^>]+>$', '</b>') is not None,  # True
        re.search(r'^</[^>]+>$', '</section>') is not None,  # True
        re.search(r'^</[^>]+>$', '</>') is not None,  # False

        # 匹配 Self-closing Tag
        re.search(r'^<[^>]+/>$', '<img />') is not None,  # True
        re.search(r'^<[^>]+/>$', '<br/>') is not None,  # True
        re.search(r'^<[^>]+/>$', '</>') is not None,  # False

        # 匹配所有 Tag
        re.search(r'^<[^>]+>$', '<video>') is not None,  # True
        re.search(r'^<[^>]+>$', '</video>') is not None,  # True
        re.search(r'^<[^>]+>$', '<br />') is not None,  # True
        re.search(r'^<[^>]+>$', '<img/>') is not None,  # True
        re.search(r'^<[^>]+>$', '<p>') is not None,  # True
    ]
    for i in range(len(items)):
        print(items[i])


if __name__ == "__main__":
    solution()
