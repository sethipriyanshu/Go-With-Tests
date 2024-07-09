package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
func BenchmarkRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
