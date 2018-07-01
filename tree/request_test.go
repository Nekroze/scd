package tree

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDirectoryRequest(t *testing.T) {
	type args struct {
		userInput string
	}
	tests := []struct {
		name    string
		args    args
		wantNew DirectoryRequest
		wantErr bool
	}{
		{
			name: "nearby directory",
			args: args{
				userInput: "tree",
			},
			wantNew: DirectoryRequest{
				url: &url.URL{
					Path: "tree",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNew, err := NewDirectoryRequest(tt.args.userInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDirectoryRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantNew.url, gotNew.url)
		})
	}
}
