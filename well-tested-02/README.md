# Naming conventions
main_test.go
package main    // same package for whitebox tests
package main_test   // add _test suffix to package for blackbox tests
func TestFoo(t *testing.T){
}

# Writing tests
* Reporting Test Failures
Immediate failure
t.FailNow()
t.Fatal(args ...interface{})
t.FatalIf(format string, args ...interface{})
Non-immediate failure
t.Fail()
t.Error(args ...interface{})
t.ErrorIf(format string, args ...interface{})

# Running tests
* go test
* go test {pkg1} {pkg2} // go test net/http
* go test ./... // Run tests in current package and descendants
* go test -v    // Generate verbose output
* go test -run {regexp} // Run only tests matching {regexp}

# Viewing test coverage
* go test -cover
* go test -coverprofile cover.out   // 
* go tool cover -func cover.out
* go tool cover -html cover.out // output html ui
* go test -coverprofile count.out -covermode count

# Table-driven tests
func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		name   string
		expect string
	}{
		{name: "Gopher", expect: "Hello, Gopher!\n"},
		{name: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.name)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got %q", s.name, got, s.expect)
		}
	}
}

# Other useful function
* Log and Logf
* Helper
* Skip, Skipf, and SkipNow
* Run
* Parallel