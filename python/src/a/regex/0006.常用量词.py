# coding=utf-8

import re


def solution() -> None:
    """
    * == {0,}  -> 可能出现，也可能不出现，出现次数没有上限；
    + == {1,}  -> 至少出现 1 次，出现次数没有上限；
    ? == {0,1} -> 至多出现 1 次，也可能不出现；
    """
    regex = r'^<[^>]+>$'
    open_tag = r'^<[^/][^>]*>$'
    close_tag = r'^</[^>]+>$'
    self_closing_tag = r'^<[^>]+/>$'
    items = [
        re.search(r'^travell?er$', 'traveller') is not None,  # True
        re.search(r'^travell?er$', 'traveler') is not None,  # True
        re.search(regex, '<body>') is not None,  # True
        re.search(regex, '</table>') is not None,  # True
        re.search(regex, '<>') is not None,  # False
        re.search(r'^\"[^\"]*\"$', "\"some\"") is not None,  # True
        re.search(r'^\"[^\"]*\"$', "\"\"") is not None,  # True

        # 匹配 Open Tag
        re.search(open_tag, '<b>') is not None,  # True
        re.search(open_tag, '<article>') is not None,  # True
        re.search(open_tag, '<>') is not None,  # False

        # 匹配 Close Tag
        re.search(close_tag, '</b>') is not None,  # True
        re.search(close_tag, '</section>') is not None,  # True
        re.search(close_tag, '</>') is not None,  # False

        # 匹配 Self-closing Tag
        re.search(self_closing_tag, '<img />') is not None,  # True
        re.search(self_closing_tag, '<br/>') is not None,  # True
        re.search(self_closing_tag, '</>') is not None,  # False

        # 匹配所有 Tag
        re.search(regex, '<video>') is not None,  # True
        re.search(regex, '</video>') is not None,  # True
        re.search(regex, '<br />') is not None,  # True
        re.search(regex, '<img/>') is not None,  # True
        re.search(regex, '<p>') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
