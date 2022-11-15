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


## KDE  

### SDDM
Default configuration file: ```/usr/lib/sddm/sddm.conf.d/default.conf```  
Configuration change: ```/etc/sddm.conf.d/```  
Autologin: ```/etc/sddm.conf.d/autologin.conf```  



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

Process Status command, display running processes
CMD     | Description
--      | --
`ps`    | start
`ps -u` | more information
`ps -A` | display all processes

### Kill Processes
CMD                 | Description
`kill (-9)[pid]`    | kill a process
`kill -L`           | list all signals
