### **Module 3: Basic Data Handling and Manipulation**

#### **Lesson 1: Introduction to Data Structures (Arrays, Slices, Maps)**

**Content:**
- **Understanding Data Structures in Go:**
  - Data structures are critical for storing, organizing, and manipulating data efficiently in any programming task, including AI.

- **Arrays:**
  - An array is a fixed-size sequence of elements of the same type.
  - **Declaration:** 
    ```go
    var arr [5]int
    arr = [5]int{1, 2, 3, 4, 5}
    ```
  - **Key Points:**
    - The size of the array is fixed and cannot be changed.
    - Arrays are zero-indexed.

- **Slices:**
  - A slice is a dynamically-sized, flexible view into the elements of an array. Slices are much more commonly used than arrays in Go.
  - **Declaration:**
    ```go
    var slice []int
    slice = []int{1, 2, 3, 4, 5}
    ```
  - **Key Points:**
    - Slices are references to arrays.
    - Slices have a length (the number of elements) and a capacity (the number of elements in the underlying array, starting from the first element in the slice).
    - Slices can be resized by appending elements.

- **Maps:**
  - A map is an unordered collection of key-value pairs, where each key is unique.
  - **Declaration:**
    ```go
    var m map[string]int
    m = map[string]int{"one": 1, "two": 2, "three": 3}
    ```
  - **Key Points:**
    - Maps are used to store data in a key-value format.
    - Maps provide fast lookups, additions, and deletions.

- **Practical Application in AI:**
  - Arrays and slices can be used to store and manipulate datasets, vectors, or matrices.
  - Maps can be useful for implementing algorithms that require fast lookups, such as storing weights in a neural network or tracking visited nodes in a graph search algorithm.

#### **Lesson 2: Reading and Writing Data in Go (CSV, JSON)**

**Content:**
- **Reading and Writing CSV Files:**
  - CSV (Comma-Separated Values) files are a common format for storing tabular data.
  - **Reading a CSV File:**
    ```go
    import (
        "encoding/csv"
        "os"
        "fmt"
    )

    func readCSV(filename string) ([][]string, error) {
        file, err := os.Open(filename)
        if err != nil {
            return nil, err
        }
        defer file.Close()

        reader := csv.NewReader(file)
        records, err := reader.ReadAll()
        if err != nil {
            return nil, err
        }
        return records, nil
    }
    ```
  - **Writing to a CSV File:**
    ```go
    import (
        "encoding/csv"
        "os"
    )

    func writeCSV(filename string, data [][]string) error {
        file, err := os.Create(filename)
        if err != nil {
            return err
        }
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        for _, record := range data {
            if err := writer.Write(record); err != nil {
                return err
            }
        }
        return nil
    }
    ```

- **Reading and Writing JSON Files:**
  - JSON (JavaScript Object Notation) is a lightweight data interchange format that's easy for humans to read and write and easy for machines to parse and generate.
  - **Reading a JSON File:**
    ```go
    import (
        "encoding/json"
        "os"
        "fmt"
    )

    func readJSON(filename string, v interface{}) error {
        file, err := os.Open(filename)
        if err != nil {
            return err
        }
        defer file.Close()

        decoder := json.NewDecoder(file)
        if err := decoder.Decode(v); err != nil {
            return err
        }
        return nil
    }
    ```
  - **Writing to a JSON File:**
    ```go
    import (
        "encoding/json"
        "os"
    )

    func writeJSON(filename string, v interface{}) error {
        file, err := os.Create(filename)
        if err != nil {
            return err
        }
        defer file.Close()

        encoder := json.NewEncoder(file)
        if err := encoder.Encode(v); err != nil {
            return err
        }
        return nil
    }
    ```

- **Practical Application in AI:**
  - Data in CSV format is often used for machine learning tasks, such as datasets with features and labels.
  - JSON format is used for configuration files, saving model parameters, or exchanging data between systems.

#### **Lesson 3: Basic Data Cleaning (Handling Missing Values, Normalization)**

