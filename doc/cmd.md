# Command

(very) Basic command

## Pacman

-S supports search/query in the sync database  
-Q supports search/query in the local database  

CMD                     | Description
-----------------       | -------------
`pacman -Rs`        | --recursive, Remove package and dependency(not required by other packages)
`pacman -Sg`        | --group, list package groups(sync)
`pacman -Si`        | --info, Extensive search, (sync)
`pacman -Ss`        | --search, Detailed search package, supports regexp(sync)
`pacman -Su`        | --sysupgrade, Update all out-of-date packages
`pacman -Sy`        | --refresh, download a fresh copy of master package database from server, pass two to force
`pacman -F <name>`  | Query the file database, -F means --files
`pacman -Q <name>`  | Search local package, query the package database
`pacman -Q`         | List all local packages
`pacman -Qc `       | --changelog
`pacman -Qd`        | --deps, list packages installed as depedencies
`pacman -Qg`        | --groups, list package groups(local)
`pacman -Qi `       | --info, Extensive search local packages
`pacman -Ql`        | --l, list all files owned by the package
`pacman -Qm`        | --foreign, list packages not found in sync database(from AUR, for example)
`pacman -Qo <file>` | --own, Search for packages that owns this file
`pacman -Qs <regexp>`| --search, Detailed search local packages, supports search by regexp
`pacman -Qt`        | --unrequired, list packages not required by ohter packages
`pacman -Qu`        | --upgrade, list out-of-date packages


Options                 | Description  
--                      | --  
--ignore pkg1,pkg2      | Applies to -S,-U, ignore upgrades even if there is one available

local database: `/var/lib/pacman/local`  
sync database: `/var/lib/pacman/sync`
pacman mirrorlist: `/etc/pacman.d/mirrorlist`  
`pacman -Syyu` to sync modified mirrolist  
pacman configuration file: `/etc/pacman.conf`

If there are signature which is unknown trust, try
`pacman-key --refresh-keys` or `pacman -Sy archlinux-keyring`


### Cache Clean

Cache directory in `/var/cache/pacman/pkg/`

Use `paccache` from `pacman-contrib`

| Cmd            | Description                                                      |
|----------------+------------------------------------------------------------------|
| paccache -ruk0 | Remove all uninstalled package cache                             |
| paccache -r    | Remove package cache, keep cache of latest 3 versions by default |
| paccache -rk2  | 2 specifies keep latest 2 versions only                          |

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


| CMD      | Desc            |
|----------|-----------------|
| `groups` | Show all groups |


### keyboard

| CMD                            | Description         |
|--------------------------------|---------------------|
| `setxkbmap -print -verbose 10` | see X keyboard settings |


## KDE  

### SDDM
Default configuration file: ```/usr/lib/sddm/sddm.conf.d/default.conf```  
Configuration change: ```/etc/sddm.conf.d/```  
Autologin: ```/etc/sddm.conf.d/autologin.conf```  

## Swap Esp and Caps  

`setxkbmap -option caps:swapescape`. Add this to `.xinitrc`
Note that is requires $DISPLAY available

Alternative ways, this set swap permanently:

`loadkeys keys.conf`,


```text
keycode 1 = Caps_Lock
keycode 58 = Escape
```

Linux version of VSC doen't respect remapping with xkbmap. To solve it, 
add `"keyboard.dispath": "keyCode"` in `settings.json`

##  Processes

### Show Processes   

`top` or `htop`  Display processes and memory use
`up/down` to navigate  
`k` to kill  

On top right, it shows info of command `uptime` and number of tasks,threads,kernel threads. 
And how many are running. See explanation of this command for details.

Press `F5` to see process tree(Which is parent proc and which is child proc)
By default, `htop` shows all threads. Press`<S-h>` to toggle thread and `<S-k>` to toggle kernel thread.

Attribute | Description  
----------| ----------
VIRT      | Amount of virtual memory
RES       | Resident size. Amount of physical memory  
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
`ps aux` | 

Example: 
```
root      3339  0.0  0.0 119428   972 pts/19   S+   16:26   0:00 grep --color=auto cmd
root     28020 26.3  0.0 1440780 55712 pts/17  Sl+  15:59   7:05 ./cmd
```
15:59 means the start time of this process. If it started at 24 hours ago, the format will be like "Oct18"

7:05 means the accumulated cpu time (user + system) for this process is 7 minutes and 5 seconds. The format is MMM:SS.
It's not the total running time of the proces up to now.

