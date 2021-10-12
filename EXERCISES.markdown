## Learn about the domain: Carries Cars

Carries Cars is a car-sharing service that started operating in Amsterdam 6 months ago. 
The service operates 300 vehicles that are spread out over the compact city center. 
All vehicles are fully electric vehicles because of demands by the city as part of granting the free-floating fleet operating agreement.

After the customers reserved a car, they have 20 minutes to reach it. 
If the vehicle is in good shape then they accept a rental agreement using a smartphone app. 
From that point onwards they can start driving around. 
Pricing is straightforward: you pay per minute (e.g. € 0.24) and are allowed a maximum (e.g. 250 kilometers) mileage.


## Exercise 1: Analyze existing code

Take 15 minutes to look around in the example project. What do you notice?

- Capture 3 things that you like about the example project.
    - The package/domain/module separation. Money doesn't need to know priceEngine exists
    - How the tests are written + their testing coverage
    - Well established interfaces with really clear behaviours
- List 3 questions on things that have you puzzled.
    - Why no having one go.mod considering both packages are part of the same app, also considering priceEngine depends on money package
    - On money.go, on the Equals function, why we are not comparing the amount and ISO code from the money using the functions we are calling from other?
    - Why we don't call duration as VerifiedDuration to make it more explicit that is a verified duration?
    - [Additional question, sorry!] Why do not make use of pointer receivers on trustedMoney functions, and change the data on existent pointers instead of creating new trustedMoney?

## Exercise 2: Enrich the Pricing Engine with the ability to extend the Reservation

Some customers complain that 20 minutes isn't enough to reach the vehicle. 
For those situations it should be possible to extend the Reservation of the vehicle. 
In exchange for that Carries Cars charges € 0.09 per minute.

Capture this behavior in one or more tests
Commit and push the tests
Implement the functionality within the domain model
Look for refactoring opportunities. Commit each refactor as a separate commit and push it


## Exercise 3: Enrich the Pricing Engine with the ability exceed the included mileage

Every pay-as-you-go rental includes a maximum of 250 kilometers of mileage. 
Beyond that customers should be charged an additional price of € 0.19 per kilometer.

Capture the price increase for exceeding the included mileage in one or more tests
Commit and push the tests
Implement the functionality within the domain model
Look for refactoring opportunities. Commit each refactor as a separate commit and push it
