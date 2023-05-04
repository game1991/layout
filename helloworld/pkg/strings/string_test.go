package strings

import "testing"

func TestCombineWithDelimiter(t *testing.T) {
	type args struct {
		req       []any
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试字符串拼接数字",
			args: args{
				req:       []any{-1, 0, 9},
				delimiter: ",",
			},
			want: "-1,0,9",
		},
		{
			name: "测试字符串拼接数字uint32",
			args: args{
				req:       []any{3, 1, 2},
				delimiter: ",",
			},
			want: "3,1,2",
		},
		{
			name: "测试字符串拼接浮点",
			args: args{
				req:       []any{1.0, 2.6, 1.4},
				delimiter: ";",
			},
			want: "1;2.6;1.4",
		},
		{
			name: "测试字符串",
			args: args{
				req:       []any{"a", "b", "c"},
				delimiter: "=",
			},
			want: "a=b=c",
		},
		{
			name: "测试字符串拼接布尔",
			args: args{
				req:       []any{true, false, false},
				delimiter: ".",
			},
			want: "true.false.false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineWithDelimiter(tt.args.req, tt.args.delimiter); got != tt.want {
				t.Errorf("CombineWithDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
