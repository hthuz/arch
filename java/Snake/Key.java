
import java.util.HashSet;
import java.awt.event.*;
public enum Key {
    UP(KeyEvent.VK_UP),
    DOWN(KeyEvent.VK_DOWN),
    LEFT(KeyEvent.VK_LEFT),
    RIGHT(KeyEvent.VK_RIGHT);

    private static HashSet<Integer> keyset = new HashSet<Integer>();
    private int keycode;

    Key(int e){
        this.keycode = e;
    }

    public static void add(int e) {
        keyset.add(e);
    }
    public static void remove(int e) {
        keyset.remove(e);
    }
    public boolean use() {
        return keyset.contains(this.keycode);
    }
    public static void print() {
        for(int key : keyset)
            System.out.print(key);
        System.out.print("\n");
    }
}
