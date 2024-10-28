This is a go port of my [RPN_Calculator](https://github.com/danomagnum/RPN_Calculator) project.  A python version upgrade broke the gui app so instead of fixing it I took the opportunity to learn how to do a gui in go with fyne.

The actual program you'll want to build is in the rpngui directory.

The description for that project is below with a few obvious changes to match what this version supports.  

Not everything below is actually supported (yet) though.:


Once a week or so I end up needing to do a quick calculation while I'm sitting at my computer.  There are a zillion options for doing this, so it would seem easy to find something that worked really well as a general purpose calculator.

But I wasn't happy with what I found.  I definitely wanted it to use [RPN](https://en.wikipedia.org/wiki/Reverse_Polish_notation), be able to handle functions, and show the current stack.  I also wanted variables, to be able to load programs from files, and to run in the terminal.

[This website](http://www.linuxfocus.org/English/January2004/article319.shtml) lists a bunch of RPN calculators.  A few of them are real close to doing what I wanted, but none of them were quite right. So I decided to roll my own.

Here are the built in operations.

### Math

All math operations pop the last two items from the stack, then push the result.  The math works as expected when one of the operands is a float, you get a float result

Initial stack: 
| Position| Initial ||  ``` + ```  |  ``` - ```  |  ``` * ```  |  ``` / ```  |  ``` % ```  |  ``` ^ ```  |
|-----------------|---------|-|-------|-------|-------|-------|-------|-------|
| infix equivalent|         | |#1 + #0|#1 - #0|#1 * #0|#1 / #0|#1 % #0|#1 ^ #0|
|#4        |    1    ||     |     |     |     |     |     |
|#3        |    2    ||  1  |  1  |  1  |  1  |  1  |  1  |
|#2        |    3    ||  2  |  2  |  2  |  2  |  2  |  2  |
|#1        |    4    ||  3  |  3  |  3  |  3  |  3  |  3  |
|#0        |    5    ||  9  | -1  |  20 |  0  |  1  | 1024|


### Stack Control

Stack control operations move items around on the stack.  dup, over, tuck, and pick add an item to the stack.  Drop removes one. The backtick ``` ```` can be used as a shorthand for drop.

| Position| Initial || ``` swap ```| ``` dup ```| ``` rot ```| ``` over ```| ``` tuck ```| ``` 3 pick ```| ``` 3 roll ```| ``` drop ```|
|---------|---------|-|------------|------------|------------|-------------|-------------|---------------|---------------|-------------|
|#4        |         ||     |1   |    |1    |1    |1      |       |     |
|#3        |    1    ||1    |2   |1   |2    |2    |2      |2      |     |
|#2        |    2    ||2    |3   |3   |3    |4    |3      |3      |1    |
|#1        |    3    ||4    |4   |4   |4    |3    |4      |4      |2    |
|#0        |    4    ||3    |4   |2   |3    |4    |1      |1      |3    |


### Comparisons

The comparisons would take the infix form of #0 &lt;operator&gt; #1, and return 1 if true and 0 if false.

| Position| Initial || ``` == ```| ``` > ```| ``` < ```| ``` >= ```| ``` <= ```|
|---------|---------|-|----------|----------|----------|-----------|-----------|
|#4        |         ||     |     |     |     |     |
|#3        |    1    ||     |     |     |     |     |
|#2        |    2    ||  1  |  1  |  1  |  1  |  1  |
|#1        |    3    ||  2  |  2  |  2  |  2  |  2  |
|#0        |    3    ||  1  |  0  |  0  |  1  |  1  |

| Position| Initial || ``` == ```| ``` > ```| ``` < ```| ``` >= ```| ``` <= ```|
|---------|---------|-|----------|----------|----------|-----------|-----------|
|#4        |         ||     |     |     |     |     |
|#3        |    1    ||     |     |     |     |     |
|#2        |    2    ||  1  |  1  |  1  |  1  |  1  |
|#1        |    3    ||  2  |  2  |  2  |  2  |  2  |
|#0        |    4    ||  0  |  1  |  0  |  1  |  0  |

| Position| Initial || ``` == ```| ``` > ```| ``` < ```| ``` >= ```| ``` <= ```|
|---------|---------|-|----------|----------|----------|-----------|-----------|
|#4        |         ||     |     |     |     |     |
|#3        |    1    ||     |     |     |     |     |
|#2        |    2    ||  1  |  1  |  1  |  1  |  1  |
|#1        |    3    ||  2  |  2  |  2  |  2  |  2  |
|#0        |    2    ||  0  |  0  |  1  |  0  |  1  |


### Conditionals

The two conditionals work similarly, but pop a different number of parameters.

They resolve to true if and only if the ? value is 1 or 1.0.  IfTrue or IfFalse can be any data type.  If they are substacks, they will be immediately executed.  To anonymously return a substacks, a simple nested substack will suffice.  (ex:``` [ 4 ] ``` will return 4 from the if block, but ``` [ [ 4 ] ] ``` will return the substack [ 4 ].

| Position| Initial ||  ``` if ```|
|---------|---------|-|-----------|
|#4        |   N/A   || N/A |
|#3        |   N/A   || N/A |
|#2        |   N/A   || N/A |
|#1        |    ?    || N/A |
|#0        | IfTrue  || IfTrue|

| Position| Initial || ``` ifelse ```|
|---------|---------|-|-----------|
|#4        |   N/A   || N/A |
|#3        |   N/A   || N/A |
|#2        |    ?    || N/A |
|#1        |  IfTrue || N/A |
|#0        | IfFalse || Result|

### Substacks

"Substacks" can be though of as functions, lists, routines, and objects.
To execute a substack, use the ```!``` operator.  To "instantiate" a substack, use the ``` !!``` operator  This effectively executes the substack just like ! but instead of pushing new values onto the main stack, it creates a new substack holding them.  If for some reason an error occurs during function execution, the stack is reverted to its state before execution started.

| Position | Initial ||  ``` ! ```|
|---------|---------|-|-----------|
|#4        |         ||           |
|#3        |    3    ||  3        |
|#2        |    2    ||  2        |
|#1        |    1    ||  1        |
|#0        |  [ 3 ]  ||  3        |

| Position | Initial ||  ``` !! ```|
|---------|---------|-|-----------|
|#4        |         ||           |
|#3        |    3    ||  3        |
|#2        |    2    ||  2        |
|#1        |    1    ||  1        |
|#0        |  [ 3 ]  ||  [ 3 ]    |

You can create substacks on the fly by using the ``` group``` operator.  It pops the last item on the stack and groups that many items into a substack.  You can also grow a substack using the ``` append``` and ``` prepend``` operators.  The ``` cat``` operator combines two substacks together into one if needed.

| Position | Initial ||  ``` group ```|  ``` append ``` |  ``` prepend ``` |  ``` 7 8 cat ``` |  ``` cat ```      |
|----------|---------|-|--------------|-----------------|-------------------|----------------|-------------------|
|#5        |    6    ||               |                 |                  |                  |                   |
|#4        |    5    ||               |                 |                  |                  |                   |
|#3        |    4    ||               |                 |                  |                  |                   |
|#2        |    2    ||  6            |                 |                  |                  |                   |
|#1        |    1    ||  5            |  6              |                  | [ 6 5 2 1 5 ]    |                   |
|#0        |    3    ||  [ 4 2 1]     |  [ 4 2 1 5 ]    | [ 6 5 2 1 5 ]    | [ 7 8 ]          | [ 6 5 2 1 5 7 8 ] |

Note that ``` group``` and ```cat``` look similar, but group preserves and embeds lists while cat joins them.

You can access items in a substack using the ``` \ ``` operator.  You can similarly get the size of a substack using ``` \size ```.

You can use \ to get items from an array.

The accessed item can be a variable on the stack.
| Position | Initial  ||  ``` 0 \ ```  |  ``` drop 1 \ ``` |  ``` drop 2 \ ``` |  ``` drop 4 \ ``` |
|----------|---------|-|---------------|-------------------|-------------------|-------------------|
|#5        |          ||               |                   |                   |                   |
|#4        |          ||               |                   |                   |                   |
|#3        |          ||               |                   |                   |                   |
|#2        |          ||               |                   |                   |                   |
|#1        |          ||  [ 1 2 3]     |  [ 1 2 3]         | [ 1 2 3]          | [ 1 2 3]          |
|#0        | [ 1 2 3] ||  1            |  2                | 3                 | NULL              |


Or part of the instruction
| Position | Initial  ||  ``` \2 ```   |
|---------|---------|-|-----------|
|#5        |          ||               |
|#4        |          ||               |
|#3        |          ||               |
|#2        |          ||               |
|#1        |          ||  [ 1 2 3]     |
|#0        | [ 1 2 3] ||  3            |

Or it can be used with the SIZE instruction
| Position | Initial       ||  ``` \size ```    |
|---------|---------|-|-----------|
|#5        |               ||                   |
|#4        |               ||                   |
|#3        |               ||                   |
|#2        |               ||                   |
|#1        |               ||  [ 1 2 3 2 1]     |
|#0        | [ 1 2 3 2 1 ] ||  5                |


You can also use ``` \varname ``` to read a variable name out of a substack that has been instatiated with ``` !!```.
| Position | Initial       ||  ``` !! \a ```    |
|---------|---------|-|-----------|
|#5        |               ||                   |
|#4        |               ||                   |
|#3        |               ||                   |
|#2        |               ||                   |
|#1        |               ||  [ 1 2 3 a =]     |
|#0        | [ 1 2 3 a = ] ||  3 a =            |


When executed, substacks run in a sub-interpreter where they have access to any variables previously defined and can pop items from their parent interpreters stack.  If a variable does not exist in their parent (or their parent's parent recursively) then a temporary variable is created that gets unbound to a constant when the execution exits.  Whatever is in the stack at the end of the substack's execution gets pushed onto the end of its parents stack (or a sub-stack with ``` !! ```), with all variables not in the parent scope resolved to their values and functions made anonymous.  Variables can be declared as local only by prefixing the variable name with a **$**.  Variables defined in this way will not show up in the parent or any child interpreters.

Anything that can go out of a substack can go inside of one, including substacks.  Since the parent interpreters variables are available, recursion is possible.  For example, here is a program that calculates the factorial of the last number on the stack.

The provided interface can automatically load commands at startup.  It will read in all the files from the *auto_functions_directory* which is defined in the settings file.  All files in this directory with extension ".rpn" are read in and interpreted line by line.  Because they are dumped directly into the interpreter, it is recommended that everything in the files be assigned to a variable and then dropped off the stack so you start with a blank stack, but new pre-defined variables.  A few of these are included for convenience such as ``` all ``` which returns 1 if there are items on the stack and 0 if there are not and ``` fact``` which returns the factorial of the last number on the stack.  This is also where you will find ``` foreach ``` and ``` reverse ``` which operate on lists as you might expect.

```plain

[
 dup     # copy the item we are factorializing
 1 -     # decrement it 
 0 <     # until it reaches zero
 [       # if greater than 0
  dup    # set up the value we are passing to the next recursion
  1 -    # it is one less than the starting value
  fact ! # recurse
  *      # multiply the result of recursion by the original value
  ]
 if
]
fact =   # set up the function name so we can recurse to it
drop

```

### Looping

There is a single looping construct, ``` while```.

While takes a function as its argument and executes it.  It then pops the last item from the stack and uses that value to determine if the function should be executed again, or if execution should end.  A 1 will continue execution, anything else will end it.  At any time during the loop, ``` break``` can be used to end the loop early. For example, ``` [ + size 1 < ] while``` can be used to add all items in the stack. ``` [ 5 == [ break ] if 1 ] while``` can be used to purge items on the stack until you reach a 5 (of course, you could use ``` [ 5 == not ] while``` to achieve the same thing. )

### Variables

To define a variable, you just enter whatever string you want as long as it does not contain any of the built in operator names or characters.  If the variable is unassigned, this will create a new variable with a value of 0. If it has already been assigned, it will push the variable onto the stack.

Assuming the variable abc has a value of 5 from a previous assignment:

| Position | Initial || ``` var ```| ``` abc ```|
|----------|---------|-|-----------|------------|
|#4        |         ||       |   3   |
|#3        |         ||   3   |   2   |
|#2        |    3    ||   2   |   1   |
|#1        |    2    ||   1   | var=0 |
|#0        |    1    || var=0 | abc=5 |

Assignment is handled with the ``` :=``` operator.  The value in position #0 must be a variable.  The value of position #1 is assigned to the variable.  If position #1 is a function, the variable name becomes the function name.

| Position | Initial ||  ``` := ```  |
|---------|---------|-|-----------|
|#4        |         ||       |
|#3        |    3    ||       |
|#2        |    2    ||  3    |
|#1        |    1    ||  2    |
|#0        |  var=0  || var=1 |

If a variable name starts with a $, it is local (and local only - called functions will no be able to access it either)

### Other

Code comments are indicated by a hash #.  Anything after the # is ignored on each parsed line.

The int command converts a variable or float to an integer.  These commands are useful to push the //value// of a non-local variable value onto the stack from a function.


### Interface

Enter puts the current line into the interpreter.

The up and down arrows move through the history.

^C will either clear the current entry or exit if there is nothing currently entered.
