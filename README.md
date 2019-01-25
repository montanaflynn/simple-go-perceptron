# Perceptron

The simplest [Perceptron](https://en.wikipedia.org/wiki/Perceptron) you'll ever see.

Ported directly from https://github.com/victorqribeiro/perceptron in JavaScript.

## How to use

Let's suppose you have the following data set:

| Height (cm) | Weight (kg) | Class (0-1) |
|-------------|-------------|-------------|
| 180         | 80          | 0           |
| 175         | 67          | 0           |
| 100         | 30          | 1           |
| 120         | 32          | 1           |

0 - adult  
1 - child

You need to process the table to this format:

```
var x = [][]float64{
	{180.0, 80.0},
	{175.0, 67.0},
	{100.0, 30.0},
	{120.0, 32.0},
}

var y = []int{0,0,1,1}
```

Then just create a new Perceptron passing the shape of the data (height and weight), the learning rate and the number of iterations. By default the learning rate is set to 0.01 and the number of iterations is set to 10.

```
p := perceptron.New(len(x[0]), 0.2, 100)
```

Call the fit function

```
p.Fit(x,y);
```

And you're all set to make predictions

```
p.Predict(float64{178.0, 70.0})
```

Super simple.
