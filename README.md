# Golang ast-tool

Tool to help parsing your own golang source-code (using the abstract-syntax-tree tools from the standard library) into an intermediate representation.
From this intermediate representation, we can easily generate boring and error-phrone boilerplate source-code.

This first implementation focuses on essing the work related to event-sourcing:
- Describe which events belong to a specific aggregate
- Type-strong boiler-plate code to build an aggregate from individual events
- Type-strong boiler-plate code to wrap and unwrap events into an envelope so that it can be eeasily stored and emitted

In a leter version, I would like to add JAX-RS style annotations to describe rest-services.

## Preparation
    go get github.com/MarcGrol/astTools
    cd ${GOPATH/src/github.com/MarcGrol/astTools
    go install

## Event-sourcing related annotations:

A regular golang struct definition with our own "+event"-annotation. 
    
    // {"Annotation":"Event","With":{"Aggregate":"Tour"}}
    type TourEtappeCreated struct {
        ...
    }        

This annotation is used to trigger code-generation. See [./example/example.go](./example/example.go)

## Http-server related annotations ("jax-rs"-like):

    // {"Annotation":"RestService","With":{"Path":"/person"}}
    type Service struct {
       ...
    }
    
    // {"Annotation":"RestOperation","With":{"Method":"GET", "Path":"/person/:uid"}}
    func (s Service) getPerson(uid string) (Person,error) {
        ...
    }        

## Http-client related annotations ("jax-rs"-like):

    // {"Annotation":"RestClient","With":{"Path":"/person"}}
    type RestClient interface {
          
          // {"Annotation":"RestClientMethod","With":{"Method":"GET", "Path":"/person"}}
          getPersons([]Person,error)
          
          // {"Annotation":"RestClientMethod","With":{"Method":"GET", "Path":"/person/:uid"}}
          getPerson(uid string) (p Person,error)
          
          // {"Annotation":"RestClientMethod","With":{"Method":"POST", "Path":"/person"}}
          createPerson(p Person) (p Person,error)
          
          // {"Annotation":"RestClientMethod","With":{"Method":"PUT", "Path":"/person/:uid"}}
          modifyPerson(p Person) (p Person,error)

          // {"Annotation":"RestClientMethod","With":{"Method":"DELETE", "Path":"/person/:uid"}}
          removePerson( uid string ) (error)
    }

### command:
    cd ${GOPATH/src/github.com/MarcGrol/astTools/
    ${GOPATH}/bin/astTools -input-dir ./example/

Observe that [wrappers.go](./example/wrappers.go) and [aggregates.go](./example/aggregates.go) have been created in [example](example/)

## Example integrated in tool-chain

We use the "go:generate" mechanism to trigger our astTools. See [example.go](./example/example.go).

    //go:generate astTools -input-dir .

### command:
    cd ${GOPATH/src/github.com/MarcGrol/astTools/example
    rm wrappers.go aggregates.go
    go generate
    
Observe that [wrappers.go](./example/wrappers.go) and [aggregates.go](./example/aggregates.go) have been created in [example]( example/)
