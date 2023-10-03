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
