package main

func main()  {

}

func twoSum2(nums []int, target int) []int {

	if nums == nil {
		return nil
	}
	for i:=0;i<len(nums);i++ {
		temp:= target - nums[i]
		if target <= 0 {
			continue
		}
		for j:=i+1;j<len(nums);j++{
			if nums[j] - temp == 0 {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum3(nums []int, target int) []int {

	if nums == nil {
		return nil
	}

	sumMap := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		temp:= target - nums[i]
		if index, ok := sumMap[temp]; ok {
			return []int{i, index}
		} else {
			sumMap[nums[i]] = i
		}
	}

	return nil
}