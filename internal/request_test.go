package internal

import (
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func Test_handledownload(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *http.Response
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				url: "https://golangcode.com/logo.svg",
			},
			want: &http.Response{
				StatusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handledownload(tt.args.url); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("handledownload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createfile(t *testing.T) {
	outFile, err := ioutil.TempFile("", "go_dl_temp_file")
	if err != nil {
		t.Fatal("Coudn't create the output file")
	}
	outFile.Close()
	// We just want to use this temp filename, so we delete the file,
	// otherwise downloader creates a new file
	os.Remove(outFile.Name())
}
