package bloomrepo

type KeyFiter interface {
	Add([]byte)
	Contains([]byte) bool
}