### Kill Processes
CMD                 | Description
--                  | --
`kill (-9)[pid]`    | kill a process. -9 is SIGKILL, by default
`kill -L`           | list all signals
`killall [pname]`     | kill a process


### nproc

`nproc` returns number of cores in CPU.

### Uptime
Information from `/proc/uptime` and `/proc/loadavg`

Content of `/proc/loadavg`
```shell
0.74 0.56 0.42 1/417 316088
```
Average of number of processes being executed or in waiting state of last 1,5,15 minutes. 1 task running and 417 tasks in total. Last process id is 316088.

Note that one core in CPU can only execute one process at a time.
So if there's only one core in CPU, average load of 1 means 1 process was running on the CPU over last 1 minute, which means CPU was fully utilized.  
If load average is 0.4 for 5 minute period, it means CPU was idle by 60%.
If load average is 3.35 for 15 minute period, it means 2.35 process were waiting for CPU.

If the CPU has two cores, 
1.00 for 1 minute period means 1 core was being used one 1 core was idle.
0.40 for 5 minute period means CPU was idle by 160% on average.
3.35 for 15 minute minutes means 1.35 processes were waiting for CPU.


```shell
$ uptime
 13:10:42 up  4:38,  1 user,  load average: 0.76, 1.19, 0.95
```
which says system has been up for 4h38m.
Load average: average system load over a period of time.

### Process tree

Command `pstree` or `F5` in `htop`



## Git

### Configuration
Path                                                       | Description
--                                                         | --
`/etc/gitconfig``git config --system`                      | config to all users and repos
`~/.gitconfig/ ~/.config/git/config` `git config --global` | config personally to this user and all your repos
`.git/config` in repo `git config --local`                 | specifict to this single repository(default)
`git config --list --show-origin`                          | view all settings
`git config <key>`                                         | check a specific key
`git config --global user.name "sth"`                      |
`git config --global user.email sth`                       | `git config user.name ` to override this setting in a specific repo
`git config --global core.editor vim`                      | set editor


### Useful commands

| Commands                               | Description              |
|----------------------------------------|--------------------------|
| `git branch -r`                        | Show all remote branches |
| `git push origin --delete branch_name` | Delete remote branch     |
| 'git show <commit hash>'              | Show details of a commit |
| 'git push -f origin <branch> '       | force remote commit to be same as local(used to reset remote commit) |




## xrandr  

Used to manage output of a screen including resolution and refresh rate
Can be used to enable multiple monitors.  

Commands                                                                 | Description
--                                                                       | --
xrandr (-q)                                                              | List information of monitors
xrandr --output <MONITOR NAME> --auto                                    | Turn on monitor using default config
xrandr --output <MONITOR NAME> --off                                     | Turn off monitor
xrandr --output <MONITOR NAME> --mode <resolution> --rate <refresh rate> | Change the resolution and refresh rate


## pamixer  
Aka amixer. But can be configured with pacman. A volume level controller

## disk
check disk usage: `df -h`

## curl
Client URL
Used to transfer data to/from a server
`curl [options] [URL...]`


| Command          | Description                                                  |
|------------------|--------------------------------------------------------------|
| curl URL         | Display content of a website(Get request to server)          |
| curl -o FILE URL | Save the downloaded file in local file                       |
| curl -O URL      | Save the downloaded file in a file with the same name as URL |
| curl -i URL      | Get the header of response                                   |
| curl -I URL      | Only get the header of response                              |
| curl -v URL      | verbose, show the header for both request and response       |
| curl -X POST     | change from GET to POST, -X is the same as --request         |

Example

1. Save the website, Note that `https//` can't be ignored
```bash
curl -o kernel.html https://www.kernel.org
```
2. Download a file in the website. Sed txt and pdf example here
```bash
curl -O https://www.gnu.org/software/sed/manual/sed.txt
curl -O https://www.gnu.org/software/sed/manual/sed.pdf
```
3. Redirects, no https/ (-I displays only the request header)
```bash
curl -I www.kernel.org
```
The command will say 301 moved permanently. The useful Location is
presented as `https://www.kernel.org` with Location key in the output

Use `--location (-L)` to redo the request with new position. This will represent the result
```bash
curl -L www.kernel.org
```

4. Combined together
```bash
curl -o -L kernel.html www.kernel.org
```

## cat  

