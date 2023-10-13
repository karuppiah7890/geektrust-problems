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

- Should we have driver ID in `Driver`? and rider ID in `Rider`?

I mean, if we are returning rider, or driver somewhere, as of now there's no ID as part of
the rider or driver. The ID is stored at the ride sharing app level, as part of the map's
key. Hmm. Maybe it's a good idea to add ID to the fields of rider and driver, this way, a
rider or a driver is a standalone thing

- Should we really have `AddDriverInput` and `AddRiderInput`? Or just use `Driver` and `Rider`?
Hmm. We can just use `Driver` as the input for `AddDriver` and `Rider` as input for `AddRider`.
Why do this? A few good things - tests can use `Driver` and `Rider` directly or we might end
up creating a separate `Driver` and `Rider` in test package with ID too as field, but if the
previous point is done then we can reuse the `Driver` and `Rider` in `pkg` as the `Driver`
and `Rider` itself will have ID. A few bad things about reusing `Driver` and `Rider` instead
of `AddDriverInput` and `AddRiderInput` is - later, when input changes, we might have to
have separate representation for input and separate representation at ride sharing app
level. Also, the outside interface should be stable and not keep changing, hmm. If we
merge them and use `Driver` for input and for storage, then it could change later and
break things and we will have to make breaking changes. Depends on what kind of features
and fields come up later and how we structure the input for it and how we would look at
it from ride sharing app store level. Hmm

- What are the different error outputs?

`INVALID_RIDE`

`RIDE_NOT_COMPLETED`

- When do the different error outputs happen?

When `START_RIDE` is called with `N` (1 <= N <= 5) but the `MATCH` has fewer than N matches, or
when the driver is not available, or the given ride ID already exists (ride ID has to be unique),
then print `INVALID_RIDE`

When `STOP_RIDE` is called but the given ride ID does not exist or if the ride has already
been stopped, then print `INVALID_RIDE`

When `BILL` is called but the ride ID does not exist, then print `INVALID_RIDE`

When `BILL` is called but the ride for which `BILL` has been called has not completed
using `STOP_RIDE`, then `RIDE_NOT_COMPLETED` error comes

- What does driver is not available mean?

There are no drivers nearby in the 5 km distance. Or there are no drivers who are available to
drive since they are occupied with existing rides, though they are within the 5 km distance

The question does not talk about this though - about driver(s) being occupied with existing ride(s)

- How to store all the rides?

All rides can be stored a list of rides represented with a `map` in golang with ride ID as key.
This will help us ensure that the ride ID is unique. Once a ride is stopped, we gotta remove the
ride ID and the ride from the list of rides.

We will also store the ride ID in the driver and rider, in case they are on a ride, or else it
will be empty. We can use a pointer to string and have `nil` for no ride and have a pointer to
a string (ride id) for cases when they are on a ride

- Matching rides - give all outputs at `pkg` level? All drivers within 5km distance,
instead of just 5, and then cap it to 5 at `main` level? Hmm 🤔💭🤨🧐. We can start
with this. Later we can customize it if needed and cap it at `pkg` level itself

And even the 5km distance - take it as an input at `pkg` level? Hmm, makes sense

- Should the field be named `RiderId` or just `ID`, for the `AddRiderInput` struct?
Same thought for `DriverId` vs `ID` for `AddDriverInput` struct

For now this okay. We can change it later if needed. This brings us back to the
question of merging `AddRiderInput` and `Rider` and just using `Rider`
instead of two separate structs, and merging `AddDriverInput` and
`Driver` and just using `Driver` instead of two separate structs

- Should we remember the matches done as part of `MATCH` command? or just run
`MATCH` on demand? 🤔 But if we run match on demand, we may notice a different
output. For example if a driver from old output started ride, they may not
be seen in the new on-demand match output. Also, in such a case, where on-demand
match is run the user may end up choosing a driver ID that they didn't intend to
because they chose from old output and we are starting ride on a driver from new
output that's not shown to the user. If we remember the old output,
then we can get the value from it while doing `START_RIDE`, but if the driver
is not available because they started another ride, then we should give
`INVALID_RIDE` as output I guess?

Yeah, I think we should just remember the old output for the corresponding rider,
so that we know which driver to choose when they say `START_RIDE`, instead of
running another `MATCH` behind the scenes during the processing of `START_RIDE`,
which will cause unnecessary confusion for the user using the system

- Testing

Apart from unit testing, run some integration tests to check if things work end to end,
and also to test the logic inside `main` package. Or else write tests at `main` package
level too! :) Unit tests. Integration and E2E tests can be a few, compared to unit tests.
That way, the tests are maintainable and easy to change when code changes

- When and where to use pointers and where not to use pointers?

Mostly I have used pointers to reduce memory usage - as a lot of times structs are
being used in function arguments and passed around, and this would cause copying
the struct everytime if it's a struct, so a pointer to struct makes the copy
light weight as it has to copy only the memory address / reference to the struct
and not the whole struct itself

