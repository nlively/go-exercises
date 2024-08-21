### **Module 2: Linear Algebra and Probability Basics**

#### **Lesson 1: Vectors, Matrices, and Their Operations**

**Content:**
- **Introduction to Vectors:**
  - A vector is an ordered list of numbers (scalars). Vectors are fundamental in machine learning, as they represent data points, features, weights, etc.
  - **Notation:** A vector can be denoted as \(\mathbf{v} = [v_1, v_2, \dots, v_n]\) where each \(v_i\) is a scalar.
  - **Example:** A 3-dimensional vector might be \(\mathbf{v} = [1, 2, 3]\).

- **Basic Vector Operations:**
  - **Addition:** Vectors of the same size can be added element-wise.
    - Example: \([1, 2, 3] + [4, 5, 6] = [5, 7, 9]\)
  - **Scalar Multiplication:** A vector can be multiplied by a scalar (a single number).
    - Example: \(2 \times [1, 2, 3] = [2, 4, 6]\)
  - **Dot Product:** The dot product of two vectors is the sum of the products of their corresponding elements.
    - Example: \([1, 2, 3] \cdot [4, 5, 6] = 1*4 + 2*5 + 3*6 = 32\)

- **Introduction to Matrices:**
  - A matrix is a rectangular array of numbers arranged in rows and columns. Matrices are used to represent datasets, transformations, and more in machine learning.
  - **Notation:** A matrix can be denoted as \(A\), with elements \(a_{ij}\) where \(i\) is the row index and \(j\) is the column index.
  - **Example:** A 2x3 matrix might be:
    \[
    A = \begin{bmatrix}
    1 & 2 & 3 \\
    4 & 5 & 6
    \end{bmatrix}
    \]

- **Basic Matrix Operations:**
  - **Addition:** Matrices of the same size can be added element-wise.
    - Example:
    \[
    \begin{bmatrix}
    1 & 2 & 3 \\
    4 & 5 & 6
    \end{bmatrix}
    +
    \begin{bmatrix}
    7 & 8 & 9 \\
    10 & 11 & 12
    \end{bmatrix}
    =
    \begin{bmatrix}
    8 & 10 & 12 \\
    14 & 16 & 18
    \end{bmatrix}
    \]
  - **Scalar Multiplication:** Each element of a matrix can be multiplied by a scalar.
    - Example:
    \[
    2 \times
    \begin{bmatrix}
    1 & 2 & 3 \\
    4 & 5 & 6
    \end{bmatrix}
    =
    \begin{bmatrix}
    2 & 4 & 6 \\
    8 & 10 & 12
    \end{bmatrix}
    \]
  - **Matrix Multiplication:** The product of two matrices is obtained by multiplying the rows of the first matrix by the columns of the second.
    - Example (Matrix \(A\) is 2x3 and Matrix \(B\) is 3x2):
    \[
    A \times B =
    \begin{bmatrix}
    1 & 2 & 3 \\
    4 & 5 & 6
    \end{bmatrix}
    \times
    \begin{bmatrix}
    7 & 8 \\
    9 & 10 \\
    11 & 12
    \end{bmatrix}
    =
    \begin{bmatrix}
    58 & 64 \\
    139 & 154
    \end{bmatrix}
    \]

#### **Lesson 2: Probability Theory Basics (Random Variables, Distributions)**

**Content:**
- **Introduction to Probability:**
  - Probability is the measure of the likelihood that an event will occur. It is a fundamental concept in AI, especially in fields like machine learning and statistics.

- **Random Variables:**
  - A random variable is a variable that can take on different values, each associated with a probability. Random variables are used to model uncertain outcomes.
  - **Types of Random Variables:**
    - **Discrete:** Takes on a finite number of distinct values (e.g., the roll of a die).
    - **Continuous:** Can take any value within a range (e.g., the height of a person).

- **Probability Distributions:**
  - A probability distribution describes how the values of a random variable are distributed.
  - **Discrete Probability Distribution:** Gives the probability of each possible value of a discrete random variable.
    - Example: The probability distribution of a fair die is \(P(X=x) = \frac{1}{6}\) for \(x \in \{1, 2, 3, 4, 5, 6\}\).
  - **Continuous Probability Distribution:** Described by a probability density function (PDF), where the area under the curve represents the probability.
    - Example: The normal distribution (Gaussian distribution) is a common continuous distribution characterized by its mean and variance.

- **Bayes’ Theorem:**
  - Bayes' Theorem describes the probability of an event, based on prior knowledge of conditions that might be related to the event. It is a crucial concept in AI, particularly in machine learning models that involve probabilities.
  - **Formula:** 
    \[
    P(A|B) = \frac{P(B|A) \cdot P(A)}{P(B)}
    \]
  - **Example:** Calculating the probability of a patient having a disease given a positive test result.