CMD                    | Description
--                     | --
`cat filename`         |
`cat -n filename`      | show with line numbers
`cat > newfile`        | create new file
`cat src1 src2 > dest` | copy content from srcfile to dest
`cat src >> dest`      | append file1 to file2
`cat -s filename`      | no repeated empty line

## Find Command

`find [where to start searching] [expression to find] [-options] [what to find]`


| Cmd                                       | Description                          |
|-------------------------------------------|--------------------------------------|
| find . -name sample.txt                   | Find sample.txt in current directory |
| find . -name *.txt                        | Find *.txt in current directory      |
| find . -name sample.txt -exec rm -i {} \; | Confirm if delete or not             |
| find * -not -type f -regex ".*sth.*\|.*sth2.*" -delete   | Find all files in cur dir and subdir that do not match required regex pattern (and delete them using pipeline)|

Find using multiple conditions, e.g.
```
find . -name \( -name '*.sv' -o -name '*.v' \) -not -name '*parameters.sv'
```
`-o` means `-or`, `-a` means `-and`. Not the use of `\( \)` to enforce precedance, which is necessary.


## Sed Command  
Sed stands for stream editor. Used for inserting, deleting, searching or replacing sth in file and print the result to stdout. It WON'T modify the original file

Commands                    | Description  
--                          | --  
sed 's/unix/linux/' FILE    | Replace unix with linux in the file, only replace the first occurence of the pattern in each line
sed 's/unix/linux/2' FILE   | Replace the second occurence of unix in the file
sed 's/unix/linux/g' FILE   | g:Global. Replace all occurences  
sed 's/unix/linux/3g' FILE  | Replace from 3th occurence to all occurences **in a line**
echo "have a nice day" | sed -E 's/([a-z])/\(\1\)/g' | Add parenthesis to every letter. -E to enable extended regular expression so that () doesn't need backslash before it. \1 to represent the matched pattern
sed '3 s/unix/linux/' FILE  | Replace only the third line
sed '1,3 s/unix/linux/' FILE| Replace from the first line to the third line
sed '2,$ s/unix/linux/' FILE| $ means last line, replace from second line to the last line
sed '5d' FILE               | Delete 5th line of the file
sed '3,6d' FILE             | Delete 3rd to 6th line of the file
sed '/abc/d' FILE           | Delete pattern matching line
sed 'G' FILE                | Insert a blank line after each line
sed '/love/G' FILE          | Insert a blank line after each line with matching pattern love


s: substitution  
/: delimiter


## Awk Command

awk is used for manipulating data.  
`awk [options] 'selection criteria' <input file>`

Note:  
Record: Each record is separated by a record separator, by default the separator is newline
Field: Each field is separated by a field separator, by default the separator is whitespace


Built-in Variables  | Description  
--                  | --  
NR                  | current number of input records/lines
NF                  | current number of fields(separate by whitespace) in current line
FS                  | Current Field Separator, by default whitespace
RS                  | current Record Separator, by default newline
OFS                 | Output Field Separator, by default whitespace, when print several parameters separator with commas, it prints value of OFS between each parameter
ORS                 | Output Record Separator, by default newline. print ORS at the end of whatever it is given to print

Command                               | Description
--                                    | --
awk '{print}' FILE                    | Print the content of the file
awk '/manager/ {print}' FILE          | Print all lines matching the pattern 'manager'
awk '{print $1,$4}' FILE              | Split each line based on whitespace and print the 1st and 4th field
awk '{print NR,$0}' FILE              | Display line number, $0 stands for all contents of current record
awk 'NR==3, NR==6 {print $0}' FILE    | Display from 3rd line to 6th line
awk '{print $1 "-" $2}'               | Field 1 and 2 are separated by "-"
awk 'END {print NR}' FILE             | Count the number of lines in the file
awk 'length($0) > 10' FILE            | Print lines with length larger than 10
awk '{if($3 == "B6") print $0;}' FILE | Find string in specific column


## Cut Command
Cut out sections from each line of files and print to stdout.  
`cut OPTIONS FILES`

Options                | Description
--                     | --
-b                     | Cut by byte
-c                     | Cut by character
-f                     | Cut by Field, tab is default delimiter, use -d to specify delimiter
--complement           | Complement the result
--output-delimiter='$' | Change the delimiter to $ in the output printed result

Commands                | Description  
--                      | --
cut -b 1,2,3 file       | Display byte 1,2,3 of each line
cut -b 1-3,5-7 file     | Display byte 1-3,5-7 of each line
cut -b 1- file          | Display byte from 1st to end
cut -c 1,2,3 file       | Display character 1,2,3 of each line
cut -d " " -f 1 file    | delimiter set to space. Display field 1, i.e. first word of each line

