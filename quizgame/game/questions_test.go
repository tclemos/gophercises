package game

import "testing"

func TestQuestionsAddFirstItem(t *testing.T) {

	// Arrange
	qq := newQuestions()

	desc := "Question description"
	answer := "Right Answer"
	q := question{
		description: desc,
		answer:      answer,
	}

	// Act
	qq.add(q)

	// Assert
	if qq.head != qq.tail {
		t.Error("head is different from tail")
	}

	if qq.head.question.description != desc {
		t.Error("head has a different description")
	}

	if qq.head.question.answer != answer {
		t.Error("head has a different answer")
	}

	if qq.head.next != nil || qq.tail.next != nil {
		t.Error("head or tail has its next different from nil")
	}

	if qq.length != 1 {
		t.Errorf("invalid lenth, expected 1, found: %d", qq.length)
	}
}

func TestQuestionsAddSecondItem(t *testing.T) {

	// Arrange
	qq := newQuestions()

	q1, q2 := question{}, question{}

	// Act
	qq.add(q1)
	qq.add(q2)

	// Assert
	if qq.head == qq.tail {
		t.Error("head should be different from tail")
	}

	if qq.head.next == nil {
		t.Error("head is missing the next refference")
	}

	if qq.tail.next != nil {
		t.Error("tail next should be nil")
	}

	if qq.length != 2 {
		t.Errorf("invalid lenth, expected 2, found: %d", qq.length)
	}
}

func TestQuestionsAddNItem(t *testing.T) {

	// Arrange
	qq := newQuestions()

	q1, q2, q3, q4 := question{}, question{}, question{}, question{}

	// Act
	qq.add(q1)
	qq.add(q2)
	qq.add(q3)
	qq.add(q4)

	// Assert
	if qq.head == qq.tail {
		t.Error("head should be different from tail")
	}

	if qq.head.next == nil || qq.head.next.next == nil || qq.head.next.next.next == nil {
		t.Error("one of the nodes is missing the next refference")
	}

	if qq.tail.next != nil {
		t.Error("tail next should be nil")
	}

	if qq.length != 4 {
		t.Errorf("invalid lenth, expected 4, found: %d", qq.length)
	}

}
