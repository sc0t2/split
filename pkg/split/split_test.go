package split_test

import (
	"encoding/csv"
	"github.com/sc0t2/split/pkg/split"
	"io/ioutil"
	"os"
	"testing"
)

const DefaultChunkSize = 1000

func TestSplitCsv(t *testing.T) {
	cases := []struct {
		name      string
		want      int
		path      string
		chunkSize int
	}{
		{"test 0 row", 0, "testdata/0.csv", DefaultChunkSize},
		{"test 1 row", 1, "testdata/1.csv", DefaultChunkSize},
		{"test 10 row", 2, "testdata/10.csv", 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// create temporary directory to hold test output
			outDir, err := ioutil.TempDir("", "split")
			if err != nil {
				t.Fatal(err)
			}

			// load up the test data
			f, err := os.Open(c.path)
			if err != nil {
				t.Fatal(err)
			}
			r := csv.NewReader(f)

			// run the function
			got, err := split.Csv(r, c.chunkSize, outDir, "testing")
			if err != nil {
				t.Fatal(err)
			}

			// compare the results
			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}

			// remove the test output
			removeErr := os.RemoveAll(outDir)
			if removeErr != nil {
				t.Fatal(removeErr)
			}

			_ = f.Close()
		})
	}
}
