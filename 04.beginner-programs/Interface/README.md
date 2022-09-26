# Go Data Structures: Interfaces

Posted on Tuesday, December 1, 2009.

From https://research.swtch.com/interfaces

Go's interfaces—static, checked at compile time, dynamic when asked for—are, for me, the most exciting part of Go from a language design point of view. If I could export one feature of Go into other languages, it would be interfaces.

This post is my take on the implementation of interface values in the “gc” compilers: 6g, 8g, and 5g. Over at Airs, Ian Lance Taylor has written two posts about the implementation of interface values in gccgo. The implementations are more alike than different: the biggest difference is that this post has pictures.

Before looking at the implementation, let's get a sense of what it must support.

Usage
Go's interfaces let you use duck typing like you would in a purely dynamic language like Python but still have the compiler catch obvious mistakes like passing an int where an object with a Read method was expected, or like calling the Read method with the wrong number of arguments. To use interfaces, first define the interface type (say, ReadCloser):

type ReadCloser interface {
    Read(b []byte) (n int, err os.Error)
    Close()
}
and then define your new function as taking a ReadCloser. For example, this function calls Read repeatedly to get all the data that was requested and then calls Close:

func ReadAndClose(r ReadCloser, buf []byte) (n int, err os.Error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    r.Close()
    return
}
The code that calls ReadAndClose can pass a value of any type as long as it has Read and Close methods with the right signatures. And, unlike in languages like Python, if you pass a value with the wrong type, you get an error at compile time, not run time.

Interfaces aren't restricted to static checking, though. You can check dynamically whether a particular interface value has an additional method. For example:

type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float:
        return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}
The value any has static type interface{}, meaning no guarantee of any methods at all: it could contain any type. The “comma ok” assignment inside the if statement asks whether it is possible to convert any to an interface value of type Stringer, which has the method String. If so, the body of that statement calls the method to obtain a string to return. Otherwise, the switch picks off a few basic types before giving up. This is basically a stripped down version of what the fmt package does. (The if could be replaced by adding case Stringer: at the top of the switch, but I used a separate statement to draw attention to the check.)

As a simple example, let's consider a 64-bit integer type with a String method that prints the value in binary and a trivial Get method:

type Binary uint64

func (i Binary) String() string {
    return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
    return uint64(i)
}
A value of type Binary can be passed to ToString, which will format it using the String method, even though the program never says that Binary intends to implement Stringer. There's no need: the runtime can see that Binary has a String method, so it implements Stringer, even if the author of Binary has never heard of Stringer.

These examples show that even though all the implicit conversions are checked at compile time, explicit interface-to-interface conversions can inquire about method sets at run time. “Effective Go” has more details about and examples of how interface values can be used.

