public class Searcha2DMatrix {

    public static void main(String[] args) {

        int[][] matrix = { {1,2} ,
                {3,4}};
        
        Searcha2DMatrix sd = new Searcha2DMatrix();
        System.out.println(sd.searchMatrix(matrix,2));

    }
    public boolean searchMatrix(int[][] matrix, int target) {
        int i = 0, j = matrix[0].length - 1;
        while (i < matrix.length && j >= 0) {
            if (matrix[i][j] == target)
                return true;
            else if (matrix[i][j] > target)
                j--;
            else
                i++;
        }
        return false;
    }


}
