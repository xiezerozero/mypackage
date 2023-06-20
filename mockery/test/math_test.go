package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiezerozero/mypackage/mockery/test/mocks"
	"testing"
)

func TestCalculator_Double(t *testing.T) {
	// 设置mock
	total := new(mocks.Totaller)
	type mockCase struct {
		inputs []interface{}
		result int
	}
	mockCases := []mockCase{
		{
			inputs: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			result: 55,
		},
		{
			inputs: []interface{}{1, 2, 3, 4, 5},
			result: 15,
		},
		{
			inputs: []interface{}{1, 2, 3},
			result: 100,
		},
	}
	for _, mc := range mockCases {
		total.On("Total", mc.inputs...).Return(mc.result)
	}
	// 将mock注入计算器中,测试指定输入和输出是否符合预期
	c := &Calculator{total: total}
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "double 1-10",
			args: args{
				inputs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 110,
		},
		{
			name: "double 1-5",
			args: args{
				inputs: []int{1, 2, 3, 4, 5},
			},
			want: 30,
		},
		{
			name: "double 1-3",
			args: args{
				inputs: []int{1, 2, 3},
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, c.Double(tt.args.inputs...))
		})
	}
}

func Foo(s Stringer) string {
	return s.String()
}

func TestString(t *testing.T) {
	mockStringer := mocks.Stringer{}
	mockStringer.On("String").Return("mocked string")
	mockStringer.AssertExpectations(t)
}
