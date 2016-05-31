# Golang ast-tool

Tool to help parsing your own golang source-code from the ast (=abstract syntax tree) into an intermediate representation.
From this intermediate representation, we can easily generate boring and error-phrone boilerplate source-code.

## Preparation
    go get github.com/MarcGrol/astTools
    cd ${GOPATH/src/github.com/MarcGrol/astTools
    go install


## Example integrated in tool-chain

We use the "go:generate" mechanism to trigger our astTools. See [example.go](./example/example.go).

    //go:generate astTools -input-dir .

### command:
    cd example
    rm *Wrapper.go
    go generate

Observe that for each annotated struct wrapping functions have been created in [example]( example/)


## Raw example:

### input-file: [example.go](./example/example.go)
A regular golang struct definition with our own "+event"-annotation. 
This annotation is used to trigger code-generation

single example:

    // +event -> aggregate: tour
    type EtappeCreated struct {
	    Year                 int
	    EtappeId             int
	    EtappeDate           time.Time
	    EtappeStartLocation  strin
	    EtappeFinishLocation string
	    EtappeLength         int
	    EtappeKind           int
    }

### command:
    ${GOPATH}/bin/astTools -input-dir ./tool/example/

### result: files in dir [example]( example/)
[example/envelope.go]  (tool/example/envelope.go)

[example/TourCreatedWrapper.go]  (example/TourCreatedWrapper.go)

[example/EtappeCreatedWrapper.go] (example/EtappeCreatedWrapper.go)

[example/CyclistCreatedWrapper.go] (example/CyclistCreatedWrapper.go)

[example/EtappeResultsCreatedWrapper.go] (example/EtappeResultsCreatedWrapper.go) 

[example/GamblerCreatedWrapper.go]  (example/GamblerCreatedWrapper.go)

[example/GamblerTeamCreatedWrapper.go]  (example/GamblerTeamCreatedWrapper.go)

[example/NewsItemCreatedWrapper.go]  (example/NewsItemCreatedWrapper.go)

Each file has the following functions:

    func (s *EtappeCreated) Wrap(uid string) (*Envelope,error) {
        ....
    }
    
    func IsEtappeCreated(envelope *Envelope) bool {
        ...
    }

    func GetIfIsEtappeCreated(envelop *Envelope) (*EtappeCreated, bool) {
        ...
    }

    func UnWrapEtappeCreated(envelop *Envelope) (*EtappeCreated,error) {
        ...
    }    
    
