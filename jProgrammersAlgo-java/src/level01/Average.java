package level01;

public class Average {
    public static void main(String args[]) {
        double[] darr1 = new double[]{5, 3, 4};
        double ex1 = average(darr1); // 4.0, 12/3
        System.out.println(ex1);
    }

    static double average(double[] darr) {
        double total = 0.0;
        for(double d: darr) {
            total += d;
        }

        return total/darr.length;
    }
}
