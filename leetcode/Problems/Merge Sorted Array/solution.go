func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if m > 0 && n > 0 {
			a := nums1[m-1]
			b := nums2[n-1]

			if a > b {
				nums1[i] = a
				m--
				continue
			}

			if b > a {
				nums1[i] = b
				n--
				continue
			}

			if a == b {
				nums1[i] = a
				nums1[i-1] = b
				m--
				n--
				i--
				continue
			}
		}

		if m > 0 && n == 0 {
			nums1[i] = nums1[m-1]
			m--
			continue
		}

		if m == 0 && n > 0 {
			nums1[i] = nums2[n-1]
			n--
			continue
		}

	}

}
