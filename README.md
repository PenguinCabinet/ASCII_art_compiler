# ASCII Art Compiler
ASCII Art Compiler is the tool that output the png file, the pdf file or the html file by the ASCII art file.

![img1](https://github.com/pengincoalition/ASCII_art_compiler/raw/master/explain/img1/output.png)

# Features
* This can output the png file, the pdf file or the html file by your ASCII art file.
* This can show you the output of your ASCII art after compiling in your browser in real time.
* This is the cli tool.
* This is the single binary.

# Quick strat
## Make New project
```shell
mkdir your-project
cd your-project
acc new
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



