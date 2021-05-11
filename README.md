# ASCII Art Compiler
ASCII Art Compiler is the tool that output the png file, the pdf file or the html file by the ASCII art file.

![img1](https://github.com/pengincoalition/ASCII_art_compiler/raw/master/explain/img1/output.png)

# Features
* This can output a png file, the pdf file or the html file by the ASCII art file.
* This is the cli tool.
* This is the single binary.

# Quick strat
## Make New project
```shell
mkdir your-project
cd your-project
acc new
mv /favorite.ttf ./font/font.ttf
#You rename any ttf font file to "font.ttf" and move it to "./font" dictionary.
```

## Draw ASCII art

```shell
nvim main.aasc
#Use your favorite text editor and enjoy drawing the ASCII art!
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
* Changing design of output files.



