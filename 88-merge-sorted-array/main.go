package main

func merge1(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i >= m {
			tmp = append(tmp, nums2[j:]...)
			break
		}
		if j >= n {
			tmp = append(tmp, nums1[i:m]...)
			break
		}
		if nums1[i] < nums2[j] {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}
	copy(nums1, tmp)
}

// 从后往前方可以就地归并
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n - 1; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p] = nums2[n-1]
			n--
		} else {
			nums1[p] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

func main() {

}
