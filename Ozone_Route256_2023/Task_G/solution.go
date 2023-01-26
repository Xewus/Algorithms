// (G) Возможные друзья (20 баллов)
// Частичное решение: 10 баллов

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func readTwoInt() (x, y int) {
	fmt.Scan(&x, &y)
	return
}

type User struct {
	Id      int
	Friends map[*User]struct{}
	Recs    map[*User]int
}

func newUser(id int) *User {
	u := &User{Id: id}
	u.Friends = make(map[*User]struct{})
	u.Recs = make(map[*User]int)
	return u
}

func (u *User) addFriend(f *User) {
	if u == f {
		return
	}
	u.Friends[f] = struct{}{}
	f.Friends[u] = struct{}{}
}

func (u *User) addRecs() {
	if len(u.Friends) == 0 {
		return
	}
	for f := range u.Friends {
		for r := range f.Friends {
			if r != u && r != f {
				u.Recs[r]++
			}
		}
	}

	for f := range u.Friends {
		delete(u.Recs, f)
	}
	delete(u.Recs, u)
}

func (u *User) recommendation() string {
	u.addRecs()
	if len(u.Recs) == 0 {
		return "0 "
	}

	max_rate := 0
	ids := make([]int, 0, 5)
	for rec, rate := range u.Recs {
		if rate < max_rate {
			continue
		}
		if rate > max_rate {
			max_rate = rate
			ids = ids[:0]
		}
		ids = append(ids, rec.Id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	ans := make([]string, len(ids))
	for i, id := range ids {
		ans[i] = strconv.Itoa(id)
	}
	return strings.Join(ans, " ")
}

func main() {
	var u, f, i int
	us, ps := readTwoInt()
	friendship := make(map[int]*User, us+us/2)

	for i = 1; i <= us; i++ {
		friendship[i] = newUser(i)
	}

	for i = 1; i <= ps; i++ {
		u, f = readTwoInt()
		friendship[u].addFriend(friendship[f])
	}
	for i = 1; i <= us; i++ {
		fmt.Println(friendship[i].recommendation())
	}
}
