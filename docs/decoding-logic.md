# seedling decoding logic

## AST structure

Example markdown

```md
- example/
  - cmd/
    - main.go
  - internal/
    - domain/
      - domain.go
  - docs/
    - document.md
  - README.md
```
Relevant dump nodes and attributes (`astdump testdata/test_tree.md`)

```py
Document {
  Children: [
    List {
      Children: [
        ListItem {
          Children: [
            Paragraph {
              Children: [
                Text {
                  Value: example/
                }
              ]
            }
            List {
              Children: [
                ListItem {
                  Children: [
                    Paragraph {
                      Children: [
                        Text {
                          Value: cmd/
                        }
                      ]
                    }
                    List {
                      Children: [
                        ListItem {
                          Children: [
                            Paragraph {
                              Children: [
                                Text {
                                  Value: main.go
                                }
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
                ListItem {
                  Children: [
                    Paragraph {
                      Children: [
                        Text {
                          Value: internal/
                        }
                      ]
                    }
                    List {
                      Children: [
                        ListItem {
                          Children: [
                            Paragraph {
                              Children: [
                                Text {
                                  Value: domain/
                                }
                              ]
                            }
                            List {
                              Children: [
                                ListItem {
                                  Children: [
                                    Paragraph {
                                      Children: [
                                        Text {
                                          Value: domain.go
                                        }
                                      ]
                                    }
                                  ]
                                }
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
                ListItem {
                  Children: [
                    Paragraph {
                      Children: [
                        Text {
                          Value: docs/
                        }
                      ]
                    }
                    List {
                      Children: [
                        ListItem {
                          Children: [
                            Paragraph {
                              Children: [
                                Text {
                                  Value: document.md
                                }
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
                ListItem {
                  Children: [
                    Paragraph {
                      Children: [
                        Text {
                          Value: README.md
                        }
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```

## Traversal

`markdown.getList()` returns the top-level list, so traversal can start there.

When depth increases, the parser creates a nested `List`; its first child is the deeper `ListItem`

`ast.Walk()` visits each node twice, on **entry** and on **exit**. This is a **depth-first** traversal.
- `entering == true` before visiting children
- `entering == false` after visiting children

**BUT THIS FLATTENS THE STRUCTURE AND WOULD REQUIRE STACK SEMANTICS TO RETAIN HIERARCHY**
