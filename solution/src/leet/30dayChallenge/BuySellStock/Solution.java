class Solution {
     public int maxProfit(int[] prices) {
        int profit = 0;
        int crIndex = 0;
        while (crIndex < prices.length-1){
            int pIndex = getPurchaseIndex(prices, crIndex);
            if(pIndex == -1) break;
            int sIndex = getSellIndex(prices, pIndex+1);
            if(sIndex == -1) {
                if(prices[pIndex] < prices[prices.length-1]){
                    profit += (prices[prices.length-1] - prices[pIndex]);
                    break;
                }
            }

            profit += (prices[sIndex] - prices[pIndex]);
            crIndex = sIndex + 1;
        }
        return profit;
    }

    private int getSellIndex(int[] prices, int crIndex) {
        while (crIndex < prices.length-1){
            if(prices[crIndex+1] < prices[crIndex]) {
                return crIndex;
            }
            crIndex++;
        }
        return -1;
    }

    private  int getPurchaseIndex(int[] prices, int crIndex) {
        while (crIndex < prices.length-1){
            if(prices[crIndex+1] > prices[crIndex]) {
                return crIndex;
            }
            crIndex++;
        }
        return -1;
    }
}
