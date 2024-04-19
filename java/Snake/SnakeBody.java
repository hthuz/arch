
import java.awt.*;
import javax.swing.*;
public class SnakeBody extends BaseElement {
    
    public SnakeBody(int x, int y) {
        super(x,y);
        this.image = new ImageIcon(Parameters.RES_PATH + "Body.png").getImage();
    }
    public SnakeBody(int x, int y, String imageName) {
        super(x, y, imageName);
    }
    
}
