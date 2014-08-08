// START1 OMIT
func TestFoo(t *testing.T) {
	if actual != expected {
		t.Errorf("Expect %v but %v", expected, actual)
	}
}
// END1 OMIT

// START2 OMIT
func AverageVolume(symbol string, day time.Time) float64 {
	volumes, err := getVolumes(symbol, day) // HL
	if err != nil { return 0 }

	total := float64(0)
	for _, volume := range volumes { total += volume }
	return total / float64(len(volumes))
}

func getVolumes(symbol string, day time.Time) ([]float64, error) {
	url := "http://stock.com/volume/" + symbol + ".json?day=" + day.Format("2006-01-02")
	resp, err := http.Get(url) // HL
	if err != nil { return nil, err }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	....
}
// END2 OMIT

// START3 OMIT
var getVolumes = func(symbol string, day time.Time) ([]float64, error) { // HL
	url := "http://stock.com/volume/" + symbol + ".json?day=" + day.Format("2006-01-02")
	resp, err := http.Get(url)
	if err != nil { return nil, err }
	....
}
// END3 OMIT

// START4 OMIT
func TestProperlyGetsAllTheVolumes(t *testing.T) {
	var called bool
	now := time.Now()
	getVolumes = func(symbol string, day time.Time) ([]float64, error) { // HL
		if symbol != "spice" { t.Errorf("Expected %v to equal spice", symbol)}
		if day != now { t.Errorf("Expected %v to equal spice", symbol)}
		called = true
		return nil, nil
	}
	AverageVolume("spice", now)
	if called == false {
		t.Error("GetVolumes should have been called")
	}
}
// END4 OMIT

// START5 OMIT
func TestReturnsZeroOnError(t *testing.T) {
	getVolumes = func(symbol string, day time.Time) ([]float64, error) { // HL
		return nil, errors.Error("some error")
	}
	actual := AverageVolume("spice", time.Now())
	if actual != 0 {
		t.Errorf("Expected %f to equal 0", actual)
	}
}

// END5 OMIT

// START6 OMIT
func TestReturnsTheAverage(t *testing.T) {
	getVolumes = func(symbol string, day time.Time) ([]float64, error) { // HL
		return []float64{3,4,6,7,9}, nil
	}
	actual := AverageVolume("spice", time.Now())
	if actual != 5.8 {
		t.Errorf("Expected %f to equal 5.8", actual)
	}
}
// END6 OMIT


// START7 OMIT
package fib

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}
// END7 OMIT

// START8 OMIT
type fibTest struct {
	n        int
	expected int
}

var fibTests = []fibTest {
	{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
// END8 OMIT

// START9 OMIT
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
// END9 OMIT
