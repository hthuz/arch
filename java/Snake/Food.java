
import javax.swing.*;
import java.util.Random;
public class Food extends BaseElement{
    public Food(int x, int y){
        super(x,y);
        this.image = new ImageIcon(Parameters.RES_PATH + "Food.png").getImage();
    }

    public void reAppear(Snake... snakes){
        Random rand = new Random();
        this.setX(rand.nextInt(Parameters.GRID_WIDTH) * Parameters.BASE_SIZE);
        this.setY(rand.nextInt(Parameters.GRID_HEIGHT) * Parameters.BASE_SIZE);
        while(this.intersects(snakes)) {
            this.setX(rand.nextInt(Parameters.GRID_WIDTH) * Parameters.BASE_SIZE);
            this.setY(rand.nextInt(Parameters.GRID_HEIGHT) * Parameters.BASE_SIZE);
        }
    }

    public boolean intersects(Snake... snakes) {
        for(Snake snake : snakes)
            if(snake.intersects(this))  
                return true;
        return false;
    }

}
