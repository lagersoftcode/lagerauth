# logger

simple leveled logger.

## usage:

First import with the desired dialect:  
```
import "github.com/ozkar99/logger"
import _ "github.com/ozkar99/logger/dialects/sqlite"
// import _ "github.com/ozkar99/logger/dialects/mysql"
```

Then to use pass a `*sql.DB` connection as the first parameter:
```
logger := logger.New(db)
logger.Debug("Hello %s", "mom")
```

this will create the `logs` table as well as an entry with type `Debug` on the table.  

You can also write to aditional `io.Writer` interfaces, if you pass them on the constructor: 
```
logger := logger.New(db, os.Stdout, fileHandle, ...)
```

The constructor is variadic, the first argument needs to be a `*sql.DB` the rest of the arguments  `io.Writer` or ommited like in the first example.

## supported levels:

- Debug
- Info
- Warn
- Error
- Fatal

## todo:
 - add postgresql dialect
 - add mssql dialect