## Head Command & Tail Command
Head: Print top N numbers of file, default 10  
Tail: Print bottom N numbers of file  
It can be used as a filter  
`head OPTIONS FILE`  


Commands                | Description  
--                      | -- 
`head -n 5 file`          | Print first 5 lines of the file
`head -c 3 file`          | Print first 3 character of the file
`head -n 20 file \| tail -10`    | Print lines 10-20




### Process tree

Command `pstree` or `F5` in `htop`



## Git

### Configuration
Path                                                       | Description
--                                                         | --
`/etc/gitconfig``git config --system`                      | config to all users and repos
`~/.gitconfig/ ~/.config/git/config` `git config --global` | config personally to this user and all your repos
`.git/config` in repo `git config --local`                 | specifict to this single repository(default)
`git config --list --show-origin`                          | view all settings
`git config <key>`                                         | check a specific key
`git config --global user.name "sth"`                      |
`git config --global user.email sth`                       | `git config user.name ` to override this setting in a specific repo
`git config --global core.editor vim`                      | set editor


### Useful commands

| Commands                               | Description              |
|----------------------------------------|--------------------------|
| `git branch -r`                        | Show all remote branches |
| `git push origin --delete branch_name` | Delete remote branch     |
| 'git show <commit hash>'              | Show details of a commit |
| 'git push -f origin <branch> '       | force remote commit to be same as local(used to reset remote commit) |




## xrandr  

Used to manage output of a screen including resolution and refresh rate
Can be used to enable multiple monitors.  

Commands                                                                 | Description
--                                                                       | --
xrandr (-q)                                                              | List information of monitors
xrandr --output <MONITOR NAME> --auto                                    | Turn on monitor using default config
xrandr --output <MONITOR NAME> --off                                     | Turn off monitor
xrandr --output <MONITOR NAME> --mode <resolution> --rate <refresh rate> | Change the resolution and refresh rate


## pamixer  
Aka amixer. But can be configured with pacman. A volume level controller

## curl
Client URL
Used to transfer data to/from a server
`curl [options] [URL...]`


Command                     | Description
--                          | --
curl URL                    | Display content of a website(Get request to server)
curl -o FILE URL            | Save the downloaded file in local file
curl -O URL                 | Save the downloaded file in a file with the same name as URL

Example

1. Save the website, Note that `https//` can't be ignored
```bash
curl -o kernel.html https://www.kernel.org
```
2. Download a file in the website. Sed txt and pdf example here
```bash
curl -O https://www.gnu.org/software/sed/manual/sed.txt
curl -O https://www.gnu.org/software/sed/manual/sed.pdf
```
3. Redirects, no https/ (-I displays only the request header)
```bash
curl -I www.kernel.org
```
The command will say 301 moved permanently. The useful Location is
presented as `https://www.kernel.org` with Location key in the output

Use `--location (-L)` to redo the request with new position. This will represent the result
```bash
curl -L www.kernel.org
```

4. Combined together
```bash
curl -o -L kernel.html www.kernel.org
```

## cat  

CMD                    | Description
--                     | --
`cat filename`         |
`cat -n filename`      | show with line numbers
`cat > newfile`        | create new file
`cat src1 src2 > dest` | copy content from srcfile to dest
`cat src >> dest`      | append file1 to file2
`cat -s filename`      | no repeated empty line

## Find Command

`find [where to start searching] [expression to find] [-options] [what to find]`


| Cmd                                       | Description                          |
|-------------------------------------------|--------------------------------------|
| find . -name sample.txt                   | Find sample.txt in current directory |
| find . -name *.txt                        | Find *.txt in current directory      |
| find . -name sample.txt -exec rm -i {} \; | Confirm if delete or not             |
| find * -not -type f -regex ".*sth.*\|.*sth2.*" -delete   | Find all files in cur dir and subdir that do not match required regex pattern (and delete them using pipeline)|

Find using multiple conditions, e.g.
```
find . -name \( -name '*.sv' -o -name '*.v' \) -not -name '*parameters.sv'
```
`-o` means `-or`, `-a` means `-and`. Not the use of `\( \)` to enforce precedance, which is necessary.


## Sed Command  
Sed stands for stream editor. Used for inserting, deleting, searching or replacing sth in file and print the result to stdout. It WON'T modify the original file

