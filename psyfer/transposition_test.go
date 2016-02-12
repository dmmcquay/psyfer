package psyfer

import "testing"

func TestTransposeRailFence(t *testing.T) {
	input := "helloworld"
	expected := "hloolelwrd"
	actual := TransposeRailFence(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeRailFence:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = "123456"
	expected = "135246"
	actual = TransposeRailFence(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeRailFence:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = "1"
	expected = "1"
	actual = TransposeRailFence(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeRailFence:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = ""
	expected = ""
	actual = TransposeRailFence(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeRailFence:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
}

func TestTransposeSplit(t *testing.T) {
	input := "helloworld"
	expected := "hweolrllod"
	actual := TransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = "1234567"
	expected = "1425367"
	actual = TransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = "123"
	expected = "123"
	actual = TransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = "1"
	expected = "1"
	actual = TransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	input = ""
	expected = ""
	actual = TransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed TransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
}

func TestTransposeRandom(t *testing.T) {
	input := "123456789abcdefghijklmnopqrstuvwxyz"
	actual := TransposeSplit(input)
	if input == actual {
		t.Errorf(
			"failed TransposeRandom:\n\texpected: something random\n\t  actual: % q",
			input,
			actual,
		)
	}
}

func TestDeTransposeSplit(t *testing.T) {
	expected := "helloworld"
	input := "hweolrllod"
	actual := DeTransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed DeTransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	expected = "12"
	input = "12"
	actual = DeTransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed DeTransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	expected = "12345"
	input = "13245"
	actual = DeTransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed DeTransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	expected = "123"
	input = "123"
	actual = DeTransposeSplit(input)
	if expected != actual {
		t.Errorf(
			"failed DeTransposeSplit:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
}
