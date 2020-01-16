package game

// questions is a linked list to hold questions in a sequence
type questions struct {
	head   *questionNode
	tail   *questionNode
	length int
}

// questionNode represents a node of the Questions linked list
type questionNode struct {
	// question represents the question this node holds
	question question

	// next is the next Question in the list
	next *questionNode
}

// newQuestions creates a new instance of Questions
func newQuestions() *questions {
	return &questions{}
}

// add inserts a new question to the end of the list
func (qq *questions) add(q question) {

	newNode := &questionNode{q, nil}

	if qq.head == nil {
		qq.head = newNode
	}

	if qq.tail == nil {
		qq.tail = newNode
	} else {
		qq.tail.next = newNode
		qq.tail = newNode
	}

	qq.length++
}
