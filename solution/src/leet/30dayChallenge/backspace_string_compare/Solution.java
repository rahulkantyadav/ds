class Solution {
     public boolean backspaceCompare(String S, String T) {
        int lenT = T.length();
        int lenS = S.length();

        int indexT = lenT-1;
        int indexS = lenS-1;
        int count = 0;

        while (!(indexT < 0 && indexS < 0)) {
            indexT = getNextValidIndex(T, indexT, count);
            indexS = getNextValidIndex(S, indexS, count);

            if (indexT < 0 || indexS < 0) {
                break;
            }

            if (S.charAt(indexS) != T.charAt(indexT)) {
                return false;
            }
            indexS--;
            indexT--;
        }

        if (indexS != indexT) {
            return false;
        }
        return true;
    }

    private int getNextValidIndex(String s, int crIndex, int count) {
        if(crIndex>=0 && s.charAt(crIndex) != '#' && count <=0) return crIndex;
        if(crIndex < 0) return -1;
        if(s.charAt(crIndex) != '#'){
            crIndex = getNextValidIndex(s, crIndex-1, count-1);
        }else{
            int newCount = getConsecutiveHashCount(s, crIndex);
            crIndex = getNextValidIndex(s, crIndex - newCount, count + newCount);
        }
        return crIndex;
    }

    private int getConsecutiveHashCount(String t, int index){
    int count =0;
        for (int i = index; i >= 0; i--) {
            if(t.charAt(i) == '#'){
                count++;
            }else{
                break;
            }
        }
        return count;

    }
}