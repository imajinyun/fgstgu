from operator import attrgetter


class User:
    def __init__(self, userId):
        self.userId = userId

    def __repr__(self):
        return 'User({})'.format(self.userId)


users = [User(9), User(7), User(8)]

# [User(9), User(7), User(8)]
[User(7), User(8), User(9)]
print(users)

# [User(7), User(8), User(9)]
print(sorted(users, key=lambda u: u.userId))

# [User(7), User(8), User(9)]
print(sorted(users, key=attrgetter('userId')))
