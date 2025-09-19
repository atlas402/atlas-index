# Atlas Index

> Service discovery and indexing SDK for x402 ecosystem

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![x402](https://img.shields.io/badge/x402-Compatible-green)](https://x402.org)

Atlas Index provides real-time indexing and discovery of x402-protected services across multiple blockchains.

## Installation

### TypeScript/JavaScript

```bash
npm install @atlas402/index
```

### Python

```bash
pip install atlas-index
```

### Go

```bash
go get github.com/atlas402/index
```

### Java

```xml
<dependency>
  <groupId>com.atlas402</groupId>
  <artifactId>index</artifactId>
  <version>1.0.0</version>
</dependency>
```

## Quick Start

### TypeScript

```typescript
import { AtlasIndex } from '@atlas402/index';

const index = new AtlasIndex({
  facilitatorUrl: 'https://facilitator.payai.network',
});

const services = await index.discover({
  category: 'AI',
  network: 'base',
});
```

## Documentation

- [TypeScript SDK](./typescript/README.md)
- [Python SDK](./python/README.md)
- [Go SDK](./go/README.md)
- [Java SDK](./java/README.md)
- [API Reference](./docs/api.md)

## License

Apache 2.0


