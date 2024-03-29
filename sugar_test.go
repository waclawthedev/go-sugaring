package sugar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if !Contains(arr, "a") {
		t.Fail()
	}
	if Contains(arr, "") {
		t.Fail()
	}
}

func TestIndexesOf(t *testing.T) {
	arr := []string{"a", "b", "c", "b"}
	indexes := IndexesOf(arr, "b")
	if indexes[0] != 1 || indexes[1] != 3 || len(indexes) != 2 {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	result := Map([]int{1, 2, 3, 4}, func(k int) int {
		return k - 1
	})
	if result[0] != 0 || result[1] != 1 || result[2] != 2 || result[3] != 3 {
		t.Fail()
	}

	result = Map([]int{}, func(k int) int {
		return k - 1
	})

	if result != nil {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4}, func(k int) bool {
		if k > 2 {
			return false
		}
		return true
	})

	if result[0] != 1 || result[1] != 2 || len([]int(result)) != 2 {
		t.Fail()
	}

	result = Filter([]int{}, func(k int) bool {
		if k > 2 {
			return false
		}
		return true
	})

	if result != nil {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	result := Reduce([]int{1, 2, 3, 4}, func(a int, b int) int {
		return a + b
	})

	if result != 10 {
		t.Fail()
	}

	result = Reduce([]int{}, func(a int, b int) int {
		return a + b
	})

	if result != 0 {
		t.Fail()
	}

	result = Reduce([]int{1}, func(a int, b int) int {
		return a + b
	})

	if result != 1 {
		t.Fail()
	}
}

func TestFlatten(t *testing.T) {
	result := Flatten([][]int{{1, 2}, {3, 4}})

	if result[0] != 1 || result[1] != 2 || result[2] != 3 || result[3] != 4 {
		t.Fail()
	}

	result = Flatten([][]int{})
	if result != nil {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	result := Min([]int{1, 2, 0, 3, 4})

	if result != 0 {
		t.Fail()
	}

	result = Min([]int{})

	if result != 0 {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	result := Max([]int{1, 2, 3, 4})

	if result != 4 {
		t.Fail()
	}

	result = Max([]int{})

	if result != 0 {
		t.Fail()
	}
}

func TestKeys(t *testing.T) {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	result := Keys(m)

	if result[0] != "one" || result[1] != "two" || len(m) != len(result) {
		t.Fail()
	}

	result = Keys(make(map[string]int))
	if len(result) != 0 {
		t.Fail()
	}
}

func BenchmarkKeys(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < 100000; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Keys(m)
	}
}

func TestValues(t *testing.T) {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	result := Values(m)

	if result[0] != 1 || result[1] != 2 || len(m) != len(result) {
		t.Fail()
	}

	result = Values(make(map[string]int))
	if len(result) != 0 {
		t.Fail()
	}
}

func BenchmarkValues(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < 100000; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Keys(m)
	}
}

func Test_SplitByN(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Basic case",
			args: args{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
		},
		{
			name: "n is bigger than slice length",
			args: args{
				s: []int{1, 2, 3},
				n: 5,
			},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "n is same as slice length",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 5,
			},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "empty slice",
			args: args{
				s: []int{},
				n: 3,
			},
			want: nil,
		},
		{
			name: "n is zero",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 0,
			},
			want: nil,
		},
		{
			name: "n is negative",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: -1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByN(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByN() = %v, want %v", got, tt.want)
			}
		})
	}
}
