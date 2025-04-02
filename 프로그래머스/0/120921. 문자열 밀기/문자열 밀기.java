public class Solution{
    public int solution(String A, String B) {
        String doubledB=B+B;

        int index=doubledB.indexOf((A));
        if (index !=-1){
            return index;
        } else {
            return -1;
        
    }
}
}