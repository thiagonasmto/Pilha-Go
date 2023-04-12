package pilha

type pilha struct{
	size int
	head *No
}

type No struct{
	value int
	next *No
}

func (pilha *pilha) Init(){
	pilha.head = nil
	pilha.size = 0
}

func (pilha *pilha) push(value int){
	if pilha.head == nil {
		pilha.head = &No{value: value, next: nil}
	}else{
		newNode := &No{value: value, next: pilha.head}
		pilha.head = newNode
	}
	pilha.size++
}

func (pilha *pilha) pop(){
	if pilha.size > 0{
		pilha.head = pilha.head.next
		pilha.size--
	}
}

func (pilha *pilha) topo() int{
	return pilha.head.value
}

func (pilha *pilha) size() int{
	return pilha.size
}

func (pilha *pilha) vazia() bool{
	if pilha.head == nil{
		return true
	}else{
		return false
	}
}