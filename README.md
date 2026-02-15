
# claude-test

This sample application was built with Claude Code and runs on a locally hosted Ollama instance.  

The project was generated using [Opencode](https://github.com/anomalyco/opencode) and leverages a [local Ollama server](https://github.com/mmontes11/k8s-ai/tree/main/infrastructure/ollama) deployed in my [homelab](https://github.com/mmontes11/k8s-infrastructure).  

### Runtime Environment

NVIDIA RTX PRO Blackwell 4000 SFF 24GB VRAM
gpt-oss:20b and 128K context window 

```
NAME                   ID              SIZE     PROCESSOR    CONTEXT    UNTIL
gpt-oss:20b-ctx128k    4d2f07cece89    21 GB    100% GPU     131072     8 minutes from now
```

## Items Operations

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/items` | POST | Create a new item. Expects JSON body: `{"name": "..", "description": ".."}`. Returns the created item with its ID and creation timestamp.
| `/items` | GET | Retrieve a list of all items.
| `/items/{id}` | GET | Retrieve a single item by ID.
| `/items/{id}` | PUT | Update an existing item. Expects JSON body same as POST.
| `/items/{id}` | DELETE | Delete an item by ID.

## Collections Operations

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/collections` | POST | Create a new collection. Body: `{"name": "..", "description": ".."}`.
| `/collections` | GET | List all collections.
| `/collections/{id}` | GET | Get a collection by ID.
| `/collections/{id}` | PUT | Update collection name/description.
| `/collections/{id}` | DELETE | Delete a collection.

### Items in a Collection

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/collections/{id}/items` | POST | Add an existing item to a collection. Body: `{"item_id": 42}`.
| `/collections/{id}/items/{item_id}` | DELETE | Remove an item from the collection.
```
