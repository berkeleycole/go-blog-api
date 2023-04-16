# Blog API

## Onion Structure

MVC revolves around models and controllers, it feels more like a factory. The model is the mold, the controllers are inner workings of factory logic, and most of the time we don't use the view as much more than a router now anymore anyway. We make assembly lines of logic that create and manipulate instances of models and the model+ORM knows how it should be input into the database. 

In Onion structure, it is more like realms of responibility. It uses the separation of concerns to allow room to grow. What a thing is and does is the same concern, so that logic lives together. But how a thing is translated to the database is a separate concern and can change if the database implementation changes. Therefore the logic of I/O lives in the database folder. Similarly, how the entity is accessed from the outside should not be the concern of the entity, because that should be able to change without making any changes intrinsic to the entity. 

There is also a heirarchy of dependencies in the onion structure, easily coupled with the separation of concerns, that maintains simplicity in the most base systems and keeps complexity from growing by keeping logic in the sphere of its own influence as much as possible. You can see the heierarchy of dependency in the main.go file, where the database is the lowest level, and as such does not need to know anything about any of the other levels. The Server has to interact with business entities and outside requests and is therefore the highest level and is dependent on the other levels, though on the database only indirectly.

One thing that was confusing to me is that there was a very clear naming convention in MVC that helped dispel any mystery of what was happening. The model was a singular name, the controller was a plural version of the same, etc.. Of course with every naming system there is a time when it is confusing and a point at which you're just used to it and it makes sense. Looking at this file structure is felt hard to follow in that files of the same name existed in most folders. But it makes sense in that your foundational business entities will appear across all concerns.

### Folder breakdown

http
- handling requests and routing is an http concern, so that logic lives in the http/server file 

databases
- each database implementation lives in its own folder

business entities
- get their own root level folder
- service file is the main file but other logic files can be added to offload more complex processes

### Issues
- "writings" is a hard word to use because the singular is so easily read as a verb and not the singular noun form, so it makes some of the syntax look funny/easy to mistake the struct Writing for a method. 