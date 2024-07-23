# Makefile

## Makefile

```
targets: prerequisites
    command
    command
    command(usually commands will create the file with the same name as target)
```
Example:  
```
hello:
  echo "Hello world"
  echo "Hahaha"
```

If no arguments are provided to `make` command, it will run the first target by default. If you provide an argument, it will only run the target with the same name as argument and necessary dependencies.

Notice the differences of the following two versions:  

```
blah:
  cc blah.c -o blah  # It won't recompile after blah.c is modified
```

```
blah: blah.c
  cc blah.c -o blah  # It will recompile after blah.c is modified
```

**IMPORTANT**: Make only run blah if blah doesn't exits or blah.c is newer than blah, which means it's better to add header file in the prequisite since editing header file may affect the compilation.

A more common template

```
blah: blah.o
  cc blah.o -o blah
blah.o: blah.c
  cc -c blah.c -o blah.o
clean:
  rm -f *.o blah
```

### Variables

Variables can only be strings.  
```
files := file1 file2  # Or file = file1 file2, Note quote is not required

executable: ${files}
  commands to create this executable

file1:
  touch file1
file2:
  touch file2

```


```
all: f1.o f2.o

f1.o f2.o:
  echo $@  # $@ contains target name, i.e. f1.o f2.o is this case
```

### Wildcard 
`*`: search in the filesystem for matched filenames   
```
thing_wrong := *.o # Don't do this! '*' will not get expanded
thing_right := $(wildcard *.o)

all: one two three four

# Fails, because $(thing_wrong) is the string "*.o"
one: $(thing_wrong)

# Stays as *.o if there are no files that match this pattern :(
two: *.o

# Works as you would expect! In this case, it does nothing.
three: $(thing_right)

# Same as rule three
four: $(wildcard *.o)
```

`%`: Matches one or more characters in a string or take the stem that was matched and replaces that in a string.

### Automatic Variables  

Variables | Description  
--        | --  
$@        | Name of the target of the rule
$%        | 
$<        | Name of the first prerequisite
$?        | Names of all prerequisites newer than target
$^        | Names of all prerequisites


### Implicit Rules  

In the following example, there's no rule to make foo.o bar.o. Make will use implicit rules to make these object files. The implicit rules are based on what kind of files you have and so on.    

Implicit rule will be applied for
1) each target without recipe
2) each double-colon rule without recipe 
```
foo: foo.o bar.o
    cc -o foo foo.o bar.o
```

```
foo.o: foo.c  # Implicit rules will be used
```

Variabels | Description  
--        | --  
CC        | Program for compiling C, usually `cc`
CXX       | Program for compiling C++, usually `g++` 
CFLAGS    | Extra flags to C compiler
CXXFLAGS  | Extra flags to C++
CPPFLAGS  | Extra flags to C preprocessor
LDFLAGS   | Extra flags to compilers when they are suppopsed to invoke the linker

Compile C program, file.c->file.o:  
`$(CC) -c $(CPPFLAGS) $(CFLAGS) $^ -o $@`   
Compile C++ program, file.cpp->file.o:  
`$(CXX) -c $(CPPFLAGS) $(CXXFLAGS) $^ -o $@`  
Link object file, file.o->file:    
`$(CC) $(LDFLAGS) $^ $(LOADLIBES) $(LDLIBS) -o $@`  


### Static Pattern Rule  
Syntex:  
```
targets...: target-pattern: prereq-patterns
    commands
```
For example, the following two Makefiles are the same  
```
objects = foo.o bar.o all.o   
all:${objects}

foo.o: foo.c
bar.o: bar.c
all.o: all.c
```

```
objects = foo.o bar.o all.o
all:{Obejcts}

${objects}: %.o : %.c
```

### Target-specific variables  

One can define variables for specific targets or patterns

```
all: var = hello world

all:
    echo "var is set: ${var}"
other:
    echo "var is nothing: ${var}"
```

```
%.c: var = hello world

blah.c:
    echo "var is set: ${var}"
other:
    echo "var is nothing: ${var}"
```

### Makefile Condition  

```
foo = ok

all:
ifeq (${foo},ok)
    COMMANDS
else
    COMMANDS
endif
```

### Check if a variable is defined

```
all:
ifdef var
    COMMANDS
endif
ifndef var
    COMMANDS
endif
```
