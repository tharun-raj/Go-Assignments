
// -------------------------------------------------------------

//create struct
type Counter struct{
	count int
}

func (a Counter) add() {
	a.count++
	
}

func main(){
	var b = Counter{count: 0} //Counter{ }
	b = b.add()
	print(b)
}