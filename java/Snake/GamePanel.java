
import javax.swing.*;
import javax.swing.text.html.HTMLDocument.HTMLReader.ParagraphAction;

import java.awt.Graphics;
import java.awt.*;

public class GamePanel extends JPanel {
    private BaseElement[] elts;
    private Image background = new ImageIcon(Parameters.RES_PATH + "Background.png").getImage();

    public GamePanel(BaseElement... elts) {
        this.elts = elts;
    }
    @Override
    public void paint(Graphics g) {
        super.paint(g);
        g.drawImage(background, 0, 0, Parameters.FRAME_WIDTH, Parameters.FRAME_HEIGHT,null);

        for (BaseElement elt : elts)
            elt.drawImage(g);
    }
}