Interface Values
Languages with methods typically fall into one of two camps: prepare tables for all the method calls statically (as in C++ and Java), or do a method lookup at each call (as in Smalltalk and its many imitators, JavaScript and Python included) and add fancy caching to make that call efficient. Go sits halfway between the two: it has method tables but computes them at run time. I don't know whether Go is the first language to use this technique, but it's certainly not a common one. (I'd be interested to hear about earlier examples; leave a comment below.)

As a warmup, a value of type Binary is just a 64-bit integer made up of two 32-bit words (like in the last post, we'll assume a 32-bit machine; this time memory grows down instead of to the right):


Interface values are represented as a two-word pair giving a pointer to information about the type stored in the interface and a pointer to the associated data. Assigning b to an interface value of type Stringer sets both words of the interface value.


(The pointers contained in the interface value are gray to emphasize that they are implicit, not directly exposed to Go programs.)

The first word in the interface value points at what I call an interface table or itable (pronounced i-table; in the runtime sources, the C implementation name is Itab). The itable begins with some metadata about the types involved and then becomes a list of function pointers. Note that the itable corresponds to the interface type, not the dynamic type. In terms of our example, the itable for Stringer holding type Binary lists the methods used to satisfy Stringer, which is just String: Binary's other methods (Get) make no appearance in the itable.

The second word in the interface value points at the actual data, in this case a copy of b. The assignment var s Stringer = b makes a copy of b rather than point at b for the same reason that var c uint64 = b makes a copy: if b later changes, s and c are supposed to have the original value, not the new one. Values stored in interfaces might be arbitrarily large, but only one word is dedicated to holding the value in the interface structure, so the assignment allocates a chunk of memory on the heap and records the pointer in the one-word slot. (There's an obvious optimization when the value does fit in the slot; we'll get to that later.)

To check whether an interface value holds a particular type, as in the type switch above, the Go compiler generates code equivalent to the C expression s.tab->type to obtain the type pointer and check it against the desired type. If the types match, the value can be copied by by dereferencing s.data.

To call s.String(), the Go compiler generates code that does the equivalent of the C expression s.tab->fun[0](s.data): it calls the appropriate function pointer from the itable, passing the interface value's data word as the function's first (in this example, only) argument. You can see this code if you run 8g -S x.go (details at the bottom of this post). Note that the function in the itable is being passed the 32-bit pointer from the second word of the interface value, not the 64-bit value it points at. In general, the interface call site doesn't know the meaning of this word nor how much data it points at. Instead, the interface code arranges that the function pointers in the itable expect the 32-bit representation stored in the interface values. Thus the function pointer in this example is (*Binary).String not Binary.String.

The example we're considering is an interface with just one method. An interface with more methods would have more entries in the fun list at the bottom of the itable.

Computing the Itable
Now we know what the itables look like, but where do they come from? Go's dynamic type conversions mean that it isn't reasonable for the compiler or linker to precompute all possible itables: there are too many (interface type, concrete type) pairs, and most won't be needed. Instead, the compiler generates a type description structure for each concrete type like Binary or int or func(map[int]string). Among other metadata, the type description structure contains a list of the methods implemented by that type. Similarly, the compiler generates a (different) type description structure for each interface type like Stringer; it too contains a method list. The interface runtime computes the itable by looking for each method listed in the interface type's method table in the concrete type's method table. The runtime caches the itable after generating it, so that this correspondence need only be computed once.

In our simple example, the method table for Stringer has one method, while the table for Binary has two methods. In general there might be ni methods for the interface type and nt methods for the concrete type. The obvious search to find the mapping from interface methods to concrete methods would take O(ni × nt) time, but we can do better. By sorting the two method tables and walking them simultaneously, we can build the mapping in O(ni + nt) time instead.

Memory Optimizations
The space used by the implementation described above can be optimized in two complementary ways.

First, if the interface type involved is empty—it has no methods—then the itable serves no purpose except to hold the pointer to the original type. In this case, the itable can be dropped and the value can point at the type directly:


Whether an interface type has methods is a static property—either the type in the source code says interface{} or it says interace{ methods... }—so the compiler knows which representation is in use at each point in the program.

Second, if the value associated with the interface value can fit in a single machine word, there's no need to introduce the indirection or the heap allocation. If we define Binary32 to be like Binary but implemented as a uint32, it could be stored in an interface value by keeping the actual value in the second word:


Whether the actual value is being pointed at or inlined depends on the size of the type. The compiler arranges for the functions listed in the type's method table (which get copied into the itables) to do the right thing with the word that gets passed in. If the receiver type fits in a word, it is used directly; if not, it is dereferenced. The diagrams show this: in the Binary version far above, the method in the itable is (*Binary).String, while in the Binary32 example, the method in the itable is Binary32.String not (*Binary32).String.

Of course, empty interfaces holding word-sized (or smaller) values can take advantage of both optimizations:


Method Lookup Performance
Smalltalk and the many dynamic systems that have followed it perform a method lookup every time a method gets called. For speed, many implementations use a simple one-entry cache at each call site, often in the instruction stream itself. In a multithreaded program, these caches must be managed carefully, since multiple threads could be at the same call site simultaneously. Even once the races have been avoided, the caches would end up being a source of memory contention.

Because Go has the hint of static typing to go along with the dynamic method lookups, it can move the lookups back from the call sites to the point when the value is stored in the interface. For example, consider this code snippet:

1   var any interface{}  // initialized elsewhere
2   s := any.(Stringer)  // dynamic conversion
3   for i := 0; i < 100; i++ {
4       fmt.Println(s.String())
5   }
In Go, the itable gets computed (or found in a cache) during the assignment on line 2; the dispatch for the s.String() call executed on line 4 is a couple of memory fetches and a single indirect call instruction.

In contrast, the implementation of this program in a dynamic language like Smalltalk (or JavaScript, or Python, or ...) would do the method lookup at line 4, which in a loop repeats needless work. The cache mentioned earlier makes this less expensive than it might be, but it's still more expensive than a single indirect call instruction.

Of course, this being a blog post, I don't have any numbers to back up this discussion, but it certainly seems like the lack of memory contention would be a big win in a heavily parallel program, as is being able to move the method lookup out of tight loops. Also, I'm talking about the general architecture, not the specifics o the implementation: the latter probably has a few constant factor optimizations still available.

More Information
The interface runtime support is in $GOROOT/src/pkg/runtime/iface.c. There's much more to say about interfaces (we haven't even seen an example of a pointer receiver yet) and the type descriptors (they power reflection in addition to the interface runtime) but those will have to wait for future posts.

Code
Supporting code (x.go):

package main

import (
 "fmt"
 "strconv"
)

type Stringer interface {
 String() string
}

type Binary uint64

func (i Binary) String() string {
 return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
 return uint64(i)
}

func main() {
 b := Binary(200)
 s := Stringer(b)
 fmt.Println(s.String())
}
Selected output of 8g -S x.go:

0045 (x.go:25) LEAL    s+-24(SP),BX
0046 (x.go:25) MOVL    4(BX),BP
0047 (x.go:25) MOVL    BP,(SP)
0048 (x.go:25) MOVL    (BX),BX
0049 (x.go:25) MOVL    20(BX),BX
0050 (x.go:25) CALL    ,BX
The LEAL loads the address of s into the register BX. (The notation n(SP) describes the word in memory at SP+n. 0(SP) can be shortened to (SP).) The next two MOVL instructions fetch the value from the second word in the interface and store it as the first function call argument, 0(SP). The final two MOVL instructions fetch the itable and then the function pointer from the itable, in preparation for calling that function.


