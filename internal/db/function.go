package db

type Function func(session *Session) (int64, error)
