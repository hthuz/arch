




public class hello{
    public static void main(String[] args) {

        String body = "Hello\nWorld\nNo\nYes Haha";
        String newline = body.replace("\n","");
        int loc = body.length() - body.replace("\n","").length();
        System.out.println(newline);
        System.out.println(loc);
    }
}

