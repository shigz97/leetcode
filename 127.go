package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	leng := len(beginWord)
	var key string
	dic := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < leng; i++ {
			key = w[:i] + "*" + w[i+1:]
			dic[key] = append(dic[key], w)
		}
	}
	kmap := make(map[string][]string)
	for _, w := range wordList {
		tmp := []string{}
		for i := 0; i < leng; i++ {
			key = w[:i] + "*" + w[i+1:]
			for _, k := range dic[key] {
				if k != w {
					tmp = append(tmp, k)
				}
			}
		}
		kmap[w] = tmp
	}
	tmp := []string{}
	for i := 0; i < leng; i++ {
		key = beginWord[:i] + "*" + beginWord[i+1:]
		for _, k := range dic[key] {
			if k != beginWord {
				tmp = append(tmp, k)
			}
		}
	}
	kmap[beginWord] = tmp
	isVisted := make(map[string]bool)
	queue := []string{}
	front, last, num := 0, 1, 1
	queue = append(queue, beginWord)
	for front < len(queue) {
		tmp := queue[front]
		for _, w := range kmap[tmp] {
			if _, ok := isVisted[w]; !ok {
				if w == endWord {
					return num + 1
				}
				isVisted[w] = true
				queue = append(queue, w)
			}
		}
		front++
		if front == last {
			num++
			last = len(queue)
		}
		//fmt.Println(queue,front,last,num)
	}
	return 0
}
