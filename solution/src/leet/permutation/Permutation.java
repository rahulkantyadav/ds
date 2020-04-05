import java.util.*;

public class Permutation {

    private int crIndex = 0;

    public static void main(String[] args) {
        List<List<Integer>> list = new ArrayList<>();
        int[] arr = { 1, 2, 3 };
        int[][] ans = new Permutation().start(arr);
        for (int[] is : ans) {
            List<Integer> sublist = new ArrayList<>();
            for (int is2 : is) {
                sublist.add(is2);
            }
            list.add(sublist);
        }
        
    }

    public int[][] start(int[] arr) {
        int len = getPermutSize(arr);
        int[][] ans = new int[len][arr.length];
        permut(arr, 0, arr.length-1, ans);
        return ans;
    }

    private void permut(int[] arr, int s, int e,int[][] ans){
        if(s==e) {
            ans[crIndex++] = arr;
        }
        else{
            for (int i = s; i <= e; i++) {
                int[] newArr = swap(arr, i, s);
                permut(newArr, s+1, e, ans);
                swap(newArr, s, i);
            }
        }
        
    }

    private int[] swap(int[] arr, int a, int b) {
        int[] copy = new int[arr.length];
            int index = 0;
            for (int i : arr) {
                copy[index++] = i;
            }

        int temp = copy[a];
        copy[a] = copy[b];
        copy[b] = temp;
        return copy;
    }

    private int getPermutSize(int[] arr) {
        int len = arr.length;
        int ans = 1;
        while (len > 0) {
            ans *= len;
            len--;
        }
        return ans;
    }
}