Commands                    | Description  
--                          | --  
sed 's/unix/linux/' FILE    | Replace unix with linux in the file, only replace the first occurence of the pattern in each line
sed 's/unix/linux/2' FILE   | Replace the second occurence of unix in the file
sed 's/unix/linux/g' FILE   | g:Global. Replace all occurences  
sed 's/unix/linux/3g' FILE  | Replace from 3th occurence to all occurences **in a line**
echo "have a nice day" | sed -E 's/([a-z])/\(\1\)/g' | Add parenthesis to every letter. -E to enable extended regular expression so that () doesn't need backslash before it. \1 to represent the matched pattern
sed '3 s/unix/linux/' FILE  | Replace only the third line
sed '1,3 s/unix/linux/' FILE| Replace from the first line to the third line
sed '2,$ s/unix/linux/' FILE| $ means last line, replace from second line to the last line
sed '5d' FILE               | Delete 5th line of the file
sed '3,6d' FILE             | Delete 3rd to 6th line of the file
sed '/abc/d' FILE           | Delete pattern matching line
sed 'G' FILE                | Insert a blank line after each line
sed '/love/G' FILE          | Insert a blank line after each line with matching pattern love


s: substitution  
/: delimiter


## Awk Command

awk is used for manipulating data.  
`awk [options] 'selection criteria' <input file>`

Note:  
Record: Each record is separated by a record separator, by default the separator is newline
Field: Each field is separated by a field separator, by default the separator is whitespace


Built-in Variables  | Description  
--                  | --  
NR                  | current number of input records/lines
NF                  | current number of fields(separate by whitespace) in current line
FS                  | Current Field Separator, by default whitespace
RS                  | current Record Separator, by default newline
OFS                 | Output Field Separator, by default whitespace, when print several parameters separator with commas, it prints value of OFS between each parameter
ORS                 | Output Record Separator, by default newline. print ORS at the end of whatever it is given to print

Command                               | Description
--                                    | --
awk '{print}' FILE                    | Print the content of the file
awk '/manager/ {print}' FILE          | Print all lines matching the pattern 'manager'
awk '{print $1,$4}' FILE              | Split each line based on whitespace and print the 1st and 4th field
awk '{print NR,$0}' FILE              | Display line number, $0 stands for all contents of current record
awk 'NR==3, NR==6 {print $0}' FILE    | Display from 3rd line to 6th line
awk '{print $1 "-" $2}'               | Field 1 and 2 are separated by "-"
awk 'END {print NR}' FILE             | Count the number of lines in the file
awk 'length($0) > 10' FILE            | Print lines with length larger than 10
awk '{if($3 == "B6") print $0;}' FILE | Find string in specific column


## Cut Command
Cut out sections from each line of files and print to stdout.  
`cut OPTIONS FILES`

Options                | Description
--                     | --
-b                     | Cut by byte
-c                     | Cut by character
-f                     | Cut by Field, tab is default delimiter, use -d to specify delimiter
--complement           | Complement the result
--output-delimiter='$' | Change the delimiter to $ in the output printed result

Commands                | Description  
--                      | --
cut -b 1,2,3 file       | Display byte 1,2,3 of each line
cut -b 1-3,5-7 file     | Display byte 1-3,5-7 of each line
cut -b 1- file          | Display byte from 1st to end
cut -c 1,2,3 file       | Display character 1,2,3 of each line
cut -d " " -f 1 file    | delimiter set to space. Display field 1, i.e. first word of each line

## Head Command & Tail Command
Head: Print top N numbers of file, default 10  
Tail: Print bottom N numbers of file  
It can be used as a filter  
`head OPTIONS FILE`  


Commands                | Description  
--                      | -- 
`head -n 5 file`          | Print first 5 lines of the file
`head -c 3 file`          | Print first 3 character of the file
`head -n 20 file \| tail -10`    | Print lines 10-20

## Grep

`grep -e "sth"`, with `-e`, |, +, ? etc needs to be escaped
`grep -E "sth`, with -E, no need to escape




## nohup

no hang up

keep the process running after exiting terminal. Useful for running a process on server 

A common usage:
`nohup ./a.out > output.log 2>&1 &`



## Read system info

architecture: `uname -a`
cpu: `lscpu`
cpu usage: `uptime` 1, 5 , 15 minutes interval
disk: `df -h` or `lsblk` or `fdisk -l`
PCI peripheries: `lspci`
memory: `free -h`

Or `dmidecode` 






