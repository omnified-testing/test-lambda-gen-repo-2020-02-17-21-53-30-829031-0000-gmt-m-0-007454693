package main

	import (
		utility "github.com/omnified-testing/test-lambda-gen-repo-2020-02-17-21-53-30-829031-0000-gmt-m-0-007454693/utility"
	)
	
	func main() {
	}
	
	// Call : The executable API of the function, just calls the subpackage.
	// Input and output needs to match that of utility subpackage's Call function
	func Call(s string) string {
		return utility.Call()
	}