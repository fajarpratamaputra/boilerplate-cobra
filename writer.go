package main

type redis struct {
}

type database interface {
	writeLineup(lineup *map[int]int)
}

func newDatabase() database {
	return &redis{}
}

func (r *redis) writeLineup(lineup *map[int]int) {
	// write to redis
}
