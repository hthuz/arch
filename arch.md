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
pacman configuration file: `/etc/pacman.conf`

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
add env var in /etc/profile     | persistent env var(for this profile)

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

Apply patches 'patch --merge -i <patch.diff>'


## xorg



### xinit
Used to start window managers  
Default configuration file: `etc/X11/xinit/xinitrc`  
Configuration file in home directory :`~/.xinitrc`  
Note commands after `exec` won't be executed  

To start, run `startx` or set auto start 

#### Autostart
Add the following in `~/.bash_profile`
```
if [-z "${DISPLAY}" && ["${XDG_VTNR}" -eq 1]] then
    exec startx
fi
```


### xorg-xinput  
A tool to manage input devices
CMD                             | DEScription  
--                              | --  
xinput list                     | show devices


## libinput  
A library to handle input devices  
Default configuration file is in `/usr/share/X11/xorg.conf.d/40-libinput.conf`

CMD                             | Description
--                              | --  
libinput list-devices           | show devices

**Note**: this method can't change touchpad to natural scrolling for my laptop

## xf86-input-synaptics

Default configuration file: `/usr/share/X11/xorg.conf.d/70-synaptics.conf`  
Copy to `/etc/X11/xorg.conf.d/` and edit here

### Touchpad Natural Scorrling  
In `70-synaptics.conf` file, add following to identifer 'touchpad catchall'
```
Option "VertScrollDelta" "-111"
Option "HorizScrollDelta" "-111"
```
### Touchpad one-click confirm
```
Option "TapButton1" "1"
```

## Battery   
Check battery: acpi

Get: fetch data from specified resources  
Post: submit data to a specified resource


## VScode  
Format code: `Ctrl + Shift + I`

## User Management  
Information about all users: `/etc/passwd`
Information about all groups: `/etc/group`

Each user has a default primary group and some other secondary groups(15 max)


CMD                                | Description
--                                 | --
`id username`                      | get id of user
`useradd username`                 | add user
`passwd username`                  | assign/change passwd to a user
`usermod -u new_id username`       | change id of a user
`usermod -g new_group_id username` | change group of a user
`usermod -l new_name old_name`     | only change login name
`usermod -d new_home_dir username` | change home direcotry of a user
`usermod -r username`              | delete a user
`groups username`                  | get user's group format:(primary secondary)
`usermod -a -G sec_group username` | add a secondary group to user




## Permission  
`ls -l filename`  show permission of a file  

-rw-r--r--  

Part    | Description
--      | -- 
-       | type of file(d for directory, s for special, - for regular file)
rw-     | owner's permission to file
r--     | members of the same group's permission
r--     | all other users' permission





### keyboard

CMD                 | Description
`setxkbmap -print -verbose 10` | see X kb settings


## KDE  

### SDDM
Default configuration file: ```/usr/lib/sddm/sddm.conf.d/default.conf```  
Configuration change: ```/etc/sddm.conf.d/```  
Autologin: ```/etc/sddm.conf.d/autologin.conf```  

## Swap Esp and Caps  
`setxkbmap -option caps:swapescape`. Add this to `.xinitrc`


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


## Makefile  
```
targets: prerequisites
    command
    command
    command
```

## picom  
Default configuration file: `/etc/xdg/picom.conf`

## xwallpaper  
`xwallpaper --zoom filename`

## ranger
Copy default configuration to `~/.config/ranger`: `ranger --copy-config=all`

eile            | Description  
`rc.conf``      | startup cmds and key binginds  
`commands.py`   | commands launced with :
`rifle.conf`    | application used when given type of file is launced

## cs411 instwhoance ip address


## MariaDB
`mariadb -u root (-h) -p`
34.66.142.241
 

## cat  

CMD                    | Description
--                     | --
`cat filename`         |
`cat -n filename`      | show with line numbers
`cat > newfile`        | create new file
`cat src1 src2 > dest` | copy content from srcfile to dest
`cat src >> dest`      | append file1 to file2
`cat -s filename`      | no repeated empty line


## x86 Assembly  
```
.data                      # readable/writable
data_item:                 # data_items points to 3
.long 3,5,6,7              # declare an array of numbers with 4bytes(consecutive)

.text                      # readable/executable
 
.global _start

_start:                    # entry to instruction
    
```
.quad: 8bytes  
.word: 2bytes

```
.rept 3
.long 0
.endr           # repeat sequence of lines 3 times
```
equivalent to 
```
.long 0
.long 0
.long 0
```



## Bash

### Commands  
Every part separated is called **word** 
Arguments are separated by blankspace with command, `[]` is also a command'
`[-f file]` (wrong)   
`[ -f file ]` (correct)  

In Bash, almost everything is string (command, argumetns etc)

### Types of commands  
**Alias**: A way of shortening a cmd, only in interactive shell  
**Functions**  
**Builtins**: builtin commands 
**Keywords**: part of Bash syntax
**Executable**  



