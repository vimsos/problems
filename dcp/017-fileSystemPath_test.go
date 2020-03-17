package dcp

import "testing"

func Test_longestPath(t *testing.T) {
	type args struct {
		fs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{fs: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"}, 32},
		{"2", args{fs: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"}, 20},
		{"3", args{fs: "a.txt"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPath(tt.args.fs); got != tt.want {
				t.Errorf("longestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