#### **Lesson 3: Linear Algebra in AI (Dot Product, Matrix Multiplication)**

**Content:**
- **The Role of Linear Algebra in AI:**
  - Linear algebra is foundational in many AI algorithms, especially in machine learning and deep learning, where data is often represented as vectors and matrices.

- **Dot Product in Machine Learning:**
  - The dot product is used to measure similarity between vectors, compute projections, and as part of neural network computations.
  - **Example in ML:** Calculating the weighted sum of inputs in a linear model or neuron.

- **Matrix Multiplication in Machine Learning:**
  - Matrix multiplication is essential in operations involving transformations, data processing, and neural network computations.
  - **Example in ML:** In neural networks, matrix multiplication is used to propagate input data through layers of neurons.

- **Eigenvectors and Eigenvalues:**
  - Eigenvectors and eigenvalues play a critical role in dimensionality reduction techniques like Principal Component Analysis (PCA).
  - **Eigenvectors:** Non-zero vectors that only change by a scalar factor when a linear transformation is applied.
  - **Eigenvalues:** Scalars associated with eigenvectors that represent the factor by which the eigenvector is scaled during transformation.

#### **Practical Exercise: Implementing Basic Matrix Operations in Go**

**Objective:** Implement basic vector and matrix operations in Go to solidify your understanding of linear algebra concepts.

**Instructions:**
1. **Task:** Implement the following operations in Go:
   - Vector addition
   - Scalar multiplication of a vector
   - Dot product of two vectors
   - Matrix addition
   - Scalar multiplication of a matrix
   - Matrix multiplication

2. **Example Workflow:**
   - Create functions to perform each operation.
   - Test your functions with simple example inputs to verify correctness.

**Code Example:**

```go
package main

import (
	"fmt"
)

// Vector addition
func addVectors(v1, v2 []float64) []float64 {
	result := make([]float64, len(v1))
	for i := range v1 {
		result[i] = v1[i] + v2[i]
	}
	return result
}

// Scalar multiplication of a vector
func scalarMultiplyVector(scalar float64, v []float64) []float64 {
	result := make([]float64, len(v))
	for i := range v {
		result[i] = scalar * v[i]
	}
	return result
}

// Dot product of two vectors
func dotProduct(v1, v2 []float64) float64 {
	var result float64
	for i := range v1 {
		result += v1[i] * v2[i]
	}
	return result
}

// Matrix addition
func addMatrices(m1, m2 [][]float64) [][]float64 {
	rows := len(m1)
	cols := len(m1[0])
	result := make([][]float64, rows)
	for i := range m1 {
		result[i] = make([]float64, cols)
		for j := range m1[i] {
			result[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return result
}

// Scalar multiplication of a matrix
func scalarMultiplyMatrix(scalar float64, m [][]float64) [][]float64 {
	rows := len(m)
	cols := len(m[0])
	result := make([][]float64, rows)
	for i := range m {
		result[i] = make([]float64, cols)
		for j := range m[i] {
			result[i][j] = scalar * m[i][j]
		}
	}
	return result
}

// Matrix multiplication
func multiplyMatrices(m1, m2 [][]float64) [][]float64 {
	rowsM1 := len(m1)
	colsM1 := len(m1[0])
	colsM2 := len(m2[0])
	result := make([][]float64, rowsM1)
	for i := range result {
		result[i] = make([]float64, colsM2)
		for j := range result[i] {
			for k := 0; k < colsM1; k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return result
}

func main() {
	// Example usage
	v1 := []float64{1, 2, 3}
	v2 := []float64{4, 5, 6}
	m1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	m2 := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// Vector operations
	fmt.Println("Vector Addition:", addVectors(v1, v2))
	fmt.Println("Scalar Multiplication of Vector:", scalarMultiplyVector(2, v1))
	fmt.Println("Dot Product:", dotProduct(v1, v2))

	// Matrix operations
	fmt.Println("Matrix Addition:", addMatrices(m1, m1))
	fmt.Println("Scalar Multiplication of Matrix:", scalarMultiplyMatrix(2, m1))
	fmt.Println("Matrix Multiplication:", multiplyMatrices(m1, m2))
}
```

**What to Learn:**
- Strengthen your understanding of linear algebra concepts that are essential in AI.
- Get comfortable with implementing and manipulating vectors and matrices in Go.
- Recognize the importance of these operations in more complex AI algorithms.

This module will provide the mathematical foundation necessary for understanding and implementing AI algorithms in subsequent modules.