**Content:**
- **Data Cleaning in AI:**
  - Data cleaning is the process of preparing raw data by dealing with inconsistencies, missing values, and errors. Clean data is crucial for building reliable machine learning models.

- **Handling Missing Values:**
  - **Approaches:**
    - **Remove Missing Values:** Discard rows or columns with missing data.
    - **Imputation:** Fill in missing values with a calculated value, such as the mean, median, or mode of the column.
    - **Example in Go:**
      ```go
      func imputeMissingValues(data [][]string, column int, imputeValue string) {
          for i := range data {
              if data[i][column] == "" {
                  data[i][column] = imputeValue
              }
          }
      }
      ```

- **Normalization:**
  - Normalization (or scaling) is the process of adjusting the range of data values. It's important when features have different scales, as many machine learning algorithms are sensitive to the magnitude of the input data.
  - **Min-Max Normalization:** Rescale data to a fixed range, typically [0, 1].
    - **Formula:** 
      \[
      x' = \frac{x - \text{min}(x)}{\text{max}(x) - \text{min}(x)}
      \]
  - **Z-Score Normalization:** Rescale data to have a mean of 0 and a standard deviation of 1.
    - **Formula:** 
      \[
      x' = \frac{x - \mu}{\sigma}
      \]
  - **Example in Go:**
    ```go
    func minMaxNormalize(data []float64) []float64 {
        min := data[0]
        max := data[0]
        for _, v := range data {
            if v < min {
                min = v
            }
            if v > max {
                max = v
            }
        }
        normalized := make([]float64, len(data))
        for i, v := range data {
            normalized[i] = (v - min) / (max - min)
        }
        return normalized
    }
    ```

- **Practical Application in AI:**
  - Clean and normalized data is essential for building robust and accurate AI models. Handling missing values and ensuring consistent scales across features prevent biases and improve model performance.

#### **Practical Exercise: Data Loading and Cleaning in Go**

**Objective:** Load a dataset from a CSV file, clean the data, and perform basic normalization.

**Instructions:**
1. **Task:** Write a Go program that:
   - Reads a dataset from a CSV file.
   - Identifies and imputes missing values.
   - Normalizes the data using min-max normalization.

2. **Example Workflow:**
   - Load the CSV data into a slice of slices.
   - Identify columns with missing values and fill them with the mean or median.
   - Normalize the numeric columns using min-max normalization.
   - Print the cleaned and normalized data.

**Code Example:**

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func imputeMissingValues(data [][]string, column int, imputeValue string) {
	for i := range data {
		if data[i][column] == "" {
			data[i][column] = imputeValue
		}
	}
}

func minMaxNormalize(data []float64) []float64 {
	min := data[0]
	max := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	normalized := make([]float64, len(data))
	for i, v := range data {
		normalized[i] = (v - min) / (max - min)
	}
	return normalized
}

func main() {
	// Example CSV file with some missing data and values to normalize
	filename := "example.csv"

	// Load data
	data, err := readCSV(filename)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Impute missing values for a specific column (e.g., column 1)
	imputeMissingValues(data, 1, "0")

	// Convert a column to float and normalize it (e.g., column 2)
	var floatData []float64
	for _, row := range data {
		val, err := strconv.ParseFloat(row[2], 64)
		if err == nil {
			floatData = append(floatData, val)
		}
	}
	normalizedData := minMaxNormalize(floatData)

	// Update the normalized data back to the original dataset
	for i := range data {
		data[i][2] = strconv.FormatFloat(normalizedData[i], 'f', 6, 64)
	}

	// Print cleaned and normalized data
	for _, row := range data {
		fmt.Println(row)
	}
}
```

**What to Learn:**
- Gain hands-on experience with data loading and manipulation in Go.
- Learn techniques for handling missing data and normalizing datasets, which are critical steps in preparing data for AI models.
- Develop an understanding of how data quality impacts AI model performance.

This module will equip you with the practical skills needed to handle and preprocess data, an essential step before applying AI algorithms.