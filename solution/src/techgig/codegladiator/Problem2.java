package techgig.codegladiator;

import java.util.Arrays;
import java.util.Scanner;

public class Problem2 {

    public static void main(String args[]) throws Exception {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();

        for (int i = 0; i < T; i++) {
            int N = sc.nextInt();
            int[] teamG = new int[N];
            Integer[] otherTeam = new Integer[N];
            for (int j = 0; j < N; j++) {
                teamG[j] = sc.nextInt();
            }
            for (int j = 0; j < N; j++) {
                otherTeam[j] = sc.nextInt();
            }
            Person[] arr = new Person[1000000];
            for (int j = 0; j < 1000000; j++) {
                arr[j] = new Person();
            }

            Arrays.sort(teamG);
            Arrays.sort(otherTeam);

            int winCount = 0;
            int teamGIndex = 0;
            int otherTeamIndex = 0;

            while(teamGIndex < N && otherTeamIndex < N){
                if(teamG[teamGIndex] > otherTeam[otherTeamIndex]){
                    winCount++;
                    otherTeamIndex++;
                }
                teamGIndex++;
            }
            System.out.println(winCount);

        }


    }
}

class Person{
    String[] str;
    String[] str1;
    String[] str2;
    Person[] p1 = new Person[100];
}
