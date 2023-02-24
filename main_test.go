package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputFile, err := ioutil.TempFile("", "input.html")
	if err != nil {
		t.Fatalf("Failed to create temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	testInput := `<div>div tag</div>
<h1>h1 tag</h1>
<h2>h2 tag</h2>
<h3>h3 tag</h3>
<h4>h4 tag</h4>
<h5>h5 tag</h5>
<h6>h6 tag</h6>
<img src="" />`

	if _, err := inputFile.Write([]byte(testInput)); err != nil {
		t.Fatalf("Failed to write test input to file: %v", err)
	}

	outputFile, err := ioutil.TempFile("", "output.html")
	if err != nil {
		t.Fatalf("Failed to create temporary output file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Run the main function
	os.Args = []string{"hidgen", inputFile.Name(), outputFile.Name()}
	main()

	outputBytes, err := ioutil.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	output := string(outputBytes)

	expectedOutput := `<div>div tag</div>
<h1>h1 tag</h1>
<a id="h2 tag"></a>
<h2>h2 tag</h2>
<a id="h3 tag"></a>
<h3>h3 tag</h3>
<h4>h4 tag</h4>
<h5>h5 tag</h5>
<h6>h6 tag</h6>
<img src="" />`
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Output does not contain expected HTML: %q", expectedOutput)
	}
}
