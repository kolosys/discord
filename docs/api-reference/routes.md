# routes API

Complete API documentation for the routes package.

**Import Path:** `github.com/kolosys/discord/examples/basic/routes`

## Package Documentation



## Types

### GatewayStatus
_No documentation available_

#### Example Usage

```go
// Create a new GatewayStatus
gatewaystatus := GatewayStatus{
    Connected: true,
    Shards: 42,
}
```

#### Type Definition

```go
type GatewayStatus struct {
    Connected bool `json:"connected"`
    Shards int `json:"shards"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Connected | `bool` |  |
| Shards | `int` |  |

### HealthResponse
_No documentation available_

#### Example Usage

```go
// Create a new HealthResponse
healthresponse := HealthResponse{
    Status: "example",
}
```

#### Type Definition

```go
type HealthResponse struct {
    Status string `json:"status"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Status | `string` |  |

### StatusResponse
_No documentation available_

#### Example Usage

```go
// Create a new StatusResponse
statusresponse := StatusResponse{
    Gateway: GatewayStatus{},
}
```

#### Type Definition

```go
type StatusResponse struct {
    Gateway GatewayStatus `json:"gateway"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Gateway | `GatewayStatus` |  |

## Functions

### Register
_No documentation available_

```go
func Register(bot *discord.Bot)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `bot` | `*discord.Bot` | |

**Returns:**
None

**Example:**

```go
// Example usage of Register
result := Register(/* parameters */)
```

## External Links

- [Package Overview](../packages/routes.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/examples/basic/routes)
- [Source Code](https://github.com/kolosys/discord/tree/main/examples/basic/routes)
