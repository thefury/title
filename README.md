# Title

`title` is a small utility that converts inpuan input stringt to title case on the command line. 

## Installation

To install the "title" program, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/thefury/title.git
    ```

2. Navigate to the project directory:
    ```sh
    cd title
    ```

3. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

To build the `title` program, run the following command:

```sh
go build -o title cmd/main.go 
cp title ~/bin  # or somewhere in your path
```

## Example

```sh
title a title has a certain format

# A Title Has a Certain Format
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.