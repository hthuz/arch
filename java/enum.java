

public class Test {
    enum Key {
        UP(4),
        DOWN(3),
        LEFT(2),
        RIGHT(1);

        private int curkey;
        Key(int curkey) {
            this.curkey = curkey;
        }
        public void print() {
            System.out.println(curkey);
            curkey++;
        }

    }

    public static void main(String[] argv) {
        System.out.println("Ama");
        Key.UP.print();
        Key.LEFT.print();
        Key.UP.print();
        Key.UP.print();
    }
}
