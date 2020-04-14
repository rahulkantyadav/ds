package leet.thirtydaychallenge.happyNumber;

import java.util.HashSet;
import java.util.Set;

class HappyNumber {
     public  boolean isHappy(int n) {
        Set<Integer> numberList = new HashSet<>();
        return checkHappy(n, numberList);
    }

    private boolean checkHappy(int n, Set<Integer> numberList){
        if(n == 1 || isPower(10, n)) return true;
        if(numberList.contains(n)) return false;
        numberList.add(n);
        n = getSquaredSum(n);
        return checkHappy(n, numberList);
    }

    private boolean isPower(int x, int y)
    {
        int res1 = (int) (Math.log(y) / Math.log(x));
        double res2 = Math.log(y) / Math.log(x);
        return (res1 == res2);
    }

    private int getSquaredSum(int n) {
        int ans = 0;
        while(n != 0){
            int r = n%10;
            ans += r*r;
            n = n/10;
        }

        return ans;
    }
}