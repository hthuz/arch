
public class hello{
    public static void main(String[] args) {
        String msg = "Hello World!";
        Class <String> msg_class = msg.getClass();
        System.out.println(msg);
        System.out.println(msg.getClass());
        System.out.println(msg.getClass().getSimpleName());
        System.out.println(msg.getClass().getTypeName());

        System.out.println(msg_class);

    }
}
