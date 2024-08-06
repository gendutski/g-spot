# Dotenv Generator

This helper will recreate `.env` file by prompting question in terminal

### Tags
| Tag              | Description                      |
|------------------|----------------------------------|
| envconfig        | Create .env key field            |
| prompt           | Prompt question on terminal and save answer as value  |
| default          | Use this value if prompt not answered |
| secret           | If `true`, the input will not be displayed on the terminal |
| comment          | Add comment before this field    |

### Example
Open [example](example/example.go) folder and run:

```bash
go run example.go
```