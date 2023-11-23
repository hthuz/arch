
#include <sys/time.h>
#include <iostream>

// For sleep
#include <thread>
#include <chrono>
using namespace std;

int main()
{
    // 1 s = 1000 milliseconds = 1000,000 microseconds
    struct timeval cur_time;
    gettimeofday(&cur_time, nullptr);
    time_t milliseconds = cur_time.tv_sec * 1000 + cur_time.tv_usec / 1000;
    time_t microseconds = cur_time.tv_sec * 1000000 + cur_time.tv_usec;
    cout << "milliseconds: " << milliseconds << endl;
    cout << "microseconds: " << microseconds << endl;

    int arr[1000];
    for(int i = 0; i < 1000; i++) {
        arr[i] = i * i - i;
    }
    this_thread::sleep_for(chrono::seconds(1));

    gettimeofday(&cur_time, nullptr);
    milliseconds = cur_time.tv_sec * 1000 + cur_time.tv_usec / 1000;
    microseconds = cur_time.tv_sec * 1000000 + cur_time.tv_usec;
    cout << "milliseconds: " << milliseconds << endl;
    cout << "microseconds: " << microseconds << endl;

}
