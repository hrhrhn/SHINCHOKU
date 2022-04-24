```
mutation CreateTodo {
  createTodo(input: {
    userId: 5
    content: "implement delete function v2"
  }) {
    id
    content
  }
}

mutation {
  createUser(input: {
    name: "TEST_USER_V2"
  }) {
    id
    name
  }
}

query {
  users {
    id 
    name
    todos {
      id
      content
      done
      user {
        id
        name
      }
    }
  }
}

query {
  todos {
    id
    content
    user {
      name
    }
  }
}

query {
  user(id:5) {
    id
    name
    todos {
      id
      content
      done
    }
  }
}

query {
  todo(id:6) {
    content
    user {
      id
      name
    }
  }
}

mutation {
  deleteUser(input:4)
}

```
