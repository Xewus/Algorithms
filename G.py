# G. Возможные друзья (20 баллов)
class User():
    def __init__(self, id: int) -> None:
        self.id = id
        self.friends: set['User'] = set()
        self.recs: dict[int, int] = dict()
    
    def add_friend(self, friend: 'User'):
        self.friends.add(friend)
        friend.friends.add(self)

    def _add_recs(self):
        if not self.friends:
            return

        for friend in self.friends:
            for rec in friend.friends:
                if not(rec is friend or rec is self):
                    if self.recs.get(rec.id) is None:
                        self.recs[rec.id] = 1
                    else:
                        self.recs[rec.id] += 1

        self.recs.pop(self.id, None)
        for friend in self.friends:
            self.recs.pop(friend.id, None)

    def recommendation(self) -> str:
        self._add_recs()
        if len(self.recs) == 0:
            return '0'

        max_rate = 0
        ids = []

        for rec, rate in self.recs.items():
            if rate < max_rate:
                continue
            if rate > max_rate:
                max_rate = rate
                ids = []
            ids.append(rec)

        ids.sort()

        return ' '.join([str(i) for i in ids])
        

def sol() -> int:
    # количество пользователей и количество пар друзей
    users, pairs = map(int, input().split())
    friendships = {i: User(i) for i in range(1, users + 1)}

    for _ in range(pairs):
        user, friend = map(int, input().split())
        friend = friendships[friend]
        friendships[user].add_friend(friend)
    
    for user in friendships.values():
        print(user.recommendation())


if __name__ == '__main__':
    sol()
