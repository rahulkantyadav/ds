package leet.thirtydaychallenge.counting_elements;


import java.util.HashSet;
import java.util.Set;

class CountuingElements {
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

    public static void main(String[] args) {
        int[] arr = {1,34,5,62,3,4,1,23};
        new CountuingElements().countElements(arr);
    }
}