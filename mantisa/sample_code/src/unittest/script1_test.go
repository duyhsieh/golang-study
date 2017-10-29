package unittest
import ("testing")

func Test_foo(t* testing.T) {
    A:=1
    if A != 1 {
        t.Error("Expected:", 1, " Got:", A)
    }

    if foo_subtest1(1,2) != 3 {
        t.Error("Expected:", 3, " Got:", foo_subtest1(1,2) )
    }
    if foo_subtest2(1,2) != 3 {
        t.Error("Expected:", 3, " Got:", foo_subtest2(1,2) )
    }
}

func Test_bar(t* testing.T) {
    t.Error("err msg1", "err msg2", "err msg 3")
}

func foo_subtest1(a1 int, a2 int) int {
    return a1+a2
}
func foo_subtest2(a1 int, a2 int) int {
    return a1-a2
}
