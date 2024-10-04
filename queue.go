package main

type Directory struct {
	dir  map[string]interface{}
	path string
}

type Queue []Directory

func (que *Queue) push(data Directory) {
	*que = append(*que, data)
}

func (que *Queue) pop() Directory {
	data := (*que)[0]
	*que = (*que)[1:]
	return data
}

func (que *Queue) isEmpty() bool {
	if len(*que) == 0 {
		return true
	} else {
		return false
	}
}
