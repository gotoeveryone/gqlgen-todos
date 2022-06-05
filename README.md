# Golang with GraphQL

## Requirements

- Docker

## Run

```console
$ docker compose up
```

## Generate

```console
$ docker compose exec server go generate ./...
```

## Example

### Find todos

```graphql
query findTodos {
  todos {
    id
    text
    done
    user {
      id
    }
  }
}
```

### Create todo

```graphql
mutation createTodo {
  createTodo(input: { text: "todo1234", userId: "54321" }) {
    id
    user {
      id
    }
    text
    done
  }
}
```

### Update todo

```graphql
mutation updateTodo {
  updateTodo(
    input: {
      id: "af1310c4-0a86-47e6-8429-31bfeb7bda86"
      text: "todo123456"
      userId: "98765"
      done: false
    }
  ) {
    id
    text
    done
  }
}
```

### Delete todo

```graphql
mutation deleteTodo {
  deleteTodo(input: { id: "af1310c4-0a86-47e6-8429-31bfeb7bda86" })
}
```
