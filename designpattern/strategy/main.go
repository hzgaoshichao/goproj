package main

func main() {

	cache := initCache(lfuTyp)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	cache.setEvictionAlgo(lruTyp)

	cache.add("d", "4")

	cache.setEvictionAlgo(fifoTyp)

	cache.add("e", "5")

}
