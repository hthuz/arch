
import java.awt.event.KeyListener;

import javax.swing.*;
import java.awt.event.KeyEvent;
public class GameFrame extends JFrame {
    public GameFrame() {
        this.setTitle("Snake Game");
        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        this.setBounds(Parameters.FRAME_X,Parameters.FRAME_Y,Parameters.FRAME_WIDTH, Parameters.FRAME_HEIGHT);
        Snake snake = new Snake(Parameters.BASE_SIZE * 5, Parameters.BASE_SIZE * 5);
        Food food = new Food(Parameters.BASE_SIZE * 7, Parameters.BASE_SIZE * 7);
        GamePanel panel = new GamePanel(snake, food);
        this.add(panel);
        this.setVisible(true);

        // Keyboard Listener
        this.addKeyListener(new KeyListener() {
            @Override
            public void keyTyped(KeyEvent e){
            }
            @Override
            public void keyPressed(KeyEvent e) {
                Key.add(e.getKeyCode());
            }
            @Override
            public void keyReleased(KeyEvent e){
                Key.remove(e.getKeyCode());
            }
        });

        // Main loop
        while(true) {
            try{
                Thread.sleep(1000);
                snake.action();
                snake.eat(food);
                panel.repaint();
            }
            catch(InterruptedException e){
                System.err.println(e.toString());
            }
        }
    }

    public static void main(String[] argv) {
        GameFrame game = new GameFrame();

    }
    
}
