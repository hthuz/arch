
import java.awt.*;
import java.util.LinkedList;

public class Snake extends BaseElement{

    LinkedList<SnakeBody> BodyList;
    private int tailx, taily;
    private int score;
    private boolean alive;

    public Snake(int x, int y) {
        super(x,y);
        this.tailx = x;
        this.taily = y;
        this.score = 0;
        this.alive = true;
        BodyList = new LinkedList<SnakeBody>();
        BodyList.add(new SnakeBody(x, y,"DownHead.png"));
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
        // Move body position
        SnakeBody newhead = new SnakeBody(this.getX(), this.getY());
        BodyList.getFirst().setImage("Body.png");
        BodyList.removeLast();
        BodyList.addFirst(newhead);
        // Head Direction
        switch(this.direction) {
            case UP:
                BodyList.getFirst().setImage("UpHead.png");
                break;
            case DOWN:
                BodyList.getFirst().setImage("DownHead.png");
                break;
            case LEFT:
                BodyList.getFirst().setImage("LeftHead.png");
                break;
            case RIGHT:
                BodyList.getFirst().setImage("RightHead.png");
                break;
        }
    }
    public void eat(Food f) {
        if(!this.intersects(f))
            return;
        BodyList.add(new SnakeBody(this.tailx, this.taily));
        this.score++;
        f.reAppear();
    }
    @Override
    public boolean intersects(BaseElement elt) {
        for(SnakeBody body : BodyList) {
            if(body.intersects(elt))
                return true;
        }
        return false;
    }
    public int getScore() {
        return this.score;
    }
    // public void setScore(int score) {
    //     this.score = score;
    // }
    public boolean isAlive() {
        return this.alive;
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
