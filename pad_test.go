package main

import "testing"

func TestZero(t *testing.T) {
	res := PadWholeNumbers("test 0", 2)
	if res != "test 00" {
		t.Fail()
	}
}

func TestNine(t *testing.T) {
	res := PadWholeNumbers("test 9", 2)
	if res != "test 09" {
		t.Fail()
	}
}

func TestNumbersOnly(t *testing.T) {
	res := PadWholeNumbers("123", 4)
	if res != "0123" {
		t.Fail()
	}
}

func TestLettersOnly(t *testing.T) {
	res := PadWholeNumbers("abc", 4)
	if res != "abc" {
		t.Fail()
	}
}

func TestSingleLetter(t *testing.T) {
	res := PadWholeNumbers("a", 4)
	if res != "a" {
		t.Fail()
	}
}

func TestSingleNumber(t *testing.T) {
	res := PadWholeNumbers("1", 4)
	if res != "0001" {
		t.Fail()
	}
}

func TestPadLengthSmallerThanNumber(t *testing.T) {
	res := PadWholeNumbers("test129", 2)
	if res != "test129" {
		t.Fail()
	}
}

func TestPadLengthLargerThanNumber(t *testing.T) {
	res := PadWholeNumbers("test129", 5)
	if res != "test00129" {
		t.Fail()
	}
}

func TestNumberAtTheBegining(t *testing.T) {
	res := PadWholeNumbers("1test", 2)
	if res != "01test" {
		t.Fail()
	}
}

func TestNumberAtTheEnd(t *testing.T) {
	res := PadWholeNumbers("test1", 2)
	if res != "test01" {
		t.Fail()
	}
}
