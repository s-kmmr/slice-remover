package strslice_test

import (
	"testing"

	"github.com/s-kmmr/slice-remover/strslice"
	"golang.org/x/xerrors"
)

func Test_Remove(t *testing.T) {
	type args struct {
		slice []string
		idx   int
	}
	type testData struct {
		name      string
		args      args
		want      []string
		wantErr   error
		isWantErr bool
	}

	tests := []testData{
		{
			name: "remove a",
			args: args{
				slice: []string{"a", "b", "c", "d", "e"},
				idx:   0,
			},
			want:      []string{"b", "c", "d", "e"},
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "slice bounds out of range idx",
			args: args{
				slice: []string{"a", "b", "c", "d", "e"},
				idx:   100,
			},
			want:      []string{"a", "b", "c", "d", "e"},
			wantErr:   xerrors.New("slice bounds out of range [:100] with capacity 5"),
			isWantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := strslice.Remove(tt.args.slice, tt.args.idx)
			if (err != nil) != tt.isWantErr {
				t.Fatalf("(err != nil) = %v .tt.isWantErr = %v", (err != nil), tt.isWantErr)
			}

			if err != nil {
				if tt.wantErr == nil {
					t.Fatal("err != nil,wantErr == nil")
				}

				if err.Error() != tt.wantErr.Error() {
					t.Errorf("err = %s. wantErr = %s", err.Error(), tt.wantErr.Error())
				}
			}

			if len(tt.want) != len(ret) {
				t.Fatalf("not match length. want = %v, result = %v", len(tt.want), len(ret))
			}

			for i, v := range tt.want {
				if ret[i] != v {
					t.Errorf("not match value. want = %v, result = %v", v, ret[i])
				}
			}
		})
	}
}
