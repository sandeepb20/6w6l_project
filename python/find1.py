import argparse
import os
import sys

my_parser=argparse.ArgumentParser(prog='find',description="To search for files in a directory")

my_parser.add_argument('dir',metavar='dir',action='store',help="Directory for searching file")
my_parser.add_argument('-name',action='store',help='All files matching with name')
my_parser.add_argument('-iname',action='store',help='All files matching with name and is case insensitive')

args=my_parser.parse_args()

print(vars(args))
if args.dir:
    print(os.getcwd())

#print(os.listdir(args.dir))

file_list=os.listdir(args.dir)
if args.name:
    for file in file_list:
        if file==args.name:
            print(file)
            print('found')
if args.iname:
    print(args.iname)
    for file in file_list:
        if file.upper()==args.iname.upper():
            print(file)
            print('found')