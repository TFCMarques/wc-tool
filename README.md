## Word Count Tool:

This is a simple challenge designed to learn the Go programming language. It provides a practical way to get hands-on experience with Go by building a tool that counts lines, words, and characters in a text file.

### How to Run

To run this tool, you need to have Go installed on your machine. Once Go is installed, you can use the following commands from the terminal:

- To count lines in a file:
  ```sh
  go run ccwc.go -l <filename>
  ```
- To count words in a file:
  ```sh
  go run ccwc.go -w <filename>
  ```
- To count characters in a file:
  ```sh
  go run ccwc.go -c <filename>
  ```
- To count characters (considering Unicode characters as individual characters):
  ```sh
  go run ccwc.go -m <filename>
  ```
- If no option is specified, the tool will count lines, words, and characters:
  ```sh
  go run ccwc.go <filename>
  ```
- To read from stdin instead of a file, omit the filename:
  ```sh
  cat test.txt | go run ccwc.go <option>
  ```

For more information on the challenge, visit [codingchallenges.fyi/challenges/challenge-wc](https://codingchallenges.fyi/challenges/challenge-wc)
