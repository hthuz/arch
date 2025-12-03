# Software

Notes for some "advanced" softwares instead of just common ones

## VScode  
Format code: `Ctrl + Shift + I`

enable esc and capslock swap for vim extension:
add "keyboard.dispatch": "keyCode" in settings.json

## pandoc
Convert from one markup format to another  

| CMD                                                    | Description                                                |
| ------------------------------------------------------ | ---------------------------------------------------------- |
| pandoc -o output.html input.md                         | -o to specify the output, only produce document fragment   |
| pandoc -s -o output.html input.md                      | --standalone, produce complete document                    |
| pandoc -f markdown -t latex hello.txt                  | --from, --to, specify the input/output format              |
| pandoc sample.md -f gfm -o sample.pdf                  | --from github style markdown to                            |
| pandoc test.txt -o test.pdf                            | Create a PDF file                                          |
| pandoc -f html -t markdown www.xxx.com                 | Read from a website                                        |
| pandoc --extract-media ./img input.docx -o output.md   | Docx to md with images                                     |
| pandoc --highlight-style pygments                      | Add syntax highlighting                                    |
| pandoc --print-highlight-sytle=pygments > my.theme    | Print a theme, and edit it!                                |
| pandoc sample.md --highlight-style my.theme -o out.pdf | Use my edited theme file                                   |
| pandoc --toc --toc-depth <3> -V toc-title=<your_title> | Add table of contents with depth 3 and your content title  |
| pandoc --template eisvogel                             | Use a template, path from `~/.local/share/pandoc/templates`|

Some syntax highlight styles: `pygments, kate, monochrome, espresso, haddock, tango, zenburn`  

### Add chapter breaks

 `-H` `(--include-in-header)`, with tex file `chapter_break.tex` and content as follows 

```
\usepackage{sectsty}
\sectionfont{\clearpage}
```

With command `pandoc sample.md -f gfm -H chapter_break.tex -o sample.pdf`


### Change settings

`-V` `--variable`

```sh
#!/bin/bash
pandoc "$1" \
    -f gfm \
    --include-in-header chapter_break.tex \
    -V linkcolor:blue \
    -V geometry:a4paper \
    -V geometry:margin=2cm \
    -V mainfont="DejaVu Serif" \      # Normal text
    -V monofont="DejaVu Sans Mono" \  # Code snippet
    -V fontsize=12 \
    --pdf-engine=xelatex \
    -o "$2"
```

## Vim  

### ColorScheme
`:colorscheme <scheme>`  Change color scheme  
Options:
blue       delek     evening    morning   peachpuff   slate  
darkblue   desert    industry   murphy    ron         torte  
default    elflord   koehler    pablo     shine       zellner  


Cmd                     | Description  
--                      | -- 
`gg`                    | Go to start of file
`G`                     | Go to end of file


## ranger
Copy default configuration to `~/.config/ranger`: `ranger --copy-config=all`  
In `~/.config/ranger`, there are following files

file            | Description  
--              | --
`rc.conf`       | Keybinding and settings
`commands.py`   | commands launced with :
`rifle.conf`    | application used when given type of file is launced
`scope.sh`      | File Preview Settings


#### Console Commands  
Use `space` to select files  
Use `:` followed by commands

CMDs              | Description
--                | --
`:bulkrename`     | Rename in bulk
`:open_with`       | Open a select file with your chosen program
`:touch FILENAME` | create new file
`:mkdir FILENAME` | create new directory
`:shell COMMANDS` | run commands in a shell
`:delete`         | Delete files

Placeholder         | Description  
--                  | --
`%s`                | Currently select file
`%d`                | Current directory

#### Usual Commands

| CMDs                    | Description               |
|-------------------------|---------------------------|
| `i`                     | Preview File              |
| `r`                     | Open File                 |
| `zh/backspace` or <c-h> | toggling hidden files     |
| `gh`                    | Go to home directory      |
| `g/`                    | Go to root directory      |
| `cw`                    | Rename File               |
| `yy`                    | Copy file                 |
| `dd`                    | cut file                  |
| `dD`                    | delete file               |
| `pp`                    | paste file                |
| `z + letter`            | Changing settings         |
| `g + letter`            | Go to dir or operate tabs |
| `c + letter`            |                           |
| `y + letter`            |                           |
| `d + letter`            |                           |
| `p + letter`            | Different paste mode      |

#### Image Preview  
`set preview_images true` in `rc.conf` and change `preview_images_method`

#### PDF Preview
Uncomment the following part in "scope.sh"

### Onedrive
onedrive on linux: insync with aur. Works but slow


