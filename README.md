cx
==

cx is a modernized version of C.
cx adds no overhead to C, but adds modern conveniences
cx is built on LLVM and will run on any platform supported by LLVM


Simple:
- ultra minimalist (based upon C and GO)
- fast as C
- simple grammar

Relationship to C:
- cx is almost reverse compatible with C
- C programs can be easily ported to cx
- cx effortlessly imports C depedencies

Advanced Features
=================

Modern conveniences:
- struct reflection
- map[int]int dictionary type
- string type
- memory safe slice, array type
- channels for concurrency (based upon golang channels)

Stricter Specification than C:
- fixed width types int32,int64, float32, float64 defined in standard
- cx defines serialization of structs to byte array
- no implicit type conversions
- official coding style for language

Reflection:
- Struct Reflection. Ability to get member variables for cx structs at runtime
- Function Reflection. Ability to get name and type signature of exported module functions at runtime.

Extensible Syntax
- cx programs can extend the grammar of the language
- programs can define domain specific languages in cx

Compilation:
- cx can be run both as a scripting language and compiled
- cx modules can be recompiled and reloaded at runtime


Cx Compilation
==============

Caching Compiler:
- cx is a caching compiler.
- cx compilation is designed to be massively parrallel
- 
- cx only re modules if a file in the module has changed
- cx caches the AST of parsed modules
- cx will only recompile a module if 


cx flow control
===============

for statement:
- There is only for. There is no while, until, not for. 
- parens are optional.
- Semicolons are only needed for seperating statements
- := infers type

Note: replace ":= range" with "in"

```
//while loop
for i < 5 {
	i++
}
```

```
//for loop
for i:=0; i<5; i++ {

}
```

```
//range keyword, for key,value
for index,value := range KeyValueExample {
	
}

```

```
//optional: range keyword?
for num := range(0, 10) {
	//do something
}
```

--- 

if statement:
- parens are optional
- last statement in list is return value for if

``` 
if x == false {
	//do something
}

```

```
if x:= y+z; x==5 {
	//returns true
}
```

--- advanced, optional ---

"get list of X that match Y"

```
select x := range List; x.Id == 5 {
	
}
//can this be for statement?
//is there better way of doing selector?
```

List Selectors/Pattern Matching
```
list := select(x:= range List, x.Id =5)
//returns list, iterator/list, condition

list := for x:= range List; x.Id= 5

//use "in" for iterator instead of ":= range"?
list := select x := range List {
	
}

python:
a = [1,2,3,4,5]
b = [x for x in a if x > 3]
token: b = filter(lambda x: x > 3, a)
```

Sorting for structs 
```
list := sort x := range; x.Index 

list := sort x,y in 
``

cx program structure
====================

Cx does not have a "link" phase in compilation. Cx uses a platform indepedent structure for loading depedencies. A cx program consists of a series of modules.  

-- module type reflection --

Each module has a reflection function that defines the types exported by the module.

```
[type name], [type signature], [field names], [function pointer]
("test_struct", "{int,int}", "x,y", 0x...)
("test_struct2", "{int,test_struct2}", "x,t1", 0x...)
```

Type signatures of types exported by the module may only contain
- atomic types
- previously defined struct types
- pointers to struct types

-- module type dependencies --

A cx module contains a list of types and signatures for types required to define types in the current module. Each type must be defined down to the atomic type.
- atomic types
- previously defined struct types
- pointers to struct types

```
[module], [type name], [type signature], [field names]
("module", "struct_name", "{int,int,int}", "x,y,z")
```

-- module function reflection -- 

Each module has a reflection function that returns the functions exported by the module. The function returns a list of

```
[function name], [type signature], [field names], [function pointer]
("test_function", "{int,int},{int}", "x,y,z", 0x...)
```

-- module function depedencies --

Each module has a table of external functions used by the module.



Note
=====

Difference between 
- assert - assert invariant
- error - error

When are asserts/invariants different from errors?
- preconditions
- post-conditions

"specifications" - post conditions/preconditions
testing impossible without a "specification?"


---

Precondition
- predicate that holds before function is run
Post-condition
- predicate that holds after funtion if precondition is met

Pre: member(x,l) 

List remove(x : Elem ,l: List) {
if (x == head(l))
return tail(l);
else
return cons(head(l),remove(x,tail(l)));
}

Post: !member(x,l)

• Most useful if they are executable
– Written in the programming language itself
– A special case of asserts

Positive Assertions
Negative Assertions

---

For static analysis:

- mutation testing?
- fuzzing

---
" Codenomicon[6] (2001) and Mu Dynamics (2005) evolved fuzzing concepts to a fully stateful mutation testing platform,"

---
http://en.wikipedia.org/wiki/Notation3

@PREFIX dc: <http://purl.org/dc/elements/1.1/>.
 
<http://en.wikipedia.org/wiki/Tony_Benn>
  dc:title "Tony Benn";
  dc:publisher "Wikipedia"

