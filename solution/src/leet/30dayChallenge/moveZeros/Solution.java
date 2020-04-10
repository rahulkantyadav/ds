class Solution {
    public void moveZeroes(int[] nums) {
        int crIndex = 0;
        for (int i = 0; i < nums.length; i++) {
            if(nums[i] != 0){
                if(crIndex != i){
                    nums[crIndex] = nums[i];
                }
                crIndex++;
            }
        }
        for (int i = crIndex; i < nums.length; i++) {
            nums[i] = 0;
        }
    }
}