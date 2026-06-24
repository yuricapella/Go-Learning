# Go Syntax, Structs, Zero Value and Literal Construction

## First example and Go syntax

So let's get into this first example now.

Now this is where some of Go's syntax might seem a little foreign to you, as opposed to some of the things that you've done before.

Go really isn't trying to be novel with syntax. It's trying to leverage your experience with those curly brace languages:

- Java
- C#
- C
- Maybe even Python

Even though Python doesn't have the curly brackets, it has the same kind of structure.

But Go also has reversed a couple of things, where instead of saying type and name, it does name and type.

And Go's reason behind that is that it reads better, both verbally and in your head.

---

## Built-in types and user-defined types

Even though Go has these built-in types, the strings and numeric symbols that all languages have, a language would be fairly useless if it didn't give you the facility to construct user-defined types, or more complex user-defined types.

That's what the keyword `struct` and `type` is going to be doing here.

---

## Declaring user-defined types with `type` and `struct`

Come back to code site.

So the keyword `type` and `struct` is how we're going to declare our own user-defined types.

Go doesn't have the keyword `class`.

Go really isn't an object-oriented programming language.

It's a, I call it a data-oriented programming language.

So this is the only way that you're going to have to declare these new types.

---

## What declaring a type gives to the compiler

When you're declaring a type, you're essentially providing the compiler two pieces of information.

You're providing the compiler:

- the amount of memory that needs to be allocated for this data;
- and then the representation of this data itself.

So you can see here that this struct type named `example` is a composite type.
type example struct {
    flag bool
    counter int16
    pi float32
}

func main() {
    //declare a variable of type example set to its
    //zero value.
    var e1 example
}

You may sometimes hear the term composite type because it's composed of these three fields existing of the built-in.

I'm not teaching you, probably not teaching you anything new here.

I just want you to look at the Go syntax.

We've made it to defining a type, giving it a name, and then laying these fields out in whatever order you feel are important, with that extra type information.

---

## Constructing a value of type `example`

Then once I've got the type declaration, you can see here on line 20, where we're constructing.

If you listen to the language I use here, I really want you to kind of like zone in on the language here.

We're constructing a value of type `example` and we're naming that variable `e1`.

And you saw the keyword `var`, so we're setting it to its zero value.

So what would the zero value of `example` be?

Well, it would mean that:

- `flag` is `false`;
- `counter` is `0`;
- `pi` is `0`.

We're going to set the full value to all fields set to its zero values.

That's what we're doing here.

So we're going to construct a value of type `example`, set to its zero value, and name it `e1`.

There it is.

---

## Formatting structs

Sometimes you'll see me use a special plus operator in the print.

It just gives you a different style of formatting.

Go has three formatting styles for structs:

fmt.prinf("%+v\n", e1)
fmt.prinf("%#v\n", e1)

- this;
- this with the sharp;
- and the plus.

Play with all of them when you do the exercise.

See which one you like the best.

I like the plus the best.

It reduces the noise in the output.

We'll see in a second.

---

## Literal construction

But what if I wanted to construct a value of type `example` where it's not going to be set to its zero value?

So Go has what we call literal construction.

Literal construction is done with these curly brackets.
e2 := example{
    flag: true,
    counter: 10,
    pi: 3.141592,
}

You will probably see a lot of code in Go in your lifetime that looks like this, especially for struct types.
e2 := example{} - empty literal construction syntax

We developers, for some reason in Go, love using this empty literal construction syntax.

But I don't like using empty literal construction for a couple of reasons.

---

## Empty literal construction vs zero value

One, as I've already said, I want to use `var` for zero value.

And though empty literal construction for a struct type will give you zero value, empty literal construction doesn't always give you zero value for all the different types there are in the system.

It is for the struct type, but not necessarily for everything else.

I'll show you an exception at some point.

But listen to the words I'm using.

Listen to the word I'm using for this syntax:

Empty literal construction.

I didn't say zero literal construction.

I said empty literal construction.

These curly brackets are really for the literal construction of a value, something that you want to set, initialize it literally to.

