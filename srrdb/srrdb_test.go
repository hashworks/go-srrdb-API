package srrdb

import (
	"os"
	"strings"
	"testing"
)

func TestDownloadAndUpload(t *testing.T) {
	t.Parallel()
	bytes, err := Download("Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS")
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(bytes) == 0 {
			t.Error("Expected any data, got 0.")
		}
		if len(bytes) < 70000 {
			if strings.HasPrefix(string(bytes), "<!") {
				t.Error("Got unexpected HTML response:", string(bytes))
			}
			t.Error("Expected > 70000 bytes, got ", len(bytes), ".")
		} else {
			file, err := os.CreateTemp(os.TempDir(), "Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS.srr")
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

func TestSearch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query   string
		want    SearchResponse
		wantErr bool
	}{
		{
			name:  "Search for Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS",
			query: "Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS",
			want: SearchResponse{
				Results: []SearchResult{
					{
						Dirname:        "Harry.Potter.And.The.Chamber.Of.Secrets.2002.720p.BluRay.x264-SiNNERS",
						HasNFOResponse: "yes",
						HasSRSResponse: "yes",
						DateResponse:   "2011-07-21 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Search(tt.query)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Search() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Search() succeeded unexpectedly")
			}
			// Expect one result
			if got.ResultCount != 1 {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}

			// Validate that all struct fields match the expected values
			if got.Results[0].Dirname != tt.want.Results[0].Dirname ||
				got.Results[0].HasNFOResponse != tt.want.Results[0].HasNFOResponse ||
				got.Results[0].HasSRSResponse != tt.want.Results[0].HasSRSResponse ||
				got.Results[0].DateResponse != tt.want.Results[0].DateResponse {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
