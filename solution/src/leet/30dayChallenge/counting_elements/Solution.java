import java.util.HashSet;
import java.util.Set;
class Solution {
    public int countElements(int[] arr) {
         int count = 0;
        Set<Integer> set = new HashSet<>();

        for (int item : arr){
            set.add(item);
        }
        for (int i = 0; i < arr.length; i++) {
            if(set.contains(arr[i]+1)){
                count++;
            }
        }
        return count;
    }
}