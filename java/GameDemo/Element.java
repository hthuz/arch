
import java.awt.*;
import javax.swing.*;

public class Element {
    private int x,y,width,height;
    private Image image;
    public Element(int x, int y, int width, int height) {
        this.image = new ImageIcon("resources/profile.jpg").getImage();
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
    }
    public Image getImage() {
        return this.image;
    }
    public int getX(){
        return this.x;
    }
    public int getY(){
        return this.y;
    }
    public int getWidth(){
        return this.width;
    }
    public int getHeight(){
        return this.height;
    }
    public void setX(int x){
        this.x = x;
    }
    public void setY(int y){
        this.y = y;
    }

    // public static void main(String[] argv) {
    //     System.out.println("Hello");

    // }
    
}
