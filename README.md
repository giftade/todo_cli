# Todo CLI

A simple command-line interface (CLI) application for managing your todo list.

## Features

- Add new tasks
- List all tasks
- Mark tasks as completed
- Delete tasks

## Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/giftade/todo_cli
  ```
2. Navigate to the project directory:
  ```sh
  cd todo_cli
  ```
3. Build the application:
  ```sh
  go build
  ```

## Usage

- Add a new task:
  ```sh
  ./todo_cli add "Your new task"
  ```
- List all tasks:
  ```sh
  ./todo_cli list
  ```
- Mark a task as completed:
  ```sh
  ./todo_cli complete <task_id>
  ```
- Delete a task:
  ```sh
  ./todo_cli delete <task_id>
  ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.