# Arch Linux Note

## Pacman  

CMD                     | Description
-----------------       | -------------
```pacman -Rs```        | Remove package and dependency 
```pacman -Ss <name>``` | Search remote package
```pacman -Q <name>```  | Search local package
```pacman -Q```         | List all local packages
```pacman -Qs ```       | Detailed search local packages
```pacman -Qi <name>``` | Extensive search local packages

pacman mirrorlist: `/etc/pacman.d/mirrorlist`  
`pacman -Syyu` to sync modified mirrolist

## systemmd  

PID 1, system and service manager  
Units: service, mount, device, socket  

CMD                                   | DESCRIPTION
--                                    | --
`systemctl status`                    |
`systemctl list-units`                |
`systemctl list-units --type=service` | list service only
`systemctl start unit`                | start a unit immediately
`systemctl stop unit`                 | stop a unit immediately
`systemctl enable unit`               | start automatically at boot
`systemctl enable --now unit`         | start automatically at boot and start it now
`systemctl disable unit`              | no longer start at boot
`systemctl mask unit`                 | impossible to start
`systemctl unmask unit`               |


## Environment Variables  

CMD                             | Description
--                              | --
`printenv <var>`                | print one variable
`printenv`                      | print all environment variables
`set`                           | print all variables(including shell var etc)
`set <myvar>`                   | set a var(not env)
`export <myvar>`                | set <myvar> to env
`export MYVAR=content`          | directory set env var (temporary)
add env var in /etc/environment | persistent env var(system-wide)

Environment varialbe $PATH: list of directories to be searched when executing command  



## DWM  
Require:  st, dmenu, dwm  
Use git clone and make install
/usr/share/xsessions/dwm.desktop content:  
```
[Desktop Entry]
Encoding=UTF-8
Name=dwm
Comment=Dynamic Window Manager  
Exec=dwm
Icon=dwm
Type=xsession

```

Shortcut                    | Description  
--                          | --
shift+alt+enter             | open st
shift+alt+c                 | close window
shift+alt+q                 | quit dwm
alt+j/k                     | change window
alt+p                       | open dmenu
alt+enter                   | set current window as master window  
alt+m/t                     | change full screen
alt+1-9                     | change tag             









## KDE  

### SDDM
Default configuration file: ```/usr/lib/sddm/sddm.conf.d/default.conf```  
Configuration change: ```/etc/sddm.conf.d/```  
Autologin: ```/etc/sddm.conf.d/autologin.conf```  

## Swap Esp and Caps  
`setxkbmap -option caps:swapescape`


##  Processes
### Show Processes   
```top```  Display processes and memory use
```up/down``` to navigate  
```k``` to kill  
Attribute | Description  
----------| ----------
VIRT      | Amount of virtual memory
RES       | Amount of physical memory  
SHR       | Amount of memory shared  
S         | State of process
TIME+     | Total CPU time consumed by process

State   | Description
--      | --
D       | uninterruptible sleep
R       | running  
S       | sleeping  
T       | traced or stopped  
Z       | zombie

Process Status command, display current running processes only
CMD     | Description
--      | --
`ps`    | start
`ps -u` | more information
`ps -A` | display all processes

### Kill Processes
CMD                 | Description
--                  | --
`kill (-9)[pid]`    | kill a process
`kill -L`           | list all signals


## Vim  

### ColorScheme
`:colorscheme <scheme>`  Change color scheme  
Options:
blue       delek     evening    morning   peachpuff   slate  
darkblue   desert    industry   murphy    ron         torte  
default    elflord   koehler    pablo     shine       zellner  



