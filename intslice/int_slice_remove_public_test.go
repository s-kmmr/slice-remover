package intslice_test

import (
	"testing"

	"github.com/s-kmmr/slice-remover/intslice"
	"golang.org/x/xerrors"
)

func Test_Remove(t *testing.T) {
	type args struct {
		slice []int
		idx   uint
	}
	type testData struct {
		name      string
		args      args
		want      []int
		wantErr   error
		isWantErr bool
	}

	tests := []testData{
		{
			name: "idx is within the threshold",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   0,
			},
			want:      []int{2, 3, 4, 5},
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "idx is within the threshold(end)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   4,
			},
			want:      []int{1, 2, 3, 4},
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "idx is over the threshold(end)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   5,
			},
			want:      []int{1, 2, 3, 4, 5},
			wantErr:   xerrors.New("slice bounds out of range [:5] with capacity 5"),
			isWantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := intslice.Remove(tt.args.slice, tt.args.idx)
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
