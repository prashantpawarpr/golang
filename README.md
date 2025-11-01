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

## Ways of declaring the variable

using var

and short hand declaration
:= known as short declaration operator

### imp

=> we are using short hand declaration inside the function only
and reassignment is not allowed with short hand declaration unless we are declaring new variable along with it.

## looping in go lang

<h1>Heading</h1> understood the loop write style in go as i am comming from the javascript background
which is very convenient to write
writing array is complicated as in go we uses { } bracked and not [] bracket and we have to define the declaration method is odd

e.g
var arr [5]int = [5]int{1,2,3,4,5};
<b> need to practice the go lang array often </b>
<b> another new concept of channel is introduced which is new to me , as it is related with goroutin and concurancy i will look it later </b>
