# Knowledge

Some useful knowledge, usually some explanations

## Booting Process

BIOS->Boot Loader->Kernel Initialization->init(or systemctl) process

init will start other processes.
From the user side, example as follows(generated by `pstree`)

```shell
systemd─┬─NetworkManager───3*[{NetworkManager}]
        |-login-bash
        ├─dbus-daemon
        ├─systemd-journal
        ├─systemd-logind
        ├─systemd-udevd
        └─wpa_supplicant
```

`bash` will load `.bash_profile` and `.bashrc`.

For example, if I run `chadwm`, orders are as follows:
bash->startx(in `.bash_profile`, exec cmds in `.xinitrc`)->chadwm start script


## Permission  
`ls -l filename`  show permission of a file  

-rw-r--r--  

Part    | Description
--      | -- 
`-`       | type of file(d for directory, s for special, - for regular file)
`rw-`     | owner's permission to file
`r--`     | members of the same group's permission
`r--`     | all other users' permission


## x86 Assembly  
```text
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


## Regular Expression

Example 1:
```bash
$echo 'abc' | sed 's/./xyz/g'
xyzxyzxyz  
```
Here a dot means **any character**  

Example 2:  
```bash
$echo 'abc' | sed 's/\./xyz/g'
abc
```
Use back slash to avoid interrept dot as regular expression

Example 3:  
```bash
$echo 'a..b..c' | sed 's/[\.b]/d/g'
adddddc
$echo 'a..b..c' | sed 's/[\.b]\+/d/g'
adc
$echo 'a..b ..c' | sed 's/[\.b]\+/d/g'
ad dc
$ echo 'a..b..c' | sed 's|[\.b]\+|d|g; s|[a-c]|d|g'
ddd
```
`[\.b]`: Any literal dot or character b  
`\+`: Look for at least one, possibly more of listed characters and the matched pattern will be considered as only a single occurence. Here `\+` is not a literal `+`.  
`[a-c]`: A range of letters from a to c  


Example 4:
```bash
$echo "have a nice day" | sed 's/$/ you all/'
have a nice day you all
```
`$`: means end of line

Example 5:  
```bash
$echo "sample+" | sed 's/[a-e]+/_/g'
sampl_  # Here + is just a +
$echo "sample+" | sed -E 's/[a-e]+/_/g'
s_mpl_+ # Here + is not a +
```
-E: Enable Extended Regex.  In extended regex, + just means /+

Example 6:
```bash
$echo "abcdefghijklmnopqrstuvwxyz ABCDEFG 0123456789" | sed -E 's/([a-o]+).*([A-Z]+)/\2 \1/'  
G abcdefghijklmno 0123456789
```
`.*`:Any character, 0 or more times.It starts matching after character o and it will keep matching characters until the last A-Z are matched. This part is not captured by sed
`\2 \1`: Change the order of lower-case letters(first search pattern) and upper-case letters( second search pattern )

Example 7 (Easier version of Example 6):  
```bash
$echo "have a nice day " | sed 's/e.*e//g'
hav day
$echo "have a nice day " | sed 's/e.*//g'
hav
$echo "have a nice day " | sed 's/.*e//g'
 day
$echo "have a nice day " | sed 's/.*[a-d]//g'
y
$echo "have a nice day " | sed 's/e.*[a-d]//g'
havy 
```
Example 8:  
```bash
$echo "abcdefghijklmnopqrstuvwxyz ABCDEFG 0123456789" | sed -E 's/[^ ]*/_/'
_ ABCDEFG 0123456789
```
`[^ ]*`: Match any non-space character, 0 or more times. ^ in [] means not

Example 9:
```bash
$ echo "have a nice day" | sed -E 's/[^ ]+ [^ ]+//'
 nice day
