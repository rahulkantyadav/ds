package leet.thirtydaychallenge.contaigous_binary_sum;

import java.util.HashMap;

public class ContagiousBinarySum {


    public int findMaxLength(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>();

        int totalSum = 0;
        int maxLen = 0;
        int n = nums.length;
        for (int i = 0; i < n; i++) {
            nums[i] = (nums[i] == 0) ? -1 : 1;
        }
        for (int i = 0; i < n; i++) {
            totalSum += nums[i];
            if (totalSum == 0) {
                maxLen = i + 1;
            }


            if (map.containsKey(totalSum + n)) {
                if (maxLen < i - map.get(totalSum + n)) {
                    maxLen = i - map.get(totalSum + n);
                }
            } else {
                map.put(totalSum + n, i);
            }
        }
        for (int i = 0; i < n; i++) {
            nums[i] = (nums[i] == -1) ? 0 : 1;
        }
        return maxLen;
    }
}
