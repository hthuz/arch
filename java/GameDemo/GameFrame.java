
import java.util.Scanner;
import java.util.jar.JarEntry;
import java.awt.*;
import java.awt.RenderingHints.Key;
import java.awt.event.*;
import javax.swing.*;

public class GameFrame extends JFrame{
    public GameFrame() {
        // Meta info
        this.setTitle("My Game");
        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        this.setBounds(400, 200, 600, 400);;

        // Create element
        Element elt = new Element(50, 300, 50, 50);
        Element obstacle = new Element(0, 50, 50, 50);
        Element[] elts = {elt,obstacle};
        GamePanel panel = new GamePanel(elts);
        this.add(panel);
        this.setVisible(true);

        // Keyboard listener
        this.addKeyListener(new KeyListener() {
            @Override
            public void keyTyped(KeyEvent e){

            }
            @Override
            public void keyPressed(KeyEvent e) {
                switch(e.getKeyCode()) {
                    case KeyEvent.VK_UP: elt.setY(elt.getY() - 10); break;
                    case KeyEvent.VK_DOWN: elt.setY(elt.getY() + 10); break;
                    case KeyEvent.VK_LEFT: elt.setX(elt.getX() - 10); break;
                    case KeyEvent.VK_RIGHT: elt.setX(elt.getX() + 10); break;
                }
            }
            @Override
            public void keyReleased(KeyEvent e) {

            }
        });

        // Mouse listener
        panel.addMouseMotionListener(new MouseMotionListener() {
            @Override
            public void mouseDragged(MouseEvent e) {
                if(e.getX() >= elt.getX() && e.getX() <= elt.getX() + elt.getWidth()
                && e.getY() >= elt.getY() && e.getY() <= elt.getY() + elt.getHeight()) {
                    elt.setX(e.getX() - elt.getWidth() / 2);
                    elt.setY(e.getY() - elt.getHeight() / 2);
                }
            }
            @Override
            public void mouseMoved(MouseEvent e) {

            }
        });

        // Main loop
        while(true) {
            try {
                Thread.sleep(5);
                obstacle.setX((obstacle.getX() + 1) % 600);
                Rectangle eltRectangle = new Rectangle(elt.getX(),elt.getY(), elt.getWidth(), elt.getHeight());
                Rectangle obstacleRectangle = new Rectangle(obstacle.getX(),obstacle.getY(), obstacle.getWidth(),obstacle.getHeight());
                if (eltRectangle.intersects(obstacleRectangle))
                    break;
                panel.repaint();
            } catch (InterruptedException e) {
                System.err.print(e.toString());
            }
        }

    }
    public static void main(String[] argv) {
        GameFrame frame = new GameFrame();

    }
}
