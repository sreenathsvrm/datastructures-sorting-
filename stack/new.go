package main

func main(){

}

func bubblesort(ar[]int) []int {
	

	for i:=0;i<len(ar);i++{
		for j:=0;j<len(ar)-1;j++{
			if ar[j]>ar[j+1]{
				ar[j],ar[j+1]=ar[j+1],ar[j]
			}
		}
	}
	return ar
}

func selection(ar[]int)[]int  {

	for i:=0;i<len(ar)-1;i++{
		for j:=i+1;j<len(ar);j++{
			if ar[i]>ar[j]{
				ar[i],ar[j] = ar[j],ar[i]
			}
		}
	}
	return ar
	
}

func  insertion(ar[]int)[]int{
	for idx,val:=range ar{

		for idx>0 && val>ar[idx-1]{
			ar[idx]=ar[idx-1]
			idx--
		}
		ar[idx]=val
	}
}