
import numpy as np
import matplotlib.pyplot as plt

def f(x):
    return x**2 - x - 2

class Nonlinear:

    def draw(self):
        x = np.linspace(-5,5,20);
        plt.grid()
        plt.plot(x,f(x))
        plt.savefig("./nonlinear.png")

    def interval_bisection(self):
        a = 0
        b = 5
        step = 0
        while (b - a > 1e-3):
            step += 1
            print(f"step:{step}, a:{a}, b:{b}")
            m = a + (b - a) / 2
            if(np.sign(f(a)) == np.sign(f(m))):
                a = m
            else:
                b = m
        print(f"result: {a}")
        return


if __name__ == "__main__":
    nonlinear = Nonlinear()
    nonlinear.interval_bisection()
    nonlinear.draw()



