package main

import "testing"

func TestSayHello(t *testing.T) {
	got := sayHello("vikas")
	want := "Hello vikas\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
