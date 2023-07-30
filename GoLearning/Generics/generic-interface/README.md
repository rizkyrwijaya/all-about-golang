# Generic Interface

Using Generics it is possible to create an interface. Rather than having multiple implementation, it should be possible to have a single implementation.

HttpRepository will be your general http approach of creating a library.

The generic should be reusable for different kind of models

## Normal Repository Analysis

In this approach we are using a simple repository approach with an interface. 

For this I assume we can say that the repository is created for solely the defined model, that is specialized for it.

Based on my time trying this here are some analysis regarding this approach.

### Workflow

Regarding using this approach in a day-to-day workflow basis means that there will
be repetitive codes. Sure it might seem like a little effort, but all the little effort will pile up
in the future and there will always be unnecessary story points in need to code a new repository.

If you are very specific about test cases, it also needs to note that there will also be effort in creating a 
specific test case for each repository in the future

Lets say we keep using this approach, here are some of the minimum aspect of development needed.

- Create the Repository which consists of:
    - New Interface
    - New Model
    - New Implementation
    - (if needed) New Mock
- Create the initialization flow 

### Usage

Since it is specialized, additional custom flows can be easily integrated without changing the other
repository. I think this is one of the strong points about creating a specialized repository is we can have different kind of processing.
But Even then those kind of specialized processing can be generalized and probably we can add it to the generic flow.

## Generic Repository Approach

In this flow, we are creating a generic repository that only focuses on hitting the URL and then decode the response. Since we use Type Instatiation
we can easily reuse the decoding flow etc. With further development we can also cater inserting query Params or even other type of http repository.


### WorkFlow

Trying to create a new repository initialization is easier than I expected. As seen on app.go the expected changes are as follows

- Create new Model (can be anywhere)
- Initiate with New


### Usage

Something that would be quite tricky is using a 3rd party decoder, but there might be ways we can handle that.