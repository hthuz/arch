# DevOp Quick Ref

## Top

top - 19:48:06 up 90 days, 58 min,  4 users,  load average: 0.66, 0.84, 1.31
Tasks: 258 total,   1 running, 257 sleeping,   0 stopped,   0 zombie
%Cpu(s): 12.5 us,  2.3 sy,  0.0 ni, 82.8 id,  0.8 wa,  0.8 hi,  0.8 si,  0.0 st
MiB Mem :  31877.5 total,   2338.7 free,   8298.8 used,  22529.0 buff/cache
MiB Swap:   2000.0 total,   1254.9 free,    745.1 used.  23578.7 avail Mem 

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                                      
1408886 root      20   0 3066072 155200  21504 S  86.7   0.5  55:02.06 mevanalysis_old                                              
1056743 root      20   0 1307812  26932  10496 S  26.7   0.1 243:25.63 sandwichwatch                                                
1318564 root      20   0 1396180  98476  49024 S   6.7   0.3   0:34.25 node


```
top - 19:48:06 up 90 days, 58 min,  4 users,  load average: 0.66, 0.84, 1.31
```
system running for 90 days, 4 active users, load averge for 1 minute, 5 minute, 10 minutes

load average: average number of processes running


```
%Cpu(s): 12.5 us,  2.3 sy,  0.0 ni, 82.8 id,  0.8 wa,  0.8 hi,  0.8 si,  0.0 st
```
user occupy 12.5% of cpu, system 2.3%, idel 82.8%, waiting IO 0.8%, hardware IRQ 0.8%, software IRQ 0.8%


## ps

```
ps aux
```
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0 175992 14128 ?        Ss   May08  30:33 /usr/lib/systemd/systemd --system --deserialize 40
root           2  0.0  0.0      0     0 ?        S    May08   0:05 [kthreadd]

## systemd

In `/etc/systemd/system/myapp.service`
```bash
[Unit]
Description=My App Service
After=network.target

[Service]
ExecStart=/usr/local/bin/myapp
Restart=always
RestartSec=3
User=nobody
# Optional
WorkingDirectory=/usr/local/bin
# Optional
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```
```bash
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl start myapp
```

## journalctl

```bash
journalctl --no-pager   # without pager, just like print to stdout
journalctl -u myapp -f  # -u <unit> specified app
journalctl -n <N>       # check lastest N lines of log
journalctl -p <level>   # print all levels above) emerg 0, alert 1, crit 2, err 3, warning 4, notice 5, info 6, debug 7
journalctl -r           # check in reverse order
journalctl -b           # only this boot
journalctl --grep="pattern"
journalctl --grep="(Started|Stopped)" # check start/stop log
journalctl -u myapp | grep "Stopped" -B 10 # check 10 lines before service Stopped, useful for checking restart information
```

## firewall

yum intsall firewalld
firewall-cmd --zone=public --add-port=80/tcp --permanent
firewall-cmd --zone=public --remove-port=8080/tcp --permanent
firewall-cmd --reload
firewall-cmd --list-all

## scp

scp <src> <dst>
root@ip:my/path/to/file

scp temp.txt root@ip:/home/desktop
scp -r ./somedir root@ip:/home/desktop

## lscpu


Number of cores: 
```
CPU(s):                   8
  On-line CPU(s) list:    0-7
```
or 
```
    Model:                60
    Thread(s) per core:   2
    Core(s) per socket:   4
    Socket(s):            1

```
socket: number of physical CPUs




## shell

`Ctrl+R`: search based some string of command

Case 1: press `Enter` to execute directly
Case 2: Keep pressing `Ctrl+R` to choose eariler command 
Case 3: press `->` or `Esc` to go back to normal cmdline and modify

If move to fast , use `Ctrl+S` to search back, with config
```
stty -ixcon
```
in .bashrc

Ctrl+R: reverse search
Ctrl+S: forward search






