package main

type userIterator struct {
	index int
	user  []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.user) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.user[u.index]
		u.index++
		return user
	}
	return nil
}
