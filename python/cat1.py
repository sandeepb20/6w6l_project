import argparse
import sys
import os

my_parser=argparse.ArgumentParser(prog='cat')

my_parser.add_argument('Path',metavar='path',type=str,nargs='*',help='Adding path to list')
#my_parser.add_argument('-f',action='store',help='File to print data')

args=my_parser.parse_args()
print(vars(args))
input_path=args.Path
print(input_path[0])
with open(input_path[0],'r') as r:
    for line in r:
        print(line)