class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap();
        String[] newStrs = new String[strs.length];
        for (int i = 0; i < strs.length; i++) {
            char[] cArr = strs[i].toCharArray();
            Arrays.sort(cArr);
            newStrs[i] = new String(cArr);
            System.out.println(newStrs[i]);
        }

        for (int i = 0; i < newStrs.length; i++) {
            if (map.containsKey(newStrs[i])) {
                List<String> list = map.get(newStrs[i]);
                list.add(strs[i]);
                map.put(newStrs[i], list);
            } else {
                List<String> list = new ArrayList<>();
                list.add(strs[i]);
                map.put(newStrs[i], list);
            }
        }
        List<List<String>> list = new ArrayList<>();

        for (Map.Entry<String, List<String>> entry : map.entrySet()) {
            List<String> subList = entry.getValue();
            list.add(subList);
        }
        return list;
    }
}