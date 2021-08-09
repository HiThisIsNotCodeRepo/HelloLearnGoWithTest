# Pointers and errors

## Private outside

> In Go if a symbol (variable,types,functions etc) starts with lowercase symbol then it's private outside the package it's defined in.

## Copy by value

In go when call function or method the arguments are copied, and whatever we do with the data will not change the
original data. To change the original data we have to pass the pointer.

## Create new types from existing ones

Syntax `type MyName OriginalType`