And when you don't literally set it to something, then you hear me saying empty.

And empty is not necessarily zero value.

---

## Code consistency

So what happens is that if we're really looking for code consistency in language and syntax, then I think we have to really want it.

We want to do our best to match it up.

So a lot of developers will use this syntax.

I think it's wrong to use when you can avoid it, because again, empty literal construction isn't necessarily zero value, though it is here at this particular type.

---

## When empty literal construction is acceptable

Alright, let's come back to the now.

If you want to construct things using empty literal construction, go right ahead.

If you're consistent, again, I'm not going to complain too hard.

But you won't see me do this.

The only time you'll see me use empty literal construction, we talk about those exceptions, you might see it on a return where I'm not going to declare a variable ahead of time.

I'll just do the empty literal construction on a return.
return example{}

That's a great, maybe, use case as an example of maybe doing that syntax, needing zero value without the variable declaration.

---

## General guideline for literal construction

But it's a general guideline.

You're going to see me use the literal construction when I want to set something other than its zero value.

And if one of these fields needs to be set to its zero value, I just won't include it.

Just won't include it.

Don't need to include a field here if it's going to be set to its zero value.

It's not necessarily adding any value.

---

## Field value syntax and comma

So you can see here also the syntax of:

field colon value comma

And it's on every single line.

People complain about this comma all the time when they're first learning Go.

They're like:

Why do I need that comma if I'm at the end?

It's very JSON-like.

And Go just chose the consistency of having the comma on every line.

But one of the things I like about having the comma on every line is that if you decide you want to rearrange something, you don't have to worry about that comma missing, which happens a lot in JSON when you're shifting things around.

So just be aware that that comma is going to be needed.

---

## Question from Stephen: setting fields after construction

Alright, Erik raised his hand here.

So from Stephen:

It's the only way to set a value inside and use it to define type, not to the zero values, at literal construction?

What Steve is asking me is:

Can I do this, Bill?

Can I do zero construction and then set fields?

You could.

And I think it's a valid question.

---

## Partially constructed values

But one of the things I'm going to teach you is if you want to stay out of trouble, you don't want partially constructed values.

And with those two lines of code of doing now on line 25 and 26, it's setting you up for failure because it's partial construction.
var e example
e.flag = true


And you're going to see me later on using local variables to gather all the state I need so I can do one construction of a user-defined type.

And this is going to keep your code cleaner and it's going to keep you out of trouble.

I've seen code return partially constructed values by accident.

But it happens.

And bugs aren't caught until they're in production.

---

## Avoiding partial construction

Alright, let's go back to the code.

So if I see things like this in code, if I see things like this in code, I'm not going to like it.

I'm going to be asking you at some point to have some sort of variable.

I don't care what it is, with what you needed.

And then what we're going to do is that literal construction and just set it up.

So we don't have any issues or partial constructions going on.

These are guidelines that I would like you to follow.

They're guidelines that I follow.

And they keep me out of a lot of trouble.

And they'll keep you out of a lot of trouble.

So don't get into that literal construction until you have everything you need to do it.
variable := true
e:= example{
    flag: variable,
}

There are times when you just don't.

Sometimes that's a smell to me that we can't do it.

I've run into it before, but normally it's a smell if we can't do it.

---

## Plus formatting and sharp formatting

Alright, so you can see the construction here.

And you can see I can run it.

And you can see here's that format with the plus V.

Let me show you what the format looks like with the sharp.

You can see the sharp is throwing a lot more type information at you.

To me, that's a little bit more noise.

I tend to like just the plus.

And you can see if you don't even do any of those, little less information.

So you'll see me do plus or less information here.

No field information at all.

So you'll see me using the plus a lot.

---

## Zero value, literal construction and dot operator

Alright, so now we've got our zero value again on that struct type.

And we have our literal construction.

And literal construction is going to be used anytime we can initialize something outside of its zero state during construction.

The other thing here is the use of the dot operator.

That's nothing novel there.

You should have hopefully seen that type of syntax before.

value dot field

value dot field

value dot field

Go is not trying to be novel with any of that whatsoever.