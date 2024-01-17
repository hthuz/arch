
import java.awt.Graphics;
import java.awt.image.*;
import javax.swing.*;

public class GamePanel extends JPanel {
    private Element[] elts;

    public GamePanel(Element[] elts) {
        this.elts = elts;
    }

    @Override
    // Called when frame.setvisible is set
    public void paint(Graphics g) {
        super.paint(g);
        for(Element elt : elts)
            g.drawImage(elt.getImage(),elt.getX(),elt.getY(),elt.getWidth(), elt.getHeight(),null);
    }

}