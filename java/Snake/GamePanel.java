
import javax.swing.*;
import java.awt.Graphics;

public class GamePanel extends JPanel {
    private BaseElement[] elts;

    public GamePanel(BaseElement... elts) {
        this.elts = elts;
    }
    @Override
    public void paint(Graphics g) {
        super.paint(g);
        for (BaseElement elt : elts)
            elt.drawImage(g);
    }
}
