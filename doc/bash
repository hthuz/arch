# Bash

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

### Empty Space
Use double quote if there's whitespace belongs to one word
`var=a b` is wrong and `var="a b"` is correct


### Special Characters  
Character | Description
--        | --
`[[  ]]`  | *Test* - evaluation of a conditional expression
`{}`      | *Inline group* - cmds inside are treated as one cmd
`()`      | *Subshell group* - cmds insde are executed in a subshell
`((  ))`  | *Arithmetic expression* - +,-,* inside are treated as operator
`$((  ))` | the expression will be replaced with the reuslt of arithmetic evaluation  
`\`       | *Escape* - Prevents next character from being interpreted as a special character.

### Quotation

Double quotes can be used for strings containing spaces. Some times, it's not necessary to wrap a variable with double quotation like `"$var"`,but it is always a good idea to do so.  
Single quote will cause Bash to interpret string literally  
Back quote allow us to execute command  

```bash
echo $var         - > hello world
echo "$var"       - > hello world (a better habit)
echo '$var'       - > $var
echo '"'"$var"'"' - > "hello world" (put single quote around double quote so that it is interrrepted literally)
echo \""$var"\"   - > "hello world" (escape double quote so that it is printed)
```

### Parameter Expansion  

`$var` variable expansion
`${var}s` Seperate s from the variable
`$(command)`  command expansion  
`$((expression))`  arithmetic expansion
```
$ echo "'$USER', '$USERs','${USER}s'"
'autentico', '', 'autenticos'
```
### String Concatenation  

Example                     | Description
--                          | --
`var=$var1$var2`            | Concatenate two variables

Command                     | Description  
--                          | --  
`${#parameter}`             | length of parameter in character, if parameter is array, it returns the length of array  
`${parameter#pattern}`      | delete pattern from start, match against start, short
`${parameter##pattern}`     | delete pattern from start, use longest match from start
`${parameter%pattern}`      | delete pattern from end, match against end, short
`${parameter%%pattern}`     | delete pattern from end, use longest match from end
`${parameter/pat/string}`   | replace first pat with string, string can be null
`${parameter//pat/string}`  | replace every pat with string
`${parameter/#pat/string}`  | match against beginning, useful to add prefix
`${parameter/%pat/string}`  | match against end, useful to add suffix

```
$var=1.5.9
${var#*.}   -> 5.9
${var##*.}  -> 9
```

### Substring  
`${VAR:start_index:length}   # start_index from 0`  
`${MYSTR:11:4}`


### Command Substitution
Insert the output of one command into a second command  
Example: `cur_time=$(date)`  


### Parameters
Parameters include *variables* and *special parameters*  
variable: parameters that are denoted by a name  
special parameters: not denoted by a name  

Assign a variable  
`varname=data`  
Access data using *parameter expansion*   
`echo "$varname"`



Special Parameters | Description
--                 | --
`0`                | name/path of the script
`1,2,3`            | *Positional Parameters* arguments passed to current script
`*`                | Expands to all words to all positional parameters. If double quoted, it expands a string containing them all
`@`                | Expands to all words to all positional parameters. If doube quoted, it expands to a list containing them all
`#`                | Number of postional parameters currently set
`?`                |
`$`                | PID of current shell
`!`                | PID of command most recently executed
`_`                | last argument of last cmd executed

To expand these parameters, add $ before them

Predefined Variables            | Description  
--                              | --  
BASH_VERSIOn                    |
HOSTNAME                        |
PPID                            | PID of parent process of shell
PWD                             | current working directory
RANDOM                          |
UID                             | ID of current user
COLUMNS                         | width of terminal in characters
LINES                           | height of terminal in characters
HOME                            |
PATH                            |



### Variable Types
Command               | Type
--                    | --
`declare -a/array=()` | Array
`declare -A`          | Associative array
`declare -i `         | Integer, rarely used
`declare -r`          | Read Only



### Glob Patterns
Glob=normal character + metacharacter

Metacharacter               | Description  
--                          | --  
\*                           | Match any string
?                           | Match any single character
\[...]                       | Match any one of enclosed character

### Extended Glob  
Enable extended glob using `shopt -s extglob`

Metacharacter   | Description  
--              | -- 
?(list)         | Match zero or one occurence of pattern
*(list)         | Match zero or more occurence of pattern  
+(list)         | Match one or more occurence of pattern 
@(list)         | Match one of given patterns
!(list)         | Match anything but given patterns

list: list of (extended) globs separated by `|`  
```bash
$ ls
names.txt  tokyo.jpg  california.bmp
$ echo !(*jpg|*bmp)
names.txt
```

### Brace Expansion
```bash
$ echo th{e,a}n
then than
$ echo {/home/*,/root}/.*profile
/home/axxo/.bash_profile /home/lhunath/.profile /root/.bash_profile /root/.profile
$ echo {1..9}
1 2 3 4 5 6 7 8 9
$ echo {0,1}{0..9}
00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19
```

### Exit Status  
Each command has an exit status. Range 0-255, 0 means success. `$?` to get exit code of last command


### Control Operator (&& || )
`&&`: If the first command return exit code 0, it will execute the next command.  
`||`: If the first command doesn't return exit code 0, it will execute the next command. Useful for simple error handling
`!`: Negate

### Grouping Statements
Use `{}` to group statements and the group is considered as one statement instead of several.  
Example. Without grouping, which will execute `echo` even though the first command fails:
```bash
$ grep -q goodword "$file" && ! grep -q badword "$file" && rm "$file" || echo "Couldn't delete: $file" >&2
```
With grouping, which is correct:
```bash
$ grep -q goodword "$file" && ! grep -q badword "$file" && { rm "$file" || echo "Couldn't delete: $file" >&2; }
```

### Conditional Blocks

`if` will check the exit code of `COMMAND1`. If it's 0, it will execute the `then` part  
```bash
if COMMAND1
then COMMAND2
elif COMMAND3
then COMMAND4
else COMMAND5
fi
```

Test command:`[]` or 'test' and `[[]]`, test things and return an exit status 
`[ a = b ]` 
`[[]]` allows pattern matching while `[]` doesn't  
`= != < > ` treat their arguments as strings   
`-eq -ne -lt -gt -le -ge` treat arguments as numbers

Tests           | Description
--              | --
-e FILE         | if file exists
-f FILE         | if is regular file
-d FILE         | if is directory
-s FILE         | if file exists and not empty
-r FILE         | if file is readable
-w FILE         | if file is writable
-x FILE         | if file is executable
FILE1 -nt FILE2 | if file1 is newer than file2
FILE1 -ot FILE2 | if file1 is older than file2
-z STRING       | if string is empty
-n STRING       | if string is not empty

Tests support by `[[]]`
Tests             | Description
--                | --
STRING = PATTERN  | if string matches the glob pattern
STRING == PATTERN | same as above
STRING != PATTERN | if string doesn't match the glob pattern
STRING =~ REGEX   | if string matches the regex


### Conditional Loops

```bash
while COMMAND
do COMMAND
done

for ((i=10; i > 0; i--))
do COMMAND
done

for i in {10..1}
do COMMAND
done
```

### Choices  
```bash
case $LANG in 
    en*) COMMAND ;;
    fr*) COMMAND ;;
    de*) COMMAND ;;
    *) COMMAND ;;
esac
```

If use `;&`, the block of next code will be executed no matter if pattern for that choice matches or not  
If use `;;*`, the block of next code will be executed depending on if the condition satisfies

```bash
select choice in A B C
do
echo "Your choice $choice"
done
```

### Arrays  

#### Declare an array:  
`names=("Bob" "Peter" "Big John")`  
`names=([0]="Bob" [1]="Peter" [20]="Big John")`  
`names[0]="Bob"`  
`unset 'names[1]' #remove an element to make in sparse`
`array=("$(array[@])")  # Reindex to remove holes in sparse array`
`photos=(~/pictures/*.jpg)`  
`file=(*)`

In order to parse stream(e.g. Output of a command) into elements of array, we can use NUL byte to indicate where each element starts and ends. As output of a command often can make its output separated by NUL.

#### Print a (array) variable

```bash
$declare -p names
declare -a names=([0]="Bob,Peter,Big John")
```
```bash
echo "${file[@]}"
echo "${file[0]}"
echo "${#file[@]}"  # Length of an array
```

```bash
printf '%s\n' "${file[@]}"   # $@: Expands to a list of words, with each array element as one word. 
```

```bash
for file in "${file[@]}"
do COMMAND
done
```

Remember always to quote properly!  

```bash
$for name in ${names[@]}; do echo $name; done
Bob
Peter
Big
John

$for name in "${names[@]}"; do echo $name; done
Bob
Peter
Big John

$for name in ${names[*]}; do echo $name; done
Bob
Peter
Big
John

$for name in "${names[*]}"; do echo $name; done
Bob Peter Big John  # If you use *, it will be a string

```


#### Expanding Indices  
```bash
$echo "${!names[@]}"
0 1 2       # In this way, you get the index of an array and can possibly use a for loop iterating indices.
```

#### Associative Array
```bash
declare -A fullNames    # declare an associate array
fullNames=( ["Bob"]="Bobbe Lorry" ["Peter"]="Pettery Parker" )
declare -p fullNames
echo "${fullNames['Bob']}"
```
Context in \[...] in a indexed array is arithmetic, you can do math here without wrapping using \$((...)).  

For indexed array, index in [...] part doesn't need a dollar sign$, but for associated array, an dollar sign is required. 
```bash
$ indexedArray=( "one" "two" )

$ declare -A associativeArray=( ["foo"]="bar" ["alpha"]="omega" )

$ index=0 key="foo"

$ echo "${indexedArray[$index]}"

one

$ echo "${indexedArray[index]}"

one

$ echo "${indexedArray[index + 1]}"

two

$ echo "${associativeArray[$key]}"

bar

$ echo "${associativeArray[key]}"

$ echo "${associativeArray[key + 1]}"

```



### Command Line Arguments
```
myfile arg1 arg2 arg3  
$0     $1   $2   $3   # $@ is set of all positional parameters
```
If you `shift` $1, $1 goes away and $2 becomes $1, $3 becomes $2 and so on

### File Descriptors  
Reference to files and other resources like files.
By default, there are three FDs:  
    + stdin, FD 0, from keyboard
    + stdout, FD 1, display on monitor
    + stderr, FD 2, display on monitor

`read`: read info from *stdin* and store in a variable  
`-p`: print according message
`read -p "What's your name?" name_variable`

### Redirection  
Redirection works by changing what FD points to. For example, change FD1 points to stdout to regular file.  

Redirection         | Description  
--                  | --  
`>`                 | Output redirection, stdout will write to a file 
`<`                 | Input redirection , stdin will read from a file 
`>>`                |

By default, if you simply type `cat`, it will read input from `stdin` and print to `stdout`. You can do some changes using redirection  

`cat < story`

Use number to denote FD that will be changed
`command > file`  
`command 1> file`
`command < file`
`command 0< file`

`rm file_that_not_exist 2> errors  # store error log in a file`
```
for sth in sths
do COMMAND
done 2> errors      # redirection applies to all output to stderr inside the loop
```

Ignore error message.  
`COMMAND 2> /dev/null`

### Heredoc and Herestring

Embed short blocks of multi-line data in to the command  
Example of heredoc:
```bash
cat << CONTENT  
>This is
>Really
>Amazing
CONTENT
```

Example of herestring:
`stdin` read information from string after `<<<`
`grep proud <<< "I am really proud"` 



### FIFO  
`mkfifo FILENAME`  
Create a special file, first write, first read, read will block until write to the file. Write will block until another process read FIFO.  


### Pipe  
Connects `stdout` of one process to `stdin` of another.  
Pipe operator will create a subshell environment for each command, so variables you modify inside the second command will appear unmodified outside of it.    


### Process Substition  
`>(command)` write  
`<(command)` read, run the command inside the parentheses and give a temporary filename that you can use.
Example   
`diff <(sort file1) <(sort file2)`


### Subshell  
Create subshell explicitly using parentheses `()`  
Any variables set during subshell are not remembered  
`(cd /tmp || exit 1; date > timestamp)`  
If `cd` fails, `exit 1` will terminate the subshell instead of the interactive shell.

### Command Grouping  
use `{}` to group commands.Separate commands with semicolon and there must be a semicolon after the last command. Or you can write commands in multiple lines. Basically it has no effect. But it allows a collection of commands to be considered as a whole with regard to redirection and control flow.`if, for, while` are already compound commands which behave like command grouping.  
If you use redirection to a command group, the file FD point to will only open once and close once, instead of opening the file and closing the file for each command in the group.


### Arithmatic Evaluation  
```bash
$a=2+3         # a is '2+3'
$let a=2+3     # a is 5
$let a='2 + 3' # a is 5
$((a=2+3))     # a is 5, which is arithmatic evaluation 
$echo $((a*b)) # arithmatic substitution  
```
`((abs = (a>0) ? a : -a))`  
`if ((flag)); then COMMAND; fi      # note $ sign is not required`  


### Functions
```bash
sum(){
    echo "$1 + $2 = $(($1 + $2))"
}
$sum 1 3
1 + 3 = 4
```
```bash
#!/bin/bash
count() {
    local i         # Local variable in the function
    for ((i=1; i<=$1; i++)); do echo $i; done
    echo 'Ah, ah, ah!'
}
for ((i=1; i<=3; i++)); do count $i; done
```

To remove a variable, alias or a function:  
`unset -f myfunction`  
`unset -v myvariable`
`unalias myalias`

### Sourcing  
When a script finishes its execution, its environment is discarded.  
`. ./myscript`  
If you use the dot operation, the script will run in current shell's environment. So Current shell's variables, working directory, functions etc will be changed.


