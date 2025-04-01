public class Solution{
public int solution(int num, int k) {
        String numStr= String.valueOf(num);

        int index=numStr.indexOf(String.valueOf(String.valueOf(k)));

        if(index != -1){
            return index+1;
        } else{
            return -1;
        }
    }
}
