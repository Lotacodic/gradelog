# gradelog

A fast, lightweight, and dynamic Command-Line Interface (CLI) engine built in Go to calculate individual CGPA, map course units, and determine academic grading ranges. 

`gradelog` processes dynamic, multi-argument student course data straight from the terminal, handles input validation seamlessly, and outputs a structured breakdown of academic performance.

---

## ✨ Features

- **Dynamic Argument Parsing:** Accept any given number of courses directly in a single command execution.
- **Robust Input Validation:** Instantly catches out-of-range scores (e.g., scores > 100 or < 0) and halts execution with precise error context.
- **Modular Architecture:** Built with clean backend separation of concerns across dedicated evaluation layers.
- **Instant Classification:** Automatically maps calculated CGPA to standard academic tiers (e.g., First Class, Second Class Upper).

---

## 🛠️ Project Structure

The project utilizes an idiomatic, flat package layout where all files share the `main` package for optimal internal compilation performance:

```text
gradelog/
├── go.mod          # Module definition and dependencies
├── main.go         # Entry point: Orchestrates CLI args parsing and data pipeline
├── checkgrade.go   # Evaluation logic converting numeric scores to letter grades
├── gradepoint.go   # Maps letter grades to quality points and computes GPA
└── cgparange.go    # Tiers and routes the final CGPA to standard honor classifications
```

## 🚀 Getting Started

# Prerequisites
- [Go](https://go.dev/doc/install) (1.16 or higher recommended)

# Installation
Clone the repository and navigate to the project directory:

git clone [https://github.com/Lotacodic/gradelog.git](https://github.com/Lotacodic/gradelog.git)
cd gradelog

# Running the Tool
Because the project is split into modular files within the same package, run or build the application using the directory wildcard (.):

go run . "COURSE_CODE CREDIT_UNITS SCORE" ["COURSE_CODE CREDIT_UNITS SCORE" ...]

## 📊 Usage Examples

# 1. Successful Calculation
Pass your course details as space-separated values within quoted strings.

go run . "MATH101 3 70" "ENG101 2 70" "GST101 1 54" "FEG101 3 78" "CHE101 3 47"

# Output:
70 -> A -> 5 -> 15
70 -> A -> 5 -> 10
54 -> C -> 3 -> 3
78 -> A -> 5 -> 15
47 -> D -> 2 -> 6
myCgpa: 4.08 -> Second class honours (Upper Division)

# 2. Error Handling & Validation
The tool guarantees data integrity. If any input contains an invalid score boundary, it safely handles the exception:

go run . "MATH101 3 70" "ENG101 2 70" "GST101 1 54" "FEG101 3 78" "CHE101 3 101"

# Output:
70 -> A -> 5 -> 15
70 -> A -> 5 -> 10
54 -> C -> 3 -> 3
78 -> A -> 5 -> 15
Error: Score 101 in "CHE101" is out of range (0-100)!

## ⚙️ Compilation / Deployment

To compile the engine down into a single, high-performance binary executable that can run natively without installing Go:
go build .
./gradelog "MATH101 3 70" "ENG101 2 70"

## 📝 License
Distributed under the MIT License. See LICENSE for more information.