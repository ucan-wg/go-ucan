window.BENCHMARK_DATA = {
  "lastUpdate": 1731497959834,
  "repoUrl": "https://github.com/ucan-wg/go-ucan",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "50ea43e3fa317d1e3b202c5e7ca176b35e19b924",
          "message": "Merge pull request #33 from ucan-wg/continuous-bench\n\ngha: add a continuous benchmark action",
          "timestamp": "2024-10-01T17:40:50+02:00",
          "tree_id": "1e61b975dacff32071db7329167226eda6aeaa00",
          "url": "https://github.com/ucan-wg/go-ucan/commit/50ea43e3fa317d1e3b202c5e7ca176b35e19b924"
        },
        "date": 1727797732051,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 42.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.37,
            "unit": "ns/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15176,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15176,
            "unit": "ns/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111582,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111582,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54064,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54064,
            "unit": "ns/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 99378,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 99378,
            "unit": "ns/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62659,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62659,
            "unit": "ns/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 111723,
            "unit": "ns/op\t   15309 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 111723,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15309,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55260,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55260,
            "unit": "ns/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100876,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100876,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58579,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58579,
            "unit": "ns/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "19824 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6b8fbcee0adc4a289da32668616a3f49bad526a1",
          "message": "Merge pull request #34 from ucan-wg/invocation-stub\n\ntoken: add invocation partial stub",
          "timestamp": "2024-10-02T11:02:32+02:00",
          "tree_id": "f11c633892fe52609f43e85a407013ded5a0ce7d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6b8fbcee0adc4a289da32668616a3f49bad526a1"
        },
        "date": 1727859793240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 41.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.8,
            "unit": "ns/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15132,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15132,
            "unit": "ns/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109556,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109556,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53087,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53087,
            "unit": "ns/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97627,
            "unit": "ns/op\t   14414 B/op\t     305 allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97627,
            "unit": "ns/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14414,
            "unit": "B/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58314,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58314,
            "unit": "ns/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108815,
            "unit": "ns/op\t   15310 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108815,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15310,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55681,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55681,
            "unit": "ns/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97985,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97985,
            "unit": "ns/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57119,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57119,
            "unit": "ns/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20989 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8615f6c72bc7d2d1fbcf9e084721332ce557d56e",
          "message": "Merge pull request #35 from ucan-wg/cid-proposal\n\ntoken: don't store the CID in the token, symmetric API for sealed",
          "timestamp": "2024-10-02T13:42:06+02:00",
          "tree_id": "f9db21048fde99c92acf6eafeda506da75819967",
          "url": "https://github.com/ucan-wg/go-ucan/commit/8615f6c72bc7d2d1fbcf9e084721332ce557d56e"
        },
        "date": 1727869361557,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 42.22,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.22,
            "unit": "ns/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15063,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15063,
            "unit": "ns/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108585,
            "unit": "ns/op\t   15342 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108585,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15342,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53039,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53039,
            "unit": "ns/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97280,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97280,
            "unit": "ns/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59753,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59753,
            "unit": "ns/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108276,
            "unit": "ns/op\t   15293 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108276,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15293,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54751,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54751,
            "unit": "ns/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97988,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97988,
            "unit": "ns/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58597,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58597,
            "unit": "ns/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20824 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d60fb71156bcd1c3347323d03627ba9e509ed0c7",
          "message": "Merge pull request #22 from ucan-wg/container\n\nadd a token container with serialization as CARv1 file",
          "timestamp": "2024-10-08T11:40:00+02:00",
          "tree_id": "dcaedada721ca4a53fc4ec289909644281438720",
          "url": "https://github.com/ucan-wg/go-ucan/commit/d60fb71156bcd1c3347323d03627ba9e509ed0c7"
        },
        "date": 1728380464442,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5952,
            "unit": "ns/op\t   17928 B/op\t      59 allocs/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5952,
            "unit": "ns/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17928,
            "unit": "B/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977339,
            "unit": "ns/op\t  149321 B/op\t    2631 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977339,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 149321,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2631,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 12011,
            "unit": "ns/op\t   24232 B/op\t      61 allocs/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 12011,
            "unit": "ns/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24232,
            "unit": "B/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1004959,
            "unit": "ns/op\t  151381 B/op\t    2633 allocs/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1004959,
            "unit": "ns/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 151381,
            "unit": "B/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2633,
            "unit": "allocs/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4847,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4847,
            "unit": "ns/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 947913,
            "unit": "ns/op\t  138775 B/op\t    2534 allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 947913,
            "unit": "ns/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138775,
            "unit": "B/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2534,
            "unit": "allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8938,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8938,
            "unit": "ns/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 966182,
            "unit": "ns/op\t  140846 B/op\t    2536 allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 966182,
            "unit": "ns/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140846,
            "unit": "B/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2536,
            "unit": "allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 42.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.09,
            "unit": "ns/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15020,
            "unit": "ns/op\t   14568 B/op\t     162 allocs/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15020,
            "unit": "ns/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14568,
            "unit": "B/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111332,
            "unit": "ns/op\t   15342 B/op\t     374 allocs/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111332,
            "unit": "ns/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15342,
            "unit": "B/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53708,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53708,
            "unit": "ns/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97648,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97648,
            "unit": "ns/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59582,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59582,
            "unit": "ns/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109148,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109148,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54425,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54425,
            "unit": "ns/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98852,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98852,
            "unit": "ns/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59612,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59612,
            "unit": "ns/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20325 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2a51d61b46fd3195fd3d8dc48d829f1bb0a09df9",
          "message": "Merge pull request #39 from ucan-wg/command-as-string\n\ncommand: make the type a string, for easier equality test",
          "timestamp": "2024-10-09T18:30:45+02:00",
          "tree_id": "073d1b9347bdb661e7a47a926e727f167f7af4ba",
          "url": "https://github.com/ucan-wg/go-ucan/commit/2a51d61b46fd3195fd3d8dc48d829f1bb0a09df9"
        },
        "date": 1728491495173,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5803,
            "unit": "ns/op\t   17928 B/op\t      59 allocs/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5803,
            "unit": "ns/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17928,
            "unit": "B/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 989770,
            "unit": "ns/op\t  148843 B/op\t    2621 allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 989770,
            "unit": "ns/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148843,
            "unit": "B/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11415,
            "unit": "ns/op\t   24232 B/op\t      61 allocs/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11415,
            "unit": "ns/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24232,
            "unit": "B/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1001742,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1001742,
            "unit": "ns/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4421,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4421,
            "unit": "ns/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 970388,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 970388,
            "unit": "ns/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8848,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8848,
            "unit": "ns/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 985494,
            "unit": "ns/op\t  140367 B/op\t    2526 allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 985494,
            "unit": "ns/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140367,
            "unit": "B/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 42.15,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.15,
            "unit": "ns/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15715,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15715,
            "unit": "ns/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111724,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111724,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54235,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54235,
            "unit": "ns/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 100569,
            "unit": "ns/op\t   14350 B/op\t     304 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 100569,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14350,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 60128,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 60128,
            "unit": "ns/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 111112,
            "unit": "ns/op\t   15246 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 111112,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15246,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 58007,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 58007,
            "unit": "ns/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 104613,
            "unit": "ns/op\t   14350 B/op\t     304 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 104613,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14350,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59921,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59921,
            "unit": "ns/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20181 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "100a510097e600ee542e1ca2c3162ef136cb9bdc",
          "message": "pkg/container: harden the CAR file round-trip with fuzzing",
          "timestamp": "2024-10-09T18:38:35+02:00",
          "tree_id": "de75bb5656a96302c7670c3f0b3661dbc1cf4c9f",
          "url": "https://github.com/ucan-wg/go-ucan/commit/100a510097e600ee542e1ca2c3162ef136cb9bdc"
        },
        "date": 1728491991117,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5923,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5923,
            "unit": "ns/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 949584,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 949584,
            "unit": "ns/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10962,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10962,
            "unit": "ns/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 960665,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 960665,
            "unit": "ns/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4376,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4376,
            "unit": "ns/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 932288,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 932288,
            "unit": "ns/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8582,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8582,
            "unit": "ns/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 947744,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 947744,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.83,
            "unit": "ns/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14850,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14850,
            "unit": "ns/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109006,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109006,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52049,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52049,
            "unit": "ns/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96023,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96023,
            "unit": "ns/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57831,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57831,
            "unit": "ns/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107731,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107731,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15245,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56559,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56559,
            "unit": "ns/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96374,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96374,
            "unit": "ns/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56745,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56745,
            "unit": "ns/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21183 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9051e5250bd50bcceb7dc13a898e2b518018dc84",
          "message": "Merge pull request #40 from ucan-wg/dlg-example\n\ndelegation: make the examples more examply, less testy",
          "timestamp": "2024-10-14T12:52:52+02:00",
          "tree_id": "8220a4f978b9aa6ca748991ec77c2f6092ff47e0",
          "url": "https://github.com/ucan-wg/go-ucan/commit/9051e5250bd50bcceb7dc13a898e2b518018dc84"
        },
        "date": 1728903219800,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5582,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5582,
            "unit": "ns/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 952866,
            "unit": "ns/op\t  148843 B/op\t    2621 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 952866,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148843,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11007,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11007,
            "unit": "ns/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993321,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993321,
            "unit": "ns/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4383,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4383,
            "unit": "ns/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 937988,
            "unit": "ns/op\t  138296 B/op\t    2524 allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 937988,
            "unit": "ns/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138296,
            "unit": "B/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8621,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8621,
            "unit": "ns/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 949706,
            "unit": "ns/op\t  140363 B/op\t    2526 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 949706,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140363,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.28,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.28,
            "unit": "ns/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14777,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14777,
            "unit": "ns/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108411,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108411,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54108,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54108,
            "unit": "ns/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96093,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96093,
            "unit": "ns/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57455,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57455,
            "unit": "ns/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107977,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107977,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15245,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53910,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53910,
            "unit": "ns/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97186,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97186,
            "unit": "ns/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59129,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59129,
            "unit": "ns/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21042 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "59da2d1a2c777ef19179a762331c9da5c8e1cb6b",
          "message": "Merge pull request #43 from ucan-wg/dlg-options\n\ndelegation: tune Nbf & Exp options",
          "timestamp": "2024-10-15T10:38:50+02:00",
          "tree_id": "ca4073a652fd30947906ea23822f7b6faee97f7b",
          "url": "https://github.com/ucan-wg/go-ucan/commit/59da2d1a2c777ef19179a762331c9da5c8e1cb6b"
        },
        "date": 1728981584660,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5613,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5613,
            "unit": "ns/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 947017,
            "unit": "ns/op\t  148837 B/op\t    2621 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 947017,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148837,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11078,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11078,
            "unit": "ns/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 958854,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 958854,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4440,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4440,
            "unit": "ns/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 930774,
            "unit": "ns/op\t  138297 B/op\t    2524 allocs/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 930774,
            "unit": "ns/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138297,
            "unit": "B/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9257,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9257,
            "unit": "ns/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 947392,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 947392,
            "unit": "ns/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.27,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.27,
            "unit": "ns/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14789,
            "unit": "ns/op\t   14568 B/op\t     162 allocs/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14789,
            "unit": "ns/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14568,
            "unit": "B/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108417,
            "unit": "ns/op\t   15293 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108417,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15293,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52198,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52198,
            "unit": "ns/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95978,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95978,
            "unit": "ns/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62242,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62242,
            "unit": "ns/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108363,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108363,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15245,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53682,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53682,
            "unit": "ns/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96766,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96766,
            "unit": "ns/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57477,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57477,
            "unit": "ns/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21172 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "aa4ad2fc10b7d8ee85d3ad2b02c8c222d363fa9f",
          "message": "Merge pull request #42 from ucan-wg/fluent-policy\n\npolicy: fluent construction",
          "timestamp": "2024-10-15T13:06:56+02:00",
          "tree_id": "211b9a7091724ac5f9b821b963feca59e552adaf",
          "url": "https://github.com/ucan-wg/go-ucan/commit/aa4ad2fc10b7d8ee85d3ad2b02c8c222d363fa9f"
        },
        "date": 1728990465642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5888,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5888,
            "unit": "ns/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 951101,
            "unit": "ns/op\t  148838 B/op\t    2621 allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 951101,
            "unit": "ns/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148838,
            "unit": "B/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10978,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10978,
            "unit": "ns/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 956728,
            "unit": "ns/op\t  150901 B/op\t    2623 allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 956728,
            "unit": "ns/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150901,
            "unit": "B/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4366,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4366,
            "unit": "ns/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 936983,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 936983,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8634,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8634,
            "unit": "ns/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 946874,
            "unit": "ns/op\t  140366 B/op\t    2526 allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 946874,
            "unit": "ns/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140366,
            "unit": "B/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.91,
            "unit": "ns/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14741,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14741,
            "unit": "ns/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107071,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107071,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 51822,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 51822,
            "unit": "ns/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95255,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95255,
            "unit": "ns/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57052,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57052,
            "unit": "ns/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 106667,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 106667,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53429,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53429,
            "unit": "ns/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100585,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100585,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56287,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56287,
            "unit": "ns/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21322 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "030db7ec0d512e30e4383827883c79b43f2f7695",
          "message": "invocation: fix comment",
          "timestamp": "2024-10-15T15:41:14+02:00",
          "tree_id": "149146e49e7e235dfd93b1116108b9a36b86e534",
          "url": "https://github.com/ucan-wg/go-ucan/commit/030db7ec0d512e30e4383827883c79b43f2f7695"
        },
        "date": 1728999736985,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5679,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5679,
            "unit": "ns/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 971518,
            "unit": "ns/op\t  148839 B/op\t    2621 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 971518,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148839,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11107,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11107,
            "unit": "ns/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 989070,
            "unit": "ns/op\t  150907 B/op\t    2623 allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 989070,
            "unit": "ns/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150907,
            "unit": "B/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4413,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4413,
            "unit": "ns/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 944400,
            "unit": "ns/op\t  138296 B/op\t    2524 allocs/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 944400,
            "unit": "ns/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138296,
            "unit": "B/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9234,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9234,
            "unit": "ns/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 957328,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 957328,
            "unit": "ns/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.55,
            "unit": "ns/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14756,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14756,
            "unit": "ns/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109203,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109203,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54092,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54092,
            "unit": "ns/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97363,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97363,
            "unit": "ns/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61286,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61286,
            "unit": "ns/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109122,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109122,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55359,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55359,
            "unit": "ns/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98148,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98148,
            "unit": "ns/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59429,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59429,
            "unit": "ns/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20269 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "c372f0793cb8dcf9ebc2b3fa3f2aa6dad5b27cb2",
          "message": "policy: match is now a method of Policy",
          "timestamp": "2024-10-15T16:37:24+02:00",
          "tree_id": "f72781e6d15a9a54d8ed1762acfc5db0cd60db85",
          "url": "https://github.com/ucan-wg/go-ucan/commit/c372f0793cb8dcf9ebc2b3fa3f2aa6dad5b27cb2"
        },
        "date": 1729003128006,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5688,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5688,
            "unit": "ns/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 955089,
            "unit": "ns/op\t  148839 B/op\t    2621 allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 955089,
            "unit": "ns/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148839,
            "unit": "B/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10993,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10993,
            "unit": "ns/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 982875,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 982875,
            "unit": "ns/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4498,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4498,
            "unit": "ns/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 937746,
            "unit": "ns/op\t  138297 B/op\t    2524 allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 937746,
            "unit": "ns/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138297,
            "unit": "B/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8675,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8675,
            "unit": "ns/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 949442,
            "unit": "ns/op\t  140361 B/op\t    2526 allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 949442,
            "unit": "ns/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140361,
            "unit": "B/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14722,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14722,
            "unit": "ns/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107593,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107593,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52069,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52069,
            "unit": "ns/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98689,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98689,
            "unit": "ns/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57357,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57357,
            "unit": "ns/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109376,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109376,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53617,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53617,
            "unit": "ns/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96071,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96071,
            "unit": "ns/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59489,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59489,
            "unit": "ns/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21157 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "2ad3aeb6da7f545626e353b25a9d79dea44a4b24",
          "message": "policy: match is now a method of Policy",
          "timestamp": "2024-10-15T16:53:06+02:00",
          "tree_id": "6c78a5b7b1a6756a695778e9d1cb6c967672c5cf",
          "url": "https://github.com/ucan-wg/go-ucan/commit/2ad3aeb6da7f545626e353b25a9d79dea44a4b24"
        },
        "date": 1729004044599,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5731,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5731,
            "unit": "ns/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 975900,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 975900,
            "unit": "ns/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11080,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11080,
            "unit": "ns/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 976309,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 976309,
            "unit": "ns/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4454,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4454,
            "unit": "ns/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 944615,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 944615,
            "unit": "ns/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8726,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8726,
            "unit": "ns/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 979660,
            "unit": "ns/op\t  140369 B/op\t    2526 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 979660,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140369,
            "unit": "B/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 42.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.07,
            "unit": "ns/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15088,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15088,
            "unit": "ns/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109992,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109992,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53913,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53913,
            "unit": "ns/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98474,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98474,
            "unit": "ns/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61232,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61232,
            "unit": "ns/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109719,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109719,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54266,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54266,
            "unit": "ns/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97451,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97451,
            "unit": "ns/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57107,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57107,
            "unit": "ns/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20964 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "081d382028a720206c4e3303e66c13bb52bf3121",
          "message": "selector: Select is now a method",
          "timestamp": "2024-10-15T17:26:49+02:00",
          "tree_id": "11f8c90bc4951df8944c09e35088ce910c088889",
          "url": "https://github.com/ucan-wg/go-ucan/commit/081d382028a720206c4e3303e66c13bb52bf3121"
        },
        "date": 1729006072713,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5705,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5705,
            "unit": "ns/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 970735,
            "unit": "ns/op\t  148844 B/op\t    2621 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 970735,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148844,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11792,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11792,
            "unit": "ns/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 991614,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 991614,
            "unit": "ns/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4490,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4490,
            "unit": "ns/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 956458,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 956458,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8816,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8816,
            "unit": "ns/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 971607,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 971607,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.97,
            "unit": "ns/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15179,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15179,
            "unit": "ns/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 114668,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 114668,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53816,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53816,
            "unit": "ns/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97870,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97870,
            "unit": "ns/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59186,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59186,
            "unit": "ns/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109594,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109594,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56489,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56489,
            "unit": "ns/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98847,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98847,
            "unit": "ns/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57884,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57884,
            "unit": "ns/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20911 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ac1b03f1444eaf3af93c8b86e1d95a9198d38c47",
          "message": "Merge pull request #44 from ucan-wg/policy-filtering\n\npolicy: add a way to filter policies with a path",
          "timestamp": "2024-10-16T14:29:57+02:00",
          "tree_id": "948cc83b040d756898427494ea3470c2b262e58f",
          "url": "https://github.com/ucan-wg/go-ucan/commit/ac1b03f1444eaf3af93c8b86e1d95a9198d38c47"
        },
        "date": 1729081846942,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5614,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5614,
            "unit": "ns/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 959235,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 959235,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11037,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11037,
            "unit": "ns/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 963119,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 963119,
            "unit": "ns/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4394,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4394,
            "unit": "ns/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 947706,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 947706,
            "unit": "ns/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9153,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9153,
            "unit": "ns/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 957519,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 957519,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.97,
            "unit": "ns/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14707,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14707,
            "unit": "ns/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107604,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107604,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53869,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53869,
            "unit": "ns/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95874,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95874,
            "unit": "ns/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59756,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59756,
            "unit": "ns/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107983,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107983,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56414,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56414,
            "unit": "ns/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96853,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96853,
            "unit": "ns/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57020,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57020,
            "unit": "ns/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21172 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "deaf9c4fe9e0734d66393949b592c8327abfcefc",
          "message": "Merge pull request #46 from ucan-wg/rework-policies\n\nselector: rework to match the spec, cleanup lots of edge cases",
          "timestamp": "2024-10-24T13:01:10+02:00",
          "tree_id": "b5b06019217dbadeea9e486b5e923e33a085fef7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/deaf9c4fe9e0734d66393949b592c8327abfcefc"
        },
        "date": 1729767727464,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5632,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5632,
            "unit": "ns/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 952301,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 952301,
            "unit": "ns/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10997,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10997,
            "unit": "ns/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 985050,
            "unit": "ns/op\t  150904 B/op\t    2623 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 985050,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150904,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4707,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4707,
            "unit": "ns/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 933504,
            "unit": "ns/op\t  138298 B/op\t    2524 allocs/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 933504,
            "unit": "ns/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138298,
            "unit": "B/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8678,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8678,
            "unit": "ns/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 946096,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 946096,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 51.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 51.49,
            "unit": "ns/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14663,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14663,
            "unit": "ns/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107494,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107494,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52664,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52664,
            "unit": "ns/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95247,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95247,
            "unit": "ns/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 56965,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 56965,
            "unit": "ns/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 106933,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 106933,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53582,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53582,
            "unit": "ns/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96193,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96193,
            "unit": "ns/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56403,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56403,
            "unit": "ns/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21259 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e76354fb0af6f16dc52eb59bba633553a5708600",
          "message": "Merge pull request #47 from ucan-wg/simplify-literal\n\nliteral: simplify package with built-in functions",
          "timestamp": "2024-10-24T14:33:33+02:00",
          "tree_id": "90c5fc08c871717a1ecf701938729c7c6b7f99cc",
          "url": "https://github.com/ucan-wg/go-ucan/commit/e76354fb0af6f16dc52eb59bba633553a5708600"
        },
        "date": 1729773257322,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5518,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5518,
            "unit": "ns/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 957058,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 957058,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10900,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10900,
            "unit": "ns/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 968791,
            "unit": "ns/op\t  150905 B/op\t    2623 allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 968791,
            "unit": "ns/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150905,
            "unit": "B/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4372,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4372,
            "unit": "ns/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 942197,
            "unit": "ns/op\t  138299 B/op\t    2524 allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 942197,
            "unit": "ns/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138299,
            "unit": "B/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8578,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8578,
            "unit": "ns/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 986028,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 986028,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.61,
            "unit": "ns/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14815,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14815,
            "unit": "ns/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 106841,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 106841,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 51852,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 51852,
            "unit": "ns/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95579,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95579,
            "unit": "ns/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57033,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57033,
            "unit": "ns/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 113419,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 113419,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53821,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53821,
            "unit": "ns/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96271,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96271,
            "unit": "ns/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56893,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56893,
            "unit": "ns/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21272 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "3b2ea8d435c96b5ff658632eda66ea891e45bf79",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:39:39+02:00",
          "tree_id": "ded3bfe341aa722d216b350809779eb703a5b24b",
          "url": "https://github.com/ucan-wg/go-ucan/commit/3b2ea8d435c96b5ff658632eda66ea891e45bf79"
        },
        "date": 1729773628258,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5651,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5651,
            "unit": "ns/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 969208,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 969208,
            "unit": "ns/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10980,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10980,
            "unit": "ns/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 998789,
            "unit": "ns/op\t  150911 B/op\t    2623 allocs/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 998789,
            "unit": "ns/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150911,
            "unit": "B/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4467,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4467,
            "unit": "ns/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 952555,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 952555,
            "unit": "ns/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8826,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8826,
            "unit": "ns/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 971687,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 971687,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.73,
            "unit": "ns/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15048,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15048,
            "unit": "ns/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107942,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107942,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53135,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53135,
            "unit": "ns/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97377,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97377,
            "unit": "ns/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59199,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59199,
            "unit": "ns/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 114890,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 114890,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55697,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55697,
            "unit": "ns/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99755,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99755,
            "unit": "ns/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58381,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58381,
            "unit": "ns/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20480 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "02c6ca24c2e2cd33f49caa7222f44479add4a55e",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:40:43+02:00",
          "tree_id": "cdc0233761e38aac7d8745bbe321d8f322d4199d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/02c6ca24c2e2cd33f49caa7222f44479add4a55e"
        },
        "date": 1729773694781,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5679,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5679,
            "unit": "ns/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 1008470,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 1008470,
            "unit": "ns/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11236,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11236,
            "unit": "ns/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 989735,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 989735,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4628,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4628,
            "unit": "ns/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 962737,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 962737,
            "unit": "ns/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8763,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8763,
            "unit": "ns/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 972654,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 972654,
            "unit": "ns/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.57,
            "unit": "ns/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14838,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14838,
            "unit": "ns/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107788,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107788,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53024,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53024,
            "unit": "ns/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98729,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98729,
            "unit": "ns/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58767,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58767,
            "unit": "ns/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109070,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109070,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54997,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54997,
            "unit": "ns/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98109,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98109,
            "unit": "ns/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57987,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57987,
            "unit": "ns/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20814 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "6d0fbd4d5a7fc79f4fb6b0b1d269605d20b96d41",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:46:01+02:00",
          "tree_id": "e004fe55b7d123bc41ded35b971ebf2e589f6bf7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6d0fbd4d5a7fc79f4fb6b0b1d269605d20b96d41"
        },
        "date": 1729774008849,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5628,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5628,
            "unit": "ns/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977863,
            "unit": "ns/op\t  148836 B/op\t    2621 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977863,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148836,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11221,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11221,
            "unit": "ns/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993765,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993765,
            "unit": "ns/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4633,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4633,
            "unit": "ns/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 958853,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 958853,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8809,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8809,
            "unit": "ns/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 975863,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 975863,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.77,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.77,
            "unit": "ns/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14920,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14920,
            "unit": "ns/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108993,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108993,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52934,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52934,
            "unit": "ns/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97380,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97380,
            "unit": "ns/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58022,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58022,
            "unit": "ns/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112049,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112049,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55029,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55029,
            "unit": "ns/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99160,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99160,
            "unit": "ns/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57986,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57986,
            "unit": "ns/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20545 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6b72799818563bb1cb9b3bf33804acfd28ac2e43",
          "message": "Merge pull request #41 from ucan-wg/streamline-did\n\ndid: simplify public API, add missing required algorithms",
          "timestamp": "2024-10-24T17:09:10+02:00",
          "tree_id": "58e1c2c45cf6801639fc31f9b75d00669f7f034d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6b72799818563bb1cb9b3bf33804acfd28ac2e43"
        },
        "date": 1729782609110,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5549,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5549,
            "unit": "ns/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 956066,
            "unit": "ns/op\t  149240 B/op\t    2651 allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 956066,
            "unit": "ns/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 149240,
            "unit": "B/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11006,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11006,
            "unit": "ns/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1021168,
            "unit": "ns/op\t  151303 B/op\t    2653 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1021168,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 151303,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4405,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4405,
            "unit": "ns/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 953307,
            "unit": "ns/op\t  138699 B/op\t    2554 allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 953307,
            "unit": "ns/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138699,
            "unit": "B/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8565,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8565,
            "unit": "ns/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 950713,
            "unit": "ns/op\t  140765 B/op\t    2556 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 950713,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140765,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.42,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.42,
            "unit": "ns/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14540,
            "unit": "ns/op\t   14570 B/op\t     162 allocs/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14540,
            "unit": "ns/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14570,
            "unit": "B/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107116,
            "unit": "ns/op\t   15141 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107116,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15141,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53403,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53403,
            "unit": "ns/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95577,
            "unit": "ns/op\t   14197 B/op\t     303 allocs/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95577,
            "unit": "ns/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14197,
            "unit": "B/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57505,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57505,
            "unit": "ns/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108297,
            "unit": "ns/op\t   15094 B/op\t     371 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108297,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15094,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 371,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53944,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53944,
            "unit": "ns/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97081,
            "unit": "ns/op\t   14197 B/op\t     303 allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97081,
            "unit": "ns/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14197,
            "unit": "B/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56892,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56892,
            "unit": "ns/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20676 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "smoyer64@gmail.com",
            "name": "Steve Moyer",
            "username": "smoyer64"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "fcb527cc5264c3bca1d51f3459bceaa0ec5266f6",
          "message": "Merge pull request #50 from ucan-wg/fix/meta-optional-in-delegation\n\nfix: change meta optional in `delegation` `Token`, model and schema",
          "timestamp": "2024-10-24T18:10:30-04:00",
          "tree_id": "128e19731ee69bd3b3a8d0a9873e37817f94a956",
          "url": "https://github.com/ucan-wg/go-ucan/commit/fcb527cc5264c3bca1d51f3459bceaa0ec5266f6"
        },
        "date": 1729807875014,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5685,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5685,
            "unit": "ns/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 1000665,
            "unit": "ns/op\t  147963 B/op\t    2651 allocs/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 1000665,
            "unit": "ns/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147963,
            "unit": "B/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 12179,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 12179,
            "unit": "ns/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1028881,
            "unit": "ns/op\t  150024 B/op\t    2653 allocs/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1028881,
            "unit": "ns/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150024,
            "unit": "B/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4947,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4947,
            "unit": "ns/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 983043,
            "unit": "ns/op\t  137421 B/op\t    2554 allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 983043,
            "unit": "ns/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137421,
            "unit": "B/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8826,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8826,
            "unit": "ns/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 996012,
            "unit": "ns/op\t  139487 B/op\t    2556 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 996012,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139487,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.82,
            "unit": "ns/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15286,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15286,
            "unit": "ns/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 112039,
            "unit": "ns/op\t   15014 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 112039,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15014,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54962,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54962,
            "unit": "ns/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 102751,
            "unit": "ns/op\t   14069 B/op\t     303 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 102751,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14069,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61268,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61268,
            "unit": "ns/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112745,
            "unit": "ns/op\t   14966 B/op\t     371 allocs/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112745,
            "unit": "ns/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 14966,
            "unit": "B/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 371,
            "unit": "allocs/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56948,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56948,
            "unit": "ns/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 102304,
            "unit": "ns/op\t   14069 B/op\t     303 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 102304,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14069,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59726,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59726,
            "unit": "ns/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "19818 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "cfb4446a05c0dd10a83788ef7d2faf5bea71da4a",
          "message": "Merge pull request #48 from ucan-wg/pol-partial\n\npolicy: implement partial matching, to evaluate in multiple steps with fail early",
          "timestamp": "2024-11-05T16:31:52+01:00",
          "tree_id": "4367dcfc03af7bf87f10cb2227d107b43d0a28dc",
          "url": "https://github.com/ucan-wg/go-ucan/commit/cfb4446a05c0dd10a83788ef7d2faf5bea71da4a"
        },
        "date": 1730820773108,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5650,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "214743 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5650,
            "unit": "ns/op",
            "extra": "214743 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "214743 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "214743 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 968419,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1227 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 968419,
            "unit": "ns/op",
            "extra": "1227 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1227 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1227 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11252,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "105832 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11252,
            "unit": "ns/op",
            "extra": "105832 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "105832 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "105832 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1032936,
            "unit": "ns/op\t  150028 B/op\t    2653 allocs/op",
            "extra": "1201 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1032936,
            "unit": "ns/op",
            "extra": "1201 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150028,
            "unit": "B/op",
            "extra": "1201 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1201 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4431,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "269520 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4431,
            "unit": "ns/op",
            "extra": "269520 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "269520 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "269520 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 960632,
            "unit": "ns/op\t  137424 B/op\t    2554 allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 960632,
            "unit": "ns/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137424,
            "unit": "B/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8774,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137506 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8774,
            "unit": "ns/op",
            "extra": "137506 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137506 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137506 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 988058,
            "unit": "ns/op\t  139484 B/op\t    2556 allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 988058,
            "unit": "ns/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139484,
            "unit": "B/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28605874 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.58,
            "unit": "ns/op",
            "extra": "28605874 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28605874 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28605874 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14653,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81644 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14653,
            "unit": "ns/op",
            "extra": "81644 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81644 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81644 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108600,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108600,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54298,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22708 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54298,
            "unit": "ns/op",
            "extra": "22708 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22708 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22708 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96347,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12488 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96347,
            "unit": "ns/op",
            "extra": "12488 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12488 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12488 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57834,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20721 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57834,
            "unit": "ns/op",
            "extra": "20721 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20721 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20721 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107472,
            "unit": "ns/op\t   15125 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107472,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53955,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22002 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53955,
            "unit": "ns/op",
            "extra": "22002 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22002 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22002 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97022,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12360 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97022,
            "unit": "ns/op",
            "extra": "12360 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12360 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12360 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57103,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57103,
            "unit": "ns/op",
            "extra": "20818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20818 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "41b8600fbcac330aa59abc9ea8c1abef4c160db9",
          "message": "Merge pull request #52 from ucan-wg/meta-readonly\n\nmeta: make a read-only version to enforce token immutability",
          "timestamp": "2024-11-06T15:28:19+01:00",
          "tree_id": "48e1f80db3fa4ad482e99d75ba65b2d916772c36",
          "url": "https://github.com/ucan-wg/go-ucan/commit/41b8600fbcac330aa59abc9ea8c1abef4c160db9"
        },
        "date": 1730903349280,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5662,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "214466 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5662,
            "unit": "ns/op",
            "extra": "214466 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "214466 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "214466 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 987037,
            "unit": "ns/op\t  147963 B/op\t    2651 allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 987037,
            "unit": "ns/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147963,
            "unit": "B/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11107,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "105382 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11107,
            "unit": "ns/op",
            "extra": "105382 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "105382 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "105382 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 994813,
            "unit": "ns/op\t  150026 B/op\t    2653 allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 994813,
            "unit": "ns/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150026,
            "unit": "B/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4448,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "264915 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4448,
            "unit": "ns/op",
            "extra": "264915 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "264915 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "264915 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 953402,
            "unit": "ns/op\t  137418 B/op\t    2554 allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 953402,
            "unit": "ns/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137418,
            "unit": "B/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8694,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136900 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8694,
            "unit": "ns/op",
            "extra": "136900 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136900 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136900 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 968425,
            "unit": "ns/op\t  139484 B/op\t    2556 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 968425,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139484,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28481461 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.65,
            "unit": "ns/op",
            "extra": "28481461 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28481461 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28481461 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14656,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81628 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14656,
            "unit": "ns/op",
            "extra": "81628 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81628 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81628 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109244,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109244,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 55905,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22113 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 55905,
            "unit": "ns/op",
            "extra": "22113 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22113 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22113 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97418,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12267 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97418,
            "unit": "ns/op",
            "extra": "12267 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12267 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12267 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 60140,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20230 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 60140,
            "unit": "ns/op",
            "extra": "20230 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20230 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20230 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112048,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112048,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54778,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21735 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54778,
            "unit": "ns/op",
            "extra": "21735 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21735 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21735 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98093,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12194 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98093,
            "unit": "ns/op",
            "extra": "12194 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12194 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12194 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58654,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20766 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58654,
            "unit": "ns/op",
            "extra": "20766 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20766 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20766 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "c9f3a6033ad519502365d100e871e328c01eca95",
          "message": "delegation: minor fix around meta",
          "timestamp": "2024-11-06T16:43:57+01:00",
          "tree_id": "22652698736e51bfc713ecfd91b29e7641bd0505",
          "url": "https://github.com/ucan-wg/go-ucan/commit/c9f3a6033ad519502365d100e871e328c01eca95"
        },
        "date": 1730907899120,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5601,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "205020 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5601,
            "unit": "ns/op",
            "extra": "205020 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "205020 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "205020 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 948098,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 948098,
            "unit": "ns/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10954,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107612 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10954,
            "unit": "ns/op",
            "extra": "107612 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107612 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107612 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 960044,
            "unit": "ns/op\t  150025 B/op\t    2653 allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 960044,
            "unit": "ns/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150025,
            "unit": "B/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4381,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "274740 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4381,
            "unit": "ns/op",
            "extra": "274740 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "274740 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "274740 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 933247,
            "unit": "ns/op\t  137419 B/op\t    2554 allocs/op",
            "extra": "1249 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 933247,
            "unit": "ns/op",
            "extra": "1249 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137419,
            "unit": "B/op",
            "extra": "1249 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1249 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9119,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "138147 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9119,
            "unit": "ns/op",
            "extra": "138147 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "138147 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "138147 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 947780,
            "unit": "ns/op\t  139485 B/op\t    2556 allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 947780,
            "unit": "ns/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139485,
            "unit": "B/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.63,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28704544 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.63,
            "unit": "ns/op",
            "extra": "28704544 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28704544 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28704544 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14514,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81993 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14514,
            "unit": "ns/op",
            "extra": "81993 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81993 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81993 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107505,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107505,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52201,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22834 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52201,
            "unit": "ns/op",
            "extra": "22834 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22834 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22834 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95673,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95673,
            "unit": "ns/op",
            "extra": "12552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57577,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "18856 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57577,
            "unit": "ns/op",
            "extra": "18856 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "18856 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "18856 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107456,
            "unit": "ns/op\t   15125 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107456,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53818,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21687 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53818,
            "unit": "ns/op",
            "extra": "21687 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21687 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21687 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97006,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12469 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97006,
            "unit": "ns/op",
            "extra": "12469 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12469 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12469 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56801,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20688 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56801,
            "unit": "ns/op",
            "extra": "20688 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20688 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20688 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "884d63a68929039c48a4aad2d5c75fdfba900046",
          "message": "Merge pull request #53 from ucan-wg/test-literal\n\nliteral: add test suite",
          "timestamp": "2024-11-06T16:45:19+01:00",
          "tree_id": "5ddb8ef3ea5244fc4a5ce21b638440879c0a00dd",
          "url": "https://github.com/ucan-wg/go-ucan/commit/884d63a68929039c48a4aad2d5c75fdfba900046"
        },
        "date": 1730907964393,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5731,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "208888 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5731,
            "unit": "ns/op",
            "extra": "208888 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "208888 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "208888 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 958821,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 958821,
            "unit": "ns/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11189,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "105817 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11189,
            "unit": "ns/op",
            "extra": "105817 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "105817 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "105817 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 981056,
            "unit": "ns/op\t  150020 B/op\t    2653 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 981056,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150020,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4552,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "260373 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4552,
            "unit": "ns/op",
            "extra": "260373 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "260373 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "260373 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 981069,
            "unit": "ns/op\t  137421 B/op\t    2554 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 981069,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137421,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8837,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134336 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8837,
            "unit": "ns/op",
            "extra": "134336 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134336 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134336 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 968532,
            "unit": "ns/op\t  139480 B/op\t    2556 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 968532,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139480,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 43.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26899196 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 43.85,
            "unit": "ns/op",
            "extra": "26899196 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26899196 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26899196 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14759,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81253 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14759,
            "unit": "ns/op",
            "extra": "81253 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81253 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81253 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 110370,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 110370,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53906,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22212 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53906,
            "unit": "ns/op",
            "extra": "22212 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22212 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22212 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 101902,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12117 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 101902,
            "unit": "ns/op",
            "extra": "12117 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12117 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12117 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59015,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20174 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59015,
            "unit": "ns/op",
            "extra": "20174 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20174 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20174 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 110888,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 110888,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55374,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55374,
            "unit": "ns/op",
            "extra": "21632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 102131,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 102131,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58327,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58327,
            "unit": "ns/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20324 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "cddade46701497d5950be289436ddb3fda883f93",
          "message": "Merge pull request #55 from ucan-wg/flacky-test\n\nliteral: fix flacky test",
          "timestamp": "2024-11-07T12:07:43+01:00",
          "tree_id": "b480b811e2c6ac7ac4f55d2ccf9f6ebe187867c7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/cddade46701497d5950be289436ddb3fda883f93"
        },
        "date": 1730977713213,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5603,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "215985 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5603,
            "unit": "ns/op",
            "extra": "215985 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "215985 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "215985 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 953424,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 953424,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11595,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108970 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11595,
            "unit": "ns/op",
            "extra": "108970 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108970 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108970 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 967007,
            "unit": "ns/op\t  150023 B/op\t    2653 allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 967007,
            "unit": "ns/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150023,
            "unit": "B/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4417,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4417,
            "unit": "ns/op",
            "extra": "268404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 936008,
            "unit": "ns/op\t  137418 B/op\t    2554 allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 936008,
            "unit": "ns/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137418,
            "unit": "B/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8654,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "129285 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8654,
            "unit": "ns/op",
            "extra": "129285 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "129285 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "129285 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 953493,
            "unit": "ns/op\t  139482 B/op\t    2556 allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 953493,
            "unit": "ns/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139482,
            "unit": "B/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.23,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23164525 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.23,
            "unit": "ns/op",
            "extra": "23164525 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23164525 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23164525 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14665,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81859 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14665,
            "unit": "ns/op",
            "extra": "81859 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81859 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81859 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 110799,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 110799,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52286,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22720 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52286,
            "unit": "ns/op",
            "extra": "22720 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22720 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22720 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96122,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12468 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96122,
            "unit": "ns/op",
            "extra": "12468 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12468 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12468 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57620,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20295 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57620,
            "unit": "ns/op",
            "extra": "20295 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20295 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20295 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109331,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109331,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53756,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53756,
            "unit": "ns/op",
            "extra": "22266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97617,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97617,
            "unit": "ns/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56694,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "21070 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56694,
            "unit": "ns/op",
            "extra": "21070 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "21070 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21070 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1e5ecdc20543ac059cd479682f03849430778f3a",
          "message": "Merge pull request #56 from ucan-wg/match-return\n\npolicy: make Match also return the failing statement",
          "timestamp": "2024-11-07T15:40:57+01:00",
          "tree_id": "db918afb20318295e0f885cbdd138a5b2b9f8c31",
          "url": "https://github.com/ucan-wg/go-ucan/commit/1e5ecdc20543ac059cd479682f03849430778f3a"
        },
        "date": 1730990507143,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5715,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "205635 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5715,
            "unit": "ns/op",
            "extra": "205635 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "205635 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "205635 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 1007281,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 1007281,
            "unit": "ns/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11164,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107386 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11164,
            "unit": "ns/op",
            "extra": "107386 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107386 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107386 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 997128,
            "unit": "ns/op\t  150029 B/op\t    2653 allocs/op",
            "extra": "1186 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 997128,
            "unit": "ns/op",
            "extra": "1186 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150029,
            "unit": "B/op",
            "extra": "1186 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1186 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4447,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268014 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4447,
            "unit": "ns/op",
            "extra": "268014 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268014 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268014 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 961217,
            "unit": "ns/op\t  137422 B/op\t    2554 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 961217,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137422,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8776,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134124 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8776,
            "unit": "ns/op",
            "extra": "134124 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134124 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134124 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 980973,
            "unit": "ns/op\t  139485 B/op\t    2556 allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 980973,
            "unit": "ns/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139485,
            "unit": "B/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24296736 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.72,
            "unit": "ns/op",
            "extra": "24296736 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24296736 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24296736 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15037,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "79773 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15037,
            "unit": "ns/op",
            "extra": "79773 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "79773 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "79773 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109412,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109412,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53827,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22340 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53827,
            "unit": "ns/op",
            "extra": "22340 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22340 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22340 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97070,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97070,
            "unit": "ns/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58189,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20448 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58189,
            "unit": "ns/op",
            "extra": "20448 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20448 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20448 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108801,
            "unit": "ns/op\t   15125 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108801,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 57789,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21872 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 57789,
            "unit": "ns/op",
            "extra": "21872 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21872 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21872 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97888,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97888,
            "unit": "ns/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57871,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20774 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57871,
            "unit": "ns/op",
            "extra": "20774 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20774 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20774 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "smoyer64@gmail.com",
            "name": "Steve Moyer",
            "username": "smoyer64"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3c705ca1508ca8aa7f756193125ae7d2aa0632ee",
          "message": "Merge pull request #49 from ucan-wg/feat/complete-invocation-stub\n\nfeat: complete invocation stub",
          "timestamp": "2024-11-07T13:51:47-05:00",
          "tree_id": "15c22ef1c4882d6aecd41611e2215a0fef63bd8e",
          "url": "https://github.com/ucan-wg/go-ucan/commit/3c705ca1508ca8aa7f756193125ae7d2aa0632ee"
        },
        "date": 1731005574207,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5672,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "211268 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5672,
            "unit": "ns/op",
            "extra": "211268 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "211268 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "211268 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 976288,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 976288,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11152,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "104923 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11152,
            "unit": "ns/op",
            "extra": "104923 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "104923 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "104923 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993487,
            "unit": "ns/op\t  150025 B/op\t    2653 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993487,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150025,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4441,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "266629 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4441,
            "unit": "ns/op",
            "extra": "266629 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "266629 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "266629 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 1007685,
            "unit": "ns/op\t  137424 B/op\t    2554 allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 1007685,
            "unit": "ns/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137424,
            "unit": "B/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9277,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134631 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9277,
            "unit": "ns/op",
            "extra": "134631 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134631 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134631 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 982709,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 982709,
            "unit": "ns/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24458014 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.87,
            "unit": "ns/op",
            "extra": "24458014 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24458014 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24458014 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15275,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "79532 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15275,
            "unit": "ns/op",
            "extra": "79532 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "79532 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "79532 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 110257,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 110257,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53980,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22177 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53980,
            "unit": "ns/op",
            "extra": "22177 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22177 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22177 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 99180,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12243 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 99180,
            "unit": "ns/op",
            "extra": "12243 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12243 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12243 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58923,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "18776 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58923,
            "unit": "ns/op",
            "extra": "18776 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "18776 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "18776 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 110030,
            "unit": "ns/op\t   15125 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 110030,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55201,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21780 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55201,
            "unit": "ns/op",
            "extra": "21780 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21780 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21780 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98256,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98256,
            "unit": "ns/op",
            "extra": "12176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57603,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20917 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57603,
            "unit": "ns/op",
            "extra": "20917 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20917 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20917 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "633b3d210ae150419d050de5fb603702fd68c4b3",
          "message": "token: move nonce generation to a shared space",
          "timestamp": "2024-11-12T10:38:25+01:00",
          "tree_id": "7532984ac6f4adffabf883f51b7045e5d2aa71a7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/633b3d210ae150419d050de5fb603702fd68c4b3"
        },
        "date": 1731404384301,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5673,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "215245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5673,
            "unit": "ns/op",
            "extra": "215245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "215245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "215245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 955863,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 955863,
            "unit": "ns/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11083,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107532 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11083,
            "unit": "ns/op",
            "extra": "107532 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107532 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107532 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 977710,
            "unit": "ns/op\t  150022 B/op\t    2653 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 977710,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150022,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4437,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "263451 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4437,
            "unit": "ns/op",
            "extra": "263451 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "263451 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "263451 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 949512,
            "unit": "ns/op\t  137418 B/op\t    2554 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 949512,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137418,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9256,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135304 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9256,
            "unit": "ns/op",
            "extra": "135304 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135304 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135304 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 958121,
            "unit": "ns/op\t  139485 B/op\t    2556 allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 958121,
            "unit": "ns/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139485,
            "unit": "B/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.69,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24194378 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.69,
            "unit": "ns/op",
            "extra": "24194378 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24194378 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24194378 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14769,
            "unit": "ns/op\t   14576 B/op\t     163 allocs/op",
            "extra": "79600 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14769,
            "unit": "ns/op",
            "extra": "79600 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14576,
            "unit": "B/op",
            "extra": "79600 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "79600 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109218,
            "unit": "ns/op\t   15173 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109218,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15173,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52860,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22809 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52860,
            "unit": "ns/op",
            "extra": "22809 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22809 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22809 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96671,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12319 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96671,
            "unit": "ns/op",
            "extra": "12319 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12319 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12319 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57476,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "18631 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57476,
            "unit": "ns/op",
            "extra": "18631 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "18631 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "18631 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108603,
            "unit": "ns/op\t   15125 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108603,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56580,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56580,
            "unit": "ns/op",
            "extra": "22176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22176 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97648,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97648,
            "unit": "ns/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56964,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "21036 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56964,
            "unit": "ns/op",
            "extra": "21036 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "21036 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21036 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "17a1d54b6ff25f90ca76e53b3a461a64a99afc21",
          "message": "Merge pull request #57 from ucan-wg/simplify-args\n\nargs: simplify API + code",
          "timestamp": "2024-11-12T15:08:14+01:00",
          "tree_id": "8a474158342566a1312406a255d3c9314c60d3fb",
          "url": "https://github.com/ucan-wg/go-ucan/commit/17a1d54b6ff25f90ca76e53b3a461a64a99afc21"
        },
        "date": 1731420547144,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5553,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "212005 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5553,
            "unit": "ns/op",
            "extra": "212005 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "212005 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "212005 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977591,
            "unit": "ns/op\t  147961 B/op\t    2651 allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977591,
            "unit": "ns/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147961,
            "unit": "B/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11010,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108118 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11010,
            "unit": "ns/op",
            "extra": "108118 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108118 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108118 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 978536,
            "unit": "ns/op\t  150024 B/op\t    2653 allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 978536,
            "unit": "ns/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150024,
            "unit": "B/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4452,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261672 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4452,
            "unit": "ns/op",
            "extra": "261672 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261672 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261672 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 961046,
            "unit": "ns/op\t  137417 B/op\t    2554 allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 961046,
            "unit": "ns/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137417,
            "unit": "B/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8968,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "112093 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8968,
            "unit": "ns/op",
            "extra": "112093 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "112093 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "112093 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 960450,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 960450,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.13,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24719923 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.13,
            "unit": "ns/op",
            "extra": "24719923 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24719923 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24719923 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.23,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "51022893 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.23,
            "unit": "ns/op",
            "extra": "51022893 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "51022893 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "51022893 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 40.56,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28598168 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 40.56,
            "unit": "ns/op",
            "extra": "28598168 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28598168 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28598168 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 86.88,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "13992120 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 86.88,
            "unit": "ns/op",
            "extra": "13992120 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "13992120 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13992120 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1927,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "563296 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1927,
            "unit": "ns/op",
            "extra": "563296 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "563296 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "563296 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14952,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "77191 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14952,
            "unit": "ns/op",
            "extra": "77191 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "77191 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "77191 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 110800,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 110800,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52762,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22596 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52762,
            "unit": "ns/op",
            "extra": "22596 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22596 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22596 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 99423,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12316 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 99423,
            "unit": "ns/op",
            "extra": "12316 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12316 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12316 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58136,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "19132 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58136,
            "unit": "ns/op",
            "extra": "19132 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "19132 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "19132 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109420,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109420,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54542,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22102 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54542,
            "unit": "ns/op",
            "extra": "22102 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22102 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22102 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100633,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12133 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100633,
            "unit": "ns/op",
            "extra": "12133 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12133 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12133 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57239,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "21057 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57239,
            "unit": "ns/op",
            "extra": "21057 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "21057 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21057 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "fabio.bozzo@gmail.com",
            "name": "Fabio Bozzo",
            "username": "fabiobozzo"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6298fa28bd1e71659a3db2387d298839dd0b92be",
          "message": "Merge pull request #51 from ucan-wg/v1-meta-encryption\n\nfeat(meta): values symmetric encryption",
          "timestamp": "2024-11-12T16:48:04+01:00",
          "tree_id": "1eabcb7aa100556613d343f24df5cd3ce840d544",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6298fa28bd1e71659a3db2387d298839dd0b92be"
        },
        "date": 1731426541296,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5625,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "200802 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5625,
            "unit": "ns/op",
            "extra": "200802 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "200802 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "200802 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977635,
            "unit": "ns/op\t  147959 B/op\t    2651 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977635,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147959,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11651,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107623 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11651,
            "unit": "ns/op",
            "extra": "107623 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107623 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107623 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993785,
            "unit": "ns/op\t  150027 B/op\t    2653 allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993785,
            "unit": "ns/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150027,
            "unit": "B/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4481,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "257025 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4481,
            "unit": "ns/op",
            "extra": "257025 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "257025 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "257025 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 961937,
            "unit": "ns/op\t  137420 B/op\t    2554 allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 961937,
            "unit": "ns/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137420,
            "unit": "B/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8981,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134628 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8981,
            "unit": "ns/op",
            "extra": "134628 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134628 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134628 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 970724,
            "unit": "ns/op\t  139485 B/op\t    2556 allocs/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 970724,
            "unit": "ns/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139485,
            "unit": "B/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1214 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24340980 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.07,
            "unit": "ns/op",
            "extra": "24340980 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24340980 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24340980 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.28,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "51226170 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.28,
            "unit": "ns/op",
            "extra": "51226170 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "51226170 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "51226170 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 41.93,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "27988922 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 41.93,
            "unit": "ns/op",
            "extra": "27988922 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "27988922 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "27988922 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.52,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "13934396 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.52,
            "unit": "ns/op",
            "extra": "13934396 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "13934396 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13934396 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1939,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "561182 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1939,
            "unit": "ns/op",
            "extra": "561182 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "561182 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "561182 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14961,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "80592 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14961,
            "unit": "ns/op",
            "extra": "80592 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "80592 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "80592 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 110049,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 110049,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53654,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22364 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53654,
            "unit": "ns/op",
            "extra": "22364 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22364 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22364 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98569,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12236 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98569,
            "unit": "ns/op",
            "extra": "12236 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12236 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12236 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59993,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "19980 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59993,
            "unit": "ns/op",
            "extra": "19980 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "19980 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "19980 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 116631,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 116631,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56188,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21595 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56188,
            "unit": "ns/op",
            "extra": "21595 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21595 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21595 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99910,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12105 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99910,
            "unit": "ns/op",
            "extra": "12105 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12105 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12105 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59477,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20440 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59477,
            "unit": "ns/op",
            "extra": "20440 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20440 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20440 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "be185a8496f9d2e6ed1d983fbe20457092dc78e0",
          "message": "Merge pull request #58 from ucan-wg/readme\n\nadd OK readme",
          "timestamp": "2024-11-12T17:19:49+01:00",
          "tree_id": "6ff90fd7f079e79e1bdfeb4b6ba8a365e1f743cf",
          "url": "https://github.com/ucan-wg/go-ucan/commit/be185a8496f9d2e6ed1d983fbe20457092dc78e0"
        },
        "date": 1731428445328,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5539,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "216639 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5539,
            "unit": "ns/op",
            "extra": "216639 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "216639 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "216639 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 949497,
            "unit": "ns/op\t  147964 B/op\t    2651 allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 949497,
            "unit": "ns/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147964,
            "unit": "B/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11502,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11502,
            "unit": "ns/op",
            "extra": "109764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 962835,
            "unit": "ns/op\t  150027 B/op\t    2653 allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 962835,
            "unit": "ns/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150027,
            "unit": "B/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4408,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "272424 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4408,
            "unit": "ns/op",
            "extra": "272424 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "272424 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "272424 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 934147,
            "unit": "ns/op\t  137419 B/op\t    2554 allocs/op",
            "extra": "1248 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 934147,
            "unit": "ns/op",
            "extra": "1248 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137419,
            "unit": "B/op",
            "extra": "1248 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1248 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8608,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "139202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8608,
            "unit": "ns/op",
            "extra": "139202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "139202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "139202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 951117,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 951117,
            "unit": "ns/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.27,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24485276 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.27,
            "unit": "ns/op",
            "extra": "24485276 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24485276 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24485276 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.42,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "50379027 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.42,
            "unit": "ns/op",
            "extra": "50379027 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "50379027 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "50379027 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 42.59,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28433714 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 42.59,
            "unit": "ns/op",
            "extra": "28433714 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28433714 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28433714 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.5,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "13958404 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.5,
            "unit": "ns/op",
            "extra": "13958404 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "13958404 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13958404 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1919,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "590360 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1919,
            "unit": "ns/op",
            "extra": "590360 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "590360 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "590360 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14755,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81256 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14755,
            "unit": "ns/op",
            "extra": "81256 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81256 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81256 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107285,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107285,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52460,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22677 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52460,
            "unit": "ns/op",
            "extra": "22677 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22677 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22677 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95614,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12513 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95614,
            "unit": "ns/op",
            "extra": "12513 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12513 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12513 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62555,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20564 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62555,
            "unit": "ns/op",
            "extra": "20564 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20564 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20564 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107847,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107847,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54127,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22160 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54127,
            "unit": "ns/op",
            "extra": "22160 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22160 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22160 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 101129,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 101129,
            "unit": "ns/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57765,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20931 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57765,
            "unit": "ns/op",
            "extra": "20931 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20931 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20931 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "257ee16b66f3460e3e1872477094519e33b97638",
          "message": "improve readme",
          "timestamp": "2024-11-12T17:44:28+01:00",
          "tree_id": "ae0f3af8b3852a4fb228bbc8157ea58256ece0d4",
          "url": "https://github.com/ucan-wg/go-ucan/commit/257ee16b66f3460e3e1872477094519e33b97638"
        },
        "date": 1731429943984,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5631,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "217174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5631,
            "unit": "ns/op",
            "extra": "217174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "217174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "217174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 972210,
            "unit": "ns/op\t  147961 B/op\t    2651 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 972210,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147961,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11151,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108328 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11151,
            "unit": "ns/op",
            "extra": "108328 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108328 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108328 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 980523,
            "unit": "ns/op\t  150026 B/op\t    2653 allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 980523,
            "unit": "ns/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150026,
            "unit": "B/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4853,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "262873 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4853,
            "unit": "ns/op",
            "extra": "262873 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "262873 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "262873 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 965683,
            "unit": "ns/op\t  137419 B/op\t    2554 allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 965683,
            "unit": "ns/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137419,
            "unit": "B/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8945,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134805 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8945,
            "unit": "ns/op",
            "extra": "134805 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134805 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134805 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 964836,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 964836,
            "unit": "ns/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.19,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24346251 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.19,
            "unit": "ns/op",
            "extra": "24346251 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24346251 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24346251 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.33,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "50503003 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.33,
            "unit": "ns/op",
            "extra": "50503003 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "50503003 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "50503003 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 40.95,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28724144 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 40.95,
            "unit": "ns/op",
            "extra": "28724144 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28724144 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28724144 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.52,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "13400526 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.52,
            "unit": "ns/op",
            "extra": "13400526 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "13400526 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13400526 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1988,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "566822 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1988,
            "unit": "ns/op",
            "extra": "566822 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "566822 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "566822 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15666,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "80625 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15666,
            "unit": "ns/op",
            "extra": "80625 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "80625 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "80625 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109196,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109196,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53372,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22406 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53372,
            "unit": "ns/op",
            "extra": "22406 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22406 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22406 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98011,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12297 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98011,
            "unit": "ns/op",
            "extra": "12297 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12297 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12297 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58710,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20244 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58710,
            "unit": "ns/op",
            "extra": "20244 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20244 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20244 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109097,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109097,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55252,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "20330 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55252,
            "unit": "ns/op",
            "extra": "20330 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "20330 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "20330 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 101963,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 101963,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57959,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20683 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57959,
            "unit": "ns/op",
            "extra": "20683 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20683 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20683 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "0edeb733980ee1f223f419556ab0a66070888428",
          "message": "improve readme",
          "timestamp": "2024-11-12T17:46:50+01:00",
          "tree_id": "ade4067a45a6865eb1ece5f410584062b490008d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/0edeb733980ee1f223f419556ab0a66070888428"
        },
        "date": 1731430070744,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5596,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "213613 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5596,
            "unit": "ns/op",
            "extra": "213613 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "213613 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "213613 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 952135,
            "unit": "ns/op\t  147962 B/op\t    2651 allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 952135,
            "unit": "ns/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147962,
            "unit": "B/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11577,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109371 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11577,
            "unit": "ns/op",
            "extra": "109371 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109371 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109371 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 969081,
            "unit": "ns/op\t  150020 B/op\t    2653 allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 969081,
            "unit": "ns/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150020,
            "unit": "B/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4440,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "271208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4440,
            "unit": "ns/op",
            "extra": "271208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "271208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "271208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 938494,
            "unit": "ns/op\t  137420 B/op\t    2554 allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 938494,
            "unit": "ns/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137420,
            "unit": "B/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1251 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8648,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "133711 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8648,
            "unit": "ns/op",
            "extra": "133711 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "133711 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "133711 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 949694,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 949694,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23707435 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.12,
            "unit": "ns/op",
            "extra": "23707435 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23707435 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23707435 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.33,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "49955412 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.33,
            "unit": "ns/op",
            "extra": "49955412 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "49955412 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "49955412 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 42.56,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28532989 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 42.56,
            "unit": "ns/op",
            "extra": "28532989 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28532989 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28532989 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.63,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "13989776 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.63,
            "unit": "ns/op",
            "extra": "13989776 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "13989776 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13989776 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1922,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "584119 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1922,
            "unit": "ns/op",
            "extra": "584119 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "584119 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "584119 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14828,
            "unit": "ns/op\t   14576 B/op\t     163 allocs/op",
            "extra": "80150 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14828,
            "unit": "ns/op",
            "extra": "80150 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14576,
            "unit": "B/op",
            "extra": "80150 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "80150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108347,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108347,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52598,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22699 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52598,
            "unit": "ns/op",
            "extra": "22699 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22699 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22699 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96528,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12379 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96528,
            "unit": "ns/op",
            "extra": "12379 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12379 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12379 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61070,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61070,
            "unit": "ns/op",
            "extra": "20552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20552 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112897,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112897,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54515,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54515,
            "unit": "ns/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97708,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97708,
            "unit": "ns/op",
            "extra": "12262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59412,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59412,
            "unit": "ns/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20934 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "0e3869c3c048a9b30804d40371c2e567147b7a0e",
          "message": "improve readme",
          "timestamp": "2024-11-12T18:06:00+01:00",
          "tree_id": "b6f581800a3c5599adfb653699b84d21773a0877",
          "url": "https://github.com/ucan-wg/go-ucan/commit/0e3869c3c048a9b30804d40371c2e567147b7a0e"
        },
        "date": 1731431220423,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5572,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "197296 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5572,
            "unit": "ns/op",
            "extra": "197296 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "197296 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "197296 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 957751,
            "unit": "ns/op\t  147960 B/op\t    2651 allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 957751,
            "unit": "ns/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147960,
            "unit": "B/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11533,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108024 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11533,
            "unit": "ns/op",
            "extra": "108024 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108024 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108024 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 974273,
            "unit": "ns/op\t  150026 B/op\t    2653 allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 974273,
            "unit": "ns/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150026,
            "unit": "B/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4428,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "264578 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4428,
            "unit": "ns/op",
            "extra": "264578 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "264578 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "264578 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 944065,
            "unit": "ns/op\t  137423 B/op\t    2554 allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 944065,
            "unit": "ns/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137423,
            "unit": "B/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8641,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136930 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8641,
            "unit": "ns/op",
            "extra": "136930 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136930 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136930 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 955254,
            "unit": "ns/op\t  139485 B/op\t    2556 allocs/op",
            "extra": "1232 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 955254,
            "unit": "ns/op",
            "extra": "1232 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139485,
            "unit": "B/op",
            "extra": "1232 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 50.29,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24256152 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 50.29,
            "unit": "ns/op",
            "extra": "24256152 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24256152 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24256152 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.26,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "50199864 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.26,
            "unit": "ns/op",
            "extra": "50199864 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "50199864 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "50199864 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 43.16,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28298312 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 43.16,
            "unit": "ns/op",
            "extra": "28298312 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28298312 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28298312 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.77,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "14009782 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.77,
            "unit": "ns/op",
            "extra": "14009782 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "14009782 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14009782 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1927,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "562909 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1927,
            "unit": "ns/op",
            "extra": "562909 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "562909 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "562909 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14778,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "81018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14778,
            "unit": "ns/op",
            "extra": "81018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "81018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "81018 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108184,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108184,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52959,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22689 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52959,
            "unit": "ns/op",
            "extra": "22689 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22689 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22689 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98234,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98234,
            "unit": "ns/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12417 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 60668,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20599 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 60668,
            "unit": "ns/op",
            "extra": "20599 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20599 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20599 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 113041,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 113041,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54834,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21787 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54834,
            "unit": "ns/op",
            "extra": "21787 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21787 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21787 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97677,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12366 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97677,
            "unit": "ns/op",
            "extra": "12366 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12366 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12366 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58911,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58911,
            "unit": "ns/op",
            "extra": "20632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20632 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20632 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "c577d73f3ee7c8c62717be547730675851329b21",
          "message": "improve readme",
          "timestamp": "2024-11-12T18:15:45+01:00",
          "tree_id": "a4b291fd612ae7ea1a9e862ed3d0ad0ddfc9d3f0",
          "url": "https://github.com/ucan-wg/go-ucan/commit/c577d73f3ee7c8c62717be547730675851329b21"
        },
        "date": 1731431807319,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5819,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "204536 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5819,
            "unit": "ns/op",
            "extra": "204536 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "204536 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "204536 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 973808,
            "unit": "ns/op\t  147961 B/op\t    2651 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 973808,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147961,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11006,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108444 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11006,
            "unit": "ns/op",
            "extra": "108444 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108444 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108444 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 972439,
            "unit": "ns/op\t  150027 B/op\t    2653 allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 972439,
            "unit": "ns/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150027,
            "unit": "B/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4398,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "266851 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4398,
            "unit": "ns/op",
            "extra": "266851 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "266851 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "266851 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 935914,
            "unit": "ns/op\t  137417 B/op\t    2554 allocs/op",
            "extra": "1266 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 935914,
            "unit": "ns/op",
            "extra": "1266 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137417,
            "unit": "B/op",
            "extra": "1266 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1266 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8707,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "139544 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8707,
            "unit": "ns/op",
            "extra": "139544 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "139544 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "139544 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 956305,
            "unit": "ns/op\t  139484 B/op\t    2556 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 956305,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139484,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 51.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24504134 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 51.98,
            "unit": "ns/op",
            "extra": "24504134 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24504134 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24504134 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.28,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "51295749 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.28,
            "unit": "ns/op",
            "extra": "51295749 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "51295749 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "51295749 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 40.66,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "27699307 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 40.66,
            "unit": "ns/op",
            "extra": "27699307 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "27699307 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "27699307 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.15,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "14151626 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.15,
            "unit": "ns/op",
            "extra": "14151626 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "14151626 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14151626 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1938,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "566084 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1938,
            "unit": "ns/op",
            "extra": "566084 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "566084 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "566084 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15021,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15021,
            "unit": "ns/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108525,
            "unit": "ns/op\t   15174 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108525,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15174,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54208,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22602 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54208,
            "unit": "ns/op",
            "extra": "22602 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22602 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22602 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96680,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12055 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96680,
            "unit": "ns/op",
            "extra": "12055 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12055 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12055 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61061,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "19286 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61061,
            "unit": "ns/op",
            "extra": "19286 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "19286 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "19286 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107924,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107924,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54644,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "22010 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54644,
            "unit": "ns/op",
            "extra": "22010 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "22010 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22010 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99442,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12348 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99442,
            "unit": "ns/op",
            "extra": "12348 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12348 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12348 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57643,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "20806 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57643,
            "unit": "ns/op",
            "extra": "20806 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "20806 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20806 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9057cbcba62c2396c080b0153582609e68f68817",
          "message": "Merge pull request #59 from ucan-wg/command-covers\n\ncommand: add Covers() for attenuation test, fix incorrect Segments()",
          "timestamp": "2024-11-13T12:38:24+01:00",
          "tree_id": "bb1d937e023d359f199ed3c1b9b5401e8575e623",
          "url": "https://github.com/ucan-wg/go-ucan/commit/9057cbcba62c2396c080b0153582609e68f68817"
        },
        "date": 1731497959300,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5625,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "214816 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5625,
            "unit": "ns/op",
            "extra": "214816 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "214816 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "214816 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 990217,
            "unit": "ns/op\t  147962 B/op\t    2651 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 990217,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147962,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11101,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108070 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11101,
            "unit": "ns/op",
            "extra": "108070 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108070 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108070 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 988374,
            "unit": "ns/op\t  150021 B/op\t    2653 allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 988374,
            "unit": "ns/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150021,
            "unit": "B/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4442,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "264174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4442,
            "unit": "ns/op",
            "extra": "264174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "264174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "264174 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 946342,
            "unit": "ns/op\t  137418 B/op\t    2554 allocs/op",
            "extra": "1215 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 946342,
            "unit": "ns/op",
            "extra": "1215 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137418,
            "unit": "B/op",
            "extra": "1215 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1215 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8578,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8578,
            "unit": "ns/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 954184,
            "unit": "ns/op\t  139483 B/op\t    2556 allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 954184,
            "unit": "ns/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139483,
            "unit": "B/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 52.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24011649 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 52.03,
            "unit": "ns/op",
            "extra": "24011649 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24011649 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24011649 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool",
            "value": 23.33,
            "unit": "ns/op\t       1 B/op\t       1 allocs/op",
            "extra": "51454186 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - ns/op",
            "value": 23.33,
            "unit": "ns/op",
            "extra": "51454186 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - B/op",
            "value": 1,
            "unit": "B/op",
            "extra": "51454186 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bool - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "51454186 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string",
            "value": 41.75,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "28554639 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - ns/op",
            "value": 41.75,
            "unit": "ns/op",
            "extra": "28554639 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "28554639 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/string - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "28554639 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes",
            "value": 84.5,
            "unit": "ns/op\t      52 B/op\t       3 allocs/op",
            "extra": "14043487 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - ns/op",
            "value": 84.5,
            "unit": "ns/op",
            "extra": "14043487 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - B/op",
            "value": 52,
            "unit": "B/op",
            "extra": "14043487 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/bytes - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14043487 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map",
            "value": 1924,
            "unit": "ns/op\t    2248 B/op\t      37 allocs/op",
            "extra": "586272 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - ns/op",
            "value": 1924,
            "unit": "ns/op",
            "extra": "586272 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "586272 times\n4 procs"
          },
          {
            "name": "BenchmarkAny/map - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "586272 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14833,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "80365 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14833,
            "unit": "ns/op",
            "extra": "80365 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "80365 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "80365 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107738,
            "unit": "ns/op\t   15173 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107738,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15173,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52240,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "22897 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52240,
            "unit": "ns/op",
            "extra": "22897 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "22897 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22897 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 101861,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "12474 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 101861,
            "unit": "ns/op",
            "extra": "12474 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "12474 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12474 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57721,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57721,
            "unit": "ns/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107733,
            "unit": "ns/op\t   15126 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107733,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54292,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21981 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54292,
            "unit": "ns/op",
            "extra": "21981 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21981 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21981 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100405,
            "unit": "ns/op\t   14229 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100405,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14229,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56906,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "21134 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56906,
            "unit": "ns/op",
            "extra": "21134 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "21134 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21134 times\n4 procs"
          }
        ]
      }
    ]
  }
}