```
`[^ ]+ [^ ]+` means the pattern of no-space space no-space  

**Summary**

Metacharacters | Example         | Description
--             | --              | --
[]             | [a-m],[0-9AF-Z] | A set of characters
[^ ]           | [^A-Za-z]       | Outside the selected range
.              | he..o           | Any character
\*              | he.*o*          | Zero or more occurences, usually followed by other metacharacters like .*
\+              | he.+o           | One or more occurences
?              | he.?o           | Zero or one occurences
{}             | he.{2}o         | Exactly specified number of occurences
^              | ^hello          | Start of the string
$              | planet$         | End of the string
`|`            | `fall|stays`    | Either or
\d             |                 | One digit
\D             |                 | One non-digit
\s             |                 | One whitespace
\S             |                 | One non-whitespace
\b             |                 | Backspace character
\n             |                 | Newline character
\t             |                 | Tab character


### Greedy match and Non-greedy match

Taking Python re as an example, `.*` is greedy match, trying to match the longest possible sequence of characters that are allowed to match   
`.*?` is non-greedy match, trying to match shortest possible sequence of characters that are allowed to match

For example
```
text = "Hello World"
H.*o => Hello Wo
H.*?o => Hello
```


## disk
Check disk usage: `df -h`

## RAM

- RAM
    - DRAM (Dynamic RAM) uses capacitor. Needs refreshing. High capacity, speed low, cost low. The refreshing takes time so it slows down. Used as main memory
        - ADRAM (Asynchronous DRAM). No synchronization so there may delay in memory response.
        - SDRAM (Synchronous DRAM). System clock Synchronizes memory access.
    - SRAM (Static RAM), uses flip-flop. Long term storage. Low capacity, speed high, cost high. Used as cache

Dynamic - Refreshing - Slow


## LaTex

### Commonly used package

tikz: create graphic elements  
graphicx: include figures


### Common fonts

| Commands                                  | Description               |
|-------------------------------------------|---------------------------|
| \renewcommand{\familydefault}{\rmdefault} | Set default font to roman |
| \renewcommand{\familydefault}{\sfdefault} | Set default font to sans  |


| Font                   | Command   | Switch Command | Decription       |
|------------------------|-----------|----------------|------------------|
| roman serif            | \textrm{} | \rmfamily      | Default for text |
| sans serif             | \textsf{} | \sffamily      |                  |
| typewritter(monospace) | \texttt{} | \ttfamily      |                  |



## Escape Character  

Many characters have special meaning. To treat this character as normal characer, add \
Shell as example, which I often confuse with other things.  

| Command                     | Output                                                          |
|-----------------------------|-----------------------------------------------------------------|
| `mkdir double dir`          | Will create two directories                                     |
| `mkdir double\ dir`         | Will create one directory named 'double dir', '\ ' escape space |
| `echo "\"`                  | Won't echo anything                                             |
| `echo "\\"`                 | print `\`                                                       |
| `echo "called "amazing""`   | print "called amazing"                                          |
| `echo "called \"amazing\""` | print "called "amazing"                                         |
| `echo what's`               | Won't work                                                      |
| `echo what\'s`              | Print "what's"                                                  |
| `echo "what's"`             | Print "what's"                                                  |
| `grep -R that's .`          | Won't work                                                      |
| `grep -R that\'s .`         | Works properly
| `grep -R "that's" .`        | Works properly                                                  |



## Architecture ISA

1. x86: Intel & AMD
	Processors include Intel 8086, 80386 etc
	Intel 80386 is 32bits. Aka x86_32
	64 bit processor with x86 architecture: aka x86_64 or x64 or AMD64


2. ARM: Apple
3. RISV-V:

i386 or 386: microprocessor Intel 80386
i686 : 32 bit Intel x86 arch


## Boot

BIOS:two mode
	- Legacy, MBR partition type
	- UEFI, GPT disk partition type, EFI: system partition in GPT (aka ESP)


## File Systems

- FAT (File Allocation Table): classical in Windows. FAT16, FAT32 are variants of FAT
- NFTS: most widely used in Windows
- ext4: in Linux

