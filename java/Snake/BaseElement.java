

import java.awt.*;
import javax.swing.*;

public class BaseElement {
    protected int x,y,width,height;
    protected int xStep, yStep;
    protected Direction direction;
    protected Image image;

    public BaseElement() {
        this.x = 0;
        this.y = 0;
        this.xStep = Parameters.BASE_SIZE;
        this.yStep = Parameters.BASE_SIZE;
        this.width = Parameters.BASE_SIZE;
        this.height = Parameters.BASE_SIZE;
        this.direction = Direction.DOWN;
        this.image = new ImageIcon(Parameters.RES_PATH + "Background.png").getImage();
    }
    public BaseElement(int x, int y)  {
        this.x = x;
        this.y = y;
        this.xStep = Parameters.BASE_SIZE;
        this.yStep = Parameters.BASE_SIZE;
        this.width = Parameters.BASE_SIZE;
        this.height = Parameters.BASE_SIZE;
        this.direction = Direction.DOWN;
        this.image = new ImageIcon(Parameters.RES_PATH + "Background.png").getImage();       
    }
    public void setX(int x) {
        this.x = x;
    }
    public void setY(int y) {
        this.y = y;
    }
    public void setDirection(Direction d) {
        this.direction = d;
    }
    public int getX() {
        return this.x;
    }
    public int getY() {
        return this.y;
    }
    public int getWidth() {
        return this.width;
    }
    public int getHeight() {
        return this.height;
    }
    public Image getImage() {
        return this.image;
    }
    public Direction getDirection() {
        return this.direction;
    }
    public void drawImage(Graphics g) {
        g.drawImage(this.image, this.x, this.y, this.width, this.height, null);
    }
    public void action() {
        this.changeDirection();
        this.moveX();
        this.moveY();
    }
    public void moveX() {
        if(this.direction.isRIGHT())
            this.x += this.xStep;
        if(this.direction.isLEFT())
            this.x -= this.xStep;
    }
    public void moveY() {
        if(this.direction.isDOWN())
            this.y += this.yStep;
        if(this.direction.isUP())
            this.y -= this.yStep;
    }
    public void changeDirection() {
        if(Key.UP.use())
            this.direction = Direction.UP;
        if(Key.LEFT.use())
            this.direction = Direction.LEFT;
        if(Key.RIGHT.use())
            this.direction = Direction.RIGHT;
        if(Key.DOWN.use())
            this.direction = Direction.DOWN;
    }

    public boolean intersects(BaseElement e) {
        Rectangle rect = new Rectangle(this.x,this.y,this.width,this.height);
        return rect.intersects(e.getX(), e.getY(), e.getWidth(), e.getHeight());
    }

    // public static void main(String[] argv) {
    //     BaseElement elt = new BaseElement();
    //     System.out.println(elt);
    // }

}