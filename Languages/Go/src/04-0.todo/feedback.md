
# Feedback on the Go To-Do CLI Project

This document provides a review of the Go to-do list project. It covers both suggestions for improving the existing code and broader concepts that distinguish a hobbyist project from a professional-grade application.

This is a great start! The project is functional and demonstrates a good understanding of Go fundamentals, including file I/O, JSON marshaling, and basic package organization. The feedback below is intended to help you take the next step in your journey as a software engineer.

---

## 1. Biggest Differences from a Professional Project

Here are the key areas that separate this project from what you'd typically find in a professional environment.

### 1.1. Testing is Essential
**The single most significant difference is the complete lack of automated tests.**

*   **Why it's a problem:** Without tests, every change (a bug fix, a new feature) requires manual testing of the entire application and risks introducing new bugs (regressions). This is not scalable.
*   **How to improve:**
    *   **Unit Tests:** Create `_test.go` files for each of your packages. Write tests for your "pure" logic functions like `logic.FindList`, `logic.FindTask`, and the functions in `configs`. Use Go's built-in `testing` package.
    *   **Integration Tests:** You can write tests that execute the commands and verify the state of the JSON files, or even use `os/exec` to run the compiled binary and check its output.

### 1.2. CLI Parsing and User Experience (UX)
The current command-line interface relies on manually parsing `os.Args`. This is fragile and leads to a poor user experience.

*   **Why it's a problem:** Manual parsing makes it hard to handle flags (e.g., `-v` for verbose), subcommands, and optional arguments consistently. It also means you have to write your own help text, which is often incomplete.
*   **How to improve:** Adopt a mature CLI library.
    *   **[Cobra](https://github.com/spf13/cobra):** The most popular choice in the Go ecosystem (used by Docker, kubectl, etc.). It makes it easy to build professional CLIs with commands, subcommands, flags, and auto-generated help and shell-completion.
    *   **[urfave/cli](https://github.com/urfave/cli):** Another excellent and widely-used library.

### 1.3. Error Handling Strategy
The current error handling mixes concerns. Functions often print an error to the console *and* return it.

*   **Why it's a problem:** A library function (`logic.Unmarshing`) should not decide how an error is presented to the user. It should return the error and let the caller (the `main` function or command handler) decide whether to log it, print a user-friendly message, or exit.
*   **How to improve:**
    *   **Return, Don't Print:** In lower-level packages (`logic`, `configs`), remove `fmt.Printf` calls for errors. Just return the error.
    *   **Wrap Errors for Context:** When you get an error from another function, wrap it to provide context. This makes debugging much easier.
        ```go
        // Before
        if err != nil {
            return config, err 
        }

        // After (Professional Practice)
        if err != nil {
            return config, fmt.Errorf("could not read config file at %s: %w", path, err)
        }
        ```
    *   **Centralize User-Facing Output:** The `main` package should be the only place that prints messages and errors to the user.

### 1.4. Project Structure
While the package separation is a good start, professional Go projects often follow a more standardized layout.

*   **[Standard Go Project Layout](https://github.com/golang-standards/project-layout):**
    *   `internal/`: For all your application-specific code. Code in `internal` can't be imported by other projects, which clearly defines your public API (or lack thereof for a binary). Your `commands`, `logic`, `configs`, and `templates` packages would live here.
    *   `cmd/`: The entry point for your application. You would have `cmd/todo/main.go`. This makes it clear where compilation starts and allows you to have multiple binaries in one project.

---

## 2. Specific Code Improvements

### 2.1. Critical Bug in Task Identification
There is a fundamental bug in how tasks are identified and manipulated.

*   **The Bug:** The `commands.Done` and `commands.RemoveTask` functions accept a `taskName` string. However, they pass this string to `logic.FindTask`, which immediately tries to convert it to an integer (`strconv.Atoi`). This will fail unless the task's name is a number. Furthermore, the `ID` field in the `Task` struct is never assigned a value when a new task is created in `commands.Add`.
*   **Solution:**
    1.  **Generate IDs:** When creating a new task, assign it a unique ID. This could be a simple incrementing integer or a more robust UUID. You'll need to read the existing tasks to find the highest current ID.
    2.  **Use IDs for Manipulation:** Change the `done` and `rm` commands to operate on the task ID, not the name. Names can be duplicated, but IDs must be unique.
    3.  **Update `FindTask`:** It should correctly accept an `int` ID.

### 2.2. Use Enums for States
The task `State` uses "magic numbers" (`0` for pending, `1` for done). This makes the code harder to read.

*   **How to improve:** Use `iota` to create a typed enum.
    ```go
    // In templates/templates.go
    type TaskState int

    const (
        StatePending TaskState = iota // 0
        StateDone                     // 1
    )

    type Task struct {
        // ...
        State TaskState
    }
    ```

### 2.3. Logging vs. Standard Output
The code mixes debug logging (`fmt.Printf("debug...")`), user-facing messages, and error messages.

*   **How to improve:**
    *   Use Go's new structured logger, **`slog`** (available since Go 1.21). It allows you to log at different levels (Debug, Info, Warn, Error) and can output structured JSON, which is invaluable in production environments.
    *   Reserve `fmt.Println` and `fmt.Printf` only for printing direct output to the user (e.g., the task list from the `ls` command).

### 2.4. Refactor `logic` and `configs` Packages
The `Unmarshing` and `Marshing` functions are very similar in both packages. You could create a generic data layer.

*   **Suggestion:** Create an `internal/store` package that handles reading and writing JSON data. It could have functions like:
    ```go
    func ReadJSON(filePath string, v interface{}) error
    func WriteJSON(filePath string, v interface{}) error
    ```
    Your `configs` and `commands` packages could then use this store, reducing code duplication.

### 2.5. Inconsistent Command Definitions
The command handling in `main.go` has some inconsistencies. For example, `rm` can take an optional list argument, but `add` cannot. A CLI library like Cobra (mentioned above) would help enforce consistency.

---

## Summary of Recommendations

1.  **Write Tests:** Start now. This is the most important step you can take.
2.  **Adopt a CLI Library:** Refactor `main.go` to use `cobra` or `urfave/cli`.
3.  **Fix the Task ID Bug:** Implement unique ID generation and use IDs for all task operations.
4.  **Refine Error Handling:** Centralize user-facing error messages in `main` and use error wrapping.
5.  **Use a Structured Logger:** Replace `fmt.Printf` for debugging/logging with `slog`.
6.  **Adopt a Standard Project Layout:** Consider moving your code into `internal/` and `cmd/`.
7.  **Use Enums:** Replace magic numbers for `State` with a typed constant.
