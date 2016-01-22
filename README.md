#marsalo

[![Build Status](https://travis-ci.org/Knetic/marsalo.svg?branch=master)](https://travis-ci.org/Knetic/marsalo)
[![Godoc](https://godoc.org/github.com/Knetic/marsalo?status.png)](https://godoc.org/github.com/Knetic/marsalo)

Marsalo is a small library that makes one-liner unmarshalling of http responses easy. The library isn't complicated, it's just a convenience library for code that tends to get written over and over.

Marsalo is Esperanto for "marshal".

##Usage


Often you'll have an HTTP request or response whose body will need to be unmarshalled into a real object model. Generally, you don't care if the request body is XML or JSON, you just want the parsed content thereof. This library encapsulates that logic of unmarshalling a request or response body from the appropriate transport type.

Assume you have the following object model (just as an example):

    type SomeObjectModel struct {

      name string `json:"name" xml:"name"`
      location string `json:"location" xml:"location"`
      entries int `json:"entries" xml:"entries"`
    }

And let's say you write a client to an HTTP API which may return either XML or JSON (or maybe you can pick, whatever), but the object represented by the response is the same.

    func HandleResponse(response \*http.Response) {

      var structuredResponse SomeObjectModel

      err := marsalo.UnmarshalResponse(response, &structuredResponse)
      if(err != nil) {
        // some parsing error
        return
      }
    }

    // Otherwise, structuredResponse now contains the unmarshalled XML or JSON body of [response]

`marsalo` works similarly for writing a webserver which needs to unmarshal the request from a client:

    func HandleRequest(request \*http.Request) {

      var structuredResponse SomeObjectModel

      err := marsalo.UnmarshalRequest(request, &structuredResponse)
      if(err != nil) {
        // some parsing error
        return
      }
    }

###Branching

I use green masters, and heavily develop with private feature branches. Full releases are pinned and unchangeable, representing the best available version with the best documentation and test coverage. Master branch, however, should always have all tests pass and implementations considered "working", even if it's just a first pass. Master should never panic.

###License

This project is licensed under the MIT general use license. You're free to integrate, fork, and play with this code as you feel fit without consulting the author, as long as you provide proper credit to the author in your works.

###Activity

If this repository hasn't been updated in a while, it's probably because I don't have any outstanding issues to work on - it's not because I've abandoned the project. If you have questions, issues, or patches; I'm completely open to pull requests, issues opened on github, or emails from out of the blue.
