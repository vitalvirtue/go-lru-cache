
# LRU Cache with Golang

This project is an implementation of an **LRU (Least Recently Used)** cache in Go. The LRU cache is a caching strategy that removes the least recently used item from memory to allow new items to be added. 

## 📜 LRU Cache Logic

An **LRU Cache** keeps the most recently used elements and evicts the least used elements from memory when it's full. This strategy is particularly useful when working with limited memory or when there's a need to quickly access frequently used data.

### How Does LRU Cache Work?

1. **Accessing or Adding Data (Set/Get)**:
   - If the requested key exists in the cache, the value is returned and the element is marked as "most recently used."
   - When a new key is added, it is added to the beginning of the list, and if the cache is full, the oldest element is removed.
2. **Using Doubly Linked List and HashMap**:
   - A hashmap (for fast access) and a doubly linked list (to manage the order of nodes) are used to provide **O(1)** complexity for adding and removing data.
3. **Capacity Control**:
   - When the capacity is reached, the oldest element (at the end of the list) is removed. This keeps the most recent and frequently used items prioritized.

## 🚀 Implementing LRU Cache with Go

In Golang, we implement the LRU cache using a `map` to quickly access nodes by key and a `doubly linked list` to manage the order of nodes. This structure allows us to achieve **O(1)** time complexity for inserting and removing data.

- **Node Structure**: Each node contains a key-value pair and pointers to the previous and next nodes to manage the order.
- **LRU Cache Structure**: The cache capacity, hashmap, and head-tail pointers are the main components.

## 🗂️ Project Structure

```bash
lru-cache-project/
├── cmd/
│   └── main.go                  # Entry point of the application
├── pkg/
│   └── cache/
│       ├── lru.go               # Main LRU cache implementation
│       ├── node.go              # Linked list nodes definition
│       └── cache.go             # General cache interface and settings
├── internal/
│   ├── config/
│   │   └── config.go            # Contains cache configuration
│   └── utils/
│       └── logger.go            # Basic logging functions
└── tests/
    ├── cache_test.go            # Unit tests for cache
    └── lru_test.go              # Tests for LRU cache
```

## ⚙️ Installation and Run

1. **Clone the Project**:
   ```bash
   git clone https://github.com/username/lru-cache-project.git
   cd lru-cache-project
   ```

2. **Install Dependencies**:
   The project uses Go modules. Load dependencies with `go mod` commands:
   ```bash
   go mod tidy
   ```

3. **Run the Project**:
   To run the application, execute:
   ```bash
   go run cmd/main.go
   ```

### Running with Docker

If you want to run the project in Docker, you can build an image and start it as a container.

1. **Build Docker Image**:
   ```bash
   docker build -t lru-cache-project .
   ```

2. **Start Docker Container**:
   ```bash
   docker run --rm lru-cache-project
   ```

## 🧪 Testing

Unit tests are located in the `tests/` directory. To run tests, use:

```bash
go test ./tests/...
```

These tests ensure the correct behavior of the LRU cache and verify that the oldest elements are removed when the capacity is exceeded.

## 📄 Example Usage

Here is a sample usage in `main.go`:

```go
package main

import (
    "fmt"
    "lru-cache-project/internal/config"
    "lru-cache-project/internal/utils"
    "lru-cache-project/pkg/cache"
)

func main() {
    cfg := config.NewConfig()
    lruCache := cache.NewLRUCache(cfg.CacheCapacity)
    lruCache.Set("A", "1")
    utils.Log("Added key A with value 1")
    lruCache.Set("B", "2")
    utils.Log("Added key B with value 2")
    
    if value, found := lruCache.Get("A"); found {
        fmt.Println("Cache 'A':", value)
    }
    
    lruCache.Set("C", "3") // B should be evicted
    utils.Log("Added key C with value 3")
    
    if _, found := lruCache.Get("B"); !found {
        fmt.Println("Cache 'B' has been evicted.")
    }
}
```

## 📝 Development Notes

- **LRU Algorithm**: Implemented using a combination of doubly linked list and hashmap.
- **Capacity Control**: Configurable through the `config/config.go` file.
- **Docker Support**: Easily deployable as a container with Dockerfile.

## 📌 Contributing

If you would like to contribute, feel free to submit a **pull request** or open an **issue**.

## 🏷️ License

This project is licensed under the MIT License.
