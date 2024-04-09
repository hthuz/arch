
import java.util.Timer;
import java.util.TimerTask;
public abstract class Utils {
    
    public static void task(int period, ITimer task) {
        Timer timer = new Timer();
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                if(GameStatus.gameover) {
                    timer.cancel();
                }
                task.run();
            }
        }, 0, period);
    }
}