- Where to compute the driver match? Where to do the sorting?

For now I have computed the driver match inside `pkg` and also done the sorting.
As part of the inputs I get the radius and also the max drivers asked by the
caller. I think it's okay to do the sorting - especially the distance sort,
apart from the lexicographical sort when distance is equal. Distance sort
makes sense, as closer drivers are a better match for the rider, and
showing that as one of the first few options is a good idea. Also,
doing that distance sort inside `pkg` makes sense, and getting
the K (Max number of drivers) minimum distances and giving the
driver id of those drivers - that also makes sense to be in `pkg`,
as sending back all drivers in ascending order of distance
doesn't make sense.

Fow now I'll stick with all the implementation being in `pkg`.

I'm just wondering if I should send back matched drvier to the caller
and then caller can get driver ID from that, or if I should just give IDs
of the drivers. Giving back matched driver would be nice - since
it will have driver ID and also driver location and also distance
of driver from rider, all of it calculated and also in proper
order. Hmm. Something to think about later. For now I'll leave
it as is

EDIT: If I'm sending back matched drivers list to the caller, do I have to
do the sorting? because matched drivers list will have distance, apart from
driver id and driver location. If the caller was a mobile app and  was
calling server, then the mobile app (the client) would get all the details
to show the closest driver's location and how far they are and also their IDs.
Actually, just giving location is also enough, but however the server does
the calculation of finding the closest drivers to the rider and has those
details as part of the computation / calculation, so they can just send it
instead of client redoing that calculation. Also, ride sharing app or
server doing the closest drivers calculation is best and giving the max
drivers is best, because the other way would be to just give all the
driver details and client having to do this calculation - which would be
heavy if there are too many drivers in the drivers list. This would make
the client a thick client - a very compute heavy client. We would rather
want a dumb client, thin, and one which doesn't have to do heavy
computation and server will take care of doing all the heavy lifting. So yeah,
some of the calculation like - closest in radius, max drivers, should all be
in ride sharing app inside `pkg`. Do we send matched drivers or just driver IDs
is now the question. Hmm 🤔

For now I'm leaving it, or else I will have to do some extra work in the test
and compare a list of matched drivers type instead of just a list of strings.
I also gotta do a small computation to extract driver ID alone from the list
of matched drivers in `main`. I'll leave it for now. It's not hard to make this
change later though :) if required

- Lot of types have been exported by using capital letter in the first
letter of the type name. Do we need to export all of them? For example,
we export `Driver` in `pkg`, we also export `Rider`, `MatchedDriver`
in `pkg`, but these types are not used outside `pkg` as of now

I'm gonna leave this for now. But it's something to think about later.

- Where to store the latest list of IDs of matched drivers for a given rider,
so that the rider can choose from that list by giving the number N to get the
Nth driver from that list.

Initially I thought of storing this list data inside ride sharing app struct
as part of riders. But I don't think that makes sense. Why keep latest
list of matched drivers of riders over there? It's okay, it can be at
`main` as part of the caller program, that will remember the previous
match output and based on that output, it can also understand the next
input for start ride for the same rider. Down falls? Some rider data
is present at `main` level, and some are at ride sharing app level,
but that's okay. Ride sharing app should have basic rider information
like rider ID, rider location so that it can do matching of rider
and driver. But ride sharing app does not have to remember the list
of matched drivers. One thing to note though - where this list of
matched drivers is stored, it should be cleared once the rider
starts a ride and it should be overwritten when the rider sends a
match request. Hmm

I guess it's kind of like showing a list of options on the client side
and the client storing the list, though the server created the list of
options but server doesn't store this list of options. Once the client
choose from the list of options, using a number than the option, the
option is then sent to the server. So the server does not need to be
involved in the logic of choosing an option from the list of options

Here, in our case, ride sharing app is kind of the server and all
implementation is in `pkg` and `main` is kind of the client

- What else will the ride sharing app store? The ride sharing app will
store  the rides that are happening and their status - over or not.
In order to understand the current ride the driver is on, the
ride sharing app will also store the ride id (not the whole ride here)
as part of the driver details. The same is done for the rider too :)

- How to handle errors at caller side when the list of errors are known
at implementation side? as a black box implementation? or just handle
based on the known information?

---

TODO:
- move some of the functions and methods in `main` to something inside `cmd`. For example `context` can go into a package under `cmd` called `ctx`
- move `driver`, `rider`, `ride` into packages of their own under `pkg`. It makes more sense
- Too many `panic`s all over the code. Check if the usage is correct and if and where
  `recover` would be required
- rename `context` in `main` package to `database`? or `db`?
    - as it's more of a database storing stuff
- Simplify code at `main` package level
    - `start_ride.go` is very messy
- Better error handling for different scenarios
- Integration testing
