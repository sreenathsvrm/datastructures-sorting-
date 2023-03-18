package main

type maxheap struct{
	maxheap []int
}


func (m*maxheap) heapify(arr[]int){

	m.maxheap = append(m.maxheap, arr...)


}

func (m *maxheap)paraid(childid int) int {
	return (childid/2)-1
}

func (m *maxheap)leftid(paraid int) int{
	return (paraid*2)+1
}

func (m *maxheap)rightid(paraid int) int {
	return (paraid*2)+2
}

func (m *maxheap)shift(paraid int)  {
	endid:=len(m.maxheap)-1
	leftid:=m.leftid(paraid)

	for leftid<=endid{
     idshift:=leftid

	 rightid:=m.rightid(paraid)

	 if rightid<endid&&m.maxheap[rightid]>m.maxheap[leftid]{
		idshift=rightid
       if 

	 }
	}
}