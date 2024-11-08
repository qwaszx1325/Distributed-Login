## Usage

Create a DlErr


``` go
func foo() *dlerr.DlError {
    // Case 1.
   dlErr := kgserr.New(dlerr.InternalServerError, "Your error message")

   // Case 2: New a dlerr with other error
   otherErr := someService()
   dlErr = dlerr.New(dlerr.InternalServerError, "Your error message",otherErr)

   return dlErr
}
```

Compare with dlCode

``` go
func foo() {
    dlErr := dlerr.New(dlerr.InternalServerError, "Your error message")
    if dlErr.Code() == dlerr.InternalServerError {
        // Do something...
    }
}
```