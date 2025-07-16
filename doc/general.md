# General

General staff

## Environment Variables  

| CMD                             | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `printenv <var>`                | print one variable                                                          |
| `printenv`                      | print all environment variables                                             |
| `set`                           | print all variables(including shell var etc)                                |
| `unset <var>`                       | Delete a environment variable |
| `set <myvar>`                   | set a var(not env)                                                          |
| `export <myvar>`                | set <myvar> to env, only change env var of my child processes               |
| `export MYVAR=content`          | directly set env var (temporary)                                            |
| add env var in /etc/environment | persistent env var(system-wide)                                             |
| add env var in /etc/profile     | persistent env var(user-wide), need to source this file         |
| add script in /etc/profile.d    | /etc/profile.d is a dir containing script export PATH                       |
| add env var in /etc/bashrc      | env var for bash shell only                                                 |
| add env var in ~/.bash_profile  | env var for bash shell only, this file is executed only once on sys startup |

Environment varialbe $PATH: list of directories to be searched when executing command  

In command-line, you can add `VAR=value` in front of the command to temporarily change the environmental variable for this command. Example:`LANG=C ls .`

To remove one directory in $PATH, run
```bash
PATH=$(echo "$PATH" | sed -e 's/:\/home\/autentico\/sthsth$//')
```


## DWM  

Better to Add:  st, dmenu, dwm  
Use git clone and make install
/usr/share/xsessions/dwm.desktop content:  
```txt
[Desktop Entry]
Encoding=UTF-8
Name=dwm
Comment=Dynamic Window Manager  
Exec=dwm
Icon=dwm
Type=xsession
```

Default Shortcut                    | Description  
--                          | --
shift+alt+enter             | open st
shift+alt+c                 | close window
shift+alt+q                 | quit dwm
alt+j/k                     | change window
alt+p                       | open dmenu
alt+enter                   | set current window as master window  
alt+m/t                     | change full screen
alt+1-9                     | change tag             
alt+f                       | change current window to floating mode
alt + m                     | change current window to full screen
alt + t                     | recover mode
alt + i                     | Change layout of clients to equal mode

Apply patches 'patch --merge -i <patch.diff>'

## xorg

Check if it's X11 or wayland: `echo $XDG_SESSION_TYPE`

### xinit
Used to start window managers  
Default configuration file: `/etc/X11/xinit/xinitrc`  
Configuration file in home directory :`~/.xinitrc`  
Note commands after `exec` won't be executed  

To start, run `startx` or set auto start 

#### Autostart
Add the following in `~/.bash_profile`
```txt
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
`sudo pacman -S xf86-input-synaptics`

Use synaptics driver for touchpad.

Synaptics is a human interface hardware and software manufacturer.

Default configuration file: `/usr/share/X11/xorg.conf.d/70-synaptics.conf`  
Copy to `/etc/X11/xorg.conf.d/` and edit here

### Touchpad Natural Scorrling  
In `/etc/X11/xorg.conf.d/70-synaptics.conf` file, add following to identifer 'touchpad catchall'
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
Or using the command `cat /sys/class/power_supply/BAT0/capacity`


## Font  

System-wide font directory:'/usr/local/share/fonts'  
System-wide font directory:`/usr/share/fonts`, for pacman, shouldn't be modifies manually

ttf is a kind of font

Make sure you have `pacman -S fontconfig` installed.
And then you can use `fc-list` to list all fonts

## Setfont

`setfont` to set font in cli mode.
Example as follows:
`setfont /usr/share/kbd/consolefonts`



## picom  

Default configuration file: `/etc/xdg/picom.conf`

## xwallpaper  
`xwallpaper --zoom filename`

## PDF  
 application/pdf)
    pdftoppm -f 1 -l 1 \


## imagemagick
Image viewer in terminal

`display filename`


## MariaDB
`mariadb -u root (-h) -p`

## fcitx5



## Pip

Pip externally-managed-environment error:
Remove `/usr/lib/python3.x/EXTERNALLY-MANAGED`

## python

pip install path: `~/.local/lib/python<version>/site-packages/`
if some packages provide a binary executable, its path is `~/.local/bin`


## Du

Estimate file space usage
```
du -sh
```

| Common flags | Description                     |
|--------------|---------------------------------|
| -a           | Count for all files             |
| -h           | --human-readable                |
| -s           | --summarize. Only display total |



## Markdown

Some common markdown supported language
- bash(*.sh)
- c(*.c,*.h)
- cpp(*.cpp)
- css
- html
- ini(*.ini)
- java
- js(*.js)
- lua(*.lua)
- make(Makefile)
- python
- sql(*.sql)
- tex(*.tex)
- text(*.txt)
- vim(*.vim,.vimrc)


## Trash

directory of Trash: `~/.local/share/Trash/`


## C 
C headfer files located in `/usr/include/`


## lspci

| CMD        | Comment                                   |
|------------|-------------------------------------------|
| `lspci`    | list pci devices                          |
| `lspci -k` | show kernel drivers handling each devices |


## cisco

dependency for cisco:
`libsoup`
`webkit2git`




## XDG

X Desktop Group, specifications for interoperability
XDG Base Directory Specification (XDG BDS)


User directories

| Vars            | Desc                             | Default            |
|-----------------|----------------------------------|--------------------|
| XDG_CONFIG_HOME | user-specific configuration      | $HOME/.config      |
| XDG_CACHE_HOME  | user-specific cached data        | $HOME/.cache       |
| XDG_DATA_HOEME  | user-specific data files         | $HOME/.local/share |
| XDG_STATE_HOME  | user-specific state files        | $HOME/.local/state |
| XDG_RUNTIME_DIR | user-specific sockets, pipes etc | /run/user/$UID     |

System directories

| Vars            | Desc         | Default                     |
|-----------------|--------------|-----------------------------|
| XDG_DATA_DIRS   | data files   | /usr/local/share:/usr/share |
| XDG_CONFIG_DIRS | config files | /etc/xdg                    |



