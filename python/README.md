# PYTHON

Here are some python scripts for implementing [command line interfaces ,with argparse](https://realpython.com/command-line-interfaces-python-argparse/). 

[Official Doc. (argparse)](https://docs.python.org/3/library/argparse.html)

[Source code (argparse)](https://github.com/python/cpython/blob/3.9/Lib/argparse.py)

### 1) [cat](https://linuxize.com/post/linux-cat-command/) 
The name of the _cat_ command comes from its functionality to concatenate files. It can read, concatenate, and write file contents to the standard output.
* (i) [cat1.py](https://github.com/sandeepb20/6w6l_project/blob/main/python/cat1.py):This script will display the contents of the path file on the terminal.
* (ii) [cat2.py](https://github.com/sandeepb20/6w6l_project/blob/main/python/cat2.py):This will copy the contents of input file to output file. **-a** will append output file ,whereas **-o** will overwrite output file.
    
### 2) [find](https://linuxize.com/post/how-to-find-files-in-linux-using-the-command-line/) 
_Find_ searches for files and directories in a directory hierarchy based on a user given expression and can perform user-specified action on each matched file.
* (i) [find1.py](https://github.com/sandeepb20/6w6l_project/blob/main/python/cat1.py):This can be used to find whether the specified file exist in gived directory. here **-name** is case sensitive ,whereas **-iname** is case insensitive.

### 3) [grep](https://linuxize.com/post/how-to-use-grep-command-to-search-files-in-linux/)
The _grep_ command stands for “global regular expression print”.
_grep_ searches one or more input files for lines that match a given pattern and writes each matching line to standard output.
