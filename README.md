# Learning

First created different variable there are three ways of defining the variable

var variable_name data_type=value; everything is defined in one line
var variable = value; it will inherit the value

variable_name := value;

## local , global and local preceeded variables

every variable has scope exists withing the block

if variable is defined outside the block it is global variable

and if we have two variables with same name one inside the block

and one outside the block then the variable inside the block is local variable

## type convertion

to type convert in go lang we have to provide the type name in front of the variable
e.g
var a int = 12;
var b float32 = 123.3;
than to type convert
use

int(b)==>> this will convert float32 to int
