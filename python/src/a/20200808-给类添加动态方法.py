"""
@see https://medium.com/@mgarod/dynamically-add-a-method-to-a-class-in-python-c49204b85bd6
"""
from functools import wraps


class Test:
    pass


def add_method(cls):
    def decorator(func):
        @wraps(func)
        def wrapper(self, *args, **kwargs):
            return func(*args, **kwargs)
        setattr(cls, func.__name__, wrapper)
        return func
    return decorator


test = Test()
try:
    test.foo()
except AttributeError as ae:
    # 'Test' object has no attribute 'foo'
    print(f'Exception caught: {ae}')

try:
    test.bar('The quick brown for jumped over the lazy dog.')
except AttributeError as ae:
    # 'Test' object has no attribute 'bar'
    print(f'Exception caught: {ae}')


@add_method(Test)
def foo():
    print('Hello World')


@add_method(Test)
def bar(s: str):
    print(f'Message: {s}')


test.foo()
test.bar('The quick brown for jumped over the lazy dog.')
print(test.foo)
print(test.bar)

# foo and bar are still usable as functions
foo()
bar('The quick brown for jumped over the lazy dog.')
print(foo)
print(bar)

"""
Output:
Exception caught: 'Test' object has no attribute 'foo'
Exception caught: 'Test' object has no attribute 'bar'
Hello World
Message: The quick brown for jumped over the lazy dog.
<bound method foo of <__main__.Test object at 0x107254610>>
<bound method bar of <__main__.Test object at 0x107254610>>
Hello World
Message: The quick brown for jumped over the lazy dog.
<function foo at 0x1071ee9e0>
<function bar at 0x1072090e0>
"""
