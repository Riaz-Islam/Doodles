package main

/*

type dsu struct {
	par  []int
	size []int
}

func dsuInit(n int) *dsu {
	d := &dsu{
		par:  make([]int, n+1),
		size: make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		d.par[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *dsu) Find(u int) int {
	if d.par[u] == u {
		return u
	}
	d.par[u] = d.Find(d.par[u])
	return d.par[u]
}

func (d *dsu) Join(u, v int) bool {
	u = d.Find(u)
	v = d.Find(v)
	if u == v {
		return false
	}
	d.par[v] = u
	d.size[u] += d.size[v]
	return true
}

func (d *dsu) InSameSet(u, v int) bool {
	u = d.Find(u)
	v = d.Find(v)
	return u == v
}

*/
