#ASCII Art Compiler
ASCII Art Compiler is the tool that output a png file, a pdf file or a html file by a ASCII art file.

![img1](https://github.com/pengincoalition/ASCII_art_compiler/raw/master/explain/img1/output.png)

# Features
* This can output pdf, png or html.
* This is cli tool.
* This is a single binary.

# Quick strat
## Build
We will make executable files.
```shell
git clone https://github.com/pengincoalition/ASCII_art_compiler
cd ASCII_art_compiler
make
#You move "acc" to the executable path.
```
## New project
```
acc new
#You rename any ttf font file to "font.ttf" and move it to "font" dictionary.
```

## Build
```shell 
acc build --type image
#output.png
acc build --type pdf
#output.pdf
acc build --type html
#output.html
```

# Futures
* Making real time view.
* Making executable files.
 

