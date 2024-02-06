
import javax.swing.*;
public class Food extends BaseElement{
    public Food(int x, int y){
        super(x,y);
        this.image = new ImageIcon(Parameters.RES_PATH + "Food.png").getImage();
    }
    public void reAppear(){
        this.setX(this.getX() + Parameters.BASE_SIZE * 2);
        this.setY(this.getY() + Parameters.BASE_SIZE * 2);
    }
}
