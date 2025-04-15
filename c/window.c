
#include <stdio.h>
#include <sys/types.h>
#include <sys/ioctl.h>
#include <unistd.h>
#include <termios.h>
#include <fcntl.h>

int main() {

	struct winsize win;
	// int fd;
	// fd = open();
	ioctl(STDIN_FILENO, TIOCGWINSZ, &win);
	printf("%d\n",win.ws_col);
	printf("%d\n",win.ws_row);

}
