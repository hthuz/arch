
import java.awt.*;
import java.util.LinkedList;

public class Snake extends BaseElement{

    LinkedList<SnakeBody> BodyList;
    private int tailx, taily;

    public Snake(int x, int y) {
        super(x,y);
        this.tailx = x;
        this.taily = y;
        BodyList = new LinkedList<SnakeBody>();
        BodyList.add(new SnakeBody(x, y));
    }

    @Override
    public void drawImage(Graphics g) {
        for(SnakeBody body : BodyList) 
            body.drawImage(g);
    }
    @Override
    public void action() {
        // Before move, record tail position
        this.tailx = BodyList.getLast().getX();
        this.taily = BodyList.getLast().getY();
        super.action();
        SnakeBody newhead = new SnakeBody(this.getX(), this.getY());
        BodyList.removeLast();
        BodyList.addFirst(newhead);
    }
    public void eat(Food f) {
        if(!this.intersects(f))
            return;
        BodyList.add(new SnakeBody(this.tailx, this.taily));

        f.reAppear();
    }

    // @Override
    // public void moveX() {
    //     BodyList.getFirst().moveX();
    // }
    // @Override
    // public void moveY() {
    //     BodyList.getFirst().moveY();
    // }

}
