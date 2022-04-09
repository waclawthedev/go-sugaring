package sugar

import "testing"

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

	if result[0] != "one" || result[1] != "two" {
		t.Fail()
	}

	result = Keys(make(map[string]int))
	if result != nil {
		t.Fail()
	}
}

func TestValues(t *testing.T) {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	result := Values(m)

	if result[0] != 1 || result[1] != 2 {
		t.Fail()
	}

	result = Values(make(map[string]int))
	if result != nil {
		t.Fail()
	}
}
