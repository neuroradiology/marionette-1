package file

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestExists ensures we report on file existence.
func TestExists(t *testing.T) {

	// Create a file, and ensure it exists
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatalf("create a temporary file failed")
	}

	// Does it exist
	res := Exists(tmpfile.Name())
	if !res {
		t.Fatalf("after creating a temporary file it doesnt exist")
	}

	// Remove the file
	os.Remove(tmpfile.Name()) // clean up

	// Does it exist, still?
	res = Exists(tmpfile.Name())
	if res {
		t.Fatalf("after removing a temporary file it still exists")
	}
}

// TestHash tests our hashing function
func TestHash(t *testing.T) {

	type Test struct {
		input  string
		output string
	}

	tests := []Test{{input: "hello", output: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{input: "steve", output: "9ce5770b3bb4b2a1d59be2d97e34379cd192299f"},
	}

	for _, test := range tests {

		// Create a file with the given content
		tmpfile, err := ioutil.TempFile("", "example")
		if err != nil {
			t.Fatalf("create a temporary file failed")
		}

		// Write the input
		_, err = tmpfile.Write([]byte(test.input))
		if err != nil {
			t.Fatalf("error writing temporary file")
		}

		out := ""
		out, err = HashFile(tmpfile.Name())
		if err != nil {
			t.Fatalf("failed to hash file")
		}

		if out != test.output {
			t.Fatalf("invalid hash %s != %s", out, test.output)
		}

		os.Remove(tmpfile.Name())
	}

	// Hashing a missing file should fail
	_, err := HashFile("/this/does/not/exist")
	if err == nil {
		t.Fatalf("should have seen an error, didn't")
	}
}
