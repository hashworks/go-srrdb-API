package srrdb

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	result, err := Search("Harry Potter")
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(result.Results) == 0 {
			t.Error("Expected any results, got 0.")
		} else {
			if result.Results[0].Dirname == "" {
				t.Error("Didn't expected the first dirname to be empty.")
			}
		}
	}
}

func TestDownloadAndUpload(t *testing.T) {
	t.Parallel()
	bytes, err := Download("Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS")
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(bytes) == 0 {
			t.Error("Expected any data, got 0.")
		} else {
			file, err := ioutil.TempFile(os.TempDir(), "Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS.srr")
			if err != nil {
				t.Error(err.Error())
			} else {
				file.Write(bytes)
				file.Close()
				response, err := UploadSRRs([]string{file.Name()}, nil)
				if err != nil {
					t.Error(err.Error())
				} else {
					if len(response.Files) == 0 {
						t.Error("Expected any results, got 0.")
					}
				}
			}
		}
	}
}
