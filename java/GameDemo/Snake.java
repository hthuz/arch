public class Snake {
    private int x;
    private int y;
    public Snake(int _x, int _y) {
        x = _x;
        y = _y;
    }
    public int getX() {
        return x;
    }
    public int getY() {
        return y;
    }
    public void printPosition() {
        System.err.println(String.format("x=%d,y=%d", x,y));
    }

    public void updatePosition(String direction) {
        switch (direction) {
            case "up":
                y--;
                break;
            case "down":
                y++;
                break;
            case "left":
                x--;
                break;
            case "right":
                x++;
                break;
        }
    }
}
