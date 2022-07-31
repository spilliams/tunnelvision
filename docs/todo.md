# todo

## cli structure

The below is just a sketch. A spitball.

```
tunnelvision
├── backend         commands that operate on a terraform backend--in this case,
│   │               I only know about s3 buckets with dynamodb lock tables.
│   │               I will build it so it can be extended later.
│   ├── graph       graph a backend (show grouping and relative size of
│   │               statefiles within)
│   ├── list        print a tree of the backend's contents (e.g. s3 objects)
│   ├── show        print definition info about the backend (e.g. s3 bucket)
│   └── search TERM search all statefiles in the backend for a term
├── completion      generate the autocompletion script for the specified shell
├── file            commands that operate on a single file
│   └── graph       graph a file
├── help            help about any command
├── module          commands pertaining to a single terraform module
│   └── graph       graph a terraform module
├── root            commands pertaining to a single terraform root
│   ├── graph       graph a terraform root
│   └── migrate     move a root from one location in the backend to another
└── stack           commands that operate on a stack of some kind (terragrunt,
    │               terraspace, tau, etc)
    └── graph       graph a stack
```

Not sure yet how this will reconcile with third-party terraform wrappers like
terragrunt, terraspace, tau, etc.

## coverage

- I would like the coverage report to be pretty like go's cover tool does.
- I would like the coverage report to show me which variables I've used in testing
- I would like the coverage report to drill into modules as long as they're local
- I would like the coverage of an integration test to show me which outputs I've used in testing
- I would like the coverage to show which data resources have been used
