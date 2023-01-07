# log

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [How to log?](#how-to-log)
* [Code examples](#code-examples)

## General info
This is a simple library enhanced with emojis and padding alongside with time of when the log occured. Used to log errors, values, arrays and maps in a cool way for Go developers.
	
## Technologies
Project is created with:
* <a href="https://github.com/golang/go">Go programming language</a>.
	
## Setup
To install this library, run the following commands:

```
$ go mod init "whatEverProjectName"
$ go get -u github.com/promiselb/log
```
## How to log?
1- All messages containings this mark sign ❗, are considered errors.

2- These two pistol-fingers 👉 👈 used to indicate values of variables.

3- This success sign ✅ indicate that the variable printed is not a zero value, while this one indicate it is ❌. It's used
To see if some messages are recieved.

4- The mysterious `T.H x` notaion represent the <i>Threat Level</i> where x ∈ [0:5], a <i>threat level</i> must be used 
to indicate how much an error is dangerous (0 no risk, max 5 fatal error and everything in between is a warning or not critical)

5- Lastly, the ☢ radiation sign (prefer call it nuke sign) indicate that a fatal error has occured 
and will exit the programm immediately after printing the error. (Nukes your program)
## Code Examples

### Prepare the envirnoment
To get it working in your file, copy the following:
``` Go
package main

import (
	w  "github.com/promiselb/log"
)

var l = w.NewLogger(nil, "Alex", 0)

func main() {

}
```
Use an alias (w in this case) because to not get confused with the default `log` package.

### Print individual values
To print variables and errors, write in your `main` function:
``` Go
x := 5
name := "promiselb"
err := errors.New("error here")

l.PrintValue("x", x)
l.PrintValue("name", name)
l.PrintError(err, 1)
```
Output:
```
Alex 2023/01/07 :
✅       { 👉 x              : 5               👈 }       T.H. 0️⃣
Alex 2023/01/07 :
✅       { 👉 name           : promiselb       👈 }       T.H. 0️⃣
Alex 2023/01/07 :
❗       { ❗: error here!     ❗ }			   T.H. 1️⃣ 
```
or this:
``` Go
var l = w.NewLogger(nil, "Alex", 0)
p := l.MakePrinter("main")
x := 5
name := "promiselb"
err := errors.New("error here")

p.PrintValue("x", x)
p.PrintValue("name", name)
p.PrintError(err, 1)
```
Output:
```
Alex main :
✅       { 👉 x              : 5               👈 }       T.H. 0️⃣
Alex main :
✅       { 👉 name           : promiselb       👈 }       T.H. 0️⃣
Alex main :
❗       { ❗: error here      ❗ }                       	   T.H. 1️⃣  
```
The outputs look similiars. However, the diffirence is the 1st logs have the time when it occured without the function name,
while 2nd logs are completely the opposite, the printer must be scoped to function to know in which function this x = v. And
to get rid of the flag thingy if if not needed.

### Print slices
``` Go
var l = w.NewLogger(nil, "Alex", 0)

fruits := []string{"apple", "mango", "orange", "bananas"}
w.PrintArray(l, "fruits", fruits)
```
Output:
```
Alex 2023/01/07 : fruits:
        ------------------------------
         | 0   |apple               | 
        ------------------------------
         | 1   |mango               | 
        ------------------------------
         | 2   |orange              | 
        ------------------------------
         | 3   |bananas             | 
        ------------------------------
```
It's not used as method because we can't use generics with method functions signature.
### Print maps
``` Go
nc main() {
	var l = w.NewLogger(nil, "Alex", 0)

	idCard := map[string]interface{}{
		"name": "promiselb",
		"age":  16.5,
		"dev":  true,
	}
	w.PrintMap(l, "idCard", idCard)
```
Output:
```
	------------------------------
         | dev |true                | 
        ------------------------------
         | name|promiselb           | 
        ------------------------------
         | age |16.5                | 
        ------------------------------
```

Happy coding!
