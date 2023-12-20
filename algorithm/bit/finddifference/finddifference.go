package finddifference

import (
	"math/bits"
	"sort"
)

/**
给定两个字符串s和t，它们只包含小写字母。
字符串t由字符串s随机重排，然后在随机位置添加一个字母。
请找出在t中被添加的字母。
*/

func FindDifference(s, t string) string {
	chars := []rune(s + t)
	var difference rune
	for _, v := range chars {
		difference ^= v //数字异或
	}
	return string(difference)
}

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。 找
出那个只出现了一次的元素。
// 数字异或得到的就是哪个数字
*/

func SingleNum(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

/**
给 你 一 个 整 数 数 组 arr 。 请 你 将 数 组 中 的 元 素 按 照 其 二 进 制 表 示 中 数 字 1 的 数 目 升 序 排
序。
如果存在多个数字二进制中1的数目相同，则必须将它们按照数值大小升序排列。
*/
type BitValue struct {
	Value    int //数字值
	BitCount int //bitCount值(二进制中1的数量)
}

type BitValues []BitValue

func (bs BitValues) Nums() []int {
	nums := make([]int, 0, len(bs))
	for _, b := range bs {
		nums = append(nums, b.Value)
	}
	return nums
}

func SortByBit(nums []int) []int {
	bitValues := make(BitValues, 0, len(nums))
	for _, num := range nums {
		bitValues = append(bitValues, BitValue{
			Value:    num,
			BitCount: bits.OnesCount64(uint64(num)),
		})
	}
	sort.Slice(bitValues, func(i, j int) bool {
		if bitValues[i].BitCount != bitValues[j].BitCount {
			return bitValues[i].BitCount < bitValues[j].BitCount
		}
		return bitValues[i].Value <= bitValues[j].Value
	})
	return bitValues.Nums()
}

/**
给一个函数f( x )，返回一个不小于x的最小的2的n次方。描述的比较绕口，举个例子。
f( 7 )=8
f( 9 )=16
f( 30 )=32
f( 64 )=64

1. 将2的1-n次方依次和num比较,得到第一个>=他的数
*/
func LowPower1(num int) int {
	res := 1
	for res < num {
		res = res << 1
	}
	return res
}

func LowPower2(x int) int {
	x -= 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	//最后再加1返回
	return x + 1
}

/**
给定一个整数数组 n u m s ，其中恰好有两个元素只出现一次，其余所有元素均出现两
次。找出只出现一次的那两个元素。
a & (-a) 可以获得a最低的非0位。

将数组按最低位为1的与运算强行分为2组(每组数字只有一个出现一次, 就好处理了)
*/
func TwoSingleNums(nums []int) []int {
	var (
		bitmask int
		res     = []int{0, 0}
	)
	for _, num := range nums {
		bitmask ^= num
	}
	// bitmask 获得的是两个出现一次的数字的异或值
	// 获取最后一位非0的值
	bitmask &= -bitmask
	// 这两个数字在这一位上面一个是1一个是0, 可以将nums按这位是1和0分为2
	// 分别异或,就可以分别得到2个出现一次的数字
	for _, n := range nums {
		if bitmask&n == 0 {
			res[0] = res[0] ^ n
		} else {
			res[1] = res[1] ^ n
		}
	}
	sort.Ints(res)
	return res
}

// 请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
func OneCounts(num int) int {
	var count int
	for num != 0 {
		// 判断最右边是否是1, 然后向右位移一位,直到数字为0
		count += (num & 1)
		num = num >> 1
	}
	return count
}

func OneCounts2(num int) int {
	var count int
	for num != 0 {
		num &= num - 1 //去掉最右边为1的位
		count++
	}
	return count
}

func OneCounts3(num int) int {
	var count int
	for num != 0 {
		num -= (num & (-num)) // 每次减去自己最右边为1的那一位直至为0
		count++
	}
	return count
}

/**
给定一个正整数，检查他是否为交替位二进制数：换句话说，就是他的二进制数相邻的两个位数永
不相等。
101010满足, 那么等于当前数字向右移动一位后,所有位都不同, 异或运算得到的结果就全部都是1(num)
此时num+1 就会得到2的n次方,与num与运算是0
*/
func HasAlternatingBits(num int) bool {
	num = num ^ (num >> 1)
	// 这里异或运算可以改为加法
	// num = num + (num >> 1)
	return num&(num+1) == 0
}

func HasAlternatingBits2(num int) bool {
	// 向右移动2位,再进行异或运算,那么只会剩下一位为1的,
	// 那么在和当前值-1与运算看是否=0
	num ^= (num >> 2)
	return num&(num-1) == 0
}

// 交换2个数字的值, 不借助临时变量
func Swap(nums []int, i, j int) []int {
	nums[i] = nums[i] + nums[j]
	nums[j] = nums[i] - nums[j]
	nums[i] = nums[i] - nums[j]
	return nums
}

func Swap2(nums []int, i, j int) []int {
	nums[i] = nums[i] ^ nums[j]
	nums[j] = nums[i] ^ nums[j]
	nums[i] = nums[i] ^ nums[j]
	return nums
}
