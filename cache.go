package cache

type Cache struct {
	data map[string]int
}

func (c Cache) New() Cache {
	c.data = make(map[string]int)
	return c
}

func (c Cache) Get(key string) int {
	return c.data[key]
}

func (c Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c Cache) Delete(key string) {
	delete(c.data, key)
}

func New() Cache {
	cacheInst := Cache{}
	t := cacheInst.New()
	return t
}
