# Instructions

Now we are going to look at Unit Tests and Integration Tests. While unit tests always take results from a single unit, such as a function call, integration tests may aggregate results from various parts and sources.

Our integration tests are going to pretend to be a user - like you were with Thunder Client - and create HTTP requests to the endpoints.

Our unit tests are going to test the individual functions in our ```items``` package.

Firstly from the terminal, navigate into the ```items folder``` using the command ```cd items```

Now run the following command from the terminal:
```go test -v```
(the -v is not strictly needed, but it provides more information. -v is short for -verbose)

You should see that it has executed and passed one test. Open up the file  ```items_test.go``` in the ```items``` folder.

!["TestAllproductsAreReturned"](/images/unittest.PNG "unittest")

## Task 1 - Your first unit test
Currently we only have one unit test, and it tests the function ```GetProducts``` in the items.go file. Let's create a test for the ```GetProduct``` function. Take a look at the function:

!["getProduct"](/images/getProduct.PNG "getProduct")

There are a couple of important things to consider here.
1. The function accepts one parameter, of type string
2. The function returns two things - a single Product which is a struct type, and an error.

Reading the code, we can see there is a function called ```strcov.Atoi```. This functions converts the ```id``` parameter from a string to an integer. However, if the string passed to this function is not able to be converted it returns an err. Remember in part 1 we sent a string to the ```/products/:id``` endpoint? The error was coming from this function.

Hmm. We're probably going to need more than one test. Let's start off with the simplest test the "Happy Path" test. A "happy path" test is checking that things work as we expect when provided valid data; it does not test for error handling or exceptions. They are often the tests written first.

Let's create a test called ```TestSingleProductIsReturned``` as below:

!["single"](/images/TestSingleProduct.PNG "single")

Let's not worry too much about the ```t *testing.T``` part for now. Just note that tests are written the same as functions, using the ```func``` keyword and we have named it ```TestSingleProductIsReturned```.

It's a good idea to name your tests about the expected outcome.

### Note: In Go, all tests *must* start with the word Test, and be in a file that ends with _test.go

Generally, all unit tests work on the Arrange, Act, Assert pattern. That is, test follow the following logical structure:

1. Arrange inputs and targets. Arrange steps should set up the test case. Does the test require any objects or special settings? Does it need to prep a database? Does it need to log into a web app? Handle all of these operations at the start of the test.
2. Act on the target behavior. Act steps should cover the main thing to be tested. This could be calling a function or method, calling a REST API, or interacting with a web page. Keep actions focused on the target behavior.
3. Assert expected outcomes. Act steps should elicit some sort of response. Assert steps verify the goodness or badness of that response. Sometimes, assertions are as simple as checking numeric or string values. Other times, they may require checking multiple facets of a system. Assertions will ultimately determine if the test passes or fails.

### Setting up our test
So since we know we need to call the GetProducts function, and that function returns two things, we can write the code just as we would anywhere else: 

!["getprodcall"](/images/getProduct1.png "getproductfunction")

That's us *acting* against a function, and getting the results. Since we now there are products with values 1-4, we should get a valid response back. So let's *assert* that we don't get an error in code:

!["getprodcall"](/images/getProduct2.png "getproductfunction")

Excellent, now if we get an *unexpected* error, we will fail the test and log it out.

### Note we use t.Fail() to manually mark the test as failed, and t.Log() to output the error details for debugging

Okay,so we have some code to check that we didn't get an error (remember, this is a happy path test. We are not expecting it to fail, but we never know what's going to happen...)

Now we want to assert that the function GetProduct actually retured the correct product! Again, we are going to use the conditional ```if``` statement to check the returned value:

!["getprodcall"](/images/getProduct3.png "getproductfunction")

Now, save the file, and on the command line run the tests using the command ```go test .```. Fingers crossed!

```
=== RUN   TestAllProductsAreReturned
--- PASS: TestAllProductsAreReturned (0.00s)
=== RUN   TestSingleProductIsReturned
--- PASS: TestSingleProductIsReturned (0.00s)
```

Congratulations, you've written your first unit test!

You can see here that even though we have had an entire product returned, we're only checking the ID. Hmm. We coud be missing out on bugs in the other fields of the Product struct, like the Description.

So how do we fix that?

## Task 2 - setting up expected data

One important things in testing is to explicitly set up the values, or object we are expecting to be returned and is part of our *Arrange* step, which was previously missing.

So we need to create our own instance of the Product struct, and populate it with the data we are expecting back. 

In Go, we can create an *instance* of a struct by calling it directly, like so:
```expectedProduct := Product()```