## Docker

List of frequently used commands and flags

image name should be put after the options

| command                                                           | Description                                                           |
|-------------------------------------------------------------------|-----------------------------------------------------------------------|
| docker run -d -p 80:80 --name my_container nginx                  | run container from image, -d: in background mode, --name: assign name |
| docker run -d -p 80:80 -p 81:81 -v /local/path:/app/data/data.csv | -v: mount local path                                                  |
| docker pull ubuntu                                                | pull an image                                                         |
| docker build -t my_image:latest -f Dockerfile .                   | build an image from Dockerfile. -t: tag, -f: name of dockerfile,      |
| docker build -t my_iamge:latest . --network=host                  | if network issue during build                                         |
| docker images                                                     | list images                                                           |
| docker ps -a                                                      | list containers, -a to show all containers (including inactive)       |
| docker stop my_container                                          | stop a container                                                      |
| docker start my_container                                         | start a container                                                     |
| docker rm my_contianer                                            | remove a container (won't in ps -a any more), -f: force               |
| docker rmi my_image:latest                                        | remove a image, -f: force                                             |
| docker save my_image -o my_image.tar                              | save a image to a tar (which can be loaded afterwards)                |
| docker load -i my_image.tar                                       | load a tar to a image                                                 |
| docker exec my_container ls                                       | run a command in running container                                    |
| docker exec -it my_container /bin/bash                            | -it: interactive terminal mode, so you can run bash cmds in docker    |
| docker run -it my_image /bin/bash                                 | run bash from image directly                                          |
| docker logs my_container                                          | show current logs of a container                                      |
| docker logs -f my_contiainer                                      | -f: follow. new logs will be printed as well                          |
| docker system/container/network/image prune                       | remove unused container/network/image                                 |
| docker system/container/network/image                             | show a list of commands related to system/container/network/image     |

Steps for docker push
1. docker login
2. docker tag myapp:latest username/myapp:latest
3. docker push username/myapp:latest

### Using specified network when building images, to fix the issue of no network
`docker build --network host`
or in docker-compose:

```yml
web:
    build:
        context: .
        network: host
```


### docker-compose volume:

```yml
web:
    image: xxx
    volume:
        - </host/path>:</docker/path>
        # e.g.
        - /home/user/app/data:/usr/share/nginx/html
        - ./config.conf:/etc/default.conf
```

### stop containers in docker-compose
```bash
docker-compose down
```


## java archlinux

path
`/usr/lib/jvm/java-8-openjdk`

`/usr/lib/jvm/java-11-openjdk`

`/usr/lib/jvm/java-25-openjdk`


## Maven

| Cmd                       | Desc                                                  |
|---------------------------|-------------------------------------------------------|
| `mvn compile`             | comple classes to `target/classes`                    |
| `mvn package`             | produce jar, output to `target/`                      |
| `mvn install`             | produce jar, output to `target/` and install to `.m2` |
| `mvn install -DskipTests` |                                                       |
| `mvn spring-boot:run`     | run spring boot                                       |



## kafka

require java

If installed using pacman, run the following cmds directly. They are under `/usr/bin`
Otherwise, specify the path e.g. `bin/kafka-topics.sh`

First step

```bash
KAFKA_CLUSTER_ID="$(kafka-storage.sh random-uuid)"
kafka-storage.sh format --standalone -t $KAFKA_CLUSTER_ID -c config/server.properties
# if using pacman
sudo -u kafka /usr/bin/kafka-storage.sh format -t "$(/usr/bin/kafka-storage.sh random-uuid)" -c /etc/kafka/server.properties --standalone
```


| Cmd                                                                                            | Desc            |
|------------------------------------------------------------------------------------------------|-----------------|
| `kafka-topics.sh --create --topic mytopic --bootstrap-server localhost:9092`                   |                 |
| `kafka-topics.sh --describe --bootstrap-server localhost:9092`                                 | show all topics |
| `kafka-console-producer.sh --topic mytopic --bootstrap-server localhost:9092`                  |                 |
| `kafka-console-consumer.sh --topic mytopic --bootstrap-server localhost:9092 --from-beginning` |                 |



## redis-cli

| Cmd                                          | Desc                 |
|----------------------------------------------|----------------------|
| `redis-cli -h <host> -p <port> -a <password` |                      |
| `redis-cli -n 0`                             | use database 0 (0-15) | 


inside shell
| Cmd             | Desc          |
|-----------------|---------------|
| `keys *`        | show all keys |
| `keys users:*`  | pattern match |
| `del key1 key2` |               |
| `ttl key1`      |               |

