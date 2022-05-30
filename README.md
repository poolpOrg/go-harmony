# go-harmony

WIP: THIS IS A WORK IN PROGRESS, IT IS BUGGY AND VERY INCOMPLETE AT THIS POINT

## What's go-harmony for ?

This package implements a music theory engine suitable to build other tools with,
it's not meant to be used by itself but as a foundation for other projects.

It doesn't rely on hardcoded lookup tables,
but rather uses notes and intervals to perform relative lookups.
As such,
the engine doesn't know what a C7 chord is,
but knows that a dominant seventh chord is a stack of intervals to apply on C,
so it will recompute the third, fifth and dominant seventh relative to C.
The same applies to scales:
the engine doesn't know what a C major scale is,
but will recompute the eight intervals leading to the octave.

This design choice was so that there's no need to keep track of lookup tables,
but rather implement structures and derive everything from them,
leading to much simpler and less error-prone code.

## Example
The package comes with a simple utility to test the APIs.

The `-note` command will parse its parameter and outout the note name along with its frequency, it defaults to 4th octave (C4) but can be raised or lowered:

```
% harmony -note 'C'
C 261.63
% harmony -note 'C5'
C 523.25
% harmony -note 'A'
A 440
% harmony -note 'Cb'
Cb 493.88
% harmony -note 'C#'
C# 277.18
% harmony -note 'Cbb'
Cbb 466.16
```

The `-chord` command will parse its parameter and outout the chord name along with the notes that constitute it and their frequencies:

```
% harmony -chord C
Cmaj
   C 261.63
   E 329.63
   G 392
% harmony -chord C7
C7
   C 261.63
   E 329.63
   G 392
   Bb 466.16
% harmony -chord C7b5
C7dim5
   C 261.63
   E 329.63
   Gb 369.99
   Bb 466.16
% harmony -chord Csus2
Csus2
   C 261.63
   D 293.66
   G 392
% harmony -chord Cadd9
Cadd9
   C 261.63
   E 329.63
   G 392
   D 293.66
% harmony -chord C5   
C5
   C 261.63
   G 392
% harmony -chord C6
C6
   C 261.63
   E 329.63
   G 392
   A 440
```

The `-scale` command will parse its parameter and output the notes in that scale:
```
% harmony -scale Caeolian
C
   C 261.63
   D 293.66
   Eb 311.13
   F 349.23
   G 392
   Ab 415.3
   Bb 466.16
   C 523.25
% harmony -scale Clocrian
C
   C 261.63
   Db 277.18
   Eb 311.13
   F 349.23
   Gb 369.99
   Ab 415.3
   Bb 466.16
   C 523.25
% harmony -scale Baeolian 
B
   B 493.88
   C# 277.18
   D 293.66
   E 329.63
   F# 369.99
   G 392
   A 440
   B 987.77
%
```

And finally,
for now,
the `-notes` command will build a chord from a given set of notes:
```
% harmony -notes "C,E,G"    
Cmaj
   C 261.63
   E 329.63
   G 392
% harmony -notes "C,Eb,G,B" 
C-M7
   C 261.63
   Eb 311.13
   G 392
   B 493.88
% harmony -notes "C,Eb,G,Bb"
Cmin7
   C 261.63
   Eb 311.13
   G 392
   Bb 466.16
% harmony -notes "C,F,G"    
Csus4
   C 261.63
   F 349.23
   G 392
```