At this point, our created instance ```expectedProduct``` doesn't have any data in it - things like ID will be 0, and Description will be an empty string ("").

#### Information
There are many ways to create a struct with data. One of the most obvious ways is to continue the code above and directly assign the values like so:

```
expectedProduct := Product()
expectedProduct.ID = 1
expectedProduct.Name = "Generic Name 1"
//all the other fields
```

However, another way also exists which allows us to assign values at the same time we create the instance of the Product struct:

```
expectedProduct := Product{ID=1, Name="Generic Name 1", ...}
```

Fortunately, because the current store uses *hard coded* data, we can go into the ```items.go``` file and copy the first line:

!["task2-1"](/images/task2-1.png "task2-1")

And then use this to create in our test the expectedProduct code like so:

!["task2-2"](/images/task2-2.png "task2-2")

#### Please note, you can have this all on one line, I have made it multi-line for easier reading and editing.

Now we have our ```expectedProduct``` instance, we can use this for our assertion rather than having to explicitly checking each field. Replace the ```if product.ID != 1``` section of code with the below:

!["task2-3"](/images/task2-3.png "task2-3")

Now if we run the tests again using ```go test -v``` we can see our tests still pass. But wait....how do we know it's actually testing anything?

Well fortunately, now we have our *Arrange* step, we can edit our expectedProduct values when we instantiate it. Try changing the Name, ID, Description, or indeed any values of the ```expectedProduct```. For instance, I changed the Name value to be "Generic Item 2" and then ran the tests again, causing the test to fail:

```
=== RUN   TestSingleProductIsReturned
    items_test.go:34: Expected product to be {1 Generic Item 2 A generic item we sell A longer description of the generic item we sell 56.99 } but got {1 Generic Item 1 A generic item we sell A longer description of the generic item we sell 56.99 }
--- FAIL: TestSingleProductIsReturned (0.00s)
```

#### Don't forget to change the values back and to save the file once you are finished!

Here's the full code:

!["task2-4"](/images/task2-4.png "task2-4")

# Task 3 - making things unhappy

Ok, so we have our "happy path" test. But we want to check that the error handling works if something goes wrong, or the function is given invalid data.

To find out what unhappy path tests we might need to code, we need to have a look at the ```GetProduct``` code. 

!["task3-1"](/images/task3-1.png "task3-1")

Reading the code, how many ways are there to get to the last line where it will return an error?

If your struggling, remember things you know:
1. Every time the program starts there will be exactly 4 products
2. The function appears to accept a string, but in our struct the ID property is an int, so the code has to convert the incoming string to a integer.
3. Write some strings down and be the computer - imagine the string coming in and fo line-by-line to determine the outcome.

#### Don't skip over to the next part until you have tried to imagine what's going to happen givena couple of different strings....









#### Here are a couple of inputs that will produce errors:
1. "hello" - because "hello" cannot be converted to an integer
2. "10" - because there is no product with that ID in our slice

So we probably need to write at least two unit tests.

## Create The Tests
Your task is to create the two unhappy path tests. Ensure that:
1. Each test is well-named for the outcome you are expecting
2. You assert the error returned
3. The test fails if the error is not returned
4. The test fails if the error text is different from that in the GetProduct function

#### Remember: The name of your test *must* start with the word "Test"!

# Task 4 - Table tests

Now, if you've been paying attention, you'll probably notice that there's a lot of duplication between both of the tests you have written. Really, all we have changed is the input both times to the ```GetProduct``` function, and we have two very similar looking tests!

Fortunately Go has a solution to this called ```table tests```, and it allows you to test the same code outcomes without having to write multiple separate tests. What does that mean in practice? Let's have a look at refactoring our existing tests.

To use table tests, Go requires you create a slice of struct. This struct contains your *input* and your *expected output* generally (the fields do not have to named as I have done below; you can have whatever names you want, and as many fields of whatever type as you need).

!["task4-1"](/images/task4-1.png "task4-1")

Now we need to initialise that slice of structs with data - the input and expected output values:

!["task4-2"](/images/task4-2.png "task4-2")

Now since we have a slice we can *range* over it and provide the input to the GetProduct function multiple times. And because we are looping over that slice, the input will be different every time, but we don't have to write separate tests.

!["task4-3"](/images/task4-3.png "task4-3")

Now save the file and run the tests!

#### For Your Information...I probably would *not* combine the two unit tests above into a single table driven test like this, as they test very different unhappy paths. But that's a long, nerdy discussion best suited for later.

# Solutions
## Task 3

!["task3-1a"](/images/task3-1a.png "task3-1a")

!["task3-1b"](/images/task3-1b.png "task3-1b")