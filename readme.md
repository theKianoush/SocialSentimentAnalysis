Instructions:
-------------
once you run the program you will be asked to enter the name of the .txt file
you want to use, make sure to include the .txt extenstion when writing the 
name of the file, Thanks!


Structure of Program:
---------------------
My biggest problem was trying to figure out how I was going to map
the .txt words to the words in the .csv
but I acomplished this by looping through every word in the .txt and use and if statment to see if 
the current word we are on matches any word in the .csv file once it found that we grab the second index of that
word, which is its score and then append that score to another array, since it is a string I can not add it up directly
then i converted all the elements in the array that holds the scores to floats
then i summed it all up and we get out final score



REPORT:
--------
1 (Abstract – an overview of the language).
Go is a statically type compiled language, often described as C for the 21st century.
It’s a popular choice for high performance, server-side applications.
It was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson who was
the co-inventor of the B and C programming languages.
Moreover, the Go language was intended to support network and multicore computing, and to
accelerate the coding process.
The source code is compiled into to machine code, which means it generally outperforms
interpreted languages, but Go is famous for its extremely fast compile times, made possible by
innovations to dependency analysis. 


2 (Approach - problems encountered). 
When trying to use this language for the first time, I encountered many problems.
I first thought that every word in the .txt files was supposed to be in .csv files, so I thought some words weren't
matching, but not every word in the .txt files are supposed to have a score.
And then another problem I had was, was that I was not able to remove punctuation/special characters from
the words in the .txt file, so maybe so words weren't mapped to words in the .csv file, due to punctuation
And the last problem that I spent the most amount of time on which was really stupid on my part was that, 
the scores in the .csv file are all strings so I was trying really hard to convert them to ints, but nothing was working
so then I realized that all I had to do was convert them to floats instead, which was a silly mistake


3 (New things learned about language).
The syntax was somewhat similar to python, it seemed really confusing at first, then when you start using it
its not so bad, just like any other new language. 
It seems like nothing is built in and you have to import packages to do almost everything. 
Its not my favorite language but I don't dislike it at all.


4 (Special notes).
I think the specail thing about it, is that it won't let you use certain function unless you use 
the error handling aspect of it. for example: the string convert to int function called strconv.Atoi()
only works if you use the error handling 

        newInt, err := strconv.Atoi("3")
        if err != nil {
            panic(err)
        }
    fmt.Println(newInt)


    
5(conclusions).
To be honest when I first read the assignment I didn't think I would be able to complete this, it seemed alot harder than it actually was.




OUTPUT:
---------
 make run
./main
Enter the name of the file you would like to use (with the .txt extension):
good.txt


[word: current_score]
gift: 1.89
used: -0.06
these: -0.85
product: 0.58
during: -0.59
loves: -1.01
works: 0.73
amazing: 1.26
anybody: -0.04
wanting: -0.02
start: 0.29
pieces: 0.06
arrived: 0.17
condition: -0.63
day: 0.49
compact: 1.03
pieces: 0.06
looks: 0.04
great: 1.48
set: 0.24
stand: 0.01
perfectly: 2.69
bottom: 0.24
stand: 0.01
little: -0.54
lips: 0.59
hang: 0.12
either: -0.19
side: 0.14
scale: 0.17
keep: -0.01
sliding: 0.09
around: 0.7
knock: -0.32
scale: 0.17
itself: -0.28
measures: 0.35
extremely: -0.64
scale: 0.17
does: 0.04
shut: -0.58
own: 0.93
left: 0.52
big: -1.22
issue: -0.81
kitchen: -0.69
scale: 0.17
used: -0.06
frustrating: -1.46
scale: 0.17
shut: -0.58
admit: -0.73
price: 0.71
whole: -1.4
bit: -0.72
mostly: -0.4
paying: 0.27
goes: 0.07
dollars: 0.24
stand: 0.01
goes: 0.07
reason: -0.12
product: 0.58
real: 0.29
coffee: -0.65
beans: 0.08
patience: 0.01
around: 0.7
less: -0.57
less: -0.57
still: 0.1
give: 0.13
great: 1.48
cups: -0.02
product: 0.58
order: -0.02
again: -0.43

Final Score: 4.71
Stars: 4
 
