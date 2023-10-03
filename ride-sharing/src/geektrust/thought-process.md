Every ride has a ride status field. The possible values for a ride status are - RIDE_STARTED, RIDE_STOPPED

Every ride has a few other fields like - start point (x1, y1), end point (x2, y2), time taken in minutes. All these fields can be used to calculate the bill for the ride

Every driver has a driver status. The possible values for a driver staus are - NOT_IN_RIDE, IN_RIDE - these denote if the driver is in a ride or not, so that we can accordingly consider if the driver is available for a ride or not while trying to match a rider with a driver

Aim:
- Great Readability
    - No magic numbers
    - No magic constants
    - Compact methods and functions
    - Expressive Code
- Great Maintainability
    - No Code Duplication
- Great Object Modelling
    - OOPS
    - Encapsulation
- Correctness
    - Correct Output
- Tests
    - Clear tests
    - Good amount of unit tests
    - Good test coverage
- Build
    - No syntax errors
    - Build should be clean and successful

---

- Using multiple parameters for a method vs one struct parameter with all inputs?

Multiple parameters is straight forward but it can have problems like giving input
in wrong order which can be correct syntax wise because the inputs in wrong order
are of same type and are adjacent. And golang does not have named parameter
method/function call like python. So, struct parameter is better as it would use
the struct fields along with the struct field values. It would be clear as to
what is being given as input

One struct parameter with all inputs is extendible - adding new fields won't break
the existing code and the method can deal with non existing fields in a particular
manner with default input considering the field as optional

- Using X and Y coordinates directly in `AddDriverInput` vs creating a
separate `Location` struct with X and Y coordinates in it?

Separate `Location` struct makes sense. `Location` field in `AddDriverInput` makes
things more clear as to - "Location is store the Location of the Driver" is more 
clear and intuitive than "X and Y fields are to store coordinates / Location of the
Driver"

Also, we can reuse `Location` struct everywhere where X and Y coordinates come up.
Also, `Location` ties X and Y together as a single unit, which makes sense

- `float` or `int` for `X` and `Y` coordinates in `Location`?

It's better to use `float` as there's no mention of how the coordinates will look like -
integers or numbers with decimal values

- Return errors everywhere or handle the error at each place?

For brevity and ease of reading - handle the error at each place and exit using `os.Exit()` with
non zero exit code or use `panic()`

A much better way would be to return errors everywhere when there's an error and handle the error
at the caller function - most probably handle at `main` function level or so. But handling error
everywhere would mean lot of such statements -

```golang
if err != nil {
    return err
}
```

or similar where we mention what kind of error occurred and when and maybe how, so that the caller
can fix the issue

For now, any function in `main` package will just exit directly. Any method or function in `pkg`
would return errors though. This way, `pkg` can be exposed to other programs and be used
with different interfaces and `main` package functions and methods need not be reusable
in other programs as they are CLI interface related implementation

- How to test add driver method?

Have a method to add driver and also a method to get driver?
or a method to check if driver exists? How will the get driver
method or driver exists method be tested though? using add
driver method? then that becomes a cyclic dependency. Hmm.
Or same test can be used to check both add driver and get driver. Hmm

`GetDriver` is a very basic method. `AddDriver` is also a very basic method, but more than
`GetDriver`. We can add test for `AddDriver` method alone
