
import numpy as np
import matplotlib.pyplot as plt
from matplotlib import cm

def f1(x: np.ndarray) -> int:
    A = np.array([[3,3],
                  [3,4]])
    return np.transpose(x) @ A @ x
def f2(x1: int,x2: int) -> int:
    return 3*x1**2 + 4*x2**2 + 6*x1*x2


def draw():
    x1 = np.linspace(-5,5,21)
    x2 = np.linspace(-5,5,21)
    X1,X2 = np.meshgrid(x1,x2)
    y = f2(X1,X2)

    fig,ax = plt.subplots(subplot_kw={"projection":"3d"})
    ax.plot_surface(x1,x2,y,alpha=0.5)
    plt.savefig("result.jpg")




if __name__ == "__main__":
    x = np.transpose(np.array([[1,2]]))
    f1(x)
    draw()

