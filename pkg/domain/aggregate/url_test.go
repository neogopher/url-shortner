// Package aggregate holds aggregates that combine multiple entities into full objects.
package aggregate

import (
	"testing"
)

func TestNewURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		want        URL
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Empty Path",
			args: args{
				path: "",
			},
			want:        URL{},
			wantErr:     true,
			expectedErr: ErrEmptyPath,
		},
		{
			name: "Invalid Path",
			args: args{
				path: "This is an invalid path",
			},
			want:        URL{},
			wantErr:     true,
			expectedErr: ErrInvalidPath,
		},
		{
			name: "Valid Path",
			args: args{
				path: "http://www.google.com",
			},
			want:        URL{},
			wantErr:     false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewURL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == true {
				if err != tt.expectedErr {
					t.Errorf("NewURL() error = %v, expectedErr %v", err, tt.expectedErr)
					return
				}
			}
		})
	}
}
