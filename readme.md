## A compiler for chimpanzee

- this is for programming language chimpanzee

### 1. A LEXER for lexical analysis
- SOURCE CODE -> TOKEN -> ABSTRACT SYNTAX TREE
- [] real product will attach line number, column number and filename to a token. cause then we can get error sign with related line.

- here is a simple syntax by chimpanzee
```chimpanzee
let five = 5;
let ten = 10;

let add = fn(x, y){
    x + y;
};

let result = add(five, ten);
```

- @?: token type is set for string, although using int or byte may get better performance

#### lexer step
1. input source code
2. use `nextToken()` to go through the source code. 
3. usually a `io.reader()` used here to read line numbers, and the filename. but simplify we dont use it here.
4. 


