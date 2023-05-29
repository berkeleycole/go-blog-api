# Questions

- [x] What is the third column on a struct for? 
    It's called a struct tag and in this case is used by the mongo driver to create the name of the document property, so that capitalization and namings don't have to be direct between struct and db.

- [x] How do you create methods on structs?
    You don't ...directly. You add them afterwards and put the name of the struct before the function name ```func (<struct name>) <method name> (<normal arguments>)```

- [x] Forget context... I just want to pass an instance of the DB everywhere, but how?
    Learning DDD/Onion style separation of concerns and layering for scaleable SOLID structure

- [x] Is this code making a higher order function? ```func read(writings *writings.Service) func(w http.ResponseWriter, req *http.Request) {...```
    Yes. It is a function returning another function

- [x] How do I specify the HTTP verb on a request
    Like this, ex: `.Methods("GET")`

- [x] Am I doing context right in the handlers?
    No I was not. I thought I had to create a new context var when in fact the handler already had one -- but where I was using it and how were correct at least. 
    
- [ ] What is context?? Kinda get it... but also totally don't

