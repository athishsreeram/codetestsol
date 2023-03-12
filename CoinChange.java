import java.util.Arrays;

public class CoinChange {

    public static void main(String[] args) {
        CoinChange c = new CoinChange();
        int []coin = {1,2,3};
        System.out.println(c.coinChange(coin,4));
    }

    public int coinChange(int []coins,int amount){

        int []dp= new int[amount+1];

        Arrays.fill(dp,amount+1);
        dp[0]=0;

        for(int i=1;i<=amount;i++){
            for(int coin:coins){
                if(i-coin>=0) {
                    dp[i] = Math.min(dp[i],1+dp[i-coin]);
                }
            }
        }

        return dp[amount] != amount+1 ? dp[amount]:-1;


    }

}
