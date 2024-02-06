


public class lambda {

    interface Runnable {
        int run(int a, int b);
    }
    public static void main (String[] argv) {
        Runnable runner = (a,b) -> {
            int c = a + b;
            c += b;
            return c;
        };

        System.out.println("My result " + runner.run(1,2));

    }

}
