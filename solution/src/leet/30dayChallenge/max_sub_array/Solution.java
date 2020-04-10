class Solution {
    public int maxSubArray(int[] nums) {
        int maxSum = Integer.MIN_VALUE;

        int startIndex = getFirstPositiveIndex(nums);
        if(startIndex != -1){
            int crSum = nums[startIndex];
            if(crSum >maxSum){
                maxSum = crSum;
            }
            for (int i= startIndex+1; i<nums.length; i++){
                if(crSum + nums[i] > 0){
                    crSum += nums[i];
                    if(crSum >maxSum){
                        maxSum = crSum;
                    }
                }else{
                    crSum = 0;
                    if(nums[i] > 0){
                        crSum = nums[i];
                    }
                }
            }
        }
        if(maxSum < 0){
            for (int i=0; i<nums.length; i++){
                if(nums[i] > maxSum){
                    maxSum = nums[i];
                }
            }
        }


        return maxSum;
    }
    
    private static int getFirstPositiveIndex(int[] nums) {
        for (int i =0; i< nums.length; i++){
            if(nums[i] > 0) return i;
        }
        return -1;
    }
}