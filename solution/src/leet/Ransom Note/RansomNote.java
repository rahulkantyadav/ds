import java.util.*;

public class RansomNote {

    public static void main(String[] args) {

        System.out
                .println(new RansomNote().canConstruct("bcagc", "bcag"));
    }

    public boolean canConstruct(String ransomNote, String magazine) {
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i = 0; i < magazine.length(); i++) {
            int val = 0;
            if (map.containsKey((int)magazine.charAt(i))) {
                val = map.get((int)magazine.charAt(i));
            }
            map.put((int) magazine.charAt(i), val + 1);
        }
        for (int i = 0; i < ransomNote.length(); i++) {
           if(!map.containsKey((int)ransomNote.charAt(i)) || map.get((int)ransomNote.charAt(i)) <=0){
               return false;
           }else{
            map.put((int)ransomNote.charAt(i), map.get((int)ransomNote.charAt(i)) - 1);
           }
        }
        return true;
    }
}